package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterNetworkConfigStaticIpConfigIpBlocks struct {
	// The network gateway used by the VMware Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#gateway GoogleGkeonpremVmwareAdminCluster#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#ips GoogleGkeonpremVmwareAdminCluster#ips}
	Ips interface{} `field:"required" json:"ips" yaml:"ips"`
	// The netmask used by the VMware Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#netmask GoogleGkeonpremVmwareAdminCluster#netmask}
	Netmask *string `field:"required" json:"netmask" yaml:"netmask"`
}

