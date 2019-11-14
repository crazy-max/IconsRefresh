<p align="center"><img width="100" src="https://raw.githubusercontent.com/crazy-max/IconsRefresh/master/.res/logo.png"></p>

<p align="center">
  <a href="https://github.com/crazy-max/IconsRefresh/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/IconsRefresh.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/IconsRefresh/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/IconsRefresh/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://github.com/crazy-max/IconsRefresh/actions"><img src="https://github.com/crazy-max/IconsRefresh/workflows/build/badge.svg" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/IconsRefresh"><img src="https://goreportcard.com/badge/github.com/crazy-max/IconsRefresh?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/IconsRefresh"><img src="https://img.shields.io/codacy/grade/834f62e0849c4c008dd8df69b816d2a0/master.svg?style=flat-square" alt="Code Quality"></a>
  <br /><a href="https://github.com/sponsors/crazy-max"><img src="https://img.shields.io/badge/sponsor-crazy--max-181717.svg?logo=github&style=flat-square" alt="Become a sponsor"></a>
  <a href="https://www.paypal.me/crazyws"><img src="https://img.shields.io/badge/donate-paypal-00457c.svg?logo=paypal&style=flat-square" alt="Donate Paypal"></a>
</p>

## About

Icons Refresh :paintbrush: is a program written in [Go](https://golang.org/) to refresh Desktop, Start Menu and Taskbar icons without restart Explorer on Windows.

## Download

You can download the application on the [**releases page**](https://github.com/crazy-max/IconsRefresh/releases/latest).

## Usage

Launch `IconsRefresh.exe`

## Build

> Go 1.12 or higher required

```
go mod download
go generate -v
go build -o bin/IconsRefresh.exe -v -ldflags "-v -H=windowsgui"
```

## How can I help ?

All kinds of contributions are welcome :raised_hands:! The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon: You can also support this project by [**becoming a sponsor on GitHub**](https://github.com/sponsors/crazy-max) :clap: or by making a [Paypal donation](https://www.paypal.me/crazyws) to ensure this journey continues indefinitely! :rocket:

Thanks again for your support, it is much appreciated! :pray:

## License

MIT. See `LICENSE` for more details.<br />
Icon credit to Oliver Scholtz
