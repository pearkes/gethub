# get

![](https://f.cloud.github.com/assets/846194/491605/3c2635ee-ba44-11e2-9d04-01b3922aec8a.gif)

## Overview

`get` helps you keep all of your git repositories that have GitHub
remotes up to date.

## Getting Started

The first time you run `get`, you pass it a path.

    $ get .

After authorizing with GitHub, all of your repositories will
be cloned into this path.

The next time you run a `get`, all of your new repositories will be cloned
and your existing repositories will be fetched.

It's useful if you have a lot of repos and may not have an
internet connection. Never leave home without a `get`.™

## Directory Structure

It's opinionated about how you organize your repositories.

    ├── pearkes
    │   ├── get
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

    $ get

1. Checks to see if the necessary requirements for `get` exist,
like it's `~/.getconfig` file.
2. If it needs to, asks for your credentials to talk to GitHub, and
subsequently creates a `~/.getconfig` file for future use.
3. Clones any repositories that are missing.
4. Runs `git fetch` in repositories that exist.

*Performance note:* Clones and fetches are executed in parallell

## Configuration

Configuration is stored in a `.getconfig` file in your home directory.
(`~/.getconfig`)

### Ignored Repositories or Organizations

Sometimes you don't want to retrieve that gigantic project that
someone committed `.mov` files to.

    [ignores]
    repo: icloud, facebook
    owner: adobe

## Contributing

Check out the [contributing guide](CONTRIBUTING.md).
