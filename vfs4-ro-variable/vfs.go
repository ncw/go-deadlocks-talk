package main

import (
	"path"
	"sync"
)

// File represents a file on cloud storage
type File struct {
	// Read only variables - never changed in the lifetime of the object
	name   string
	parent *Dir

	// Read write variables protected by a mutex
	mu   sync.Mutex
	size int64
}

// Path returns the full path of the File object
func (f *File) Path() string {
	// No locking needed as we only access read only variables
	return path.Join(f.parent.Path(), f.name)
}

// Size of the file object
func (f *File) Size() int64 {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.size
}

// Truncate the file to size
func (f *File) Truncate(size int64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.size = size
}

// Dir represents a directory on cloud storage
type Dir struct {
	mu     sync.Mutex
	name   string
	parent *Dir
	files  []*File
	dirs   []*Dir
}

// Path returns the full path of the Dir
func (d *Dir) Path() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.parent == nil {
		return "/"
	}
	return path.Join(d.parent.Path(), d.name)
}

// Size of all the objects in the Dir
func (d *Dir) Size() int64 {
	d.mu.Lock()
	defer d.mu.Unlock()
	var total int64
	for _, f := range d.files {
		total += f.Size()
	}
	return total
}

func process1(f *File) {
	for {
		f.Path()
	}
}

func process2(d *Dir) {
	for {
		d.Size()
	}
}

func main() {
	root := &Dir{}
	file1 := &File{name: "file1", size: 42, parent: root}
	root.files = []*File{file1}

	go process1(file1)
	go process2(root)

	select {}
}
