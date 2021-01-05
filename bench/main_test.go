
package main

import (
    "testing"
    "os"
    "fmt"
)


func BenchmarkFnv1aGolib(b *testing.B) {
	h:=GetHash_golib_system("xwforensics64") ^ uint64(6605813339339102567)
	if h != uint64(17439059603042731363) {
		fmt.Println(h)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		GetHash_golib_system("xwforensics64")
	}

}

func BenchmarkFnv1aOwn(b *testing.B) {
	h:=GetHash("xwforensics64") ^ uint64(6605813339339102567)
	if h != uint64(17439059603042731363) {
		fmt.Println(h)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		GetHash("xwforensics64")
	}

}
