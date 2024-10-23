# Go Vector Search Code Example Test PoC

This is a PoC to explore testing Go Vector Search code examples for MongoDB
documentation.

The structure of this Go project is as follows:

- `/example`: This directory contains example code, marked up with Bluehawk,
  that will be outputted to the `/generated` directory when we run the Bluehawk
  script.
- `/tests`: This directory contains the test infrastructure to actually run
  the tests by invoking the example code.

## To run the tests locally

### Create an Atlas cluster

To run these tests locally, you need an Atlas cluster with the sample data set
loaded and no search indexes. For best results, create a fresh M0 cluster, add
sample data, and save the connection string for use in the next step.

### Create a .env file

Create a file named '.env' at the root of the '/go' directory within this
project. Add your Atlas connection string as an environment value named
`ATLAS_CONNECTION_STRING`:

```
ATLAS_CONNECTION_STRING="<your-connection-string>"
```

Replace the `<your-connection-string>` placeholder with the connection
string from the Atlas cluster you created in the prior step.

### Install the dependencies

This test suite requires you to have `Golang` installed. If you do not yet
have Go installed, refer to [the Go installation page](https://go.dev/doc/install)
for details.

From the root of the `/go` directory, run the following command to install
dependencies:

```
go mod download
```

### Run the tests

In your IDE, navigate to the `/tests` directory, and find the test file you want
to run. For example, `/tests/manage-indexes/manage-indexes_test.go`.

Use your IDE to run package tests, run file tests, or run a specific test.

- Package tests: run all the tests in the `tests` directory
- File tests: run only the tests in the given file
- Individual tests: press the Run button next to the name of the test you want to run.

Alternately, if you prefer to run tests from the command line:

#### Run All Tests from the command line

From the `/tests` directory, run:

```
go test -v -p 1 ./...
```

In the above command:

- `-v` is a "verbose" flag that outputs the test output to the console
- `-p 1` is a flag that specifies only running one test at a time (parallel 1)
  to avoid collisions when creating/editing/dropping indexes
- `./...` is the way to specify to run all tests in the directories below this one

#### Run Individual Tests from the command line

```
go test -run TestName
```

## To run the tests in CI

A future unit of work will add a GitHub workflow that automatically runs these
tests in CI when someone makes a change to any file in the `/examples` directory.

## To generate tested code examples

This test suite uses [Bluehawk](https://github.com/mongodb-university/Bluehawk)
to generate code examples from the test files. Bluehawk contains functionality
to replace content that we do not want to show verbatim to users, remove test
functionality from the outputted code examples, etc.

Install Bluehawk, and then run the Bluehawk script at the root of the `/go`
directory to generate updated example files:

```
./bluehawk.sh
```