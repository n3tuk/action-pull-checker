# n3tuk Pull Request Checker

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
        uses: n3tuk/action-pull-requester@v1.0.0
```
