package limits

import (
	"github.com/JqckB/gophercloud/v2"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
