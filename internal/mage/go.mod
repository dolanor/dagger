module github.com/dagger/dagger/internal/mage

go 1.26.1

require (
	dagger.io/dagger v0.21.9
	github.com/dagger/dagger v0.9.11
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/hexops/gotextdiff v1.0.3
	github.com/magefile/mage v1.15.0
	github.com/moby/buildkit v0.13.0-beta3
	github.com/opencontainers/image-spec v1.1.0-rc5
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa
	golang.org/x/mod v0.33.0
	golang.org/x/sync v0.20.0
)

require (
	github.com/99designs/gqlgen v0.17.89 // indirect
	github.com/Khan/genqlient v0.8.1 // indirect
	github.com/adrg/xdg v0.5.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dagger/querybuilder v0.0.0-20260402040506-574a5e81cb59 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sosodev/duration v1.4.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.32 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.41.0 // indirect
	go.opentelemetry.io/otel/metric v1.41.0 // indirect
	go.opentelemetry.io/otel/trace v1.41.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
)

replace github.com/dagger/dagger => ../../
