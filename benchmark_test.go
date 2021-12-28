package lab2

import (
	"fmt"
	"math/rand"
	"testing"
)

var convertRes string

func BenchmarkPrefixToPostfix(b *testing.B) {
	const baseLen = 1
	for i := 0; i < 20; i++ {
		l := baseLen * 1 << i
		in := ""
		operators := [5]string{"+", "-", "*", "/", "^"}
		for k := 0; k < l; k++ {
			if rand.Intn(2) == 0 {
				in = operators[rand.Intn(5)] + " " + in + "0 "
			} else {
				in += operators[rand.Intn(5)] + " 0 "
			}
		}
		in += "0"
		b.Run(fmt.Sprintf("len=%d", len(in)), func(b *testing.B) {
			convertRes, _ = PrefixToPostfix(in)
		})
	}
}
