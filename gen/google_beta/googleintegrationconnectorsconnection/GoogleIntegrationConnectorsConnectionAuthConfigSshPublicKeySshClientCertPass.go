package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass struct {
	// The resource name of the secret version in the format, format as: projects/*\/secrets/*\/versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.6.0/docs/resources/google_integration_connectors_connection#secret_version GoogleIntegrationConnectorsConnection#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

