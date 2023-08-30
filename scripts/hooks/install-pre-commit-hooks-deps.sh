#!/usr/bin/env bash

set -e
set -u
set -o pipefail

# 1. Installing GUM.
if ! command -v gum > /dev/null; then
  echo "Error: gum is not installed. Please install gum first." >&2
  brew install gum
fi

gum style --foreground 212 "GUM is properly installed ⚡️"

# 2. Installing Homebrew.
if ! command -v brew > /dev/null; then
  gum style --foreground 196 "Homebrew is not installed. Please install Homebrew first."
  gum style --foreground 196 "See https://brew.sh/ for more information."

  brew install
fi

gum style --foreground 212 "Homebrew is properly installed ⚡️"

# 3. Install pre-commit framework.
if ! command -v pre-commit > /dev/null; then
  gum style --foreground 196 "pre-commit is not installed. Please install pre-commit first."
  gum style --foreground 196 "See https://pre-commit.com/ for more information."

  brew install pre-commit
else
  pre-commit autoupdate
fi

gum style --foreground 212 "pre-commit is properly installed ⚡️"

# 4. Install TFEnv
if ! command -v tfenv > /dev/null; then
  gum style --foreground 196 "tfenv is not installed. Please install tfenv first."
  gum style --foreground 196 "Installing TFenv..."
  brew install tfenv
fi

gum style --foreground 212 "tfenv is properly installed ⚡️"

#5. Installing Terraform version from TFEnv
if ! tfenv list | grep -qE '\*'; then
  gum style --foreground 196 "No Terraform version is selected. Please select a Terraform version first."
  gum style --foreground 196 "Installing Terraform version latest"
  tfenv install latest
  tfenv use latest
else
  tfenv update
fi

gum style --foreground 212 "Terraform is properly installed ⚡️"

# 6. installing TFlint
if ! command -v tflint > /dev/null; then
  gum style --foreground 196 "tflint is not installed. Please install tflint first."
  gum style --foreground 196 "Installing TFlint..."
  brew install tflint
else
  brew upgrade tflint
fi

gum style --foreground 212 "tflint is properly installed ⚡️"

# 7. Installing Gobrew as a Go version manager.
if ! command -v gobrew > /dev/null; then
  gum style --foreground 196 "gobrew is not installed. Please install gobrew first."
  gum style --foreground 196 "Installing Gobrew..."
  brew install gobrew
fi

if ! grep -q 'export PATH=' "$HOME/.zshrc" || ! grep -q 'export GOROOT=' "$HOME/.zshrc"; then
  gum style --foreground 196 "gobrew is not configured. Please configure gobrew first."
    # Add the lines at the end of the ~/.zshrc file
  echo "export PATH=\"\$HOME/.gobrew/current/bin:\$HOME/.gobrew/bin:\$PATH\"" >> "$HOME/.zshrc"
  echo "export GOROOT=\"\$HOME/.gobrew/current/go\"" >> "$HOME/.zshrc"
else
  gum style --foreground 212 "gobrew is properly configured ⚡️"
fi

gobrew install 1.19
gobrew use 1.19

gum style --foreground 212 "gobrew is properly installed ⚡️"

# 8. GolangCI
if ! command -v golangci-lint > /dev/null; then
  gum style --foreground 196 "golangci-lint is not installed. Please install golangci-lint first."
  gum style --foreground 196 "Installing golangci-lint..."
  brew install golangci-lint
else
  brew upgrade golangci-lint
fi

gum style --foreground 212 "golangci-lint is properly installed ⚡️"

# 9. ShellCheck
if ! command -v shellcheck > /dev/null; then
  gum style --foreground 196 "shellcheck is not installed. Please install shellcheck first."
  gum style --foreground 196 "Installing shellcheck..."
  brew install shellcheck
else
  brew upgrade shellcheck
fi

gum style --foreground 212 "shellcheck is properly installed ⚡️"

# 10. Installing NVM (node version manager)
if ! command -v nvm > /dev/null; then
  gum style --foreground 196 "nvm is not installed. Please install nvm first."
  gum style --foreground 196 "Installing nvm..."
  brew install nvm
else
  brew upgrade nvm
fi

if ! grep -q 'export NVM_DIR=' "$HOME/.zshrc" || ! grep -q '/opt/homebrew/opt/nvm/nvm.sh' "$HOME/.zshrc" || ! grep -q '/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm' "$HOME/.zshrc"; then
  gum style --foreground 196 "nvm is not properly configured. Please configure nvm first."
  # Add the lines at the end of the ~/.zshrc file
  {
    echo "export NVM_DIR=\"\$HOME/.nvm\""
  } >> "$HOME/.zshrc"
  echo '[ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && \. "/opt/homebrew/opt/nvm/nvm.sh"  # This loads nvm' >> "$HOME/.zshrc"
  echo '[ -s "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm" ] && \. "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion' >> "$HOME/.zshrc"
else
  gum style --foreground 212 "nvm is properly configured ⚡️"
fi

gum style --foreground 212 "nvm is properly installed ⚡️"

# 11. Installing NodeJS
if ! command -v node > /dev/null; then
  nvm install node
fi

if ! command -v markdown-link-check > /dev/null; then
  npm install -g markdown-link-check
else
  npm upgrade -g markdown-link-check
fi

gum style --foreground 212 "markdown-link-check is properly installed ⚡️"

if ! command -v terradoc > /dev/null; then
  gum style --foreground 196 "terradoc is not installed. Please install terradoc first."
  go install github.com/mineiros-io/terradoc/cmd/terradoc@latest
else
  go get -u github.com/mineiros-io/terradoc/cmd/terradoc
fi

gum style --foreground 212 "terradoc is properly installed ⚡️"
