package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	pb "github.com/FranZk9/lab3Distribuidos/proto"
	"google.golang.org/grpc"
)

var last_register_sector = ""
var last_register_base = ""
var last_register_dir = ""
var last_register_clock = ""
var serverdir = ""

func main() {
	patron, compileErr := regexp.Compile("^(AgregarBase|RenombrarBase|ActualizarValor|BorrarBase)")
	if compileErr != nil {
		fmt.Println("Error de compilacion de regex: ", compileErr)
	}

	port := ":50051"
	connS, err := grpc.Dial("dist085"+port, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS.Close()
	serviceCliente := pb.NewMessageServiceClient(connS)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresar cualquiera de los siguientes comandos: \n1)AgregarBase <nombre sector> <nombre base> [valor] \n2)RenombrarBase <nombre sector> <nombre base> <nuevo nombre> \n3)ActualizarValor <nombre sector> <nombre base> <valor> \n4)BorrarBase <nombre sector> <nombre base> \n")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		match := patron.MatchString(text)
		if match {
			input := strings.Split(text, " ")
			if input[0] == "AgregarBase" {
				nombre_sector := input[1]
				nombre_base := input[2]
				valor := "0"
				if len(input) == 4 {
					valor = input[3]
				}

				if last_register_sector == nombre_sector && last_register_base == nombre_base {
					fmt.Println("ReadYourWrites aplicado")
					serverdir = last_register_dir
				} else {
					res, err := serviceCliente.AgregarBase(context.Background(),
						&pb.CommandAddUpdate{
							NombreSector: nombre_sector,
							NombreBase:   nombre_base,
							Valor:        valor,
						})
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}
					serverdir = res.Body
				}

				fmt.Println("Conectándose a " + serverdir)
				//conexión con el servidor
				connS2, err2 := grpc.Dial(serverdir+port, grpc.WithInsecure())
				if err2 != nil {
					panic("No se pudo conectar con el servidor" + err2.Error())
				}
				defer connS2.Close()
				serviceCliente2 := pb.NewMessageServiceClient(connS2)
				res, err := serviceCliente2.AgregarBase(context.Background(),
					&pb.CommandAddUpdate{
						NombreSector: nombre_sector,
						NombreBase:   nombre_base,
						Valor:        valor,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}

				last_register_sector = nombre_sector
				last_register_base = nombre_base
				last_register_dir = serverdir
				last_register_clock = res.Body
				fmt.Println("Se añadió con exito la base " + nombre_base + " en el sector " + nombre_sector + " con valor de " + valor + " - Reloj: " + last_register_clock)
			} else if input[0] == "RenombrarBase" {
				nombre_sector := input[1]
				nombre_base := input[2]
				nuevo_nombre := input[3]

				if last_register_sector == nombre_sector && last_register_base == nombre_base {
					fmt.Println("ReadYourWrites aplicado")
					serverdir = last_register_dir
				} else {
					res, err := serviceCliente.RenombrarBase(context.Background(),
						&pb.CommandRename{
							NombreSector: nombre_sector,
							NombreBase:   nombre_base,
							NuevoNombre:  nuevo_nombre,
						})
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}
					fmt.Println(res.Body)
					serverdir = res.Body
				}

				fmt.Println("Conectándose a " + serverdir)
				connS2, err2 := grpc.Dial(serverdir+port, grpc.WithInsecure())
				if err2 != nil {
					panic("No se pudo conectar con el servidor" + err2.Error())
				}
				defer connS2.Close()
				serviceCliente2 := pb.NewMessageServiceClient(connS2)
				res, err := serviceCliente2.RenombrarBase(context.Background(),
					&pb.CommandRename{
						NombreSector: nombre_sector,
						NombreBase:   nombre_base,
						NuevoNombre:  nuevo_nombre,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				last_register_sector = nombre_sector
				last_register_base = nuevo_nombre
				last_register_dir = serverdir
				last_register_clock = res.Body
				fmt.Println("Se renombró con exito la base " + nombre_base + " a " + nuevo_nombre + " en el sector " + nombre_sector + " - Reloj: " + last_register_clock)
			} else if input[0] == "ActualizarValor" {
				nombre_sector := input[1]
				nombre_base := input[2]
				valor := input[3]

				if last_register_sector == nombre_sector && last_register_base == nombre_base {
					fmt.Println("ReadYourWrites aplicado")
					serverdir = last_register_dir
				} else {
					res, err := serviceCliente.ActualizarValor(context.Background(),
						&pb.CommandAddUpdate{
							NombreSector: nombre_sector,
							NombreBase:   nombre_base,
							Valor:        valor,
						})
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}
					fmt.Println(res.Body)
					serverdir = res.Body
				}
				connS2, err2 := grpc.Dial(serverdir+port, grpc.WithInsecure())
				if err2 != nil {
					panic("No se pudo conectar con el servidor" + err2.Error())
				}
				defer connS2.Close()
				serviceCliente2 := pb.NewMessageServiceClient(connS2)
				res, err := serviceCliente2.ActualizarValor(context.Background(),
					&pb.CommandAddUpdate{
						NombreSector: nombre_sector,
						NombreBase:   nombre_base,
						Valor:        valor,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				last_register_sector = nombre_sector
				last_register_base = nombre_base
				last_register_dir = serverdir
				last_register_clock = res.Body
				fmt.Println("Se actualizo con exito el valor de la base " + nombre_base + " en el sector " + nombre_sector + " con un nuevo valor de " + valor + " - Reloj: " + last_register_clock)
			} else if input[0] == "BorrarBase" {
				nombre_sector := input[1]
				nombre_base := input[2]
				if last_register_sector == nombre_sector && last_register_base == nombre_base {
					fmt.Println("ReadYourWrites aplicado")
					serverdir = last_register_dir
				} else {
					res, err := serviceCliente.BorrarBase(context.Background(),
						&pb.CommandDelete{
							NombreSector: nombre_sector,
							NombreBase:   nombre_base,
						})
					if err != nil {
						panic("No se puede crear el mensaje " + err.Error())
					}
					fmt.Println(res.Body)
					serverdir = res.Body
				}
				connS2, err2 := grpc.Dial(serverdir+port, grpc.WithInsecure())
				if err2 != nil {
					panic("No se pudo conectar con el servidor" + err2.Error())
				}
				defer connS2.Close()
				serviceCliente2 := pb.NewMessageServiceClient(connS2)
				res, err := serviceCliente2.BorrarBase(context.Background(),
					&pb.CommandDelete{
						NombreSector: nombre_sector,
						NombreBase:   nombre_base,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}

				last_register_sector = nombre_sector
				last_register_base = nombre_base
				last_register_dir = serverdir
				last_register_clock = res.Body
				fmt.Println("Se borró con éxito " + nombre_base + " en el sector " + nombre_sector + " - Reloj: " + last_register_clock)
			}
		} else {
			fmt.Println("Comando inválido")
		}
	}
}
