package main
// go test -bench=. atoi_test.go.
import "fmt"
import "strconv"
import "testing"

var num = 123456
var numstr = "123456"

func BenchmarkStrconvParseInt(b *testing.B) {
  num64 := int64(num)
  for i := 0; i < b.N; i++ {
    x, err := strconv.ParseInt(numstr, 10, 64)
    if x != num64 || err != nil {
      b.Error(err)
    }
  }
}

func BenchmarkAtoi(b *testing.B) {
  for i := 0; i < b.N; i++ {
    x, err := strconv.Atoi(numstr)
    if x != num || err != nil {
      b.Error(err)
    }
  }
}

func BenchmarkFmtSscan(b *testing.B) {
  for i := 0; i < b.N; i++ {
    var x int
    n, err := fmt.Sscanf(numstr, "%d", &x)
    if n != 1 || x != num || err != nil {
      b.Error(err)
    }
  }
}