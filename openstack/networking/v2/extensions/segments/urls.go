package segments

import "github.com/JqckB/gophercloud/v2"

const urlBaase = "segments"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(urlBaase)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(urlBaase, id)
}
