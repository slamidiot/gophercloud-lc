package testing

import (
	"testing"

	"gophercloud-lc/openstack/identity/v2/tenants"
	"gophercloud-lc/pagination"
	th "gophercloud-lc/testhelper"
	"gophercloud-lc/testhelper/client"
)

func TestListTenants(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantsSuccessfully(t)

	count := 0
	err := tenants.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := tenants.ExtractTenants(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedTenantSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}
