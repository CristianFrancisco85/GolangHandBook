package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Funcion con costo computacional alto
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Memory hara cache apartir de la resultados de la funcion y los guardara en un map
type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.RWMutex
}

// La funcion para hacer cache devuelve un tipo generico y un error
type Function func(key int) (interface{}, error)

// El resultado que se guardara ser aun structo del valor o su posible error
type FunctionResult struct {
	value interface{}
	err   error
}

//Contructor
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// Metodo para obtener un valor de la cache, si no existe lo calcula y lo guarda
func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]

	if !exists {

		result.value, result.err = m.f(key)
		m.cache[key] = result

	}
	m.lock.Unlock()

	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 42, 41, 42, 38, 45, 45}
	var wg sync.WaitGroup

	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("N:%d, T:%s, R:%d\n", index, time.Since(start), value)
		}(n)
	}
	wg.Wait()
}
