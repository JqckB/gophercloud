package extensions

import "github.com/JqckB/gophercloud/v2"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
