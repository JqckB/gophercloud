package common

import (
	"github.com/JqckB/gophercloud/v2"
	th "github.com/JqckB/gophercloud/v2/testhelper"
	"github.com/JqckB/gophercloud/v2/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient(fakeServer th.FakeServer) *gophercloud.ServiceClient {
	sc := client.ServiceClient(fakeServer)
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
