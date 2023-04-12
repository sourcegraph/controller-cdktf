package googlednsmanagedzone


type GoogleDnsManagedZoneServiceDirectoryConfig struct {
	// namespace block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_managed_zone#namespace GoogleDnsManagedZone#namespace}
	Namespace *GoogleDnsManagedZoneServiceDirectoryConfigNamespace `field:"required" json:"namespace" yaml:"namespace"`
}

