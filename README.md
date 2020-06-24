# MVG info

MVG info is a simple tool to fetch interruption notifications for Munich's public transport system. 
It currently outputs the data in a [BitBar]-compatible way.

## Installation

1. Install [BitBar]. If you're using [homebrew], you can use `brew cask install bitbar`.
1. Download the binary for Mac OS from the [Releases] page into your BitBar plugin directory. 

Alternatively, you can clone the repository and run `make build`

## Developing

It's a standard project with Gomodules. There are lots of more or less helpful targets in the `Makefile`. 

[BitBar]: https://getbitbar.com
[Releases]: https://github.com/muffix/mvg-info/releases/latest
[homebrew]: https://brew.sh
