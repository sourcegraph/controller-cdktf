package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigContainerdConfigRegistryHostsHostsCa struct {
	// URI for the Secret Manager secret that hosts the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#gcp_secret_manager_secret_uri GoogleContainerNodePool#gcp_secret_manager_secret_uri}
	GcpSecretManagerSecretUri *string `field:"optional" json:"gcpSecretManagerSecretUri" yaml:"gcpSecretManagerSecretUri"`
}

