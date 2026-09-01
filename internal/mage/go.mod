module github.com/dagger/dagger/internal/mage

go 1.26.3

require (
	dagger.io/dagger v0.9.11
	github.com/dagger/dagger v0.9.11
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/hexops/gotextdiff v1.0.3
	github.com/magefile/mage v1.15.0
	github.com/moby/buildkit v0.33.0-rc1
	github.com/opencontainers/image-spec v1.1.1
	golang.org/x/exp v0.0.0-20260603202125-055de637280b
	golang.org/x/mod v0.40.0
	golang.org/x/sync v0.22.0
)

require (
	github.com/99designs/gqlgen v0.17.41 // indirect
	github.com/Khan/genqlient v0.6.0 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/sosodev/duration v1.1.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.10 // indirect
	golang.org/x/sys v0.47.0 // indirect
)

replace github.com/dagger/dagger => ../../
