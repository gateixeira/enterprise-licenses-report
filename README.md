# GitHub Licenses Report

Helper tool to generate a report of licenses consumption for a GitHub Enterprise.

## Usage

Run the tool via command line:
```
$ ./enterprise-licenses-report --help
```

### Example

```
$ ./enterprise-licenses-report generate-report --enterprise <enterprise_slug> --token <source_token>
```

### Output

![image](https://user-images.githubusercontent.com/4645845/233187321-99bfb6c6-1c67-440f-936b-e68992a5d482.png)

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
