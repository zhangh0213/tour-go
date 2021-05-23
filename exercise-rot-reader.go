package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot(b byte) byte{
	if b >= 'A' && b <= 'M' || b >= 'a' && b <= 'm'{
		return (b + 13)
	}
	return (b - 13)
}

func (rot13reader rot13Reader) Read(s []byte) (n int,err error){
	n, err = rot13reader.r.Read(s)
	for i := range s{
		s[i] = rot(s[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
