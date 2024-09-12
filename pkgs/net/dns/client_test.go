package dns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func doTestResolver(t *testing.T, c *Client) {
	r := c.Resolver()

	ips, err := r.LookupIP(context.Background(), "ip", "google.com")
	assert.NoError(t, err)
	assert.True(t, len(ips) > 0)
}

func TestUDPDNS(t *testing.T) {
	c := BuildClient(&BuildClientOptions{Servers: []string{"8.8.8.8"}})
	doTestResolver(t, c)
}

func TestDoHDNS(t *testing.T) {
	c := BuildClient(&BuildClientOptions{Servers: []string{"https://1.1.1.1/dns-query"}})
	doTestResolver(t, c)
}
