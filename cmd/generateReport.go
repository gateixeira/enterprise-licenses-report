/*
Package cmd provides a command-line interface for changing GHAS settings for a given organization.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gateixeira/enterprise-licenses-report/cmd/github"
	"github.com/spf13/cobra"
)

type LicenseReport struct {
	ConsumedLicenses github.ConsumedLicenses `json:"consumed_licenses"`
	Date             string                  `json:"date"`
}

// migrateOrgCmd represents the migrateOrg command
var migrateOrgCmd = &cobra.Command{
	Use:   "generate-report",
	Short: "Generate GitHub licenses report",
	Long: `This script generates a report with the consumed licenses for a given enterprise.
	
	The historical data is stored as JSON in an issue in the provided repository.
	
	An HTML report is generated in the working directory and uploaded as an artifact if running in GitHub Actions.
	
	Example:
	
	./enterprise-licenses-report generate-report --enterprise <enterprise_slug> --token <source_token> --organization <organization> --repository <repository>`,

	Run: func(cmd *cobra.Command, args []string) {
		enterprise, _ := cmd.Flags().GetString(enterpriseFlagName)
		token, _ := cmd.Flags().GetString(tokenFlagName)
		organization, _ := cmd.Flags().GetString(organizationFlagName)
		repository, _ := cmd.Flags().GetString(repositoryFlagName)

		log.Println("Generating report for enterprise", enterprise)
		log.Println("Using organization", organization)
		log.Println("Using repository", repository)

		licenses, err := github.GetConsumedLicenses(enterprise, token)

		if err != nil {
			fmt.Println("Error getting consumed licenses")
			os.Exit(1)
		}

		report := LicenseReport{
			Date:             time.Now().Format("2006-01-02"),
			ConsumedLicenses: *licenses,
		}

		issue, err := github.GetLatestIssueWithLabel(organization, repository, "report", token)

		if err != nil {
			log.Println("Error getting latest issue with label report")
			os.Exit(1)
		}

		var reports []LicenseReport
		if issue != nil {
			json.Unmarshal([]byte(issue.Body), &reports)
		}

		reports = append(reports, report)

		jsonContent, err := json.Marshal(reports)

		if err != nil {
			fmt.Println("Error marshalling JSON")
			os.Exit(1)
		}

		labels := []string{"report"}
		github.CreateIssue(organization, repository, token, report.Date, jsonContent, labels)

		dates := make([]string, len(reports))
		consumedLicenses := make([]int, len(reports))
		purchasedLicenses := make([]int, len(reports))
		for i, report := range reports {
			dates[i] = report.Date
			consumedLicenses[i] = report.ConsumedLicenses.TotalSeatsConsumed
			purchasedLicenses[i] = report.ConsumedLicenses.TotalSeatsPurchased
		}

		var yAxis YAxis
		yAxis.Label = "Consumed Licenses"
		yAxis.Values = consumedLicenses

		var yAxis2 YAxis
		yAxis2.Label = "Purchased Licenses"
		yAxis2.Values = purchasedLicenses

		GenerateChart(enterprise, dates, yAxis, yAxis2)
	},
}

func init() {

	migrateOrgCmd.Flags().String(enterpriseFlagName, "", "The slug of the enterprise.")
	migrateOrgCmd.Flags().String(tokenFlagName, "", "The token to authenticate with the GitHub API.")
	migrateOrgCmd.Flags().String(repositoryFlagName, "", "The repository to store the report in.")
	migrateOrgCmd.Flags().String(organizationFlagName, "", "The organization that owns the repository.")

	if os.Getenv("GITHUB_ENTERPRISE_SLUG") != "" {
		migrateOrgCmd.Flags().Lookup(enterpriseFlagName).Value.Set(os.Getenv("GITHUB_ENTERPRISE_SLUG"))
	} else {
		migrateOrgCmd.MarkFlagRequired(enterpriseFlagName)
	}

	if os.Getenv("GITHUB_PAT") != "" {
		migrateOrgCmd.Flags().Lookup(tokenFlagName).Value.Set(os.Getenv("GITHUB_PAT"))
	} else {
		migrateOrgCmd.MarkFlagRequired(tokenFlagName)
	}

	if os.Getenv("GITHUB_REPOSITORY") != "" {
		migrateOrgCmd.Flags().Lookup(repositoryFlagName).Value.Set(strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1])
		migrateOrgCmd.Flags().Lookup(organizationFlagName).Value.Set(strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[0])
	} else {
		migrateOrgCmd.MarkFlagRequired(repositoryFlagName)
		migrateOrgCmd.MarkFlagRequired(organizationFlagName)
	}

	rootCmd.AddCommand(migrateOrgCmd)
}
