package cuentaPalabras

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

var Merged = make(map[string]int)

//var wg sync.WaitGroup

func Capturar_caracteres() string {

	var caracteres string
	var salto_linea string

	fmt.Println("==>Ingresar caracteres:")
	fmt.Print("=>")

	in := bufio.NewReader(os.Stdin)

	caracteres, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(">>>>>>>>>", err)
		panic("No se pudo interpretar los caracteres")
	}

	fmt.Println("==>Desea Incluir el salto de linea entre los delimitadores? S/N")
	fmt.Print("=>")
	fmt.Scanln(&salto_linea)

	if strings.ToUpper(salto_linea) != "S" && salto_linea != "" {
		caracteres = strings.ReplaceAll(caracteres, "\n", "")
	}

	return caracteres

}

func Convertir_a_byte(caracteres string) []byte {

	caracteres_bytes := make([]byte, 256)
	caracteres_bytes = []byte(caracteres)

	fmt.Println("")
	fmt.Println("==>Delimitadores designados ")

	fmt.Print("         ", caracteres_bytes)
	if strings.Contains(caracteres, "\n") {
		fmt.Printf(" (El %d: Representa salto de linea)\n", caracteres_bytes[len(caracteres_bytes)-1])
		fmt.Println("         ", caracteres)
	} else {
		fmt.Println("")
		fmt.Println("         ", caracteres)
	}

	return caracteres_bytes
}

func Procesar_archivo(key_sensitive string, hilos int, archivo string, delimitadores []byte) {

	var archivo_nombre string
	if file, error := os.Open(archivo); error != nil {
		fmt.Println(error)
		panic("No fue posible obtener el archivo!!")

	} else {

		//Al final de la funcion cerramos la conexion con el archivo
		defer func() {
			fmt.Println("==>Cerrando el archivo!")
			file.Close()
		}()

		archivo_nombre = file.Name()
		contenido := make([]byte, 25600000000)
		acumulaByte := make([]byte, 0, 2560)
		palabraByte := make([]string, 0, 100000)

		fmt.Println("")
		fmt.Println("==>Procesando archivo..")

		//Asignamos el contenido del txt en la variable contenido (en binario)
		//Devuelve dos variables, la longitud del archivo y el error (no usamos esa por eso _)
		long, _ := file.Read(contenido)

		if key_sensitive == "S" || key_sensitive == "s" {
			// Agrpa y Convierte los grupos de bytes en palabras
			for i := 0; i < long; i++ {
				if ok := slices.Index(delimitadores, contenido[i]); ok >= 0 && contenido[i-1] != 195 {
					if len(acumulaByte) > 0 {
						palabraByte = append(palabraByte, string(acumulaByte[:len(acumulaByte)]))
					}

					acumulaByte = nil

				} else {
					acumulaByte = append(acumulaByte, contenido[i])
				}

			}

			//Obtengo la ultima palabra
			palabraByte = append(palabraByte, string(acumulaByte[:len(acumulaByte)]))
		} else {
			//Agrpa y Convierte los grupos de bytes en palabras
			for i := 0; i < long; i++ {
				if ok := slices.Index(delimitadores, contenido[i]); ok >= 0 && contenido[i-1] != 195 {
					if len(acumulaByte) > 0 {

						palabraByte = append(palabraByte, strings.ToUpper(string(acumulaByte[:len(acumulaByte)])))
					}

					acumulaByte = nil

				} else {
					acumulaByte = append(acumulaByte, contenido[i])
				}
			}
			//Obtengo la ultima palabra
			palabraByte = append(palabraByte, strings.ToUpper(string(acumulaByte[:len(acumulaByte)])))
		}

		//imprimo el total de palabras encontradas
		fmt.Println("===> Total de palabras encontradas en el archivo:", len(palabraByte))

		//Divido el archivo para procesarlo en los hilos indicados
		//cada hilo integra los resultados en Merged
		size := len(palabraByte)
		var j int

		if len(palabraByte)%hilos == 0 {
			size = len(palabraByte) / hilos
		} else {
			size = len(palabraByte)/hilos + 1
		}
		bucle := 1
		//wg.Add(hilos)
		for h := 0; h < len(palabraByte); h += size {
			j += size
			if j > len(palabraByte) {
				j = len(palabraByte)
			}
			fmt.Println(bucle, " Desde", h, "hasta", j)
			bucle = bucle + 1
			//go contar_palabras(palabraByte[h:j])
			contar_palabras(palabraByte[h:j])

		}
		//wg.Wait()

		//imprimir en pantalla
		fmt.Println("")
		fmt.Printf("===> Se procesaron %d palabras \n", len(palabraByte))
		//Exportar los resultados
		exportar_resultados(archivo_nombre, Merged)

	}
}

