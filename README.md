# Oh My Scan

One Paragraph of project description goes here.

```
   ___  _      __  __        ___
  / _ \| |_   |  \/  |_  _  / __| __ __ _ _ _
 | (_) | ' \  | |\/| | || | \__ \/ _/ _` | ' \
  \___/|_||_| |_|  |_|\_, | |___/\__\__,_|_||_|
                      |__/
```

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
    - [How to use](#how-to-use)
      - [Download a scan](#download-a-scan)
      - [Merge a scan](#merge-a-scan)
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

### How to use

#### Download a scan

To download a scan in multiple images, please run this command:

```sh
ohmyscan download --platform PLATFORM_NAME --manga MANGA_NAME --number NUMBER
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan download --platform scan-op --manga doubt --number 1
```

#### Merge a scan

After that, if you want to merge all images of a scan in one PDF file, please do:

```sh
ohmyscan merge --name MANGA_NAME --number NUMBER
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan merge --name doubt --number 1
```

## Authors

**Alexandre Figueiredo** - _Initial work_ - [thelittlebigbot](https://github.com/thelittlebigbot).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
