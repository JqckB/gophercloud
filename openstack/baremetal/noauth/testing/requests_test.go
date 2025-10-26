package testing

import (
	"testing"

	"github.com/JqckB/gophercloud/v2/openstack/baremetal/noauth"
	th "github.com/JqckB/gophercloud/v2/testhelper"
)

func TestNoAuth(t *testing.T) {
	noauthClient, err := noauth.NewBareMetalNoAuth(noauth.EndpointOpts{
		IronicEndpoint: "http://ironic:6385/v1",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", noauthClient.TokenID)
}
