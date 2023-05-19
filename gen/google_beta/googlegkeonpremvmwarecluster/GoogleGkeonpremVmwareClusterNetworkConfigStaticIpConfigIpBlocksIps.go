package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterNetworkConfigStaticIpConfigIpBlocksIps struct {
	// Hostname of the machine. VM's name will be used if this field is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_gkeonprem_vmware_cluster#hostname GoogleGkeonpremVmwareCluster#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_gkeonprem_vmware_cluster#ip GoogleGkeonpremVmwareCluster#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

