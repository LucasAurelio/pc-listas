package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  var vezes int

  fmt.Print("Digite o número de réplicas desejadas: ")
  fmt.Scan(&vezes)
  gateway(vezes)
}

func gateway(num_replicas int) int{
  first := make(chan int)

  for i := 1; i <= num_replicas ; i++ {
		go request(first)
	}

  go contador(first)

  return <-first

}

func request(first chan int){
  rand.NewSource(time.Now().UTC().UnixNano())
  min := 1
  max := 29
  chosen := rand.Intn(max) + min

  time.Sleep(time.Duration(chosen) * time.Second)
  first <- chosen
}

func contador(first chan int){
  time.Sleep(8 * time.Second)
  first <- -1
}
