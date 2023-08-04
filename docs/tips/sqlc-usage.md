# sqlc usage

```shell
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

## App Description

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "tutorial"
        out: "tutorial"
```

## Generate

```shell
sql genearte
```