package main

import (
	pk3 "GoIntermedio/abstractfactory"
	pk "GoIntermedio/myclass"
	pk2 "GoIntermedio/myclass2"
	"fmt"
	"sync"
	"time"
)

func main() {

	// POO - Primera Parte - Usando Structs

	//Se instancia y se asigna valores
	e := pk.Employee{
		Id:       100,
		Name:     "Cristian Meoño",
		Vacation: true,
	}
	fmt.Printf("%v\n", e)

	//Se crea un puntero con new y luego se asignan valores
	e2 := new(pk.Employee)
	e2.Id = 200
	e2.Name = "Cristian Francisco"
	e2.Vacation = false
	fmt.Printf("%v\n", *e2)

	//Se usa el constructor
	e3 := pk.NewEmployee(1, "Cristian Meoño Canel", false)
	e3.SetId(85)
	fmt.Printf("%v\n", *e3)
	fmt.Print("\n\n")

	// POO - Segunda Parte - Usando Interfaces y Composicion

	ftEmployee := pk2.FullTimeEmployee{}
	ftEmployee.Name = "Cristian"
	pk2.GetMessage(ftEmployee)
	fmt.Printf("%v\n", ftEmployee)

	tEmployee := pk2.TemporaryEmployee{}
	pk2.GetMessage(tEmployee)
	fmt.Print("\n\n")

	//POO - Tercera Parte - Patron de Diseño Abstract Factory

	smsFactory, _ := pk3.GetNotificationFactory("SMS")
	emailFactory, _ := pk3.GetNotificationFactory("Email")

	pk3.SendNotification(smsFactory)
	pk3.SendNotification(emailFactory)

	pk3.GetMethod(smsFactory)
	pk3.GetMethod(emailFactory)
	fmt.Print("\n\n")

	// Funciones Anonimas
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(1 * time.Second)
		fmt.Println("End")
		c <- 1
	}()

	<-c
	fmt.Print("\n\n")

	//Funciones Variadicas y retornos con nombre

	fmt.Println(sum(1, 2, 3, 4, 5))
	printName("Cristian", "Francisco", "Meoño", "Canel")
	fmt.Println(getValues(2))

	// CONCURRENCIA

	c1 := make(chan int, 1)
	c1 <- 1

	fmt.Println(<-c1)
	fmt.Print("\n\n")

	//WaitGroup

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
	fmt.Print("\n\n")

	//Usando Buffered Channels
	c2 := make(chan int, 5)
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		c2 <- 1
		wg2.Add(1)
		go doSomething2(i, &wg2, c2)
	}

	wg2.Wait()
	fmt.Print("\n\n")

	//Salida y Entrada
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
	fmt.Print("\n\n")

	// Worker Pools

	tasks := []int{20, 30, 45, 10, 40}
	nWorkers := 5

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}

}

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printName(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func sumToTest(x, y int) int {
	return x + y
}

func getMaxToTest(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func fibonacciToTest(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciToTest(n-1) + fibonacciToTest(n-2)
}

func doSomething(i int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Started %d\n", i)
	time.Sleep(2)
	fmt.Println("Finished")
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished %d\n", i)
	<-c
}

func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func Double(cIn <-chan int, cOut chan<- int) {
	for v := range cIn {
		cOut <- 2 * v
	}
	close(cOut)
}

func Print(c <-chan int) {
	for value := range c {
		fmt.Printf("#%d\n", value)
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}
