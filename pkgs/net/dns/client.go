package dns

import (
	"context"
	"net"
	"time"

	"github.com/AdguardTeam/dnsproxy/upstream"
	"github.com/miekg/dns"
)

type Client struct {
	upstreamList []upstream.Upstream
}

type BuildClientOptions struct {
	Servers   []string `json:"server"`
	Bootstrap *net.Resolver
}

func BuildClient(opts *BuildClientOptions) *Client {
	if opts == nil {
		return nil
	}

	c := Client{}

	for _, up := range opts.Servers {
		upInst, err := upstream.AddressToUpstream(up, &upstream.Options{
			// InsecureSkipVerify: true,
			Timeout:   time.Second * 5,
			Bootstrap: opts.Bootstrap,
		})

		if err != nil {
			continue
		}

		c.upstreamList = append(c.upstreamList, upInst)
	}

	return &c
}

func (c *Client) Conn() net.Conn {
	conn := dnsConn{}
	conn.roundTrip = func(ctx context.Context, req string) (string, error) {
		msg := new(dns.Msg)

		err := msg.Unpack([]byte(req))
		if err != nil {
			return "", err
		}

		reply, _, err := upstream.ExchangeParallel(c.upstreamList, msg)
		if err != nil {
			return "", err
		}

		data, err := reply.Pack()
		if err != nil {
			return "", err
		}

		return string(data), nil
	}

	return &conn
}

func (c *Client) Resolver() *net.Resolver {
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return c.Conn(), nil
		},
	}
}
