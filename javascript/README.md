# JavaScript Vector Search Code Example Test PoC

This is a PoC to explore testing JavaScript Vector Search code examples for
MongoDB documentation.

The structure of this JavaScript project is as follows:

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

Create a file named '.env' at the root of the '/javascript' directory within
this project. Add your Atlas connection string as an environment value named
`ATLAS_CONNECTION_STRING`:

```
ATLAS_CONNECTION_STRING="<your-connection-string>"
```

Replace the `<your-connection-string>` placeholder with the connection
string from the Atlas cluster you created in the prior step.

Add an `ENV` environment value whose value is `"Atlas"`. This denotes that
you are running tests against Atlas instead of a local instance. Some functionality
is not supported in local deployment, and some query results vary between the
two environments, so specifying the environment gives the test suite info about
which tests to run and which outputs to expect.

```
ENV="Atlas"
```

### Install the dependencies

This test suite requires you to have `Node.js` v20 or newer installed. If you
do not yet have Node installed, refer to
[the Node.js installation page](https://nodejs.org/en/download/package-manager)
for details.

From the root of the `/javascript` directory, run the following command to install
dependencies:

```
npm install
```

### Run the tests

#### Run All Tests from the command line

From the `/javascript` directory, run:

```
npm test
```

This invokes the following command from the `package.json` `test` key:

```
export $(xargs < .env) && jest
```

The `export $(xargs < .env)` reads the values from your `.env` file and makes
them available to Jest, the test runner. And then it invokes the Jest to
run the tests.

#### Run Test Suites from the command line

You can run all the tests in a given test suite (file).

From the `/tests` directory, run:

```
export $(xargs < .env) && jest -t '<text string from the 'describe' block you want to run>'
```

#### Run Individual Tests from the command line

You can run a single test within a given test suite (file).

From the `/tests` directory, run:

```
export $(xargs < .env) && jest -t '<text string from the 'it' block of the test you want to run>'
```

## To run the tests in CI

A future unit of work will add a GitHub workflow that automatically runs these
tests in CI when someone makes a change to any file in the `/examples` directory.

## To generate tested code examples

This test suite uses [Bluehawk](https://github.com/mongodb-university/Bluehawk)
to generate code examples from the test files. Bluehawk contains functionality
to replace content that we do not want to show verbatim to users, remove test
functionality from the outputted code examples, etc.

Install Bluehawk, and then run the Bluehawk script at the root of the `/javascript`
directory to generate updated example files:

```
./bluehawk.sh
```
