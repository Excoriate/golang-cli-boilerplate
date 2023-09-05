# Golang CLI Boilerplate 🛠️

Add your CLI description here


---
 [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE) [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

---
## Getting Started 🚦
Add your content here

---

### Usage 🎬
Add your content here
For a list of commands, run `golang-cli-boilerplate --help`.
```bash
golang-cli-boilerplate --help
```
---

## Features 🧩

* Out-of-the-box environment variable management 🌳
* add your content here
---

### Linters 🧹
This CLI includes a [MakeFile](https://www.gnu.org/software/make/) and a [TaskFile](https://taskfile.dev/#/) with the necessary (best practices) configuration to build and lint your CLI. Both tools include the same capability, so you can choose the one you prefer.
In order to enforce [precommit](https://pre-commit.com/) hooks, run:
```bash
# Using the TaskFile
task pre-commit-init

# Or, using the MakeFile
make pre-commit-init
```

---

## 🔧 Installing
The binaries are managed by [GoReleaser](https://goreleaser.com/), [Homebrew](https://brew.sh/) and [Release please](https://github.com/googleapis/release-please). For a more detailed set of instructions, see the [installation](docs/INSTALLATION.md) file. If you're a `Linux/MacOS` user, you can install it using [Homebrew](https://brew.sh/):
```bash
brew tap Excoriate/homebrew-tap
brew install golang-cli-boilerplate
```
>**NOTE**: The `brew` method requires a valid `GITHUB_TOKEN` environment variable with enough permissions to read and write into the `tap` repository.

---
## 📚 Documentation
Add your content here.

---

## Tooling 🧑‍🔧

This template is equipped with an array of tools to maintain a high standard of code quality and accelerate the development process:

* [Precommit](https://pre-commit.com/) framework for managing and maintaining multi-language pre-commit hooks
* [Taskfile](https://taskfile.dev/#/) is a simpler way to automate tasks
* [Makefile](https://www.gnu.org/software/make/) for managing build workflow
* [GolangCI-linter](https://golangci-lint.run/) for consolidated linting to improve code quality
* [GoReleaser](https://goreleaser.com/) for easy binary release management
* [Docker](https://www.docker.com/) for containerization
* [ShellCheck](https://www.shellcheck.net/) for shell script linting
* [Release please](https://github.com/googleapis/release-please) for automated releases
>**NOTE**: For pre-commit, ensure you're running `task pre-commmit-init` to add the pre-configured hooks into your `git` hooks. This will ensure that the code is linted and formatted before committing it. Also, there are other linters included (yaml, docker, shell, md, etc.).

---
## Roadmap 🗓️
Add your content here

## Contributing
Please read our [contributing guide](./CONTRIBUTING.md). All issues, pull requests and GitHub stars are welcome! Happy coding! 💻


## Community
Find me in:

- 📧 [Email](mailto:your_email@domain.com)
- 🧳 [Linkedin](https://www.linkedin.com/in/myuser/)
