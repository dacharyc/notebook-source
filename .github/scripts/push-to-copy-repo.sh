#!/usr/bin/env bash
# bash boilerplate
readonly SCRIPT_NAME="$(basename "$0")"
readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
function l { # Log a message to the terminal.
    echo
    echo -e "[$SCRIPT_NAME] ${1:-}"
}

# move to the root the notebook-copy repo
cd "./notebook-copy"
echo "Open root of notebook-copy repo"

# fetch branches from notebook-copy
git fetch
# stash currently copied notebook
git stash
# check out existing branch from notebook-copy
git checkout $BRANCH 
# overwrite any previous copied notebook changes with current ones
git checkout stash -- .

git add -A .
git config user.name github-actions
git config user.email github-actions@github.com
git commit -am "Automatic update"
git push --set-upstream origin $BRANCH

echo "Updated Jupyter Notebook file successfully pushed to notebook-copy repo"
