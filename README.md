# Oh My Scan

```
   ___  _      __  __        ___
  / _ \| |_   |  \/  |_  _  / __| __ __ _ _ _
 | (_) | ' \  | |\/| | || | \__ \/ _/ _` | ' \
  \___/|_||_| |_|  |_|\_, | |___/\__\__,_|_||_|
                      |__/
```

[![Maintenance](https://img.shields.io/badge/maintained-yes-green.svg)](https://github.com/afigueir/ohmyscan/pulse)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/afigueir/ohmyscan/blob/master/LICENSE)
[![Go report card](https://goreportcard.com/badge/github.com/afigueir/ohmyscan)](https://github.com/afigueir/ohmyscan)

## Table of contents

- [Oh My Scan](#oh-my-scan)
  - [Table of contents](#table-of-contents)
  - [Overview](#overview)
    - [About the project](#about-the-project)
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

Oh My Scan is a script developed in Go that allows you to download locally your favorite french manga scans. To carry out this project, I was inspired by the initial [go-scan-op.com-scrapper](https://github.com/SegFault42/go-scan-op.com-scrapper) project of [https://github.com/SegFault42](https://github.com/SegFault42).

### Buit with

- [Go](https://golang.org/) - The language used.

### Version

**1.0.0**

## Getting started

**WARNING: If you are on Windows, your antivirus may be panicking. But don't worry, just ignore it. If you have any question about this script, don't hesitate to contact me by email or on Twitter.**

### Downloading

To download this application, please go to [releases](https://github.com/afigueir/ohmyscan/releases/) section and choose the `ohmyscan` file if you're in a Unix system, or `ohmyscan.exe` if your are on a Windows system.

### How to use

_If you're are on a windows system, add `.exe` after `ohmyscan`._

#### Download a scan

To download a scan in multiple images, please run this command:

```sh
ohmyscan download --name MANGA_NAME --number NUMBER
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan download --name doubt --number 1
```

If you want to merge directly the scans you just downloaded, please run :

```sh
ohmyscan download --name MANGA_NAME --number NUMBER --merge BOOLEAN
```

**Example**:

```sh
# https://scan-op.cc/manga/doubt/1
ohmyscan download --name doubt --number 1 --merge true
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

**Alexandre Figueiredo** - _Initial work_ - [afigueir](https://github.com/afigueir).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
