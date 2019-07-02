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

func gateway(num_replicas int){
  first := make(chan int)

  for i := 1; i <= num_replicas ; i++ {
		go request(first)
	}

  fmt.Println("O primeiro request a acabar foi o de valor: ", <-first)

}

func request(first chan int){
  rand.NewSource(time.Now().UTC().UnixNano())
  min := 1
  max := 29
  chosen := rand.Intn(max) + min
  fmt.Println("O número da vez foi: ", chosen)

  time.Sleep(time.Duration(chosen) * time.Second)
  first <- chosen
}
