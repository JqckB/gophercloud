package services

import "github.com/JqckB/gophercloud/v2"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
