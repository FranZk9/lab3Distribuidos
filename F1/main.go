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
	"time"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo22-Laboratorio-3/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeMensajeServiceServer
}

var relojVectorial [3]int = [3]int{0, 0, 0}
var serv *grpc.Server
var logFulcrum1 = "Servidor Fulcrum 1/log.txt"

func (s *server) GetSoldados(ctx context.Context, msg *pb.CommandGet) (*pb.Mensaje, error) {
	res := buscar_valor(msg.NombreSector, msg.NombreBase)
	return &pb.Mensaje{Body: res + "-" + convertir_vector_to_string(relojVectorial) + "-Servidor Fulcrum 1"}, nil
}

func (s *server) AgregarBase(ctx context.Context, msg *pb.CommandAddUpdate) (*pb.Mensaje, error) {
	fmt.Println("AgregarBase invocado")
	var new_text = ""
	nombre_archivo_base := crear_archivo_sector(msg.NombreSector)
	new_text = msg.NombreSector + " " + msg.NombreBase + " " + msg.Valor
	escribir(new_text, nombre_archivo_base)
	new_text = "AgregarBase " + new_text
	escribir(new_text, logFulcrum1)
	aumentar_vector()

	return &pb.Mensaje{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func (s *server) RenombrarBase(ctx context.Context, msg *pb.CommandRename) (*pb.Mensaje, error) {
	fmt.Println("RenombrarBase invocado")
	var new_text = ""
	renombrar_base(msg.NombreSector, msg.NombreBase, msg.NuevoNombre)
	new_text = "RenombrarBase " + msg.NombreSector + " " + msg.NombreBase + " " + msg.NuevoNombre
	escribir(new_text, logFulcrum1)
	aumentar_vector()

	return &pb.Mensaje{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func (s *server) ActualizarValor(ctx context.Context, msg *pb.CommandAddUpdate) (*pb.Mensaje, error) {
	fmt.Println("ActualizarValor invocado")
	var new_text = ""
	actualizar_valor(msg.NombreSector, msg.NombreBase, msg.Valor)
	new_text = "ActualizarValor " + msg.NombreSector + " " + msg.NombreBase + " " + msg.Valor
	escribir(new_text, logFulcrum1)
	aumentar_vector()

	return &pb.Mensaje{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

func (s *server) BorrarBase(ctx context.Context, msg *pb.CommandDelete) (*pb.Mensaje, error) {
	fmt.Println("BorrarBase invocado")
	var new_text = ""
	borrar_base(msg.NombreSector, msg.NombreBase)
	new_text = "BorrarBase " + msg.NombreSector + " " + msg.NombreBase
	escribir(new_text, logFulcrum1)
	aumentar_vector()

	return &pb.Mensaje{Body: convertir_vector_to_vanguardia(relojVectorial)}, nil
}

// funciones

func crear_archivo_sector(texto string) string {

	var arch = "Servidor Fulcrum 1/" + texto + ".txt"
	vacio := []byte("")
	err := os.WriteFile(arch, vacio, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return arch
}

func crear_archivo_log() {
	vacio := []byte("")
	err := os.WriteFile(logFulcrum1, vacio, 0644)
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
	var name = "Servidor Fulcrum 1/" + nombre_sector + ".txt"
	archivo, err := os.Open(name)
	if err != nil {
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
	var name = "Servidor Fulcrum 1/" + nombre_sector + ".txt"
	archivo, err := os.Open(name)
	if err != nil {
		//si no encuentro el valor a actualizar, debo crearlo con los valores actuales
		// seria el mismo proceso si es que intento agregar un valor nuevo
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
	var name = "Servidor Fulcrum 1/" + nombre_sector + ".txt"
	archivo, err := os.Open(name)

	if err != nil {
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
	relojVectorial[0] += 1
}

func propagar_cambios() {
	//conexion con servidor fulcrum 2
	fmt.Println("Propagando cambios al Servidor Fulcrum 2")
	port := ":50052"
	connS, err := grpc.Dial("dist087"+port, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	serviceCliente := pb.NewMessageServiceClient(connS)
	res, err := serviceCliente.Intercambio(context.Background(),
		&pb.Mensaje{
			Body: convertir_vector_to_string(relojVectorial),
		})
	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	defer connS.Close()

	copia := res.Body
	connS2, err2 := grpc.Dial("dist088"+port, grpc.WithInsecure())
	if err2 != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	serviceCliente2 := pb.NewMessageServiceClient(connS2)
	res2, err2 := serviceCliente2.Intercambio(context.Background(),
		&pb.Mensaje{
			Body: copia,
		})
	if err2 != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	defer connS2.Close()

	relojVectorial = convertir_string_to_vector(res2.Body) //convertir string a relojVectorial
	copia1 := res2.Body

	connS3, err3 := grpc.Dial("dist087"+port, grpc.WithInsecure())
	if err3 != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	serviceCliente3 := pb.NewMessageServiceClient(connS3)
	res3, err3 := serviceCliente3.Intercambio(context.Background(),
		&pb.Mensaje{
			Body: "final-" + copia1,
		})
	if err3 != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	defer connS3.Close()
	res3.Body = "2"

	fmt.Println("Vector actualizado Servidor Fulcrum 1-2")
	connS4, err4 := grpc.Dial("dist088"+port, grpc.WithInsecure())
	if err3 != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS4.Close()
	serviceCliente4 := pb.NewMessageServiceClient(connS4)
	res4, err4 := serviceCliente4.Intercambio(context.Background(),
		&pb.Mensaje{
			Body: "final-" + copia1,
		})
	if err4 != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	res4.Body = "2"
	connS4.Close()
	fmt.Println("Vector actualizado Servidor Fulcrum 1-3")
	crear_archivo_log() //vacia el archivo de log anterior
	fmt.Println("Cambios propagados")
}

func buscar_valor(nombre_sector string, nombre_base string) string {

	var name = "Servidor Fulcrum 1/" + nombre_sector + ".txt"
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
	return "No se encontro ningun valor que haga match"
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

func consistenciaEventual() {
	for {
		select {
		case <-time.After(60 * time.Second):
			fmt.Println("Hola")
		}
	}
}

func main() {
	crear_archivo_log()
	fmt.Println("Servidor Fulcrum 1 Iniciado")
	go func() {
		for {
			time.Sleep(time.Second * 60)
			propagar_cambios()
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
