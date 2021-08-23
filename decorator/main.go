package main

import "fmt"

func main() {
	pizza := &veggeMania{}

	pizzaWithTomato := &tomatoTopping{pizza: pizza}

	pizzaWithTomatoAndCheese := &cheeseTopping{pizza: pizzaWithTomato}

	fmt.Printf("Price: %v \n", pizzaWithTomatoAndCheese.getPrice())
}
