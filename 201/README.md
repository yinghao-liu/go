# 201

link -X flag

> -X importpath.name=value
>
> Set the value of the string variable in importpath named name to value.
> This is only effective if the variable is declared in the source code either uninitialized
> or initialized to a constant string expression. -X will not work if the initializer makes
> a function call or refers to other variables.

version
`go build -ldflags="-X main.Version=v1.0.0 -X main.version=v2.0.0"`

with date
`go build -ldflags="-X main.Version=v1.0.0 -X 'main.date=$(date)'"`

with subfolder
`go build -ldflags="-X 'version/domain.DomainVersion=v1.0.0'"`

## reference

1. [Using ldflags to Set Version Information for Go Applications](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)
2. go doc cmd/link
3. go help packages
4. go help importpath
