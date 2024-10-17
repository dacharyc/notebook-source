#!/usr/bin/env bash
# bash boilerplate
set -euo pipefail # strict mode
readonly SCRIPT_NAME="$(basename "$0")"
readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
function l { # Log a message to the terminal.
    echo
    echo -e "[$SCRIPT_NAME] ${1:-}"
}

# File to copy from notebook source repo
NOTEBOOK_FILE=./notebook-source/notebooks/hello-world.ipynb

# if the file exists in the notebook source repo, copy it to the notebook copy repo
if [ -f "$NOTEBOOK_FILE" ]; then
    echo "Copying $NOTEBOOK_FILE"
    cp -R ./notebook-source/notebooks/hello-world.ipynb $DESTINATION_PATH
fi

echo "Notebook file copied to $DESTINATION_PATH"
