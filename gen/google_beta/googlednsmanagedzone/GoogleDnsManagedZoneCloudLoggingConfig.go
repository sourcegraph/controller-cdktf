package googlednsmanagedzone


type GoogleDnsManagedZoneCloudLoggingConfig struct {
	// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#enable_logging GoogleDnsManagedZone#enable_logging}
	EnableLogging interface{} `field:"required" json:"enableLogging" yaml:"enableLogging"`
}

