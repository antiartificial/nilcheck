# nilcheck

A static analysis tool for Go that detects potential nil pointer dereferences.

## Features

- Identifies pointer dereferences without nil checks
- Detects unsafe access of nil struct pointers
- Finds potential issues with slices containing nil pointers
- Analyzes method calls on nil receivers
- Caches results for better performance

## What's next
- golangci-lint support

## Installation

```
go install github.com/antiartificial/nilcheck@latest
```

## Usage

```
nilcheck [flags] [packages]
```

### Flags

- `-pkgs`: Comma-separated list of packages to analyze
- `-funcs`: Comma-separated list of functions to analyze
- `-cache`: Path to cache file (default: `.nilcheck_cache.json`)

### Examples

```
# Analyze a specific package
nilcheck -pkgs=main ./...

# Analyze specific functions
nilcheck -funcs=main.main,GetUsers ./...

# Custom cache location
nilcheck -cache=.custom_cache.json ./...
```

## Development

```
# Build
go build

# Run tests
go test -v ./...

# Clean cache
rm .nilcheck_cache.json
```

## License

[MIT](LICENSE.txt)