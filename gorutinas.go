package main

import ("fmt"
		"math/rand"
		"sync"
		"time"
	)

var wg sync.WaitGroup

func main(){
	wg.Add(2)
	fmt.Println("Iniciamos los gorutinas ...")
	//Lanazamos una gorutina
	go imprimirCantidad("A")
	//Lanzamos otra gorutina B
	go imprimirCantidad("B")
	fmt.Println("Esperamos a que finalice...")
	wg.Wait()
	fmt.Println("\n Termina el programa")

}

func imprimirCantidad(etiqueta string) {
	defer wg.Done()

	for cantidad := 1; cantidad <= 10; cantidad++{
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Cantidad: %d de %s\n", cantidad, etiqueta) 
	}
}