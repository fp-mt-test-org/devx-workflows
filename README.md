# Development Workflow Scripts

## Usage

To install into a repository:

1. From the root of your repo, execute:
```
sh -c "$(curl -fsSL https://github.com/fp-mt-test-org/devx-workflows/releases/latest/download/install-flex.sh)"
```
3. Run `flex init`

## Working on the Workflow Scripts

### Getting Started

1. Clone this repository.
2. Configure: `./scripts/user/configure-localhost.sh`

### Basics

First thing to know is that the workflow scripts are used to build and test themselves.

### Build & Unit Test

To build a new version of the workflow binaries, execute the build workflow:

    auto_install=true ./scripts/build-flex.sh

This will compile, unit test and update the binaries in the `.devx-workflows` directory.

Once you have flex built, you can then use flex to build itself:

    flex build

### Feature Test

Unit tests are great, but they don't mean the features are working!

To execute feature tests, execute the test workflow:

    flex test

### Build, Unit & Feature Test Shortcut

It's common to want to run the feature tests after the unit tests, so here's a command to do them in one step:

    flex build-test
