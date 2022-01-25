package main

import "fmt"

func main() {
	fmt.Println("✨BIENVENIDOS A LA SUBASTA FRUFRU✨")
	fmt.Println("INGRESE EL NOMBRE DEL PRODUCTO A SUBASTAR")
	var nombre_producto string
	_, err := fmt.Scanf("%s", &nombre_producto)
	if err != nil {
		panic(err)
	}

	fmt.Println("INGRESE EL VALOR DEL PRODUCTO A SUBASTAR")
	var valor_inicial int
	_, err = fmt.Scanf("%d", &valor_inicial)
	if err != nil {
		panic(err)
	}

	fmt.Println("¡¡¡ QUE INICIE LA SUBASTA !!!")
	valor_actual := valor_inicial
	var nuevo_valor int

	for {
		fmt.Println("EL VALOR ACTUAL ES $", valor_actual)
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
			} else if nuevo_valor > valor_actual {
				valor_actual = nuevo_valor
			}

		} else if sino == "no" {
			fmt.Println("🎉 VENDIDO POR $", valor_actual, "🎉")
			break
		}
	}
}
