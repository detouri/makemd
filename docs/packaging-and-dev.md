# Packaging And Development Assets

The repo includes scripts and manifests for local builds, installation, and distribution.

## Entry Points

- `cmd/makemd.go`: binary entrypoint that runs `cli.Execute()`.
- `makefile`: local developer targets for formatting, tests, builds, installs, cross-compilation, checksums, and cleanup.
- `scripts/`: helper shell scripts for common install/build flows.
- `packaging/`: release installer and package-manager manifests.

## Make Targets

The `makefile` currently defines:

- `make help`: prints the documented targets.
- `make deps`: runs `go mod tidy`.
- `make fmt`: runs `gofmt` across `cmd`, `cmdspec`, and `internal`.
- `make test`: runs `go test ./...`.
- `make build`: creates `dist/makemd` with linker flags.

The file also advertises `install`, `cross`, `checksums`, and `clean`, but those targets are not implemented in the current file contents.

## Helper Scripts

### `scripts/goinstall.sh`

Wraps:

```sh
go install github.com/your-org/makemd/cmd/makemd@latest
```

### `scripts/curl-install.sh`

Wraps the packaged installer:

```sh
curl -fsSL https://raw.githubusercontent.com/detouri/makemd/main/packaging/install.sh | sh
```

### `scripts/build.sh`

Clones a repository, runs `make build`, and checks `./dist/makemd version`.

The script still references `your-org/makemd`, so it is a scaffold rather than a polished release workflow.

## Packaging Manifests

### `packaging/install.sh`

Provides a POSIX shell installer that:

- detects OS and architecture
- downloads the correct release tarball from GitHub Releases
- extracts the archive
- installs `makemd` into `${INSTALL_DIR:-/usr/local/bin}`

The current script has a typo in the extraction directory variable (`tmldir` instead of `tmpdir`), which would need fixing before production use.

### `packaging/makemd.rb`

Homebrew formula scaffold for macOS and Linux release archives.

Current characteristics:

- version pinned to `0.2.0`
- URLs still point to `your-org/makemd`
- checksums are placeholders

### `packaging/scoop.json`

Scoop manifest for Windows release archives.

Current characteristics:

- homepage points to `detouri/makemd`
- download URLs target `v0.2.0`
- `version` still says `0.1.0`
- archive hashes are placeholders

## Repository Layout

- `internal/cli`: Cobra command implementations.
- `internal/readme`: template catalog, service layer, defaults, and default badges.
- `internal/markdown`: markdown document model, parser, editor, and render helpers.
- `packaging/`: installer/manifests intended for distribution.
- `scripts/`: convenience wrappers around install/build flows.
