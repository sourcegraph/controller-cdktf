package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterProxy struct {
	// The proxy url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_vmware_admin_cluster#url GkeonpremVmwareAdminCluster#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// A comma-separated list of IP addresses, IP address ranges, host names, and domain names that should not go through the proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_vmware_admin_cluster#no_proxy GkeonpremVmwareAdminCluster#no_proxy}
	NoProxy *string `field:"optional" json:"noProxy" yaml:"noProxy"`
}

