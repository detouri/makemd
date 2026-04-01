#!/usr/bin/env sh

set -eu

if [ "$#" -ne 1 ]; then
    echo "usage: $0 <version>" >&2
    exit 1
fi

version="$1"
version="${version#v}"

case "$version" in
    *[!0-9.]*|"")
        echo "version must contain only digits and dots" >&2
        exit 1
        ;;
esac

repo_root="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"

printf '%s\n' "$version" > "${repo_root}/VERSION"

perl -0pi -e 's/version = "[^"]+"/version = "'"$version"'"/' "${repo_root}/internal/cli/root.go"
perl -0pi -e 's/version "[^"]+"/version "'"$version"'"/' "${repo_root}/packaging/makemd.rb"
perl -0pi -e 's/"version": "[^"]+"/"version": "'"$version"'"/' "${repo_root}/packaging/scoop.json"
perl -0pi -e 's/releases\/download\/v[^\/"]+/releases\/download\/v'"$version"'/g' "${repo_root}/packaging/scoop.json"

echo "updated versioned files to ${version}"
