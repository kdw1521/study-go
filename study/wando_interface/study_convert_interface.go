package wandointerface

type Reader interface {
	Read()
}
type Closer interface {
	Close()
}

type File struct{}

func (f *File) Read() {}

// func (f *File) Close() {}

func ReadFile(reader Reader) {
	c, ok := reader.(Closer)
	if ok {
		c.Close()
	}
}
