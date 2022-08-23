# TF Module Skeleton

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This repository contains a reference terraform module

## How to Use This Repo

This repo is intended to be used as a template for any new TF module. In some cases, we're able to use our VCS to template for us. In other cases, we aren't.

### Prerequisites

- [asdf](https://github.com/asdf-vm/asdf) used for tool version management
- [repo](https://android.googlesource.com/tools/repo) used to pull in all components to create the full repo template

### Repo Init

Run the following commands to prep repo and enable all `Makefile` commands to run

```
asdf plugin add terraform
asdf plugin add tflint
asdf plugin add golang
asdf plugin add golangci-lint
asdf plugin add nodejs
asdf plugin add opa
asdf plugin add conftest
asdf plugin add pre-commit
asdf plugin add terragrunt

asdf install
```

### Templating

#### Manual Templating

This applies to systems like Azure DevOps and CodeCommit.

We need to clone the repo, rename it, and start a fresh git history to get rid of the `tf-module-skeleton` history. Below is a loose explanation of how to do this.

```
git clone <this repo's URL>
mv tf-module-skeleton tf-<whatever it is you're building>
cd tf-<whatever it is you're building>
rm -rf .git
git init
vi .git/HEAD
Change the HEAD to point to `main` instead of `master` and save the `HEAD`
```

#### Remove Educational Material

We need to clear out the example code (different from the boilerplate code). We want to save the repo structure; we don't need the contents. There are `examples`, and `tests` that apply to the boilerplate that we're not going to need as developers of new modules.

Note before you clear these things out, it's useful to actually understand what they are and why they're there. We'll be building our own as we go forward so we need to know what it is we're removing. If this isn't your first module, it's safe to fly through this. If this is your first (or your first several, even), take the time to read the code before you remove it.

```
cd path/to/this/repo
rm -rf examples/*
rm -rf README.md
mv TEMPLATED_README.md README.md
```

#### Version Control (VCS) Templating

This applies to systems like GitHub.

TBD

### Repo Setup

#### Module Configuration

- You'll need to update [`versions.tf`](./versions.tf) based on your provider needs.

## Explanation of the Skeleton

### Resources and Providers

- Terraform modules define resources that can be instantiated via provider
- Possible providers include but are not limited to `aws`, `azurerm`, and `gcp`
- In this module we generate text resources with the `random` provider

### Module Guidelines

- Each repository should have a default module in its root
  - Should have default values and be instantiable with minimal to no inputs
  - We can think of these default values as the "default example"
- A `Makefile` provides tasks for terraform module development
  - It will clone and cache `tf-module-components` and `platform-components` git repositories when a `make configure` is ran
  - For clearing cached components, it provides a `make clean` command
  - Linter config and all other tasks are defined in `tf-module-components`
- An `examples` folder contains example uses of the default and nested modules
  - There should be at least one example for each nested module
- A `tests` folder contains Go functional tests
  - Make pre-deploy tests that validate terraform plan json where applicable
  - Make post-deploy tests that validate the deployment where applicable
- Provider should be configured by the user, not the module
  - Modules only define what providers/versions are required

### Go Functional Tests

- Modules are how Go manages dependencies
- To initiate a new modules run the command: `go mod init [repo_url]`
  - It is recommended to use the absolute repository url (e.g. github.com/nexient-llc/tf-module-skeleton)
- Relative path is highly discouraged in Go, use absolute path to import a package
  - (e.g. `github.com/nexient-llc/tf-module-skeleton/[path_to_file]`)
- To update paths or versions run the command: `go get -t ./...`  go will update the dependencies accordingly
