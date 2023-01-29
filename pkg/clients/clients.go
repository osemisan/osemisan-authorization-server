package clients

import "golang.org/x/exp/slices"

type Client struct {
	Id string
	Secret string
	URIs []string
}

type Clients []Client

var OsemisanClients = Clients{
	{
		"osemisan-client-id-1",
		"osemisan-client-secret-1",
		[]string{"http://localhost:9000"},
	},
}

func (cs Clients) Find (id string) int {
	idx := slices.IndexFunc(cs, func (c Client) bool {
		return c.Id == id
	})
	return idx
}
