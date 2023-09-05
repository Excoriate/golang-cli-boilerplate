# Golang CLI Boilerplate 🛠️

This is a versatile and easy-to-use template for building a robust Golang CLI with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper). It comes with built-in AWS adapters via Golang SDK v2 and handy utilities.

---
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Excoriate/golang-cli-boilerplate)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)
[![Docker Build](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/docker-build.yml/badge.svg)](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/docker-build.yml)
[![Go Build](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/go-build.yml/badge.svg)](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/go-build.yml)
[![Go linter](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/go-ci-lint.yaml/badge.svg)](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/go-ci-lint.yaml)
[![Go unit tests](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/go-ci-test.yml/badge.svg)](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/go-ci-test.yml)
[![Lint Docker](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/docker-hadolint.yml/badge.svg)](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/docker-hadolint.yml)
[![Yamllint GitHub Actions](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/yaml-linter.yml/badge.svg)](https://github.com/Excoriate/golang-cli-boilerplate/actions/workflows/yaml-linter.yml)

---
## Getting Started 🚦
Use this repository as a [GitHub Template](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template) to create your own CLI:

- [ ] Clone this repository and start adding your command and flag definitions. Utilize pre-installed AWS adapters and included utilities to kickstart your CLI development.
- [ ] This template's CLI is called `golang-cli-boilerplate`. You can rename it by changing the name using your favourite tool, or just using our IDE with a simple `Ctrl+Shift+R` and replace all the occurrences of `golang-cli-boilerplate` with your new name (or including the org/`golang-cli-boilerplate`) if applicable.
- [ ] Configure [PreCommit](https://pre-commit.com/) hooks by running `task pre-commit-init`. This will install the pre-commit hooks in your local repository.
- [ ] Update the `LICENSE.md` file with your own license.
- [ ] Verify that everything is working as expected:
```bash
# If you want to use the MakeFile included.
make lint

# Or, if you're using Taskfile
task go-lint
```
- [ ] After this step, you should be able to run your CLI:
```bash
# TaskFile wraps the binary in a task, so you can run it like this:
task cli-run -- help

# Or directly, just ensure you're building the binary first
go build -o <my-cli> main.go

# Or, with TaskFile
task cli-build
```
>**NOTE**: This template includes a [MakeFile](Makefile) and a [TaskFile](Taskfile.yml) with the necessary (best practices) configuration to build and lint your CLI. Both tools include the same capability, so you can choose the one you prefer.

---

## 🔧 Release your CLI
The release of your CLI is done using [GoReleaser](https://goreleaser.com/). For MacOs, you can use [Homebrew](https://brew.sh/). This template already includes a `.goreleaser.yml` file with the necessary (best practices) configuration to release your CLI.
In addition, a [GitHub Action](.github/workflows/release.yml) is included to automate the release process.
```yaml
  goreleaser:
    if: needs.release-please.outputs.releases_created == 'true'
    permissions:
      contents: write
    needs:
      - release-please
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@8f4b7f84864484a7bf31766abe9204da3cbe65b3 # v3
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@4d34df0c2316fe8122ab82dc22947d607c0c91f9 # v4
        with:
          go-version: '1.20'
      - name: Download Syft
        uses: anchore/sbom-action/download-syft@422cb34a0f8b599678c41b21163ea6088edb2624 # v0.14.1
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@f82d6c1c344bcacabba2c841718984797f664a6b # v4
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{secrets.GH_HOMEBREW_TOKEN}}

```
>**NOTE**: In order to use the GitHub Action, you need to create a `GH_HOMEBREW_TOKEN` secret in your repository with enough permissions to read and write into the `tap` repository.


---
## 📚 Documentation
Documenting your CLI is relevant. This repository includes a [docs](docs/templates/) folder with a template for the documentation of your CLI. You can use it as a starting point for your own documentation. It includes:
- 📃 `README.md` with a standard structure for a CLI repository.
- 📃 `INSTALLATION.md` file with the installation instructions for your CLI.
- 📃 `CONTRIBUTING.md` file with the instructions for contributing to your CLI.
- 📃 `CODE_OF_CONDUCT.md` file with the code of conduct for your CLI.
- 📃 `LICENSE.md` file with the license for your CLI.
```bash
tree -L 3  docs/
docs/
├── about_docs.md
└── templates
    ├── CODE_OF_CONDUCT.md
    ├── CONTRIBUTING.md
    ├── INSTALLATION.md
    ├── LICENSE
    └── README.md
```
>**NOTE**: It's recommended to move these files accordingly, nevertheless it's strongly encouraged to keep a concise documentation structure, keeping the `README.md` simple, concise, and store the more detailed documentation in the `docs` folder.
For more details about the document templates, see [this](docs/about_docs.md).

---
## Features 🧩

* Out-of-the-box environment variable management 🌳
* Auto-scan host environment variables for `AWS` and `Terraform` credentials 📄
* Import env vars from dotfiles (`.env`) 📄
* Leverages built-in AWS adapters ([Golang SDK v2](https://aws.github.io/aws-sdk-go-v2/))
* Ready-to-use utilities for common tasks 🧰. See the [pkg](pkg) folder for more details.
* Built-in [Docker](https://www.docker.com/) support 🐳
* Out-of-the-box [GitHub Actions](https://docs.github.com/en/actions) workflows for CI/CD 🚀
* Built-in [PreCommit](https://pre-commit.com/) hooks for linting and formatting 🧹
* Out-of-the-box support for output data in `yaml`, `json` or `tables`🤖

### Safely share the `cliClient` to `subCommands` 🤝
On each subcommand (at the parent level, which means, those that are in the top of your `pkg`), ensure you're implementing the `GetClient` function:
```go
func GetClient(cmd *cobra.Command) *cli.Client {
	ctx := cmd.Context().Value(cliutils.GetCtxKey())
	if ctx == nil {
		log.Fatal("Unable to get the client context.")
	}
	client, ok := ctx.(*cli.Client)
	if !ok {
		log.Fatal("Unable to assert client.")
	}
	return client
}
```
### Adapters 🧩
Adapters are known also as `clients`. They can plug into the `cliClient` and provide additional functionality. This template includes a subcommand called `aws-ecs` in the `cmd/example` package. It's a subcommand that use the `aws` adapter to read the `ECS` clusters in your account. It's a good example of how to use the `cliClient` and the `aws` adapter. See [here](https://github.com/Excoriate/golang-cli-boilerplate/blob/4caff5eade39799fb3945e52d14f937251233e9a/cmd/example/aws.go#L68-L68)

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
* [ ] Add a built-in `GitHub` adapter.
* [ ] Add a CLI for quickly bootstrapping a new CLI
* [ ] Add an [OpenAI](https://openai.com/) adapter for generating CLI documentation and/or other purposes.

## Contributing
Please read our [contributing guide](./CONTRIBUTING.md). All issues, pull requests and GitHub stars are welcome! Happy coding! 💻


## Community
Find me in:

- 📧 [Email](mailto:alex_torres@outlook.com)
- 🧳 [Linkedin](https://www.linkedin.com/in/alextorresruiz/)


<a href="https://github.com/Excoriate/golang-cli-boilerplate/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Excoriate/golang-cli-boilerplate" />
</a>
