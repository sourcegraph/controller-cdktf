package googlecontainercluster


type GoogleContainerClusterMasterAuth struct {
	// client_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#client_certificate_config GoogleContainerCluster#client_certificate_config}
	ClientCertificateConfig *GoogleContainerClusterMasterAuthClientCertificateConfig `field:"required" json:"clientCertificateConfig" yaml:"clientCertificateConfig"`
}

