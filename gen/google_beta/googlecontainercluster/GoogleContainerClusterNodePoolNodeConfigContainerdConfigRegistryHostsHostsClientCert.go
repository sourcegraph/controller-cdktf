package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigContainerdConfigRegistryHostsHostsClientCert struct {
	// URI for the Secret Manager secret that hosts the client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#gcp_secret_manager_secret_uri GoogleContainerCluster#gcp_secret_manager_secret_uri}
	GcpSecretManagerSecretUri *string `field:"optional" json:"gcpSecretManagerSecretUri" yaml:"gcpSecretManagerSecretUri"`
}

