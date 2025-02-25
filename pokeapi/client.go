package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	cache *Cache
}

func NewClient() Client {
	return Client{
		cache: NewCache(3*time.Minute, 15*time.Second),
	}
}

// Retrieve data from the given URL.
// If cached, the cached data will be returned instead.
func (c *Client) Get(url string) ([]byte, error) {
	data, exists := c.cache.Get(url)
	if exists {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return []byte{}, fmt.Errorf("unsuccessful response: %s", res.Status)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	// Make uncached date really slow so it's obvious...
	//time.Sleep(time.Millisecond * 300)

	c.cache.Add(url, data)

	return data, nil
}

// Convenience method to perform GET and unmarshal (successful) response into result
func (c *Client) GetAndUnmarshal(url string, result any) error {
	data, err := c.Get(url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	return nil
}

const baseUrl = "https://pokeapi.co/api/v2"

// Prefixes path with API base URL, or returns override if not nil
func ApiUrlOrOverride(path string, override *string) string {
	if override != nil {
		return *override
	}
	return ApiUrl(path)
}

// Prefixes path with API base URL
func ApiUrl(path string) string {
	return baseUrl + path
}
