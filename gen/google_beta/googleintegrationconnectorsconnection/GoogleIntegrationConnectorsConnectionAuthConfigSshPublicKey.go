package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey struct {
	// The user account used to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#username GoogleIntegrationConnectorsConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Format of SSH Client cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#cert_type GoogleIntegrationConnectorsConnection#cert_type}
	CertType *string `field:"optional" json:"certType" yaml:"certType"`
	// ssh_client_cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#ssh_client_cert GoogleIntegrationConnectorsConnection#ssh_client_cert}
	SshClientCert *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert `field:"optional" json:"sshClientCert" yaml:"sshClientCert"`
	// ssh_client_cert_pass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#ssh_client_cert_pass GoogleIntegrationConnectorsConnection#ssh_client_cert_pass}
	SshClientCertPass *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass `field:"optional" json:"sshClientCertPass" yaml:"sshClientCertPass"`
}

