package googlednsmanagedzone


type GoogleDnsManagedZoneServiceDirectoryConfig struct {
	// namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dns_managed_zone#namespace GoogleDnsManagedZone#namespace}
	Namespace *GoogleDnsManagedZoneServiceDirectoryConfigNamespace `field:"required" json:"namespace" yaml:"namespace"`
}

