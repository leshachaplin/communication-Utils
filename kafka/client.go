package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type Client struct {
	connection *kafka.Conn
	reader     *kafka.Reader
}

// Close kafka connection
func (k *Client) Close() error {
	err := k.connection.Close()
	err = k.reader.Close()
	return err
}

// New kafka client
func New(topic string, url string, group string) (*Client, error) {

	senderMsg, err := kafka.DialLeader(context.Background(), "tcp", url, topic, 0)
	if err != nil {
		return nil, err
	}

	readerMsg := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{url},
		Topic:     topic,
		GroupID:   fmt.Sprintf("%v", group),
		Partition: 0,
	})
	return &Client{
		connection: senderMsg,
		reader:     readerMsg,
	}, nil
}

// write message to kafka
func (k *Client) WriteMessage(msg []byte) error {
	//log
	_, err := k.connection.WriteMessages(kafka.Message{
		Value: msg,
	})
	return err
}

// read message to kafka
func (k *Client) ReadMessage() ([]byte, error) {
	//log

	ctx, _ := context.WithTimeout(context.Background(), time.Second/10)
	m, err := k.reader.ReadMessage(ctx)
	return m.Value, err
}

func (k *Client) FetchMessage() (kafka.Message, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second/10)
	m, err := k.reader.FetchMessage(ctx)
	return m, err
}

func (k *Client) Commit(msg kafka.Message) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second/10)
	err := k.reader.CommitMessages(ctx, msg)
	return err
}

func (k *Client) SetOffset(offset int64) error {
	err := k.reader.SetOffset(offset)
	return err
}
