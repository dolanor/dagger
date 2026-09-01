module dagger.io/dagger

go 1.21

// retract engine releases from SDK releases
retract [v0.0.0, v0.2.36]

require (
	github.com/99designs/gqlgen v0.17.31
	github.com/Khan/genqlient v0.6.0
	github.com/adrg/xdg v0.4.0
	github.com/stretchr/testify v1.12.1
	github.com/vektah/gqlparser/v2 v2.5.6
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa
	golang.org/x/sync v0.6.0
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect

require (
	github.com/mitchellh/go-homedir v1.1.0
	golang.org/x/sys v0.14.0 // indirect
)
