github.com/aauker/fuse -- Filesystems in Go
===================================

`github.com/aauker/fuse` is a Fork of the Go library for writing FUSE userspace
filesystems.

The small change this package brings is passing the UID to the Lookup and ReadDirAll functions to enable access control to tree-level directories.

It is a from-scratch implementation of the kernel-userspace
communication protocol, and does not use the C library from the
project called FUSE. `github.com/aauker/fuse` embraces Go fully for safety and
ease of programming.

Here’s how to get going:

    go get github.com/aauker/fuse

Website: http://github.com/aauker/fuse/

Github repository: https://github.com/bazil/fuse

API docs: http://godoc.org/github.com/aauker/fuse

Our thanks to Russ Cox for his fuse library, which this project is
based on.
