wireless
========

[![Build Status](https://travis-ci.org/solarnz/wireless.svg)](https://travis-ci.org/solarnz/wireless)
[![GoDoc](https://godoc.org/github.com/solarnz/wireless?status.svg)](http://godoc.org/github.com/solarnz/wireless)

A golang library to interact with the wireless library on linux.

Dependencies
------------

- You must be building your golang application with cgo enabled (so no cross-compilation)
- libiw to be installed (libiw-dev on Ubuntu, wireless-tools on Arch)
