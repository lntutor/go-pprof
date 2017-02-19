package main

import (
	"flag"
	"fmt"
	"sync"
  "math"
)

func main() {
	n := flag.Int("n", 10, "maximum number to consider")
	flag.Parse()
  sieve := calculatePrimeNarrowLock(*n)
  fmt.Println(sieve)

  // sieve = calculatePrimeWideLock(*n)
  // fmt.Println(sieve)

  // sieve = calculatePrimeSeq(*n)
  // fmt.Println(sieve)
}

func calculatePrimeWideLock(n int) []bool{
  sieve := make([]bool, n)
  for i := 0; i < n; i++ {
    sieve[i] = true
  }
  sieve[0] = false
  sieve[1] = false

  var mu sync.Mutex
  var wg sync.WaitGroup
  wg.Add(n/2)

  for i := 0; i < n/2; i++ {
    go func(i int) {
      if IsPrime(i) {
        count := 2
        // WIDE OMIT
        mu.Lock() // HL
        for i*count < n {
          sieve[i*count] = false
          count++
        }
        mu.Unlock() // HL
      }
      wg.Done()
    }(i)
  }
  wg.Wait()
  return sieve
}

func calculatePrimeNarrowLock(n int) []bool{
  sieve := make([]bool, n)
  for i := 0; i < n; i++ {
    sieve[i] = true
  }
  sieve[0] = false
  sieve[1] = false

  // var mu sync.Mutex
  var wg sync.WaitGroup
  wg.Add(n/2)

  for i := 0; i < n/2; i++ {
    go func(i int) {
      if IsPrime(i) {
        count := 2
        // NARROW OMIT
        for i*count < n {
          // mu.Lock() // HL
          sieve[i*count] = false
          // mu.Unlock() // HL
          count++
        }
      }
      wg.Done()
    }(i)
  }
  wg.Wait()
  return sieve
}

func calculatePrimeSeq(n int) []bool{
  sieve := make([]bool, n)
  for i := 0; i < n; i++ {
    sieve[i] = true
  }
  sieve[0] = false
  sieve[1] = false

  for i := 0; i < n/2; i++ {
    if IsPrime(i) {
      count := 2
      for i*count < n {
        sieve[i*count] = false
        count++
      }
    }
  }
  return sieve
}

func calculatePrimeRace(n int) []bool{
  sieve := make([]bool, n)
  for i := 0; i < n; i++ {
    sieve[i] = true
  }
  sieve[0] = false
  sieve[1] = false

  var wg sync.WaitGroup
  wg.Add(n/2)

  for i := 0; i < n/2; i++ {
    go func(i int) {
      if IsPrime(i) {
        count := 2
        // RACE OMIT
        for i*count < n {
          // mu.Lock() // HL
          sieve[i*count] = false
          // mu.Unlock() // HL
          count++
        }
      }
      wg.Done()
    }(i)
  }
  wg.Wait()
  return sieve
}

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}
