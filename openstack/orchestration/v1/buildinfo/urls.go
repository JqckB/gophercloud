package buildinfo

import "github.com/JqckB/gophercloud/v2"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
