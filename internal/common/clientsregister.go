package common

var AllClients = []Client{}

func RegisterClient(c Client) {
	AllClients = append(AllClients, c)
}
