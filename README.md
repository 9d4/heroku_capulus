# heroku_capulus

![Deploy](https://github.com/traperwaze/heroku_capulus/workflows/Deploy/badge.svg)

A cli tool to make request to heroku or any server in an interval.

# Usage

Download the binary in [here](https://github.com/9d4/heroku_capulus/releases) based on your platform.
The configuration file is needed to run the binary. Put the config file in the same directory with the binary.
Configuration file should named as `config.json`. Customize the config based on your need.
The configuration file should at least look like below.

```json
{
    "interval": "10m",
    "urls": [
        "https://google.com",
        "https://ask.com"
        "https://duckduckgo.org"
    ],
    "timezone": "Asia/Jakarta",
    "startAt": "06:00",
    "stopAt": "19:00"
}
```

- **interval** is gap between request. How long it should wait before next request.
e.g. `10m`, `1h10m`, `15m`, `1h10m15s`

- **urls** List of urls to be requested.

- **timezone** is the timezeone that used by `startAt` and `stopAt`

- **startAt & stopAt** The tool will only run between the `startAt` and `stopAt`. Use 24 hour format.

# Dev

Install dependencies:

```
$ go get
```

Run:

```
$ go run .
```

Build:

```
$ go build
```
