## go-video

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-video)
[![Build Status](https://travis-ci.org/ferhatelmas/go-video?branch=master)](https://travis-ci.org/ferhatelmas/go-video)

> Check if a filepath is a video file.

### Install

```
go get github.com/ferhatelmas/go-video
```

### Usage

```go
import "github.com/ferhatelmas/go-video"

video.Is("src/unicorn.mp4")
//=> true

video.Is("src/unicorn.txt")
//=> false
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
