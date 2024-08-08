package main

import(
	"fmt"
)

func isPrime(x int)bool {
	if x <=1 {
		return false
	}

	for i:=2 ; i<x ; i++{
		if x%i == 0{
			return false
		}
	}
	return true
}

func main(){
	var n int
	fmt.Scan(&n)
	if isPrime(n){
		fmt.Printf("yes %d is prime\n" , n)
	} else { 
		fmt.Printf("no")
	}
}