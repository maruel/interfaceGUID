interfaceGUID
=============

interfaceGUID exposes a single function `CalculateGUID` which calculates a
unique deterministic value based on the type provided. It is to be used when
communicating with remote services to quickly assert common knowledge before
starting to communicate, for example via `encoding/gob`.

[![GoDoc](https://godoc.org/github.com/maruel/interfaceGUID?status.svg)](https://godoc.org/github.com/maruel/interfaceGUID)
[![Build Status](https://travis-ci.org/maruel/interfaceGUID.svg?branch=master)](https://travis-ci.org/maruel/interfaceGUID)
[![Coverage Status](https://img.shields.io/coveralls/maruel/interfaceGUID.svg)](https://coveralls.io/r/maruel/interfaceGUID?branch=master)


Example
-------

    type MyService interface {
      ...
    }

    guid := interfaceGUID.CalculateGUID(reflect.TypeOf((*MyService)(nil)).Elem())
    // Then compare the string with the remote service.
