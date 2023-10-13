#!/usr/bin/env bash

#
# Script Name: Air Setup Manager
#
# Author: Alex Torres (github.com/Excoriate), alex_torres@outlook.com
#
# Usage: ./air.sh [--force]
#
# Description: This bash script is designed to manage the setup process for the Air utility
# a Live reload for Go apps. The script checks whether the Air utility is installed,
# and then proceeds to create an Air configuration file (".air.toml") in the current directory.
# Existing "air.toml" files are not overwritten unless the --force flag is used.
#
# Parameters:
#    --force:    If passed, the existing air.toml will be overwritten.
#
# Examples:
#    Initial setup:              ./air.sh
#    Overwrite existing setup:   ./air.sh --force
#
# Note: The generated "air.toml" file contains some standard configuration for the Air utility.
#       Feel free to adjust it according to your needs.
#
# For further details and support, contact the author.
#
##################################################################################

set -euo pipefail

# Constants
readonly AIR_BINARY="air"
readonly AIR_TOML_FILE=".air.toml"

# Log a message
log() {
  local MESSAGE="$1"
  echo "${MESSAGE}"
}

# Check if Air binary exists
# Check if Air binary exists
check_air_installation() {
  local go_installation
  local bin_installation
  go_installation="$HOME/go/bin/air"
  bin_installation="/usr/local/bin/air"
  local paths=("$go_installation" "$bin_installation")

  # Check direct command
  if command -v ${AIR_BINARY} &> /dev/null; then
    log "${AIR_BINARY} is already installed and available in PATH."
    return
  fi

  # Check known possible paths
  for path in "${paths[@]}"; do
    if [[ -x ${path} ]]; then
         log "${AIR_BINARY} is installed at ${path} but not available in PATH."
         log "Consider adding ${path%/*} to PATH environment variable."
         return
    fi
  done

  # If air binary is not found
  log "${AIR_BINARY} is not installed. Install it via 'GO111MODULE=on go get -u github.com/cosmtrek/air'."
  exit 1
}

# Check if Air config file exists
check_air_config_file() {
  local FORCE=$1

  if [[ -f ${AIR_TOML_FILE} && "${FORCE}" == "false" ]]; then
    log "Configuration file ${AIR_TOML_FILE} already exists. Skipping creation..."
  else
    log "Creating air configuration file..."
    cat <<EOT > ${AIR_TOML_FILE}
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/main.go"
  delay = 0
  exclude_dir = ["tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = ["index.html"]
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = true

[screen]
  clear_on_rebuild = false
  keep_scroll = true

EOT

    if ! [[ -f ${AIR_TOML_FILE} ]]; then
      log "Failed to create ${AIR_TOML_FILE}"
      exit 1
    fi
  fi
}

# Main function
main() {
  local FORCE="false"

  # Args parsing
  while (( "$#" )); do
    case "$1" in
      --force)
        FORCE="true"
        shift
        ;;
      *)
        log "Illegal option!"
        exit 1
        ;;
    esac
  done

  # Check existing air installation
  check_air_installation

  # Check if air config file exists
  check_air_config_file ${FORCE}
}

# Execute the script with given args
main "$@"
