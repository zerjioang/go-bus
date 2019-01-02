package mutex

import (
	"testing"
	"time"
)

func BenchmarkMutexbusConcurrent(b *testing.B) {
	b.Run("str-to-uint32-concurrent", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			go func() {
				if strTouint32("HelloWorld") != 926844193 {
					b.Error("hash function failed")
				}
			}()
		}
		time.Sleep(10 * time.Second)
	})
}
