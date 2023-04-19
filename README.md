# GitHub Licenses Report

Helper tool to generate a report of licenses consumption for a GitHub Enterprise.

## Usage

Run the tool via command line:
```
$ ./github-licenses-report --help
```

### Example

```
$ ./github-licenses-report generate-report --enterprise <enterprise_slug> --token <source_token>
```

### Output

## Development

### Requirements

* Go 1.19+

### Build

```
make build
```

### Generate binaries

```
make compile
```

This will generate binaries for Linux, MacOS and Windows in bin/ folder.
