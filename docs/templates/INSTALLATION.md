## 🔧 How to Install

The CLI golang-cli-boilerplate provides binary distributions for every release which are generated using [GoReleaser](https://goreleaser.com/). To install it, you can use the pre-built binaries which are available for Linux, Windows, and macOS:
1. Navigate to the [Releases](https://github.com/Excoriate/golang-cli-boilerplate/releases) page.
2. Download the archive for your respective OS and architecture.
3. Extract the archive.
4. Move the `golang-cli-boilerplate` binary into your `$PATH`.

Or, based on your OS. For Mac, you can use [Homebrew](https://brew.sh/):

```bash
brew tap Excoriate/homebrew-tap https://github.com/Excoriate/homebrew-tap.git
brew install golang-cli-boilerplate
```
>**NOTE**: There are compiled binaries available for most of the common platforms, including Windows. Check the
[Releases](https://github.com/Excoriate/golang-cli-boilerplate/releases) page.

### Linux/Mac via brew

```
brew tap Excoriate/homebrew-tap
brew install golang-cli-boilerplate
```

<details>
  <summary>RPM-based installation (RedHat/CentOS/Fedora)</summary>

**32 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/Excoriate/golang-cli-boilerplate/releases/download/v0.3.14/golang-cli-boilerplate_386.rpm
  sudo rpm -ivh golang-cli-boilerplate_386.rpm
  ```
  <!---x-release-please-end-->

**64 bit:**

  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/Excoriate/golang-cli-boilerplate/releases/download/v0.3.14/golang-cli-boilerplate_amd64.rpm
  sudo rpm -ivh -i golang-cli-boilerplate_amd64.rpm
  ```
  <!---x-release-please-end-->
</details>

<details>
  <summary>DEB-based installation (Ubuntu/Debian)</summary>

**32 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/Excoriate/golang-cli-boilerplate/releases/download/v0.3.14/golang-cli-boilerplate_386.deb
  sudo dpkg -i golang-cli-boilerplate_386.deb
  ```
  <!---x-release-please-end-->
**64 bit:**

  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/Excoriate/golang-cli-boilerplate/releases/download/v0.3.14/golang-cli-boilerplate_amd64.deb
  sudo dpkg -i golang-cli-boilerplate_amd64.deb
  ```
  <!---x-release-please-end-->
</details>

<details>

  <summary>APK-based installation (Alpine)</summary>

**32 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/Excoriate/golang-cli-boilerplate/releases/download/v0.3.14/golang-cli-boilerplate_386.apk
  apk add golang-cli-boilerplate_386.apk
  ```
  <!---x-release-please-end-->
**64 bit:**
  <!---x-release-please-start-version-->
  ```
  curl -LO https://github.com/Excoriate/golang-cli-boilerplate/releases/download/v0.3.14/golang-cli-boilerplate_amd64.apk
  apk add golang-cli-boilerplate_amd64.apk
  ```
  <!---x-release-please-end-->x
</details>

<details>
  <summary>Failing Installation on WSL or Linux (missing gcc)</summary>
  When installing Homebrew on WSL or Linux, you may encounter the following error:

  ```
  ==> Installing golang-cli-boilerplate from golang-cli-boilerplate-ai/golang-cli-boilerplate Error: The following formula cannot be installed from a bottle and must be
  built from the source. golang-cli-boilerplate Install Clang or run brew install gcc.
  ```

If you install gcc as suggested, the problem will persist. Therefore, you need to install the build-essential package.
  ```
     sudo apt-get update
     sudo apt-get install build-essential
  ```
</details>


### Windows

* Download the latest Windows binaries of **golang-cli-boilerplate** from the [Release](https://github.com/Excoriate/golang-cli-boilerplate/releases)
  tab based on your system architecture.
* Extract the downloaded package to your desired location. Configure the system *path* variable with the binary location
