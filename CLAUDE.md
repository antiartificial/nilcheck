# Nilcheck Project Guidelines

## Build and Test Commands
- Build: `go build`
- Run: `go run nilcheck.go [flags]`
- Run all tests: `go test -v ./...`
- Run specific test: `go test -v -run=TestName`
- Clean cache: `rm .nilcheck_cache.json`

## Code Style Guidelines
- Standard Go formatting (use `gofmt` or `go fmt`)
- Order imports: standard library first, then third-party packages
- Error handling: always check errors and provide descriptive messages
- Variable naming: camelCase for variables, PascalCase for exported functions
- Comments: public functions should have descriptive comments
- Use pointer receivers consistently within types
- Avoid nil pointer dereferences by checking pointers before use
- Use sync.Mutex for protecting shared resources
- Slice handling: verify slices containing pointers have non-nil elements

## CLI Flags
- `-pkgs`: comma-separated list of packages to analyze (e.g., `-pkgs=main,mypkg`)
- `-funcs`: comma-separated list of functions to analyze (e.g., `-funcs=main.main,GetUsers`)
- `-cache`: path to cache file (default: `.nilcheck_cache.json`)