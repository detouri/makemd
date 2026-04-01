# CLI Reference

`makemd` is the executable entrypoint for the repo. The binary is wired in `cmd/makemd.go`, with Cobra-based commands in `internal/cli`.

## Commands

### `makemd new`

Creates a new README from a built-in template.

Supported flags:

- `--template`: required template name.
- `--title`: required project title.
- `--description`: short summary for the generated README preamble.
- `--owner`: repository owner or organization.
- `--repo`: repository name.
- `--module`: Go module path.
- `--binary`: binary or CLI name.
- `--install`: install command rendered into generated docs.
- `--run`: run command rendered into generated docs.
- `--test`: test command rendered into generated docs.
- `--coverage`: coverage command value stored in config defaults.
- `--docs-url`: external documentation URL.
- `--demo-url`: demo URL.
- `--issues-url`: issue tracker URL.
- `--ci`: CI provider, default `github-actions`.
- `--license`: license name, default `MIT`.
- `--audience`: target audience text.
- `--status`: project status badge value.
- `--go`: minimum Go version.
- `--private`: marks the project as internal/closed-source.
- `-o, --output`: output path, default `README.md`.
- `--force`: overwrites an existing file.

Behavior:

- Resolves missing values through `internal/readme.applyDerivedDefaults`.
- Validates that `template` and `title` are present.
- Uses `internal/readme.Service.Generate` to build the final markdown.

Example:

```sh
makemd new \
  --template cli \
  --title "MakeMD" \
  --owner detouri \
  --repo makemd \
  --binary makemd
```

### `makemd edit`

Edits an existing markdown file through the in-repo parser/editor.

Supported flags:

- `--file`: target markdown file, default `README.md`.
- `--set`: replace a section body using `title=<markdown>`.
- `--remove`: remove a section by title.
- `--append`: append a new section using `title=<markdown>`.
- `--title`: replace the H1 document title.
- `--toc`: rebuild the table of contents section.

Behavior:

- Parses the file into a `markdown.Document`.
- Rewrites matching sections case-insensitively.
- Supports escaped `\n` and `\t` values inside `--set` and `--append`.

Example:

```sh
makemd edit --file README.md --set "Usage=Run \`makemd --help\` first." --toc
```

### `makemd badges`

Prepends badge markdown into the README preamble.

Supported flags:

- `--file`: target markdown file, default `README.md`.
- `--add`: badge in the form `alt|imageURL|linkURL(optional)`.

Behavior:

- Converts each badge definition with `internal/markdown.Badge`.
- Inserts the rendered badge line at the top of the preamble.

Example:

```sh
makemd badges --add "CI|https://img.shields.io/badge/ci-passing-green|https://example.com"
```

### `makemd list`

Lists the built-in README templates exposed by `internal/readme.Templates()`.

Current template names:

- `app`
- `cli`
- `closed-source-lib`
- `open-source-lib`
- `package`

### `makemd version`

Prints the version configured on the root command.
