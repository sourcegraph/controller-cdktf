package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterPrivateRegistryConfig struct {
	// The registry address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_vmware_admin_cluster#address GoogleGkeonpremVmwareAdminCluster#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The CA certificate public key for private registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_vmware_admin_cluster#ca_cert GoogleGkeonpremVmwareAdminCluster#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
}

