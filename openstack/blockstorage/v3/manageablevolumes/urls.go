package manageablevolumes

import "github.com/JqckB/gophercloud/v2"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("manageable_volumes")
}
