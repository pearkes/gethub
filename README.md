# gethub

![](https://f.cloud.github.com/assets/846194/628861/8090cb46-d104-11e2-84eb-7878aa9057d0.gif)

## Overview
[![Build Status](https://api.travis-ci.org/pearkes/get.png?branch=master)](https://travis-ci.org/pearkes/gethub)

`gethub` helps you keep all of your git repositories that have GitHub
remotes up to date.

## Installation

You can download a binary, `deb` or `exe`, depending on your platform.

- darwin/386: [[binary](http://gethub.jack.ly/0.1.2/darwin_386/gethub_0.1.2_darwin_386.zip)]
- darwin/amd64: [[binary](http://gethub.jack.ly/0.1.2/darwin_amd64/gethub_0.1.2_darwin_amd64.zip)]
- linux/386: [[deb](http://gethub.jack.ly/0.1.2/linux_386/gethub_0.1.2_i386.deb)] [[binary](http://gethub.jack.ly/0.1.2/linux_386/gethub_0.1.2_linux_386.tar.gz)]
- linux/amd64: [[deb](http://gethub.jack.ly/0.1.2/linux_amd64/gethub_0.1.2_amd64.deb)] [[binary](http://gethub.jack.ly/0.1.2/linux_amd64/gethub_0.1.2_linux_amd64.tar.gz)]
- linux/arm: [[deb](http://gethub.jack.ly/0.1.2/linux_arm/gethub_0.1.2_armel.deb)] [[binary](http://gethub.jack.ly/0.1.2/linux_arm/gethub_0.1.2_linux_arm.tar.gz)]
- windows/386: [[exe](http://gethub.jack.ly/0.1.2/windows_386/gethub_0.1.2_windows_386.zip)]
- windows/amd64: [[exe](http://gethub.jack.ly/0.1.2/windows_amd64/gethub_0.1.2_windows_amd64.zip)]

To determine your platform:

    uname -sm

On Darwin, you can copy the binary to your bin:

    cp ~/path/to/gethub /usr/local/bin/

Or, if you have [Go](http://golang.org/) installed:

    go install github.com/pearkes/gethub

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
