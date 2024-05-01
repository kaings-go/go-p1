package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name string
	rightFork int
	leftFork int
}

var philosophers = [] Philosopher{
	{ name: "Socrates", leftFork: 0, rightFork: 1 },
	{ name: "Aristotle", leftFork: 1, rightFork: 2 },
	{ name: "Pascal", leftFork: 2, rightFork: 3 },
	{ name: "Locke", leftFork: 3, rightFork: 4 },
	{ name: "Plato", leftFork: 4, rightFork: 0 },
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("===========================")
	fmt.Println("The table is empty")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	seated.Done()
	// wait until all philosophers are seated - next line only will be executed once all philosophers are seated
	seated.Wait()

	for i := hunger; i > 0; i-- {
		if (philosopher.leftFork < philosopher.rightFork) {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t\t%s takes the left fork [%v].\n", philosopher.name, philosopher.leftFork)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t\t%s takes the right fork [%v].\n", philosopher.name, philosopher.rightFork)
		} else {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t\t%s takes the right fork [%v].\n", philosopher.name, philosopher.rightFork)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t\t%s takes the left fork [%v].\n", philosopher.name, philosopher.leftFork)
		}

		// forks[philosopher.leftFork].Lock()
		// fmt.Printf("\t\t%s takes the left fork [%v].\n", philosopher.name, philosopher.leftFork)
		// forks[philosopher.rightFork].Lock()
		// fmt.Printf("\t\t%s takes the right fork [%v].\n", philosopher.name, philosopher.rightFork)

		fmt.Printf("\t%s has both forks [%v]-[%v] & is eating.\n", philosopher.name, philosopher.leftFork, philosopher.rightFork)
		time.Sleep(eatTime)
		
		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		fmt.Printf("\t\t%s puts down the left fork [%v].\n", philosopher.name, philosopher.leftFork)
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("\t\t%s puts down the right fork [%v].\n", philosopher.name, philosopher.rightFork)

		fmt.Printf("\t%s puts down both forks [%v]-[%v].\n", philosopher.name, philosopher.leftFork, philosopher.rightFork)
	}

	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")
}

