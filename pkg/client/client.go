package client

import (
	"github.com/webitel/engine/discovery"
)

type Client struct {
	rabbit RabbitClient
	grpc   GrpcClient
}

// ! NewClient creates new client for logger.
// * rabbitUrl - connection string to rabbit server
// * clientId - name that will be recognized by consul
// * address - address to connect to consul server
func NewClient(rabbitUrl string, clientId, address string) (*Client, error) {
	disc, err := discovery.NewServiceDiscovery(clientId, address, func() (bool, error) { return true, nil })
	if err != nil {
		return nil, err
	}
	cli := &Client{grpc: NewGrpcClient(disc)}
	rab := NewRabbitClient(rabbitUrl, cli)
	cli.rabbit = rab
	return cli, nil
}

func (c *Client) Open() error {
	err := c.rabbit.Open()
	if err != nil {
		return err
	}
	err = c.grpc.Start()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	c.rabbit.Close()
	c.grpc.Stop()
}

func (c *Client) Rabbit() RabbitClient {
	return c.rabbit
}
