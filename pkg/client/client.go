package client

import (
	"fmt"
	"internal/utils"
	"net"
	"time"
)

const (
	UDP                  = "udp"
	DefaultPaketSize     = 1400
	DataTransferDeadline = 5 * time.Second
)

type Client struct {
	ip   string
	port int
	conn net.Conn
}

func CreateClient(ip string, port int) (c *Client, err error) {
	var client = &Client{
		ip:   ip,
		port: port,
	}

	if ok, err := utils.IsValidIp(ip); !ok {
		return nil, err
	}

	if ok, err := utils.IsValidPort(client.port); !ok {
		return nil, err
	}

	if client.conn, err = net.Dial(UDP, net.JoinHostPort(ip, fmt.Sprint(port))); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Read() ([]byte, error) {
	var buffer = make([]byte, DefaultPaketSize)

	if err := c.conn.SetReadDeadline(time.Now().Add(DataTransferDeadline)); err != nil {
		return nil, err
	}

	if _, err := c.conn.Read(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (c *Client) Write(data []byte) error {
	if err := c.conn.SetWriteDeadline(time.Now().Add(DataTransferDeadline)); err != nil {
		return err
	}

	_, err := c.conn.Write(data)

	return err
}

func (c *Client) Close() {
	c.conn.Close()
}
