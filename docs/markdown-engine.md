# Markdown Engine

The markdown layer in `internal/markdown`. It provides a small document model plus helpers for generating and editing README content.

## Core Types

### `Document`

Fields:

- `Title`: H1 text.
- `Preamble`: markdown blocks inserted between the H1 and the first section.
- `Sections`: ordered list of sections.

`Document.Render()`:

- renders the title if present
- renders preamble blocks separated by blank lines
- renders visible sections only
- clamps heading levels to `##` through `######`
- appends HTML comments from `Section.Comment`

### `Section`

Fields:

- `Level`
- `Title`
- `Body`
- `Hidden`
- `Comment`

`NewSection(title, body)` creates a level-2 section.

## Parser

`Parse(content string) Document` supports:

- the first `# ` heading as document title
- `##` and deeper headings as sections
- preamble capture before the first section
- section body preservation with original line breaks
- preamble splitting into paragraph-like blocks

## Editor Operations

`Editor` wraps a parsed `Document` and supports:

- `SetTitle`
- `AddPreamble`
- `SetSection`
- `RemoveSection`
- `AppendSection`
- `UpsertTOC`
- `InsertBadges`
- `Render`

This is the functionality used by the `makemd edit` and `makemd badges` commands.

## Rendering Helpers

The package exposes helpers for common README fragments:

- `Badge(alt, imageURL, linkURL...)`: renders badge markdown, optionally linked.
- `BlockQuote(text)`: prefixes lines with `>`.
- `CodeBlock(lang, code)`: renders fenced code blocks.
- `Bullets(items...)`: renders `-` lists.
- `Numbered(items...)`: renders numbered lists.
- `Table(headers, rows)`: renders pipe tables and escapes `|` in cells.
- `Anchor(title)`: normalizes section titles into anchor-friendly fragments.
- `TOCFromSections(sections)`: builds an in-document table of contents from sections.
- `SortSectionsByTitle(sections)`: sorts sections alphabetically by title.

## Current Limits

The parser/editor is intentionally small and optimized for README-shaped markdown, not full Markdown compliance.

Notable characteristics:

- headings below level 2 are treated as section headings once parsing has entered section mode
- comments are emitted during render, but not parsed back into `Section.Comment`
- TOC links are generated in the repo’s own anchor format, not via a full GitHub slug implementation
