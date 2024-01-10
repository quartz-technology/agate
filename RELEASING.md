# Releasing Agate

This document describes the procedure to follow when releasing a new version of `Agate`.

## Prerequisites

Prior to complete the following tasks, make sure that you have the following tools installed:
- [Changie](https://changie.dev/guide/installation/)

## Versioning

`Agate` follows the [semver](https://semver.org/) convention for releasing new versions of the 
project.

## Procedure

### 1. Document the changes

Lets consider that the current version is `v0.1.5` and the next is `v0.2.0` and
that 3 PRs have been merged since the release of `v0.1.5`

- **PR #17**: `Refactor data aggregator` by Alice.
- **PR #18**: `Update code documentation` by Bob.
- **PR #19**: `Improve data aggregator performance` by Pierre.

The developer in charge of each PR should use `changie` to document those changes right before 
merging - unless stated otherwise by the maintainer:

For Alice:
```shell
changie new --kind "Breaking" --body "Refactor data aggregator" --custom "PR=17" --custom "Author=Alice"
```

For Bob:
```shell
changie new --kind "Changed" --body "Update code documentation" --custom "PR=18" --custom "Author=Bob"
```

For Pierre:
```shell
changie new --kind "Added" --body "Improve data aggregator performance" --custom "PR=19" --custom "Author=Pierre"
```

### 2. Prepare the release

The documentation written by the developers in the previous section allows the maintainer to batch 
relevant documented 
changes for the next 
release.

To accomplish this, create a new branch `feat/vX.Y.Z-release-notes` where `vX.Y.Z` is the new
version to be released.

Then, use `changie` to batch the documented changes:
```shell
changie batch vX.Y.Z
```

Commit and push the changes made to the `CHANGELOG.md`.

Finally, merge this PR to the main branch.

### 3. Create a new release

On the main branch, the maintainer can create a new version by tagging the current latest commit:
```shell
git tag vX.Y.Z
git push origin vX.Y.Z
```

This should trigger a GitHub action to run and create a new release of the project.
