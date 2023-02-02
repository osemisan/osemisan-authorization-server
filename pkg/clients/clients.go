package clients

import "golang.org/x/exp/slices"

type Client struct {
	Id     string
	Secret string
	URIs   []string
	// セミの種類をスペースで区切った文字列
	Scope string
}

type Clients []Client

var OsemisanClients = Clients{
	{
		"osemisan-client-id-1",
		"osemisan-client-secret-1",
		[]string{"http://localhost:9000/callback"},
		"abura kuma",
	},
}

func (cs Clients) Find(id string) int {
	idx := slices.IndexFunc(cs, func(c Client) bool {
		return c.Id == id
	})
	return idx
}

func (c Client) ContainsURI(uri string) bool {
	return slices.Contains(c.URIs, uri)
}
