package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockIps struct {
	// IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#ip GoogleGkeonpremVmwareAdminCluster#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Hostname of the machine. VM's name will be used if this field is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#hostname GoogleGkeonpremVmwareAdminCluster#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

