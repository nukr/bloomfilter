package main

import (
	"fmt"

	"github.com/nukr/bloomfilter"
)

func main() {
	bf := bloomfilter.New(128, 3)
	bf.Add([]byte("aakajdlfja;ljd;lj;lakjd;lfkja;lsdjkf;aaa"))
	fmt.Println(bf.MayContain([]byte("aakajdlfja;ljd;lj;lakjd;lfkja;lsdjkf;aaa")))
	fmt.Println(bf.MayContain([]byte("will false")))
}
