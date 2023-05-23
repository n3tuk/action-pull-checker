# n3tuk Pull Request Checker

[![GitHub go-integrations Workflow Status](https://img.shields.io/github/actions/workflow/status/n3tuk/action-pull-requester/go-integrations.yaml?label=go-integrations&style=flat-square)](https://github.com/n3tuk/action-pull-requester/actions/workflows/go-integrations.yaml)
[![GitHub codeql Workflow Status](https://img.shields.io/github/actions/workflow/status/n3tuk/action-pull-requester/codeql.yaml?label=codeql&style=flat-square)](https://github.com/n3tuk/action-pull-requester/actions/workflows/codeql.yaml)
[![GitHub go-releaser Workflow Status](https://img.shields.io/github/actions/workflow/status/n3tuk/action-pull-requester/go-releaser.yaml?label=go-releaser&style=flat-square)](https://github.com/n3tuk/action-pull-requester/actions/workflows/go-releaser.yaml)
[![Codecov Status](https://codecov.io/gh/n3tuk/action-pull-requester/branch/main/graph/badge.svg?token=ZTYAZGRQG5)](https://codecov.io/gh/n3tuk/action-pull-requester)

A GitHub Action for running standard checks and automations on pull requests for
the `n3tuk` Organisation.

## Usage

You can use the [pull-requester GitHub Action][pull-requester-marketplace] in a
[GitHub Workflow][github-workflow] by configuring a YAML file in your GitHub
repository (under `.github/workflows/pull-requester.yaml`), with the following
contents:

[github-workflow]: https://help.github.com/en/articles/about-github-actions
[pull-requester-marketplace]: https://github.com/marketplace/actions/pull-requester

```yaml
---
name: Pull Requester

on:
  pull_request:
    types:
      - opened
      - reopened
      - synchronize
      - edited
      # Catch when added labels are forcefully removed
      - unlabeled
    branches:
      - main

permissions:
  contents: read
  packages: read
  issues: write
  pull-requests: write

jobs:
  pull-requester:
    runs-on: ubuntu-latest
    name: Check the Pull Request

    concurrency:
      # Ensure that GitHub runs a single concurrent job for any Pull Requester
      # event on any one pull request (i.e. github.event.number), and bias that
      # to the latest job started, which will have access to the latest settings
      group: pull-requester-${{ github.event.number }}
      cancel-in-progress: true

    steps:
      - name: Pull Requester
        uses: n3tuk/action-pull-requester@v1
```

> **Note**:
> Do **not** use the `main` branch (or any other branch) as a reference for the
> GitHub Action, as the binaries for the GitHub Action builds and releases on
> tagging, and the Action downloads those on running.

## Inputs

| Name | Description | Required | Default |
| :--- | :---------- | :------: | :------ |
| `title-minimum` | The minimum number of characters that a title should contain | `false` | `25` |
