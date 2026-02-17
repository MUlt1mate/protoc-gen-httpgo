# Example of using httpgo

This package contains generation examples and test cases for generated code

### Running

#### Run unit tests

```bash
go test ./...
```

#### Run integration tests

```bash
go run .
```

#### Run generation
run `make installdeps install` from parent directory and then

```bash
go generate ./proto/generate.go
```