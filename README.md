# Development Workflow Scripts

# What is the devx workflow tool Flex?
Flex is a CLI & config driven x-plat tool for defining and executing workflows.

Similar to tools such as Make, however it's not tightly coupled to software build tools and can be used to execute any CLI tasks, including but not limited to building and testing code.

## Usage

### Installation

To install into a repository:

1. From the root of your repo, execute:
```
bash -c "$(curl -fsSL https://github.com/fp-mt-test-org/devx-workflows/releases/latest/download/install-flex.sh)"
```
3. Run `flex init`

#### Get the Version

You can see the version of flex like so:

    flex -version

## Working on the Workflow Scripts

### Getting Started

1. Clone this repository.
2. Configure: `./scripts/setup-dev-dependencies.sh`

### Basics

First thing to know is that the workflow scripts are used to build and test themselves.

### Build & Unit Test

To build and install a new version of flex, execute the build workflow:

    auto_install=1 ./scripts/build-flex.sh

This will compile, unit test and update the binaries in the `.devx-workflows` directory.

Once you have flex built, you can then use flex to build itself:

    flex build

### Feature Test

Unit tests are great, but they don't mean the features are working!

To execute feature tests, execute the test workflow:

    flex test
