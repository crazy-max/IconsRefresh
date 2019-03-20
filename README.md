<p align="center"><img width="100" src="https://raw.githubusercontent.com/crazy-max/IconsRefresh/master/.res/logo.png"></p>

<p align="center">
  <a href="https://github.com/crazy-max/IconsRefresh/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/IconsRefresh.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/IconsRefresh/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/IconsRefresh/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://travis-ci.com/crazy-max/IconsRefresh"><img src="https://img.shields.io/travis/com/crazy-max/IconsRefresh/master.svg?style=flat-square" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/IconsRefresh"><img src="https://goreportcard.com/badge/github.com/crazy-max/IconsRefresh?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/IconsRefresh"><img src="https://img.shields.io/codacy/grade/834f62e0849c4c008dd8df69b816d2a0/master.svg?style=flat-square" alt="Code Quality"></a>
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

## License

MIT. See `LICENSE` for more details.<br />
Icon credit to Oliver Scholtz
