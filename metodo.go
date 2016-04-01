package main

import "fmt"

type jugador struct {	//Establecemos nuestra estructura
	nombre string
	goles int
	partidos int
}

func (j *jugador) average() float64 {	//definimos la función con un apuntador como entrada
	return float64(j.partidos) / float64(j.goles)

}

func main() {
	jugadores := []jugador{	//llenamos de datos nuestra estructura
		{"Carlos",20,60},
		{"Ramón",17,60},
		{"Rafa",34,60},
	}	

	for i, player := range jugadores {		//iteramos para cada jugador
        fmt.Printf("Jugador: %s  Promedio de goles por partido: %f\n", player.nombre, jugadores[i].average())
	}

}