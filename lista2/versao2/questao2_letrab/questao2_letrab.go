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

  sum := 0
  check := 0
  for i := 1; i <= num_replicas ; i++ {
		go request(first)
	}
  go contador(first)

  for i := 1; i <= num_replicas ; i++ {
		check = <-first
    if check == -1{
      sum = -1
    }else{
      sum += check
    }
	}

  return sum

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
  time.Sleep(16 * time.Second)
  first <- -1
  close(first)
}
