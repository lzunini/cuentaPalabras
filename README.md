# Descargar Paquete:

## go get github.com/lzunini/cuentaPalabras
<br>

# Usar Paquete:

## import "github.com/lzunini/cuentaPalabras"
<br>

# Funciones:

## Capturar_caracteres 

* Solicita y Captura los cacarteres ingresados en la terminal.
* Devuelve string que representa los caracteres ingresados. 

**var caracteres strin**
<br>

**caracteres = cuentaPalabras.Capturar_caracteres()**
<br>

## Convertir_a_byte

* Recibe un string y lo convierte a byte.



**caracteres := ",;¿?(). ¡!\n"**
<br>

**caracteres_bytes := make([]byte, 256)**
<br>

**caracteres_bytes = cuentaPalabras.Convertir_a_byte(caracteres)**
<br>


## Procesar_archivo

* Parametro uno: string que indica si el procesamiento sera key sensitive ("S" o "s") o no.
* Parametro dos: int que indica cantidad de split a aplicar en el archivo a procesar.
* Parametro tres: string que indica ruta y nombre del archivo a procesar (ej /home/nombre/archivo.txt).
* Parametro cuatro: array de byte que indican los delimitadores o separadores de palabras a aplicar.
* ..


Esta funcion, imprime en pantalla la cantidad total de palabras y el el top 10 de las palabras (palabra1: cantidad de ocurrencias ... palabra10: cantidad de ocurrencias)
y exporta en formato txt el resultado con el conteo de cada palabra del archivo, ordenado por ocurrencia y alfabeticamente (mismo directorio del archivo/ mismo nombre del archivo + yyyymmddhhmmss + .txt). 


**key_sensitive := "S"**
<br>

**split := 1000**
<br>

**archivo := "/home/nombre/archivo.txt"**
<br>

**caracteres := ",;¿?(). ¡!\n"**
<br>

**caracteres_bytes := make([]byte, 256)**
<br>

**caracteres_bytes = []byte(caracteres)**
<br>

**cuentaPalabras.Procesar_archivo(key_sensitive, split, archivo, caracteres_bytes)**
