# Template for Go Module

This repository provides a template for Go module.

## Prepare

1. Rename module name in *go.mod*.
2. Remove example package module.

## Development

### Formatting

Format source code.

```sh
go fmt ./...
```

### Linting

Lint source code.

```sh
go tool staticcheck ./...
```

### Testing

Execute test code.

```sh
go test ./...
```

### Updating

Update dependency tools.

```sh
go get -u tool
```

### License Checking

Check dependency license.

```sh
go tool go-licenses report ./...
```

Dump dependency licenses.

```sh
go tool go-licenses save ./... --save_path ./licenses
```

### Documentation

Generate API document.

```sh
./scripts/gen-docs.sh
```

### Building Artifacts

Build binary.

```sh
go build cmd/shell/main.go
```
