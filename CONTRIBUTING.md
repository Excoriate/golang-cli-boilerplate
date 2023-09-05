# Contributing to Golang CLI Boilerplate 🤝

Thank you for your interest in contributing to our Golang CLI Boilerplate! Our CLI serves as a robust template built on top of Cobra and Viper, coming pre-packed with utilities and AWS adapters via Golang SDK v2. This document simplifies your journey to contribute to our open-source project.

## Pre-requisites 💼

* Golang `1.20` or later
* Docker, or a compatible container engine for building container images
* Ensure you're installing [precommit](https://pre-commit.com/) hooks by running `task pre-commit-init`. This will install the pre-commit hooks in your local repository.
* Optionally, you can install [Taskfile](https://taskfile.dev/#/) to run the tasks defined in the [Taskfile.yml](./TaskFile.yml) file. However, all the Tasks within the `TaskFile.yml` are also available as Makefile targets.

## Getting Started 🎬

* Check our [open issues](https://github.com/Excoriate/golang-cli-boilerplate/issues).
* Propose a new feature or bug fix by opening a new issue for discussion.
* For beginners, look at issues labeled as [good first issue](https://github.com/Excoriate/golang-cli-boilerplate/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22).

Contributions are always welcome whether you have a new feature proposal, a bug fix, or even if you want to improve our documentation.
### Open a GitHub issue

For bug reports or requests, please submit your issue in the appropriate repository.

We advise that you open an issue and ask the
[CODEOWNERS](.github/CODEOWNERS) and community prior to starting a contribution.
This is your chance to ask questions and receive feedback before
writing (potentially wrong) code. We value the direct contact with our community
a lot, so don't hesitate to ask any questions.


## Communication ⚡

* 📧 [Email](mailto:alex_torres@outlook.com)
* 🧳 [LinkedIn](https://www.linkedin.com/in/alextorresruiz/)

Feel free to introduce yourself and discuss any issue or feature proposal.

## How to Make a Contribution 🛠️

1. Fork the repo and create a new branch from `master`.
2. Follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) to make your changes self-explanatory.
3. If you're working on an existing issue, please assign it to yourself. In case you're not a member of our organization, leave a comment to let us know.
4. Ensure your Pull Request is up-to-date with the `master` branch and all checks are passing.
5. Fill out the PR template, link the issue it addresses, and mark it as a draft until work is complete.
6. After work is complete, request a review.

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

1. Keep PRs focused and small. For multiple changes, create multiple PRs.
2. Ensure your PR is up-to-date with the `main` branch and all checks are passing.
3. Fill out the [PR template](.github/PULL_REQUEST_TEMPLATE.md) and link the issue it addresses.

## Reviewing PRs 🕵️‍♀️

1. Be constructive and assign yourself to the PR.
2. Check if all checks are passing.
3. If you are unsure about something, ask the author.
4. Try out the changes to ensure they work as intended.
5. Approve and merge the PR after ensuring it has all approvals and checks are passing.

## Building The Project 💼

You can build the project using `go build .` in the root of the repository, or if you want to create the container image, you can use `docker build -t golang-cli-boilerplate -f Dockerfile .` .

## Releasing 🏷️

We automate releases using [Release Please](https://github.com/googleapis/release-please) and [GoReleaser](https://goreleaser.com/). Your merged commits will automatically trigger this process.

## Thank you for your interest in our project! 🙌
