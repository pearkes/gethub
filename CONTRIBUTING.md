# Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Reporting a Bug

If you don't mind running the problem command (if applicable) in debug
mode and dumping the output in your GitHub issue, it'd be very helpful.

    get [<command>] --debug

## Development Environment

To add a feature, fix a bug, or to run a development build of `get`
on your machine, you'll need to have [go](http://golang.org/) installed.

You can then build a binary:

    $ go build

And run it:

    $ ./get

To install the development binary on your system:

    $ chmod +x ./get
    $ [sudo] cp get /usr/bin/

If you need help with your environment, feel free to open an issue.
