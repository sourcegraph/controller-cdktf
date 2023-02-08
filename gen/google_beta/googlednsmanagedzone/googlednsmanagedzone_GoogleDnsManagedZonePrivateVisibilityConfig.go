package googlednsmanagedzone


type GoogleDnsManagedZonePrivateVisibilityConfig struct {
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_managed_zone#networks GoogleDnsManagedZone#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
}

