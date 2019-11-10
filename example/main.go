package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thetechnick/gobetween-client-go/gobetween"
	"github.com/thetechnick/gobetween-client-go/gobetween/api"
)

func main() {
	c := gobetween.NewClient("http://localhost:8888", gobetween.WithAuth(&gobetween.BasicAuth{
		Username: "admin",
		Password: "1111",
	}))

	ctx := context.Background()
	// s, err := c.ListServers(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	server := &api.Server{
		Bind:     "0.0.0.0:8889",
		Protocol: "tcp",
		Balance:  "roundrobin",
		Discovery: &api.DiscoveryConfig{
			Kind: "static",
			StaticDiscoveryConfig: &api.StaticDiscoveryConfig{
				StaticList: []string{
					"google.de:80",
				},
			},
		},
	}

	err := c.CreateServer(ctx, "test-server", server)
	if err != nil {
		panic(err)
	}

	s, err := c.GetServer(ctx, "test-server")
	if err != nil {
		panic(err)
	}

	j, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))
}
