package websocket

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

type Server struct {
	onMsg func(msg interface{})
}

// New Websocket Server
func NewServer(onMessage func(msg interface{}), ctx context.Context, e *echo.Echo) (*Server, error) {

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

			var msg interface{}

			err := websocket.JSON.Receive(ws, msg)
			if err != nil {
				log.Errorf("message not read %s", err)
				continue
			}

			go func(msg interface{}) {
				s.onMsg(msg)
			}(msg)

			time.Sleep(time.Second)
		}
	}).ServeHTTP(wr, req)
	return nil
}

func (s *Server) send(ctx context.Context, wr http.ResponseWriter, req *http.Request, msg interface{}) error {

	websocket.Handler(func(ws *websocket.Conn) {
		fmt.Println("handle server connection")
		defer ws.Close()
		go func() {
			<-ctx.Done()
			ws.Close()
		}()
		for {

			err := websocket.JSON.Send(ws, msg)
			if err != nil {
				log.Errorf("message not read %s", err)
				continue
			}

			time.Sleep(time.Second)
		}
	}).ServeHTTP(wr, req)
	return nil
}
