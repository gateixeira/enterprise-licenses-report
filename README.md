# Enterprise Licenses Report

GitHub Action that can be used as a CLI tool to report license usage and GitHub adoption for an enterprise

## Why?

Enterprise administrators do not have an easy way to measure and visualize growth of GitHub adoption over a certain period of time.

If this tool runs scheduled as a cron job with a certain cadence, it will fetch the data and generate an HTML report as workflow artifact that takes into account historical data from previous runs

## How?

This tool fetches the licenses via API a creates a new issue for each execution against a repository (the one that the action ran against, unless specified otherwise). This new issue contains the licenses data from this execution concatenated with the licenses data from all previous executions.

It then plots a chart in an HTML file that is uploaded to the workflow run's artifacts:

![image](https://user-images.githubusercontent.com/4645845/233187321-99bfb6c6-1c67-440f-936b-e68992a5d482.png)

## Run as Action

Run the tool via command line:

```
$ ./enterprise-licenses-report --help
```

Or set it up as an Action workflow:

```yml
- name: Generate Report
  uses: gateixeira/enterprise-licenses-report@main
  env:
    GITHUB_ENTERPRISE_SLUG: "<enterprise_slug>"
    GITHUB_PAT: ${{ secrets.PAT }}
- name: Upload HTML
  uses: actions/upload-artifact@v3
  with:
    name: Upload Report
    path: ${{ github.workspace }}/*.html
    if-no-files-found: error
```

The PAT token set via secrets has to have enterprise read permissions to fetch the licenses as well as repository permission to create and update issues.

Scopes: `repo`, `read:enterprise`

### Run as CLI

To run it via command line, first do `make build` and then

```
$ ./bin/enterprise-licenses-report generate-report --enterprise <enterprise_slug> --token <source_token> --organization <organization_slug> --repository <repository_slug>
```

Note that via command line the tool does not have the context repository it is running in, as it does when running as an Action workflow. Therefore, `organization` and `repository` flags need to be provided.

## Development

### Requirements

- Go 1.19+

### Build

```
make build
```

### Generate binaries

```
make compile
```

This will generate binaries for Linux, MacOS and Windows in bin/ folder.

## Contributing

Pull requests are welcome ❤️
