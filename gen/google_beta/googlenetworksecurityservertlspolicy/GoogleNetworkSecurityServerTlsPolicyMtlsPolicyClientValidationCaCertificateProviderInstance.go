package googlenetworksecurityservertlspolicy


type GoogleNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance struct {
	// Plugin instance name, used to locate and load CertificateProvider instance configuration.
	//
	// Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#plugin_instance GoogleNetworkSecurityServerTlsPolicy#plugin_instance}
	PluginInstance *string `field:"required" json:"pluginInstance" yaml:"pluginInstance"`
}

