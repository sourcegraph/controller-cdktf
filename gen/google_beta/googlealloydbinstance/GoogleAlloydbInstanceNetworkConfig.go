package googlealloydbinstance


type GoogleAlloydbInstanceNetworkConfig struct {
	// authorized_external_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#authorized_external_networks GoogleAlloydbInstance#authorized_external_networks}
	AuthorizedExternalNetworks interface{} `field:"optional" json:"authorizedExternalNetworks" yaml:"authorizedExternalNetworks"`
	// Enabling public ip for the instance.
	//
	// If a user wishes to disable this,
	// please also clear the list of the authorized external networks set on
	// the same instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#enable_public_ip GoogleAlloydbInstance#enable_public_ip}
	EnablePublicIp interface{} `field:"optional" json:"enablePublicIp" yaml:"enablePublicIp"`
}

