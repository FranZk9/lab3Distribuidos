package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/FranZk9/lab3Distribuidos/proto"
	"google.golang.org/grpc"
)

var relojVectorial [3]int = [3]int{0, 0, 0}

var logFulcrum3 = "Servidor Fulcrum 3/log.txt"

var serv *grpc.Server
var serv2 *grpc.Server

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	coming := strings.Split(msg.Body, "-")
	if coming[0] == "final" {
		fmt.Println("Cambios propagados")
		relojVectorial = convertir_string_to_vector(coming[1])
		crear_archivo_log()
		return &pb.Message{Body: "Listo"}, nil
	} else {
		fmt.Println("Iniciando propagación")
		vector_received := convertir_string_to_vector(coming[0])
		vector_res := convertir_vector_to_string(comparacion_vectores(vector_received, relojVectorial))
		return &pb.Message{Body: vector_res}, nil
	}
}

func (s *server) GetSoldados(ctx context.Context, msg *pb.CommandGet) (*pb.Message, error) {
	res := buscar_valor(msg.NombreSector, msg.NombreBase)
	//res-reloj-minombre
	//1/1/1
	return &pb.Message{Body: res + "-" + convertir_vector_to_string(relojVectorial) + "-Servidor Servidor Fulcrum 3"}, nil
}
func (s *server) AgregarBase(ctx context.Context, msg *pb.CommandAddUpdate) (*pb.Message, error) {
	fmt.Println("AgregarBase invocado")
	var new_text = ""
	nombre_archivo_base := crear_archivo_sector(msg.NombreSector)
	new_text = msg.NombreSector + " " + msg.NombreBase + " " + msg.Valor
	escribir(new_text, nombre_archivo_base)
	new_text = "AgregarBase " + new_text //registrar en el log
	escribir(new_text, logFulcrum3)
	aumentar_vector()

	return &pb.Message{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func (s *server) RenombrarBase(ctx context.Context, msg *pb.CommandRename) (*pb.Message, error) {
	fmt.Println("RenombrarBase invocado")
	var new_text = ""
	renombrar_base(msg.NombreSector, msg.NombreBase, msg.NuevoNombre)
	new_text = "RenombrarBase " + msg.NombreSector + " " + msg.NombreBase + " " + msg.NuevoNombre
	escribir(new_text, logFulcrum3)
	aumentar_vector()

	return &pb.Message{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func (s *server) ActualizarValor(ctx context.Context, msg *pb.CommandAddUpdate) (*pb.Message, error) {
	fmt.Println("ActualizarValor invocado")
	var new_text = ""
	actualizar_valor(msg.NombreSector, msg.NombreBase, msg.Valor)
	new_text = "ActualizarValor " + msg.NombreSector + " " + msg.NombreBase + " " + msg.Valor //registrar en el log
	escribir(new_text, logFulcrum3)
	aumentar_vector()

	return &pb.Message{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func (s *server) BorrarBase(ctx context.Context, msg *pb.CommandDelete) (*pb.Message, error) {
	fmt.Println("BorrarBase invocado")
	var new_text = ""
	borrar_base(msg.NombreSector, msg.NombreBase)
	new_text = "BorrarBase " + msg.NombreSector + " " + msg.NombreBase //registrar en el log
	escribir(new_text, logFulcrum3)
	aumentar_vector()

	return &pb.Message{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func crear_archivo_sector(texto string) string {

	var arch = "Servidor Fulcrum 3/" + texto + ".txt"
	vacio := []byte("")
	err := os.WriteFile(arch, vacio, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return arch
}

func crear_archivo_log() {
	vacio := []byte("")
	err := os.WriteFile(logFulcrum3, vacio, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func escribir(texto string, archivo string) {
	texto_anterior, err := os.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	texto = texto + "\n"
	new := []byte(texto)
	texto_anterior = append(texto_anterior, new...)

	err = os.WriteFile(archivo, texto_anterior, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func sobreEscribir(texto string, archivo string) {
	_, err := os.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	texto = texto + "\n"
	new := []byte(texto)
	err = os.WriteFile(archivo, new, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func renombrar_base(nombre_sector string, nombre_base string, nuevo_nombre string) {
	var name = "Servidor Fulcrum 3/" + nombre_sector + ".txt"
	archivo, err := os.Open(name)
	if err != nil {
		//si no encuentro el archivo, debo crearlo y registrarlo renombrado con valor 0
		nombre_archivo_renom := crear_archivo_sector(nombre_sector)
		new_text := nombre_sector + " " + nuevo_nombre + " " + "0"
		escribir(new_text, nombre_archivo_renom)
		return
	}
	var texto_archivo_actualizado = ""

	fileScanner := bufio.NewScanner(archivo)
	var flag = 0
	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), " ")
		if flag > 0 {
			texto_archivo_actualizado += "\n"
		}
		retorno_sector, retorno_base, retorno_valor := lista[0], lista[1], lista[2]
		flag += 1
		if (retorno_sector == nombre_sector) && (retorno_base == nombre_base) {
			texto_archivo_actualizado += nombre_sector + " " + nuevo_nombre + " " + retorno_valor
		} else {
			texto_archivo_actualizado += retorno_sector + " " + retorno_base + " " + retorno_valor
		}
	}
	sobreEscribir(texto_archivo_actualizado, name)

}

func actualizar_valor(nombre_sector string, nombre_base string, valor string) {
	var name = "Servidor Fulcrum 3/" + nombre_sector + ".txt"
	archivo, err := os.Open(name)
	if err != nil {
		nombre_archivo_actu := crear_archivo_sector(nombre_sector)
		new_text := nombre_sector + " " + nombre_base + " " + valor
		escribir(new_text, nombre_archivo_actu)
		return

	}
	var texto_archivo_actualizado = ""

	fileScanner := bufio.NewScanner(archivo)
	var flag = 0
	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), " ")
		if flag > 0 {
			texto_archivo_actualizado += "\n"
		}
		retorno_sector, retorno_base, retorno_valor := lista[0], lista[1], lista[2]
		flag += 1
		if (retorno_sector == nombre_sector) && (retorno_base == nombre_base) {
			texto_archivo_actualizado += retorno_sector + " " + retorno_base + " " + valor
		} else {
			texto_archivo_actualizado += retorno_sector + " " + retorno_base + " " + retorno_valor
		}
	}
	sobreEscribir(texto_archivo_actualizado, name)
}

func borrar_base(nombre_sector string, nombre_base string) {
	var name = "Servidor Fulcrum 3/" + nombre_sector + ".txt"
	archivo, err := os.Open(name)

	if err != nil {
		//si no se encuentra el archivo, entonces solamente se debe registrar la accion en el log
		//la funcion que invoca esta funcion se encarga de escribirlo en el log
		return
	}

	var texto_archivo_actualizado = ""
	var flag = 0
	fileScanner := bufio.NewScanner(archivo)

	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), " ")
		if flag > 0 {
			texto_archivo_actualizado += "\n"
		}

		retorno_sector, retorno_base, retorno_valor := lista[0], lista[1], lista[2]
		flag += 1
		if (retorno_sector == nombre_sector) && (retorno_base == nombre_base) {
			continue
		} else {
			texto_archivo_actualizado += retorno_sector + " " + retorno_base + " " + retorno_valor
		}
	}
	sobreEscribir(texto_archivo_actualizado, name)
}

func aumentar_vector() {
	relojVectorial[2] += 1

}

func buscar_valor(nombre_sector string, nombre_base string) string {

	var name = "Servidor Fulcrum 3/" + nombre_sector + ".txt"
	arch, err := os.Open(name)

	if err != nil {
		retorno := "Soldados enemigos: 0" + " en el sector: " + nombre_sector + " y base: " + nombre_base
		return retorno
	}
	fileScanner := bufio.NewScanner(arch)

	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), " ")

		retorno_sector, retorno_base, retorno_valor := lista[0], lista[1], lista[2]
		if (nombre_sector == retorno_sector) && (nombre_base == retorno_base) {
			retorno2 := "Soldados enemigos: " + retorno_valor + " en el sector: " + nombre_sector + " y base: " + nombre_base
			return retorno2
		}
	}
	return "no encontre ningun valor que haga match"
}

func convertir_vector_to_string(relojVectorial [3]int) string {
	var text = ""
	text += strconv.Itoa(relojVectorial[0]) + "/"
	text += strconv.Itoa(relojVectorial[1]) + "/"
	text += strconv.Itoa(relojVectorial[2])
	return text
}
func convertir_vector_to_vanguardia(relojVectorial [3]int) string {
	var text = "["
	text += strconv.Itoa(relojVectorial[0]) + ","
	text += strconv.Itoa(relojVectorial[1]) + ","
	text += strconv.Itoa(relojVectorial[2])
	text += "]"
	return text
}

func convertir_string_to_vector(relojVectorial string) [3]int {
	clock_split := strings.Split(relojVectorial, "/")
	clock_split_1, err := strconv.Atoi(clock_split[0])
	if err != nil {
		fmt.Println("Error during conversion")
		panic("No se puede crear el mensaje " + err.Error())
	}
	clock_split_2, err := strconv.Atoi(clock_split[1])
	if err != nil {
		fmt.Println("Error during conversion")
		panic("No se puede crear el mensaje " + err.Error())
	}
	clock_split_3, err := strconv.Atoi(clock_split[2])
	if err != nil {
		fmt.Println("Error during conversion")
		panic("No se puede crear el mensaje " + err.Error())
	}
	var vectortemp [3]int = [3]int{clock_split_1, clock_split_2, clock_split_3}
	return vectortemp
}

func comparacion_vectores(vector1 [3]int, vector2 [3]int) [3]int {
	for i := 0; i < 3; i++ {
		if vector1[i] < vector2[i] {
			vector1[i] = vector2[i]
		}
	}
	return vector1
}

func main() {
	crear_archivo_log()
	fmt.Println("Servidor Fulcrum 3 Iniciado")
	go func() {
		for {
			port2 := ":50052"
			listener2, err2 := net.Listen("tcp", port2) //conexion sincrona

			if err2 != nil {
				panic("No se creo la conexion tcp" + err2.Error())
			}

			serv2 = grpc.NewServer()
			for {
				pb.RegisterMessageServiceServer(serv2, &server{})
				if err2 = serv2.Serve(listener2); err2 != nil {
					panic("No se inicio el server" + err2.Error())
				}
			}
		}
	}()
	for {
		port := ":50051"
		listener, err := net.Listen("tcp", port) //conexion sincrona

		if err != nil {
			panic("No se creo la conexion tcp" + err.Error())
		}

		serv = grpc.NewServer()
		for {
			pb.RegisterMessageServiceServer(serv, &server{})
			if err = serv.Serve(listener); err != nil {
				panic("No se inicio el server" + err.Error())
			}
		}
	}

}
