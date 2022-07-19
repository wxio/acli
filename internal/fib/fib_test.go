package fib

import (
	"testing"
)

func iterativeFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibIterative(i)
	}
}

func recursiveFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibRecursive(i)
	}
}

func channelFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibChannel(i)
	}
}

func BenchmarkIterativeFib3(b *testing.B)  { iterativeFib(3, b) }
func BenchmarkIterativeFib10(b *testing.B) { iterativeFib(10, b) }
func BenchmarkIterativeFib20(b *testing.B) { iterativeFib(20, b) }
func BenchmarkIterativeFib40(b *testing.B) { iterativeFib(40, b) }
func BenchmarkIterativeFib45(b *testing.B) { iterativeFib(45, b) }
func BenchmarkIterativeFib50(b *testing.B) { iterativeFib(50, b) }

func BenchmarkChannelFib4(b *testing.B)  { channelFib(4, b) }
func BenchmarkChannelFib10(b *testing.B) { channelFib(10, b) }
func BenchmarkChannelFib20(b *testing.B) { channelFib(20, b) }
func BenchmarkChannelFib40(b *testing.B) { channelFib(40, b) }
func BenchmarkChannelFib45(b *testing.B) { channelFib(45, b) }
func BenchmarkChannelFib50(b *testing.B) { channelFib(50, b) }
func BenchmarkChannelFib60(b *testing.B) { channelFib(60, b) }
func BenchmarkChannelFib80(b *testing.B) { channelFib(80, b) }

func BenchmarkRecursiveFib1(b *testing.B)  { recursiveFib(1, b) }
func BenchmarkRecursiveFib2(b *testing.B)  { recursiveFib(2, b) }
func BenchmarkRecursiveFib3(b *testing.B)  { recursiveFib(3, b) }
func BenchmarkRecursiveFib10(b *testing.B) { recursiveFib(10, b) }
func BenchmarkRecursiveFib20(b *testing.B) { recursiveFib(20, b) }
func BenchmarkRecursiveFib40(b *testing.B) { recursiveFib(40, b) }
func BenchmarkRecursiveFib45(b *testing.B) { recursiveFib(45, b) }
func BenchmarkRecursiveFib50(b *testing.B) { recursiveFib(50, b) }
