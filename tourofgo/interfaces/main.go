package main

import (
	_ "fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(buf []byte) (n int, err error) {
	n, err = rot.r.Read(buf)
	for i, b := range buf {
		// fmt.Println(buf[i])
		buf[i] = rot13(b)
		
	}
	return
}

func rot13(b byte) (c byte) {
	switch {
	case b >= 'A' && b <= 'Z':
		c = b +13
		// c = (b-'A'+13)%26 + 'A'
	case b >= 'a' && b <= 'z':
		// c = (b-'a'+13)%26 + 'a'
		c = b -13
	default:
		c = b
	}
	return
}

func main() {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	// r := strings.NewReader("Hello, Reader! sgs gdfgsdfgs 11111111111111")

	// b := make([]byte, 8)
	// for {
	// 	n, err := r.Read(b) // считывает информацию в b, ошибку в эрор, а количество прочитанных данных в n
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
}
