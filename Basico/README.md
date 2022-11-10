# Go Basico

- Compilado
- Creado por Google
- Pensado para la nube y la concurrencia
- Multiparadigma

# Indice

- [Basics](#Basics)
  - [Hello World](#Hello-World)
  - [Variables y Constantes](#Variables-y-Constantes)
  - [Operaciones Aritméticas](#Operaciones-Aritméticas)
- [Librería fmt](#Librería-fmt)
- [Funciones](#Funciones)
- [Control de Flujo](#Control-de-Flujo)
  - [Ciclos](#Ciclos)
  - [Condicionales](#Condicionales)
  - [Switch](#Switch)
  - [Defer, Break y Continue](#Defer-Break-y-Continue)
- [Arrays y Slices](#Arrays-y-Slices)
  - [Basics](#Basics)
  - [Recorridos](#Recorridos)
- [Maps](#Maps)
- [Structs](#Structs)
- [Modificador de Acceso](#Modificador-de-Acceso)
- [Punteros](#Punteros)
- [Goroutine](#Goroutine)
- [Channels](#Channels)
- [Multiplexeando Channels](#Multiplexeando-Channels)


# Basics

## Hello World

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")
}
```

## Variables y Constantes

```go
import "fmt"

func main() {
	//DECLARACION DE CONSTANTES
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//DECLARACION DE VARIABLES ENTERAS
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64 
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
```

## Operaciones Aritméticas

```go
import "fmt"

func main() {
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = y - x
	fmt.Println("Resta", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)
}	

```

# Librería fmt

```go
import "fmt"

func main() {
	//Declaracion de Variables
	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 1000
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Print(message)

	//Tipo dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
```

# Funciones

```go
package main

import "fmt"

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value 1 y Value 2:", value1, value2)

	value3, _ := doubleReturn(4)
	fmt.Println("Value 3:", value3)
}

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
```

# Control de Flujo

## Ciclos

```go
package main

import "fmt"

func main() {
	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	
	fmt.Println("")

	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
```

## Condicionales

```go
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 2
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	//OR
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	//Convertir texto a numero

	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
```

## Switch

```go
package main

import (
	"fmt"
)

func main() {

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Sin condicion
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor de 0")
	default:
		fmt.Println("No condicion")
	}

}
```

## Defer, Break y Continue

```go
package main

import (
	"fmt"
)

func main() {

	//Defer ejecuta la funcion hasta el final de la funcion
	defer fmt.Println("defer")
	fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y Break
	for i := 0; i < 10; i++ {
		//Continue
		if i == 2 {
			fmt.Println("Es dos")
			//continue
			break
		}
		fmt.Println(i)
	}

}
```

# Arrays y Slices

## Basics

```go
package main

import (
	"fmt"
)

func main() {

	//Array

	var array [4]int

	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	//Append slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

}
```

## Recorridos

```go
package main

import (
	"fmt"
	"strings"
)

func main() {

	slice := []string{"aMa", "hola", "amor a roma"}

	for _, valor := range slice {
		isPalindromo(valor)
	}
}

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}
```

# Maps

```go
package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
```

# Structs

```go
package main

import (
	"fmt"
)

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Toyota", year: 2020}
	fmt.Println(myCar)

	//Otra forma
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}
```

# Modificador de Acceso

```go
package mypackage

import "fmt"

// CarPublic Car de acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// CarPublic Car de acceso privado, lleva minuscula
type CarPrivate struct {
	brand string
	year  int
}

func PrintMesssage() {
	fmt.Println("Hola")
}
```

# Punteros

```go
package myclass

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(rm, dsk int, brn string) Pc {
	myPc := Pc{ram: rm, disk: dsk, brand: brn}
	return myPc
}

func (myPc Pc) FormatPrint() {
	fmt.Printf("Esta pc marca %s cuenta con una ram de %dGB y un disco de %dGB.\n", myPc.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm int) {
	myPc.ram = rm
}
```

```go
package main

import (
	pc "CursoGoPlatzi/myclass"
	"fmt"
)

func main() {
	myPc := pc.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}
```

# Goroutine

```go
import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)

	wg.Wait()

	//Funcion anonima
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second)
}

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}
```

# Channels

```go
import (
	"fmt"
)

func main() {

	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bye", c)

	fmt.Println(<-c)

}

// chan<- canal de entrada
// <-chan canal de salida
func say(text string, c chan<- string) {
	c <- text
}
```

# Multiplexeando Channels

```go
func main() {

	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	//Range y close
	close(c)
	//c <- "Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("email1", email1)
	go message("email2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}

func message(text string, c chan string) {
	c <- text
}
```


# Basics

## Hello World

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")
}
```

## Variables y Constantes

```go
import "fmt"

func main() {
	//DECLARACION DE CONSTANTES
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//DECLARACION DE VARIABLES ENTERAS
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64 
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
```

## Operaciones Aritméticas

```go
import "fmt"

func main() {
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = y - x
	fmt.Println("Resta", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)
}	

```

# Librería fmt

```go
import "fmt"

func main() {
	//Declaracion de Variables
	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 1000
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Print(message)

	//Tipo dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
```

# Funciones

```go
package main

import "fmt"

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value 1 y Value 2:", value1, value2)

	value3, _ := doubleReturn(4)
	fmt.Println("Value 3:", value3)
}

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
```

# Control de Flujo

## Ciclos

```go
package main

import "fmt"

func main() {
	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	
	fmt.Println("")

	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
```

## Condicionales

```go
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 2
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	//OR
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	//Convertir texto a numero

	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
```

## Switch

```go
package main

import (
	"fmt"
)

func main() {

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Sin condicion
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor de 0")
	default:
		fmt.Println("No condicion")
	}

}
```

## Defer, Break y Continue

```go
package main

import (
	"fmt"
)

func main() {

	//Defer ejecuta la funcion hasta el final de la funcion
	defer fmt.Println("defer")
	fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y Break
	for i := 0; i < 10; i++ {
		//Continue
		if i == 2 {
			fmt.Println("Es dos")
			//continue
			break
		}
		fmt.Println(i)
	}

}
```

# Arrays y Slices

## Basics

```go
package main

import (
	"fmt"
)

func main() {

	//Array

	var array [4]int

	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	//Append slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

}
```

## Recorridos

```go
package main

import (
	"fmt"
	"strings"
)

func main() {

	slice := []string{"aMa", "hola", "amor a roma"}

	for _, valor := range slice {
		isPalindromo(valor)
	}
}

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}
```

# Hash Table

```go
package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
```

# Structs

```go
package main

import (
	"fmt"
)

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Toyota", year: 2020}
	fmt.Println(myCar)

	//Otra forma
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}
```

# Modificador de Acceso

```go
package mypackage

import "fmt"

// CarPublic Car de acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// CarPublic Car de acceso privado, lleva minuscula
type CarPrivate struct {
	brand string
	year  int
}

func PrintMesssage() {
	fmt.Println("Hola")
}
```

# Punteros

```go
package myclass

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(rm, dsk int, brn string) Pc {
	myPc := Pc{ram: rm, disk: dsk, brand: brn}
	return myPc
}

func (myPc Pc) FormatPrint() {
	fmt.Printf("Esta pc marca %s cuenta con una ram de %dGB y un disco de %dGB.\n", myPc.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm int) {
	myPc.ram = rm
}
```

```go
package main

import (
	pc "CursoGoPlatzi/myclass"
	"fmt"
)

func main() {
	myPc := pc.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}
```

# Goroutine

```go
import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)

	wg.Wait()

	//Funcion anonima
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second)
}

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}
```

# Channels

```go
import (
	"fmt"
)

func main() {

	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bye", c)

	fmt.Println(<-c)

}

// chan<- canal de entrada
// <-chan canal de salida
func say(text string, c chan<- string) {
	c <- text
}
```

# Multiplexeando Channels

```go
func main() {

	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	//Range y close
	close(c)
	//c <- "Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("email1", email1)
	go message("email2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}

func message(text string, c chan string) {
	c <- text
}
```