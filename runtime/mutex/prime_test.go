package main

import (
	"fmt"
	"testing"
)


func TestRacer(t *testing.T) {
  calculatePrimeRace(10000)
}

func benchFuncPrime(b *testing.B, f func(int) []bool) {

	for n := 10; n <= 10000; n *= 10 {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(n)
			}
		})
	}
}


func benchFuncFactor(b *testing.B, f func(int) map[int]int) {

	for n := 10; n <= 10000; n *= 10 {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(n)
			}
		})
	}
}

func BenchmarkPrimeNarrowLock(b *testing.B) { benchFuncPrime(b, calculatePrimeNarrowLock) }
func BenchmarkPrimeWideLock(b *testing.B)   { benchFuncPrime(b, calculatePrimeWideLock) }
// func BenchmarkPrimeSq(b *testing.B)         { benchFuncPrime(b, calculatePrimeSeq) }


// func BenchmarkFactorNarrowSection(b *testing.B) { benchFuncFactor(b, countFactorsNarrowSection) }
// func BenchmarkFactorWideSection(b *testing.B)   { benchFuncFactor(b, countFactorsWideSection) }
// func BenchmarkFactorSq(b *testing.B)            { benchFuncFactor(b, countFactorsSeq) }
