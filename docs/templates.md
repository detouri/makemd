# Template System

The template system lives in `internal/readme`. It turns a `ProjectConfig` into a rendered `markdown.Document`.

## Generation Flow

1. `internal/cli/new.go` collects command flags into `readme.ProjectConfig`.
2. `readme.Service.Generate` validates the request and applies derived defaults.
3. `readme.Templates()` resolves the selected template.
4. A template builder assembles a markdown document with standard sections.
5. `markdown.Document.Render()` converts the document model into final markdown.

## ProjectConfig Fields

The generator can populate content from:

- project identity: `Title`, `Description`, `Template`, `Owner`, `Repo`, `Module`, `BinaryName`
- command snippets: `InstallCommand`, `RunCommand`, `TestCommand`, `CoverageCommand`
- links: `DocsURL`, `DemoURL`, `IssuesURL`
- metadata: `Audience`, `Status`, `MinimumVersion`, `CIProvider`, `LicenseName`
- feature toggles: `Private`, `IncludeBadges`, `IncludeTOC`, `IncludeFAQ`, `IncludeRoadmap`, `IncludeSecurity`, `IncludeContrib`, `IncludeArch`, `IncludeAPI`, `IncludeConfig`

## Derived Defaults

If callers omit values, `applyDerivedDefaults` fills them in:

- description falls back to a generic project summary.
- repo is slugified from the title.
- module defaults to `github.com/<owner>/<repo>`.
- binary name uses the first slug word.
- install/run/test/coverage commands get Go-oriented defaults.
- docs/demo/issues URLs are synthesized.
- private projects default to `Proprietary` and `internal`.
- public projects default to `MIT` and `active`.
- audience, minimum version, and CI provider also get defaults.

## Built-in Templates

### `open-source-lib`

Best for reusable public libraries.

Generated sections include:

- badges and table of contents
- features
- installation
- quick start
- usage
- API
- configuration
- examples
- testing
- roadmap
- contributing
- security
- FAQ
- license

### `closed-source-lib`

Best for internal/shared libraries.

Generated sections include:

- badges and table of contents
- features
- installation
- quick start
- usage
- API
- configuration
- examples
- testing
- compatibility
- support
- FAQ
- ownership

This template also forces `Private=true`.

### `package`

Best for smaller importable packages.

Generated sections include:

- optional badges
- table of contents
- installation
- usage
- API
- examples
- testing
- license

### `app`

Best for applications, services, daemons, and internal products.

Generated sections include:

- badges and table of contents
- features
- screenshots
- prerequisites
- installation
- configuration table
- running
- usage
- architecture
- deployment
- testing
- roadmap
- troubleshooting
- contributing
- license

### `cli`

Best for command-line tools.

Generated sections include:

- badges and table of contents
- features
- installation
- quick start
- commands table
- configuration
- shell completion
- examples
- testing
- FAQ
- license

## Badge Helpers

`internal/readme/badges.go` generates default badge markdown from project metadata.

Current behavior:

- GitHub Actions projects get CI and issues badges.
- non-public projects skip the license badge.
- all projects get a status badge.

The helper currently contains placeholder and typo-prone badge hostnames in some branches, so packaging the generated README may require cleanup.
