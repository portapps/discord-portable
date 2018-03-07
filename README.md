<p align="center"><a href="https://github.com/portapps/discord-portable" target="_blank"><img width="100" src="https://github.com/portapps/discord-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/discord-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/discord-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/discord-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/discord-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/portapps/discord-portable"><img src="https://img.shields.io/appveyor/ci/portapps/discord-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/discord-portable"><img src="https://goreportcard.com/badge/github.com/portapps/discord-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/portapps/discord-portable"><img src="https://img.shields.io/codacy/grade/46d1e15b6c984642a2f2e7932f9c119b.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://beerpay.io/portapps/portapps"><img src="https://img.shields.io/beerpay/portapps/portapps.svg?style=flat-square" alt="Beerpay"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[Discord](https://discordapp.com) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

> 💡 PTB (Public Test Build) portable version of Discord is available [here](https://github.com/portapps/discord-ptb-portable)

## Installation

There are different kinds of artifacts :

* `discord-portable-win{32,64}-x.x.x-x-setup.exe` : Full portable release of Discord as a setup. **Recommended way**!
* `discord-portable-win{32,64}-x.x.x-x.7z` : Full portable release of Discord as a 7z archive.
* `discord-portable-win{32,64}.exe` : Only the portable binary (must be renamed `discord-portable.exe`)
* `DiscordSetup-win{32,64}-x.x.x.exe` : The original setup from the [official website](https://discordapp.com/download).

### Fresh installation

Install `discord-portable-win{32,64}-x.x.x-x-setup.exe` where you want then run `discord-portable.exe`.

### App already installed

If you have already installed Discord from the original setup, do the same thing as a fresh installation and :

* Move data located in `%APPDATA%\discord\*` to `data` folder.

Run `discord-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Discord from your computer.

### Upgrade

For an upgrade, simply download and install the [latest setup](https://github.com/portapps/discord-portable/releases/latest).

## How can i help ?

All kinds of contributions are welcomed :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Beerpay](https://beerpay.io/portapps/portapps/badge.svg?style=beer-square)](https://beerpay.io/portapps/portapps)
or [![Paypal](https://cdn.rawgit.com/portapps/portapps/master/res/paypal.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
