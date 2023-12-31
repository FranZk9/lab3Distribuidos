package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo22-Laboratorio-3/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Escribe la consulta: ")

		scanner.Scan()
		estado := scanner.Text()

		if strings.ToLower(estado) == "exit" {
			break
		}

		serviceClient := pb.NewMensajeServiceClient(conn)

		res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
			Mensaje: &pb.Mensaje{
				Nombre: estado,
			},
		})

		if err != nil {
			panic("no se creo el mensaje" + err.Error())
		}

		fmt.Println("Estado enviado:", res.Mensajeid)
		fmt.Println(" ")
	}
}
