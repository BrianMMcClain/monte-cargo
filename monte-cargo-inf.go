package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
)

type Result struct {
  hits int
  total int
}

func generatePoints(batchSize int, c chan *Result) {
  r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
  for {
    hitCount := 0
    for i := 0; i < batchSize; i++ {
      x := r.Float64()
      y := r.Float64()

      inCircle := (math.Pow(x, 2) + math.Pow(y, 2)) < 1.0
      if inCircle {
        hitCount += 1
      }
    }

    res := Result{hits: hitCount, total: batchSize}
    c <- &res
  }
}

func estimateAndPrintPi(hits float64, total float64) float64 {
  est := (float64(hits) / float64(total)) * 4.0
  fmt.Printf("\r                                                                  \r%v (%d data points)", est, int(total))
  return est
}

func main() {

  batchSize := 10000000
  numCores := 4

  c1 := make(chan *Result)
  time.Sleep(100 * time.Millisecond)
  c2 := make(chan *Result)
  time.Sleep(100 * time.Millisecond)
  c3 := make(chan *Result)
  time.Sleep(100 * time.Millisecond)
  c4 := make(chan *Result)

  go generatePoints(batchSize, c1)
  go generatePoints(batchSize, c2)
  go generatePoints(batchSize, c3)
  go generatePoints(batchSize, c4)

  fmt.Printf("Calculating pi with %v processors\n", numCores)

  total := 0
  hits := 0

  for {
    totalBatch := batchSize * numCores
    for i := 0; i < totalBatch; i++ {
      select {
      case res := <- c1:
        hits += res.hits
        total += res.total
        estimateAndPrintPi(float64(hits), float64(total))
      case res := <- c2:
        hits += res.hits
        total += res.total
        estimateAndPrintPi(float64(hits), float64(total))
      case res := <- c3:
        hits += res.hits
        total += res.total
        estimateAndPrintPi(float64(hits), float64(total))
      case res := <- c4:
        hits += res.hits
        total += res.total
        estimateAndPrintPi(float64(hits), float64(total))
      }
    }
  }
}