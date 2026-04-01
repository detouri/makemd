#!/usr/bin/env sh

set -eu

REPO="detouri/makemd"
BINARY="makemd"
VERSION="${VERSION:-latest}"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

os() {
    uname | tr '[:upper:]' '[:lower:]'
}

arch() {
    case "$(uname -m)" in
        x86_64|amd64) echo "amd64";;
        arm64|aarch64) echo "arm64";;
        *) echo "unsupported";;
    esac
}

OS="$(os)"
ARCH="$(arch)"

[ "$ARCH" = "unsupported" ] && echo "unsupported architecture" && exit 1

if [ "$VERSION" = "latest" ]; then
    URL="https://github.com/${REPO}/releases/latest/download/${BINARY}_${OS}_${ARCH}.tar.gz"
else
    URL="https://github.com/${REPO}/releases/${VERSION}/download/${BINARY}_${OS}_${ARCH}.tar.gz"
fi

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT INT TERM

curl -fsSL "$URL" -o "$tmpdir/${BINARY}.tar.gz"
tar -xzf "$tmpdir/${BINARY}.tar.gz" -C "$tmpdir"
install "$tmpdir/${BINARY}" "${INSTALL_DIR}/${BINARY}"
echo "installed ${BINARY} to ${INSTALL_DIR}/${BINARY}"
