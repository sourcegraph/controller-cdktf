package googlednsmanagedzone


type GoogleDnsManagedZoneForwardingConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_managed_zone#target_name_servers GoogleDnsManagedZone#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

