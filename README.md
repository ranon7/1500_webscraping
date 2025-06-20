# 1500_webscraping

Collection of Go command line tools for web scraping [1500chan](https://1500chan.org).

## Installation

- Install [Go](https://go.dev/doc/install) in your OS
- Run `go install github.com/ranon7/1500_webscraping@latest`

## Media scraping

Downloads all media files from a thread and saves on a local folder.

### Sample usage

Downloads all jpg and png files from a /b/ thread to the current directory.

```shell
1500_webscraping media_scrap --board=b --thread=14067101 --formats=jpg,png --location=. --verbose --m=30
```

### Command line arguments

| argument | required | default | description                                                                                            | example                                                           |
|----------|----------|---------|--------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------|
| board    | yes      |         | Valid board name from 1500chan.org                                                                     | b                                                                 |
| thread   | yes      |         | Thread number must exist in the selected board                                                         | 37601                                                             |
| formats  | yes      |         | A list of file formats as a comma separated string                                                     | jpg,mp4,webm                                                      |
| location | yes      |         | The location where the media will be saved into your file system. If it doesn't exist it'll be created | /home/anon/archive or just "." (expands to the current directory) |
| verbose  | no       | false   | Enable detailed logs. Disabled by default                                                              | true                                                              |
| m        | no       | 30      | Controls the parallism - how many files being downloaded at a time                                     | 1000                                                              |

### Output folder structure

Files will be downloaded under the `location`/`board`/`thread`/ folder. Those values are supplied as command line arguments. The path structure will be created if it does not exist.

## Setup for local development

- clone
- install go
- run `go get .`

## Contact

<ranon7@protonmail.com>

## Contribute

- fork
- open pr
