package main

import "fmt"

func main() {
	fmt.Println("✨BIENVENIDOS A LA SUBASTA FRUFRU✨")
	fmt.Println("INGRESE EL NOMBRE DEL PRODUCTO A SUBASTAR")
	var nombre string
	_, err := fmt.Scanf("%s", &nombre)
	if err != nil {
		panic(err)
	}

	fmt.Println("INGRESE EL VALOR DEL PRODUCTO A SUBASTAR")
	var valor_inicial int
	_, err = fmt.Scanf("%d", &valor_inicial)
	if err != nil {
		panic(err)
	}

	fmt.Println("!!! QUE INICIE LA SUBASTA")
	valor_actual := valor_inicial
	fmt.Println("EL VALOR ACTUAL ES", valor_actual)
	var nuevo_valor int
	valor_actualizado := nuevo_valor

	for {
		fmt.Println("¿ALGUIEN DA MÁS?")
		var sino string
		_, err := fmt.Scanf("%s", &sino)
		if err != nil {
			panic(err)
		}

		if sino == "si" {
			fmt.Println("INGRESE EL NUEVO VALOR A OFERTAR")
			_, err := fmt.Scanf("%d", &nuevo_valor)
			if err != nil {
				panic(err)
			}

			if nuevo_valor <= valor_actual {
				fmt.Println("EL VALOR OFERTADO NO PUEDE SER MENOR O IGUAL AL VALOR ESTABLECIDO")
				fmt.Println("EL VALOR ACTUAL ES", valor_actual)
				fmt.Println("INGRESE EL NUEVO VALOR A OFERTAR")
				_, err := fmt.Scanf("%d", &nuevo_valor)
				if err != nil {
					panic(err)
				}
			} else if nuevo_valor > valor_actual {
				fmt.Println("EL VALOR ACTUAL ES", valor_actualizado)
				fmt.Println("¿ALGUIEN DA MÁS?")
				_, err := fmt.Scanf("%s", &sino)
				if err != nil {
					panic(err)
				}
			}

		} else if sino == "no" {
			fmt.Println("VENDIDO POR", valor_actual)
			break
		}
	}
}
