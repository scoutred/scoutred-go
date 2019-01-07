package client

func New(key string) *Client {
	return &Client{
		Key: key,
	}
}

// Client has various methods for fetching resources from Scoutred's API
type Client struct {
	// Key is the API key for the instance of the client
	Key string
}
