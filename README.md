This is just an experiment to learn some things about github actions. Do not use.

```yaml
name: Run Supply Chain Linters
on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main
jobs:
  tests:
    runs-on: ubuntu-latest
    env:
      GO111MODULE: on
    steps:
      - name: Checkout Source
        uses: actions/checkout@v4
      - name: Run Supply Chain Linter
        uses: illume/supply-chain-linters@main
        with:
          args: ./...
```
