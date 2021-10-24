package vfs

import "path"

// File represents a file on cloud storage
type File struct {
	name   string
	size   int64
	parent *Dir
}

// Path returns the full path of the File object
func (f *File) Path() string {
	return path.Join(f.parent.Path(), f.name)
}

// Size of the file object
func (f *File) Size() int64 {
	return f.size
}

// Truncate the file to size
func (f *File) Truncate(size int64) {
	f.size = size
}

// Dir represents a directory on cloud storage
type Dir struct {
	name   string
	parent *Dir
	files  []*File
	dirs   []*Dir
}

// Path returns the full path of the Dir
func (d *Dir) Path() string {
	if d.parent == nil {
		return "/"
	}
	return path.Join(d.parent.Path(), d.name)
}

// Size of all the objects in the Dir
func (d *Dir) Size() int64 {
	var total int64
	for _, f := range d.files {
		total += f.Size()
	}
	return total
}
