// +build acceptance

package v3

import (
	"testing"

	"gophercloud-lc/acceptance/clients"
	"gophercloud-lc/acceptance/tools"
	"gophercloud-lc/openstack/identity/v3/services"
)

func TestServicesList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	allPages, err := services.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list services: %v", err)
	}

	allServices, err := services.ExtractServices(allPages)
	if err != nil {
		t.Fatalf("Unable to extract services: %v", err)
	}

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}

}
