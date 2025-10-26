package apiversions

import (
	"strings"

	"github.com/JqckB/gophercloud/v2"
	"github.com/JqckB/gophercloud/v2/openstack/utils"
)

func listURL(c *gophercloud.ServiceClient) string {
	baseEndpoint, _ := utils.BaseEndpoint(c.Endpoint)
	endpoint := strings.TrimRight(baseEndpoint, "/") + "/"
	return endpoint
}
