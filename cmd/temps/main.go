package main

import (
	"TSP_general/temps"

	"math/rand"
	"time"
)

func main(){
 rand.Seed(time.Now().UnixNano())
    temps.Garantie(15)

}
