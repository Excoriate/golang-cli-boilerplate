# Contributing to [Your CLI Name] 🤝

Thank you for your interest in contributing to our [Your CLI Name]! Your contributions play a vital role in enhancing the quality and performance of our tool.

## Pre-requisites 💼

* Minimal Golang version `1.20` or later
* Docker or any compatible container engine for building container images
* Installed `pre-commit` hooks by running `pre-commit install`. This will install the pre-commit hooks in your local repository.
* [Taskfile](https://taskfile.dev/#/) optionally, for running the tasks defined in `Taskfile.yml`. All tasks from `Taskfile.yml` are also available as Makefile targets.

## Getting Started 🎬

* Have a peek at the [open issues](https://github.com/your-github-username/your-repo/issues).
* Suggest a new feature or report a bug by creating a new issue.
* For beginners, look for issues labeled as [good first issue](https://github.com/your-github-username/your-repo/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22).

Your assistance and contributions are always welcome, whether it's bug fixing, new feature suggestions, or documentation improvements.

## Communication ⚡

* [Your preferred communication channel]
* Feel free to get in touch and discuss any issue or feature proposal.

## How to Contribute 🛠️

1. Fork the repo and create a new branch from `main`.
2. Make sure your commit messages adhere to [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) standards for clarity.
3. If you're working on an issue, please assign it to yourself.
4. Your Pull Request should be updated with the `main` branch, and must pass all checks.
5. Ensure that the PR template is completed, and the issue it addresses is linked.
6. Mark your PR as draft until work completion, thereafter, request a review.

## Semantic commits 📝
We use [Semantic Commits](https://www.conventionalcommits.org/en/v1.0.0/) to make it easier to understand what a commit does and to build pretty changelogs. Please use the following prefixes for your commits:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `chores`: Changes to the build process or auxiliary tools and libraries such as documentation generation
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests or correcting existing tests
- `ci`: Changes to our CI configuration files and scripts

An example for this could be:
```
git commit -m "docs: add a new section to the README"
```

## Releasing 🏷️
Releases are done using [Release Please](https://github.com/googleapis/release-please) and [GoReleaser](https://goreleaser.com/). The workflow looks like this:

* A PR is merged to the `main` branch:
  * Release please is triggered, creates or updates a new release PR
  * This is done with every merge to main, the current release PR is updated every time

* Merging the 'release please' PR to `main`:
  * Release please is triggered, creates a new release and updates the changelog based on the commit messages
  * GoReleaser is triggered, builds the binaries and attaches them to the release
  * Containers are created and pushed to the container registry

>**NOTE**: With the next relevant merge, a new release PR will be created and the process starts again

## Building 🛠️
Building the binary is as simple as running `go build .` in the root of the repository. If you want to build the container image, you can run `docker build -t golang-ci-boilerplate -f ./Dockerfile .` in the root of the repository or use any of the Makefile targets or Taskfile tasks.
```bash
# Build the binary
go build .

# Build the container image
docker build -t golang-ci-boilerplate -f ./Dockerfile .

# Build the binary and container image
make build
task cli-build

# Build the docker container
make docker-build
task docker-build
```



## Making a Pull Request 📝

1. Individual PRs for distinct changes will help in better management.
2. Update your PR with the latest changes in `master` branch to keep it conflict-free.
3. Fill out the [PR template](.github/PULL_REQUEST_TEMPLATE.md) and link the associated issue.

## Reviewing PRs 🕵️‍♀️

1. Be constructive and assign yourself to the PR.
2. Confirm all checks are passing.
3. Clarify your doubts with the author.
4. Test the changes to ensure they work as intended.
5. Once all approvals are there and checks are passing, approve and merge the PR.

## Building The Project 💼

Build the project using `go build .` in the repository root, or to create a container image, use `docker build -t your-cli-name -f Dockerfile .`.

## Releasing 🏷️

We use [Release Please](https://github.com/googleapis/release-please) and [GoReleaser](https://goreleaser.com/) to automate releases. Your merged commits will trigger this process.

## We appreciate your contributions! 🙌
