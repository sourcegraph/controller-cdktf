package googlecontainercluster


type GoogleContainerClusterMasterAuthClientCertificateConfig struct {
	// Whether client certificate authorization is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#issue_client_certificate GoogleContainerCluster#issue_client_certificate}
	IssueClientCertificate interface{} `field:"required" json:"issueClientCertificate" yaml:"issueClientCertificate"`
}

