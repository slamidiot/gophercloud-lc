package diskconfig

import "gophercloud-lc/openstack/compute/v2/servers"

type ServerWithDiskConfig struct {
	servers.Server
	DiskConfig DiskConfig `json:"OS-DCF:diskConfig"`
}

func (s ServerWithDiskConfig) ToServerCreateResult() (m map[string]interface{}) {
	m["OS-DCF:diskConfig"] = s.DiskConfig
	return
}

type CreateServerResultBuilder interface {
	ToServerCreateResult() map[string]interface{}
}
