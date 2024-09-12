# Golang DNS Client

Allow change go default dns client with dnsproxy

## Example 

```go
package main

import ( 
    "net"
    "github.com/5aaee9/utils/pkgs/net/dns"
)


func init() {
    // Use cloudflare dns in remain golang net resolve
    c := dns.BuildClient(&dns.BuildClientOptions{Servers: []string{"https://1.1.1.1/dns-query"}})
    net.DefaultResolver = c.Resolver()
}

```