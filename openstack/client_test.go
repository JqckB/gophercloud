package openstack

import tokens2 "github.com/JqckB/gophercloud/v2/openstack/identity/v2/tokens"

var _ tokens2.AuthOptionsBuilder = &v2TokenNoReauth{}
