package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	pb "github.com/FranZk9/lab3Distribuidos/proto"
	"google.golang.org/grpc"
)

var serv *grpc.Server
var port = ":50051"

type server struct {
	pb.UnsafeMensajeServiceServer
}

func (s *server) GetSoldados(ctx context.Context, msg *pb.CommandGet) (*pb.Message, error) {
	nServidor := rand.Intn(3) + 1
	name := ""
	dir := ""
	if nServidor == 1 {
		dir = "dist085"
		name = "Servidor Fulcrum 1"
	} else if nServidor == 2 {
		dir = "dist086"
		name = "Servidor Fulcrum 2"
	} else if nServidor == 3 {
		dir = "dist087"
		name = "Servidor Fulcrum 3"
	}
	connS, err := grpc.Dial(dir+port, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS.Close()
	fmt.Println("GetSoldados: Conectándose a " + name + " (" + dir + ")")
	serviceCliente := pb.NewMessageServiceClient(connS)
	res, err := serviceCliente.GetSoldados(context.Background(),
		&pb.CommandGet{
			NombreSector: msg.NombreSector,
			NombreBase:   msg.NombreBase,
		})
	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}
	return &pb.Mensaje{Body: res.Body}, nil
}

func (s *server) AgregarBase(ctx context.Context, msg *pb.CommandAddUpdate) (*pb.Mensaje, error) {
	nServidor := rand.Intn(3) + 1
	dir := ""
	if nServidor == 1 {
		dir = "dist085"
	} else if nServidor == 2 {
		dir = "dist086"
	} else if nServidor == 3 {
		dir = "dist087"
	}
	fmt.Println("Agregar Base invocado")
	return &pb.Mensaje{Body: dir}, nil
}
func (s *server) RenombrarBase(ctx context.Context, msg *pb.CommandRename) (*pb.Mensaje, error) {
	nServidor := rand.Intn(3) + 1
	dir := ""
	if nServidor == 1 {
		dir = "dist085"
	} else if nServidor == 2 {
		dir = "dist086"
	} else if nServidor == 3 {
		dir = "dist087"
	}
	fmt.Println("Renombrar Base invocado")
	return &pb.Mensaje{Body: dir}, nil
}
func (s *server) ActualizarValor(ctx context.Context, msg *pb.CommandAddUpdate) (*pb.Mensaje, error) {
	nServidor := rand.Intn(3) + 1
	dir := ""
	if nServidor == 1 {
		dir = "dist085"
	} else if nServidor == 2 {
		dir = "dist086"
	} else if nServidor == 3 {
		dir = "dist087"
	}
	fmt.Println("Actualizar Valor invocado")
	return &pb.Mensaje{Body: dir}, nil
}
func (s *server) BorrarBase(ctx context.Context, msg *pb.CommandDelete) (*pb.Mensaje, error) {
	nServidor := rand.Intn(3) + 1
	dir := ""
	if nServidor == 1 {
		dir = "dist085"
	} else if nServidor == 2 {
		dir = "dist086"
	} else if nServidor == 3 {
		dir = "dist087"
	}
	fmt.Println("Borrar Base invocado")
	return &pb.Mensaje{Body: dir}, nil
}

//----------------------------------------------------------------------------------------------------------------

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {

	fmt.Println("Solicitud de " + req.Mensaje.Nombre + " recibida, mensaje enviado: " + req.Mensaje.Nombre)

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	defer conn.Close()

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
		Mensaje: &pb.Mensaje{
			Nombre: req.Mensaje.Nombre,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	fmt.Println("Estado enviado:", res.Mensajeid)

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func (s *server) CreateLista(ctx context.Context, req *pb.ConsultarLista) (*pb.RespuestaLista, error) {

	fmt.Println("Solicitud de " + req.Estado.Nombre + " recibida")

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	defer conn.Close()

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.CreateLista(context.Background(), &pb.ConsultarLista{
		Estado: &pb.Estado{
			Nombre: req.Estado.Nombre,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	return &pb.RespuestaLista{
		Estadoid: res.Estadoid,
	}, nil

}

func main() {
	fmt.Println("Broker-Luna iniciado")
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("No se creo la conexion tcp " + err.Error())
	}

	serv := grpc.NewServer()

	pb.RegisterMensajeServiceServer(serv, &server{})

	if err = serv.Serve(listner); err != nil {
		panic("No se inicio el server " + err.Error())
	}
}
