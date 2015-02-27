package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
)

func generatePoints(count int, c chan int) {
  r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
  hitCount := 0
  for i := 0; i <= count; i++ {
    x := r.Float64()
    y := r.Float64()

    inCircle := (math.Pow(x, 2) + math.Pow(y, 2)) < 1.0
    if inCircle {
      hitCount += 1
    }
  }

  c <- hitCount
}

func main() {

  pointsPerCore := 1000000000
  numCores := 4

  c1 := make(chan int)
  c2 := make(chan int)
  c3 := make(chan int)
  c4 := make(chan int)

  go generatePoints(pointsPerCore, c1)
  go generatePoints(pointsPerCore, c2)
  go generatePoints(pointsPerCore, c3)
  go generatePoints(pointsPerCore, c4)

  fmt.Printf("Calculating pi with %v data points\n", pointsPerCore * numCores)

  total := 0
  expectedTotal := pointsPerCore * numCores
  hits := 0

  rand.Seed(time.Now().UTC().UnixNano())

  for total < expectedTotal {
    select {
    case res := <- c1:
      hits += res
      total += pointsPerCore
    case res := <- c2:
      hits += res
      total += pointsPerCore
    case res := <- c3:
      hits += res
      total += pointsPerCore
    case res := <- c4:
      hits += res
      total += pointsPerCore
    }
  }

  est := (float64(hits) / float64(total)) * 4.0

  fmt.Printf("%v\n", est)
}