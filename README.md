# vbump

## Installation

`brew install calm/vbump/vbump`

## Usage

After setting up git in a repo (`git init`) but before having tagged the repo with a version, run `vbump init` to initialize the repo with tag `v0.0.1`

After that, you can use one of the following commands to bump the version:

```
vbump major
vbump minor
vbump patch
```

These will bump the major, minor and patch number of the semver accordingly, and push that up to git as a release tag. It will device the current latest version by parsing the latest tag as a semver.

For example, if the latest tag is `v1.7.9`:

```
vbump major > v2.0.0
vbump minor > v1.8.0
vbump patch > v1.7.10
```
