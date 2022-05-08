package main

import (
    "testing"
    "strconv"
)
//go test -bench=. atoi_check.go

func Atoi (s string) int {
    var (
        n uint64
        i int
        v byte
    )   
    for ; i < len(s); i++ {
        d := s[i]
        if '0' <= d && d <= '9' {
            v = d - '0'
        } else if 'a' <= d && d <= 'z' {
            v = d - 'a' + 10
        } else if 'A' <= d && d <= 'Z' {
            v = d - 'A' + 10
        } else {
            n = 0; break        
        }
        n *= uint64(10) 
        n += uint64(v)
    }
    return int(n)
}

func BenchmarkAtoi(b *testing.B) {
    for i := 0; i < b.N; i++ {
        in := Atoi("9999")
        _ = in
    }   
}

func BenchmarkStrconvAtoi(b *testing.B) {
    for i := 0; i < b.N; i++ {
        in, _ := strconv.Atoi("9999")
        _ = in
    }   
}