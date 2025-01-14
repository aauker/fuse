package fstestutil

import (
	"context"
	"os"
	"syscall"

	"github.com/aauker/fuse"
	"github.com/aauker/fuse/fs"
)

// SimpleFS is a trivial FS that just implements the Root method.
type SimpleFS struct {
	Node fs.Node
}

var _ fs.FS = SimpleFS{}

func (f SimpleFS) Root() (fs.Node, error) {
	return f.Node, nil
}

// File can be embedded in a struct to make it look like a file.
type File struct{}

func (f File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = 0o666
	return nil
}

// Dir can be embedded in a struct to make it look like a directory.
type Dir struct{}

func (f Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0o777
	return nil
}

// ChildMap is a directory with child nodes looked up from a map.
type ChildMap map[string]fs.Node

var _ fs.Node = (*ChildMap)(nil)
var _ fs.NodeStringLookuper = (*ChildMap)(nil)

func (f *ChildMap) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0o777
	return nil
}

func (f *ChildMap) Lookup(ctx context.Context, name string) (fs.Node, error) {
	child, ok := (*f)[name]
	if !ok {
		return nil, syscall.ENOENT
	}
	return child, nil
}
