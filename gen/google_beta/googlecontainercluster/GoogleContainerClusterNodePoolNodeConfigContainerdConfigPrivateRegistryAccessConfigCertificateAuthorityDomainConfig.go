package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfig struct {
	// List of fully-qualified-domain-names. IPv4s and port specification are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#fqdns GoogleContainerCluster#fqdns}
	Fqdns *[]*string `field:"required" json:"fqdns" yaml:"fqdns"`
	// gcp_secret_manager_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#gcp_secret_manager_certificate_config GoogleContainerCluster#gcp_secret_manager_certificate_config}
	GcpSecretManagerCertificateConfig *GoogleContainerClusterNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfig `field:"required" json:"gcpSecretManagerCertificateConfig" yaml:"gcpSecretManagerCertificateConfig"`
}

