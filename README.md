# MVG info

MVG info is a simple tool to fetch interruption notifications for Munich's public transport system. 
It currently outputs the data in a [BitBar]-compatible way.

![](https://user-images.githubusercontent.com/12208771/85635106-67641180-b67d-11ea-82a9-9530a68c1138.png)

## Installation

1. Install [BitBar]. If you're using [homebrew], you can use `brew cask install bitbar`.
1. Install the plugin by doing one of the following options:
    1. Copy the following link into your browser's address bar and confirm: `bitbar://openPlugin?src=https://github.com/muffix/mvg-info/releases/latest/download/mvginfo.10m` 
    1. Download the plugin from the [Releases] page into your BitBar plugin directory
    1. Clone the repository, run `make build` to compile, and then copy it into your BitBar plugin directory
    
The plugin updates every 10 minutes by default. This can be changed (e.g. to every 5 minutes) by renaming the binary 
in the BitBar plugin directory from `mvginfo.10m` to `mvginfo.5m`. 

## Developing

It's a standard project with Gomodules. There are lots of more or less helpful targets in the `Makefile`. 

[BitBar]: https://getbitbar.com
[Releases]: https://github.com/muffix/mvg-info/releases/latest
[homebrew]: https://brew.sh
