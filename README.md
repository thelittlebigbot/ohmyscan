# Oh My Scan

One Paragraph of project description goes here.

## Table of contents

- [Oh My Scan](#oh-my-scan)
  - [Table of contents](#table-of-contents)
  - [Overview](#overview)
    - [About the project](#about-the-project)
    - [Status](#status)
    - [Buit with](#buit-with)
    - [Version](#version)
  - [Getting started](#getting-started)
    - [Downloading](#downloading)
    - [Running](#running)
      - [Example](#example)
  - [Authors](#authors)
  - [License](#license)

## Overview

### About the project

Oh My Scan is a script developed in Go that allows you to download locally your favorite french scans. To carry out this project, I was inspired by the initial [go-scan-op.com-scrapper](https://github.com/SegFault42/go-scan-op.com-scrapper) project of [https://github.com/SegFault42](https://github.com/SegFault42).

### Status

**WIP**

### Buit with

- [Go](https://golang.org/) - The language used.

### Version

**0.1.0**

## Getting started

### Downloading

```sh
git clone https://github.com/thelittlebigbot/ohmyscan
```

### Running

```sh
ohmyscan download --platform PLATFORM_NAME --manga MANGA_NAME --number NUMBER
```

#### Example

```sh
ohmyscan download --platform scan-op --manga doubt --number 1
```

## Authors

**Alexandre Figueiredo** - _Initial work_ - [thelittlebigbot](https://github.com/thelittlebigbot).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
