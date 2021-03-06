package testing

import (
	"testing"

	"gophercloud-lc/openstack/cdn/v1/base"
	th "gophercloud-lc/testhelper"
	fake "gophercloud-lc/testhelper/client"
)

func TestGetHomeDocument(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := base.Get(fake.ServiceClient()).Extract()
	th.CheckNoErr(t, err)

	expected := base.HomeDocument{
		"rel/cdn": map[string]interface{}{
			"href-template": "services{?marker,limit}",
			"href-vars": map[string]interface{}{
				"marker": "param/marker",
				"limit":  "param/limit",
			},
			"hints": map[string]interface{}{
				"allow": []string{"GET"},
				"formats": map[string]interface{}{
					"application/json": map[string]interface{}{},
				},
			},
		},
	}
	th.CheckDeepEquals(t, expected, *actual)
}

func TestPing(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePingSuccessfully(t)

	err := base.Ping(fake.ServiceClient()).ExtractErr()
	th.CheckNoErr(t, err)
}
