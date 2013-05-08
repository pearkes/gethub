# get

![Demo Gif]()

## Overview

`get` helps you keep all of your git repositories that have GitHub
remotes up to date.

It's opinionated about how you organize your repositories. When
you run `get` the first time, it will build you something like this:

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

It's really useful if you have a lot of repos and an inconsistent network
connection. Before hopping on a plane, train or automobile, run: `get`.

## The Command

    get

1. Checks to see if the necessary requirements for `get` exist
2. If it needs to, asks for your credentials to talk to GitHub
3. Clones any missing repositories
4. Runs `git fetch` in repostories that exist

This is done in parallel as much as possible to speed things up.

## Configuration

Configuration is stored in a `.getconfig` file in your home directory.
(`~/.getconfig`)

### Ignored Repositories or Organizations

Sometimes you don't want to retrieve that gigantic project that
someone committed `.mov` files to.

    repo-ignore: icloud, facebook
    owner-ignore: adobe

### Binary Paths

`get` uses `curl` and `git`. It'll use whatever is in your `PATH`, but
that can be overidden.

    curl-path: /usr/local/bin/curl
    git-path:  /usr/local/bin/git

## Contributing

Check out the [contributing guide]().
