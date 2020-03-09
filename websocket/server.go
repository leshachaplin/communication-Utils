package websocket

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
)

type Server struct {
	onMsg func(msg string)
}

// New Websocket Server
func NewServer(onMessage func(msg string), ctx context.Context, e *echo.Echo) (*Server, error) {

	w := &Server{
		onMsg: onMessage,
	}

	e.GET("/role", func(c echo.Context) error {
		return w.listen(ctx, c.Response(), c.Request())
	})
	return w, nil
}

// websocket handler function
func (s *Server) listen(ctx context.Context, wr http.ResponseWriter, req *http.Request) error {

	websocket.Handler(func(ws *websocket.Conn) {
		fmt.Println("handle server connection")
		defer ws.Close()
		go func() {
			<-ctx.Done()
			ws.Close()
		}()
		for {
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Errorf("message not read %s", err)
				continue
			}

			s.onMsg(msg)
		}
	}).ServeHTTP(wr, req)
	return nil
}
