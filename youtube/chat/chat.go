package chat

import(
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	log.Panicf("Receiced message body form client: %s", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}