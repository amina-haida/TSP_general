package main 

import ("TSP_general/donnees"
"time"
"fmt")

func main(){

	depart := "FR"
dicoVilles := donnees.Ville_aléatoire(15)

start := time.Now()
exhaustive.Exhaustive(depart, dicoVilles)
fmt.Println(time.Since(start))

}