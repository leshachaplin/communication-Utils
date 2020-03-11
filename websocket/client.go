package websocket

import (
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

func (c *Client) Read(msg []byte) ([]byte, error, int) {
	i, err := c.conn.Read(msg)
	if err != nil {
		return nil, err, i
	}
	return msg, nil, -1
}

//write message to websocket
func (c *Client) Write(mes []byte) error {
	if _, err := c.conn.Write(mes); err != nil {
		return err
	}
	return nil
}
