package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// : completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		// : simular trabajo
		time.Sleep(500 * time.Millisecond)

		ch <- x * x
	}()
	return ch
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// : crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(3)
	f3 := asyncCuadrado(4)

	// : Opción 1: esperar cada futuro secuencialmente
	fmt.Println("Secuencial:")
	fmt.Println(<-f1)
	fmt.Println(<-f2)
	fmt.Println(<-f3)

	// : Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.

	f4 := asyncCuadrado(5)
	f5 := asyncCuadrado(6)
	f6 := asyncCuadrado(7)

	fmt.Println("Fan-in:")
	for v := range fanIn(f4, f5, f6) {
		fmt.Println(v)
	}
}
