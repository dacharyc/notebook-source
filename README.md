# Code Example Source Repo - Demo

This repository represents the source for _all_ code examples in MongoDB
documentation. The code examples in this repo will be written in
test suites, and the tested code will be outputted to a `generated`
directory for inclusion in other docs sets.

Alongside the tested code, we may include the source for other code used
in our documentation. For this proof-of-concept, we have a `notebooks`
directory that contains Jupyter Notebooks demonstrating Atlas Vector Search
functionality.

## Workflow

As the source repo, any updates to code in the documentation should
occur in _this_ repository. GitHub Workflows, Actions, and other tools
will run to automatically push the updated code out to other destinations,
based on need.

For this proof-of-concept, updates to the Jupyter Notebooks in the `notebooks`
directory automatically get pushed out to the [notebook-copy](https://github.com/dacharyc/notebook-copy)
repository. The workflows in the `.github` directory copy the files and push
them to the destination repository.

In the `notebook-copy` repo, the notebooks are displayed as a top-level item.
They can have their own README, and the file structure matches the file
structure from _within_ the `notebooks` folder in this repository. Copying
this directory out to a separate repository makes it easy for users who are
interested in Jupyter Notebooks to clone the repo and consume the content,
_without_ all of the additional examples and infrastructure that will live
alongside this content in the source repository.

Updates to any part of this repository _outside_ of the `notebooks` directory
do not kick off any workflows. In the real source repo, there will be many
additional directories that each have their own workflows to run automated
tests and push out updated code example files to destination repositories.

## Jupyter Notebooks

Currently, the Jupyter Notebooks are manually created and maintained independent
of code examples in the documentation.

Updates to these notebooks must occur independently but in step with updates
to the documentation.

If we find we are regularly updating or adding Python examples, we can invest
the time to write tooling to parse Jupyter Notebook files and extract the
Python code (or vice versa?) to single source the Notebooks and Python examples
in the documentation. That will be a more time-intensive process, so would
be a separate future unit of work scoped independently.
