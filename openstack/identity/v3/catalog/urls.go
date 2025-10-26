package catalog

import "github.com/JqckB/gophercloud/v2"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}
