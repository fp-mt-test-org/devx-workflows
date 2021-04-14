# Development Workflow Scripts

Looking for [user documentation](https://flexport.atlassian.net/wiki/spaces/IN/pages/1595900187/Workflow+Scripts+Current+state)?

## Working on the Workflow Scripts

### Getting Started

1. Clone this repository.
2. Configure: `./scripts/configure-localhost.sh`

### Basics

First thing to know is that the workflow scripts are used to build and test themselves.

### Build & Unit Test

To build a new version of the workflow binaries, execute the build workflow:

    ./scripts/build-flex.sh

This will compile, unit test and update the binaries in the `./devx-workflows` directory.

### Feature Test

Unit tests are great, but they don't mean the features are working!

To execute feature tests, execute the test workflow:

    flex test

### Build, Unit & Feature Test Shortcut

It's common to want to run the feature tests after the unit tests, so here's a command to do them in one step:

    flex build-test
