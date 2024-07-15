package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfig struct {
	// URI for the secret that hosts a certificate. Must be in the format 'projects/PROJECT_NUM/secrets/SECRET_NAME/versions/VERSION_OR_LATEST'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#secret_uri GoogleContainerNodePool#secret_uri}
	SecretUri *string `field:"required" json:"secretUri" yaml:"secretUri"`
}

