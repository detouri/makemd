<a id="readme-top"></a>

<!--
    Project Shields

    https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- Project Logo --->
<br />
<div style="margin-bottom: 54px" align="center">
  <a href="https://github.com/detouri/makemd">
    <img style="aspect-ratio: 4/5;" src="images/logo.png" alt="Logo" height="320">
  </a>

  <h3 align="center">MakeMD</h3>

  <p align="center">
    Small cli tool to create README.md files for your project
    <br />
    <a href="https://github.com/detouri/makemd">
        <strong>Explore the docs »</strong>
    </a>
    <br />
    <br />
    <a href="https://github.com/detouri/makemd">
        View Demo
    </a>
    &middot;
    <a href="https://github.com/detouri/makemd/issues/new?labels=bug&template=bug-report---.md">
        Report Bug
    </a>
    &middot;
    <a href="https://github.com/detouri/makemd/issues/new?labels=enhancement&template=feature-request---.md">
        Request Feature
    </a>
  </p>
</div>

[![CI](https://github.com/detouri/makemd/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/detouri/makemd/actions/workflows/ci.yml?query=branch%3Amain)
[![Build Matrix](https://github.com/detouri/makemd/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/detouri/makemd/actions/workflows/build.yml?query=branch%3Amain)
[![Coverage](https://raw.githubusercontent.com/detouri/makemd/main/.github/badges/coverage.svg)](https://github.com/detouri/makemd/actions/workflows/ci.yml?query=branch%3Amain)

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#-community--donations">Community & Donations</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

There are many great README templates available on GitHub; however, I wanted to create a cli tool that would generate a readme based on the type of project I was working on, and also organise the readme with clear sections.

Here's why:
* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like creating a README from scratch
* Readme files should be clear documentation of the project and easy to follow

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## 🏃‍♂️ Getting Started

This is an example of how to get started with `makemd`.

<!-- TODO: Add a professional get staerted into for this project calleed makemd -->

### Prerequisites

Install one of the following first:

- [Go](https://go.dev/dl/) 1.26+ for `go install`
- [Homebrew](https://brew.sh/) for the formula-based install
- [Scoop](https://scoop.sh/) for the Windows manifest
- `curl` and `tar` for the release installer script

### 📦 Installation

Choose the packaging option that matches your environment.

#### Go install

```sh
go install github.com/detouri/makemd/cmd/makemd@latest
```

#### Homebrew

From a cloned checkout of this repository:

```sh
brew install ./packaging/makemd.rb
```

#### Scoop

From a cloned checkout of this repository:

```powershell
scoop install .\packaging\scoop.json
```

#### Shell installer

Install the latest published release on macOS or Linux:

```sh
curl -fsSL https://raw.githubusercontent.com/detouri/makemd/main/packaging/install.sh | sh
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## ▶ Usage

Use the local docs below for the current command surface, template behavior, markdown engine, and packaging assets in this repository:

- [CLI reference](docs/cli.md)
- [Template system](docs/templates.md)
- [Markdown engine](docs/markdown-engine.md)
- [Packaging and development assets](docs/packaging-and-dev.md)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 🛠 Development

Development in this repository is split across three main packages:

- `internal/readme` contains the README templates, generation service, defaults, and badge helpers.
- `internal/markdown` contains the markdown document model, parser, editor, and rendering helpers.
- `internal/cli` contains the command definitions for creating, editing, listing, and updating README content.

### Local setup

1. Install Go `1.26.1` or newer.
2. Clone the repository and enter the project directory.
3. Sync dependencies:

```sh
go mod tidy
```

### Day-to-day workflow

1. Make changes in the package closest to the behavior you want to update.
2. Format touched Go files:

```sh
gofmt -w ./cmd ./internal
```

3. Run the test suite:

```sh
go test ./...
```

### Working areas

- Template content belongs in `internal/readme/templates.go`.
- README generation flow and defaults belong in `internal/readme`.
- Markdown parsing, section editing, and table-of-contents logic belong in `internal/markdown`.
- Command flags and CLI wiring belong in `internal/cli`.
- Installer and packaging assets live in `scripts/` and `packaging/`.

### Release guide

Use this flow for every shipped release.

1. Pick the next version, for example `0.4.1`.
2. Sync all versioned files in one pass:

```sh
./scripts/set-version.sh 0.4.1
```

3. Review the changed files:
   `VERSION`, `internal/cli/root.go`, `packaging/makemd.rb`, and `packaging/scoop.json`.
4. Commit the version bump on a branch and open a pull request to `main`.
5. Merge the pull request.
   Every push to `main` runs CI, refreshes the coverage badge, builds all OS/arch binaries, and publishes a rolling prerelease from the latest `main` commit.
6. After the merged commit is on `main`, create and push the release tag:

```sh
git tag v0.4.1
git push origin v0.4.1
```

7. The tagged release workflow verifies that the tag points at `main`, rebuilds all release binaries, and publishes the GitHub Release assets for:
   `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`, and `windows/arm64`.

The rolling prerelease is for validating the latest `main` build. The final installable release is created only from a pushed `v*` tag.

<!-- ROADMAP -->
## 🗺 Roadmap

- [x] Add Changelog 🗒
- [x] Add back to top links 🔗
- [ ] Multi-language Support 🌐
    - [ ] French 🇫🇷
    - [ ] German 🇩🇪
    - [ ] Chinese 🇨🇳
    - [ ] Japanese 🇯🇵

See the [open issues](https://github.com/detouri/makemd/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## 💻 Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You may also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/<feature-name>`)
3. Commit your Changes (`git commit -m 'Add some <feature-name>'`)
4. Push to the Branch (`git push origin feature/<feature-name>`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- COMMUNITY AND DONATIONS -->
## 🤑 Community & Donations

If you want to support MakeMD, there are a few useful ways to help:

* Star the project on GitHub: [detouri/makemd](https://github.com/detouri/makemd)
* Share the project with other developers who write READMEs often
* Open issues, suggest features, or contribute code
<!-- * Reach out via [@detouri](https://twitter.com/detouri) if you'd like to discuss sponsorship or donations -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## 🔉 Contact

Your Name - [@detouri](https://twitter.com/detouri)

Project Link: [https://github.com/detouri/makemd](https://github.com/detouri/makemd)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## 🏆 Acknowledgments

A few resources used and may be helpful to you!

* [Markdown styling](https://www.markdownguide.org/basic-syntax/#reference-style-links)
* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Img Shields](https://shields.io)
* [List of README examples](https://github.com/matiassingers/awesome-readme)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
