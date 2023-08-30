#!/usr/bin/env bash
#
# Script to build CLI and move binary to the correct location

# variables
declare BINARY_NAME

# Log a message
log() {
  local MESSAGE=${1}
  echo "${MESSAGE}"
}

# Remove binary file if it exists
remove_old_binary() {
  local binary_full_path
  binary_full_path=$1

  if [[ -f ${binary_full_path} ]]; then
    log "Removing old binary..."
    rm "${binary_full_path}"
  fi
}

# Build binary
build_binary() {
  local binary_full_path
  binary_full_path=$BINARY_NAME

  log "Building binary in path ${binary_full_path}..."
  if ! go build -o "${binary_full_path}"; then
      log "Failed to build binary"
      exit 1
  fi

  log "Binary built successfully"
}

# Add binary to .gitiignore
add_to_gitignore_if_not_exist(){
    local gitignore_file
    gitignore_file=".gitignore"

    if [[ ! -f ${gitignore_file} ]]; then
        log "Failed to find .gitignore file"
        exit 1
    fi

    if [[ $(grep -c "${BINARY_NAME}" "${gitignore_file}") -eq 0 ]]; then
        log "Adding binary to .gitignore file..."
        echo "${BINARY_NAME}" >> "${gitignore_file}"
    fi
}

main() {
  BINARY_NAME=$1

  if [[ -z ${BINARY_NAME} ]]; then
    log "No binary name provided, cannot continue"
    exit 1
  fi

  remove_old_binary "${BINARY_NAME}"
  build_binary
  add_to_gitignore_if_not_exist
}

main "$1"
