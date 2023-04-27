/*
Package cmd provides a command-line interface for changing GHAS settings for a given organization.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gateixeira/github-licenses-report/cmd/github"
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
	
	The historical data is stored in a JSON file called reports.json.
	
	The report is generated in a chart stored as an html file called report.html.
	
	Example:
	
	./github-licenses-report generate-report --enterprise <enterprise_slug> --token <source_token>`,

	Run: func(cmd *cobra.Command, args []string) {
		enterprise, _ := cmd.Flags().GetString(enterpriseFlagName)
		token, _ := cmd.Flags().GetString(tokenFlagName)
		organization, _ := cmd.Flags().GetString(organizationFlagName)
		repository, _ := cmd.Flags().GetString(repositoryFlagName)

		licenses := github.GetConsumedLicenses(enterprise, token)

		report := LicenseReport{
			Date:             time.Now().Format("2006-01-02"),
			ConsumedLicenses: *licenses,
		}

		fileContent, err := github.ReadFile(organization, repository, "reports.json", token)
		if err != nil {
			fmt.Println("Error reading file")
			os.Exit(1)
		}

		var reports []LicenseReport
		json.Unmarshal(fileContent, &reports)

		reports = append(reports, report)

		jsonContent, err := json.Marshal(reports)

		if err != nil {
			fmt.Println("Error marshalling JSON")
			os.Exit(1)
		}

		err = github.UpdateFile(organization, repository, "reports.json", token, jsonContent, "Update reports.json")
		if err != nil {
			fmt.Println("Error writing file")
			os.Exit(1)
		}

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

		file, err := os.Open("report.html")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		buffer := bufio.NewReader(file)
		fileInfo, _ := file.Stat()
		var size int64 = fileInfo.Size()
		bytes := make([]byte, size)
		buffer.Read(bytes)

		err = github.UpdateFile(organization, repository, "report.html", token, bytes, "Update report.html")
		if err != nil {
			fmt.Println("Error writing file")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateOrgCmd)
}
