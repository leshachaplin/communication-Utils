package websocket

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Client struct {
	origin string
	url    string
	conn   websocket.Conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// new websocket client
func NewClient(origin, url string) (*Client, error) {

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		return nil, err
	}

	c := &Client{
		origin: origin,
		url:    url,
		conn:   *ws,
	}

	return c, nil
}

func (c *Client) Read() (string, error) {
	return "", fmt.Errorf("not implemented")
}

//write message to websocket
func (c *Client) Write(mes string) error {
	if _, err := c.conn.Write([]byte(mes)); err != nil {
		return err
	}
	return nil
}
