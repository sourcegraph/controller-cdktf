package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionSslConfig struct {
	// Enum for controlling the SSL Type (TLS/MTLS) Possible values: ["TLS", "MTLS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#type GoogleIntegrationConnectorsConnection#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#additional_variable GoogleIntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_certificate GoogleIntegrationConnectorsConnection#client_certificate}
	ClientCertificate *GoogleIntegrationConnectorsConnectionSslConfigClientCertificate `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// Type of Client Cert (PEM/JKS/.. etc.) Possible values: ["PEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_cert_type GoogleIntegrationConnectorsConnection#client_cert_type}
	ClientCertType *string `field:"optional" json:"clientCertType" yaml:"clientCertType"`
	// client_private_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_private_key GoogleIntegrationConnectorsConnection#client_private_key}
	ClientPrivateKey *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKey `field:"optional" json:"clientPrivateKey" yaml:"clientPrivateKey"`
	// client_private_key_pass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_private_key_pass GoogleIntegrationConnectorsConnection#client_private_key_pass}
	ClientPrivateKeyPass *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPass `field:"optional" json:"clientPrivateKeyPass" yaml:"clientPrivateKeyPass"`
	// private_server_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#private_server_certificate GoogleIntegrationConnectorsConnection#private_server_certificate}
	PrivateServerCertificate *GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificate `field:"optional" json:"privateServerCertificate" yaml:"privateServerCertificate"`
	// Type of Server Cert (PEM/JKS/.. etc.) Possible values: ["PEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#server_cert_type GoogleIntegrationConnectorsConnection#server_cert_type}
	ServerCertType *string `field:"optional" json:"serverCertType" yaml:"serverCertType"`
	// Enum for Trust Model Possible values: ["PUBLIC", "PRIVATE", "INSECURE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#trust_model GoogleIntegrationConnectorsConnection#trust_model}
	TrustModel *string `field:"optional" json:"trustModel" yaml:"trustModel"`
	// Bool for enabling SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#use_ssl GoogleIntegrationConnectorsConnection#use_ssl}
	UseSsl interface{} `field:"optional" json:"useSsl" yaml:"useSsl"`
}

