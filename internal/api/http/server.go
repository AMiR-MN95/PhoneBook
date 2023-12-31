package http

import (
	"PhoneBook/pkg/token"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	token  token.Token
	app    *fiber.App
}

func New(lg *zap.Logger, token token.Token) *Server {
	s := &Server{logger: lg, token: token}

	s.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	v1 := s.app.Group("apo/v1")

	auth := v1.Group("auth")
	auth.Post("/register", s.register)
	auth.Post("/login", s.login)

	contacts := v1.Group("contacts", s.fetchUserId)
	contacts.Get("/", s.needsAuthentication)

	return s
}

func (server *Server) Server(port int) error {
	addr := fmt.Sprintf(":%d", port)
	if err := server.app.Listen(addr); err != nil {
		server.logger.Error("error resolving server", zap.Error(err))

		return err
	}

	return nil
}
