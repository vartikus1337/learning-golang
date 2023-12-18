package main

import (
   "fmt"
   "sync"
)

const N = 10

func main() {
   m := make(map[int]int)

   wg := &sync.WaitGroup{}
   mu := &sync.Mutex{}
   wg.Add(N)
   for i := 0; i < N; i++ {
      go func() {
         defer wg.Done()
         mu.Lock()
         m[i] = i
         mu.Unlock()
      }()
   }
   wg.Wait()

   i := 0
   for a := range m {
      fmt.Println("m", i, " - ", a)
      i++
   }
}

func main2() {
	m := make(map[int]int)

   wg := &sync.WaitGroup{}
   mu := &sync.Mutex{}
   wg.Add(N)
   for i := 0; i < N; i++ {
      go func(i int) {
         defer wg.Done()
         mu.Lock()
         m[i] = i
         mu.Unlock()
      }(i)
   }
   wg.Wait()

   i := 0
   for a := range m {
      fmt.Println("m", i, " - ", a)
      i++
   }
}