package test

import (
	"compress/flate"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkPoolGC(b *testing.B) {
	b.ReportAllocs()
	var a, z [1000]*flate.Writer
	p := sync.Pool{New: func() interface{} { return &flate.Writer{} }}
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*flate.Writer)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}
		a = z
		runtime.GC()
	}
}
