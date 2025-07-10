package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r   io.Reader
	key byte
}

func (self rot13Reader) Read(v []byte) (int, error) {
	data := make([]byte, 128)
	n, _ := self.r.Read(data)
	data = data[:n]
	for i := range data {
		if unicode.IsUpper(rune(data[i])) {
			v[i] = (data[i]-byte('A')+self.key)%(26) + byte('A')
		} else {
			v[i] = (data[i]-byte('a')+self.key)%(26) + byte('a')
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s, 1}
	io.Copy(os.Stdout, &r)
}
