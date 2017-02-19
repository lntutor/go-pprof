package main

import (
    "bytes"
    "testing"
)

func BenchmarkConcat(b *testing.B) {
    var str string
    //Use b.ResetTimer() if we'd like to exclude some init code from our timing
    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        str += "x"
    }
}

func BenchmarkBuffer(b *testing.B) {
    var buffer bytes.Buffer
    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        buffer.WriteString("x")
    }
}
