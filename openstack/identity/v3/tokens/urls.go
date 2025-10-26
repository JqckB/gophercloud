package tokens

import "github.com/JqckB/gophercloud/v2"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
