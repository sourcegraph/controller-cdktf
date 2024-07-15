package googlenetworksecurityservertlspolicy


type GoogleNetworkSecurityServerTlsPolicyServerCertificate struct {
	// certificate_provider_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#certificate_provider_instance GoogleNetworkSecurityServerTlsPolicy#certificate_provider_instance}
	CertificateProviderInstance *GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance `field:"optional" json:"certificateProviderInstance" yaml:"certificateProviderInstance"`
	// grpc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#grpc_endpoint GoogleNetworkSecurityServerTlsPolicy#grpc_endpoint}
	GrpcEndpoint *GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint `field:"optional" json:"grpcEndpoint" yaml:"grpcEndpoint"`
}

