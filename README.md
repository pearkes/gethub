# gethub

![](https://f.cloud.github.com/assets/846194/618342/7f7cc24a-ceb2-11e2-9bdb-0eb19f0dd552.gif)

## Overview
[![Build Status](https://travis-ci.org/pearkes/get.png?branch=master)](https://travis-ci.org/pearkes/get)

`gethub` helps you keep all of your git repositories that have GitHub
remotes up to date.

## Installation

Please see the [issue](https://github.com/pearkes/get/issues/2).

*This is temporary until binaries are properly compiled for common platforms.*

## Getting Started

    $ gethub authorize

This asks you where you want to clone your repositories as well
as creating an OAuth token for future GitHub requests.

The next time you run a `gethub`, all of your new repositories
will be cloned and your existing repositories will be fetched.

It's useful if you have a lot of repos and may not have an
internet connection.

Never leave home without running `gethub`.

## Directory Structure

It's opinionated about how you organize your repositories.

    ├── pearkes
    │   ├── gethub
    │   ├── tugboat
    │   └── jack.ly
    ├── mitchellh
    │   └── vagrant
    ├── amadeus
    │   └── html7
    ├── someorg
    │   └── bigproject
    └── someotherorg
        └── biggerproject

Basically, your repositories will be name-spaced according
to who the owner is on GitHub.

## Behind the Curtain

    $ gethub

1. Checks to see if the necessary requirements for `gethub` exist,
like it's `~/.gethubconfig` file.
2. Makes sure the path to your repositories looks ok.
3. Clones any repositories that are missing.
4. Runs `git fetch` in repositories that exist.

## Configuration

Configuration is stored in a `.gethubconfig` file in your home directory.
(`~/.gethubconfig`)

### Ignored Repositories

Sometimes you don't want to retrieve that gigantic project that
someone committed `.mov` files to.

    [ignores]
    repo: icloud, facebook
    owner: adobe

## Contributing

Check out the [contributing guide](CONTRIBUTING.md).
