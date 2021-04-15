# Development Workflow Scripts

## Usage

To install into a repository:

1. From the root of your repo, execute:

    sh -c "$(curl -fsSL https://github.com/fp-mt-test-org/devx-workflows/releases/latest/download/install-flex.sh)"


```
VERSION=0.2.4 bash -c 'mkdir -p .devx-workflows/darwin_amd64 && curl -LX GET "https://github.com/fp-mt-test-org/devx-workflows/releases/download/v${VERSION}/devx-workflows_${VERSION}_darwin_amd64.tar.gz" | tar -xvz -C .devx-workflows/darwin_amd64'
```

```
VERSION=0.2.4 bash -c 'mkdir -p .devx-workflows/linux_amd64 && curl -LX GET "https://github.com/fp-mt-test-org/devx-workflows/releases/download/v${VERSION}/devx-workflows_${VERSION}_linux_amd64.tar.gz" | tar -xvz -C .devx-workflows/linux_amd64'
```

2. Add alias flex="./.devx-workflow/darwin_amd64/flex" to your ~/.zshrc / ~/.bashrc so you can just run flex [command]. For more info on how to use flex locally from your repo, run flex help. 

3. Run `flex init`

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
