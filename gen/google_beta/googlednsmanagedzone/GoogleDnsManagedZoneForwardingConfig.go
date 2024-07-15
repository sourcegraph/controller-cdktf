package googlednsmanagedzone


type GoogleDnsManagedZoneForwardingConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#target_name_servers GoogleDnsManagedZone#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

