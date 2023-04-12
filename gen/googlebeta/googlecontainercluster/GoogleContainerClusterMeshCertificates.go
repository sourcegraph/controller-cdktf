package googlecontainercluster


type GoogleContainerClusterMeshCertificates struct {
	// When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enable_certificates GoogleContainerCluster#enable_certificates}
	EnableCertificates interface{} `field:"required" json:"enableCertificates" yaml:"enableCertificates"`
}

