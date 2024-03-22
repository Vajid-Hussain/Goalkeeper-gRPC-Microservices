package auth

import (
	"fmt"
	"log"

	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/auth/pb"
	"github.com/vajid-hussain/grpc-microservice-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

// tlsConfig:=&tls.

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	// b, _ := ioutil.ReadFile("./../ca.cert")
	// cp := x509.NewCertPool()
	// if !cp.AppendCertsFromPEM(b) {
	// 	fmt.Println("credentials: failed to append certificates")
	// }

	// tlsConfig := &tls.Config{
	// 	InsecureSkipVerify: true,
	// 	RootCAs: cp,
	// }

	cred, err := credentials.NewClientTLSFromFile("./service.pem", "")
	// cred, err := credentials.NewClientTLSFromFile("./../service.pem", "")
	if err != nil {
		log.Fatalf("could not process the credentials: %v", err)
	}

	cc, err := grpc.Dial(c.Auth_port, grpc.WithTransportCredentials(cred))
	// fmt.Println(c.Auth_port, "\n", cc)
	if err != nil {
		fmt.Println("cound not connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
