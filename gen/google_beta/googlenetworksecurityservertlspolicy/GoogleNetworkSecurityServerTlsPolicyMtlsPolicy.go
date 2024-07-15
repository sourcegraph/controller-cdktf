package googlenetworksecurityservertlspolicy


type GoogleNetworkSecurityServerTlsPolicyMtlsPolicy struct {
	// client_validation_ca block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#client_validation_ca GoogleNetworkSecurityServerTlsPolicy#client_validation_ca}
	ClientValidationCa interface{} `field:"optional" json:"clientValidationCa" yaml:"clientValidationCa"`
	// When the client presents an invalid certificate or no certificate to the load balancer, the clientValidationMode specifies how the client connection is handled.
	//
	// Required if the policy is to be used with the external HTTPS load balancing. For Traffic Director it must be empty. Possible values: ["CLIENT_VALIDATION_MODE_UNSPECIFIED", "ALLOW_INVALID_OR_MISSING_CLIENT_CERT", "REJECT_INVALID"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#client_validation_mode GoogleNetworkSecurityServerTlsPolicy#client_validation_mode}
	ClientValidationMode *string `field:"optional" json:"clientValidationMode" yaml:"clientValidationMode"`
	// Reference to the TrustConfig from certificatemanager.googleapis.com namespace. If specified, the chain validation will be performed against certificates configured in the given TrustConfig. Allowed only if the policy is to be used with external HTTPS load balancers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#client_validation_trust_config GoogleNetworkSecurityServerTlsPolicy#client_validation_trust_config}
	ClientValidationTrustConfig *string `field:"optional" json:"clientValidationTrustConfig" yaml:"clientValidationTrustConfig"`
}