func contar_palabras(palabraByte []string) /*map[string]int */ {
	lista := make(map[string]int)
	var palabra string
	ocurrencias := make([]string, 0, 100000)
	contador := 0

	palabraByteFijo := make([]string, 0, 100000)
	palabraByteFijo = palabraByte

	//busca palabra por palabra del archivo  y va contando las ocurrencias
	//una vez contada la palabra, la misma se retira del slicen
	for g := 0; g < len(palabraByteFijo); g++ {
		palabra = palabraByteFijo[g]
		contador = 0
		for h := 0; h < len(palabraByte); h++ {
			if palabraByte[h] == palabra {
				contador = contador + 1
				if h == 0 {
					palabraByte = append(palabraByte[1:1], palabraByte[h+1:]...)
					h--
				} else {
					palabraByte = append(palabraByte[:h], palabraByte[h+1:]...)
					h--
				}
			}

		}

		if contador != 0 {
			lista[palabra] = contador
			palabra = palabra + ": " + strconv.Itoa(contador)
			ocurrencias = append(ocurrencias, palabra)

		}
		if len(palabraByte) == 0 {
			if palabra == ": 1" {
				ocurrencias = append(ocurrencias, "(Esta última ocurrencia corresponde a delimitadores al final del archivo)")
			}
			break
		}

	}
	//Integro los resultados de los hilos
	for keyb, valueb := range lista {
		Merged[keyb] = Merged[keyb] + valueb
	}

	//wg.Done()

}

func exportar_resultados(archivo string, lista map[string]int) {
	type resultado struct {
		palabra  string
		cantidad int
	}
	var mi_resultado []resultado
	var resultado_imprimir []resultado

	var export_result string
	var mensaje string

	for k, v := range lista {
		mi_resultado = append(mi_resultado, resultado{k, v})
	}

	sort.Slice(mi_resultado, func(i, j int) bool {
		if mi_resultado[i].cantidad == mi_resultado[j].cantidad {

			return mi_resultado[i].palabra > mi_resultado[j].palabra

		} else {
			return mi_resultado[i].cantidad > mi_resultado[j].cantidad
		}

	})
	//Imprime en pantalla el top 10 de palabras mas encontradas

	if len(mi_resultado) < 10 {
		resultado_imprimir = mi_resultado[0:len(mi_resultado)]
		mensaje = "===> Resultado del conteo: "
	} else {
		resultado_imprimir = mi_resultado[0:9]
		mensaje = "==> Se expone el top 10 de palabras mas encontradas "
	}
	fmt.Println(mensaje, resultado_imprimir)

	//Crea archivo para resultados
	t := time.Now()
	fecha := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	archivo = archivo + fecha + ".txt"

	f, err1 := os.Create(archivo)

	if err1 != nil {
		fmt.Println("Problemas para crear el archivo de resultados.", err1)
		panic("Error al intentar crear archivo")
	}
	fmt.Println("")
	fmt.Println("==>Exportando resultados..")
	for _, resultado := range mi_resultado {

		//fmt.Printf("%d, %s\n", resultado.cantidad, resultado.palabra)
		export_result = strconv.Itoa(resultado.cantidad) + " " + resultado.palabra + " \n"

		_, err2 := f.WriteString(export_result)

		if err2 != nil {
			fmt.Println("Problemas para escribir el resultado.", err2, export_result)
			panic("Error al intentar escrivir archivo.")
		}
	}
	fmt.Println("==>Se exportaron los resultados en el archivo: ", archivo)

}
