package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfig struct {
	// auxiliary_services_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#auxiliary_services_config GoogleDataprocCluster#auxiliary_services_config}
	AuxiliaryServicesConfig *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig `field:"optional" json:"auxiliaryServicesConfig" yaml:"auxiliaryServicesConfig"`
	// kubernetes_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#kubernetes_cluster_config GoogleDataprocCluster#kubernetes_cluster_config}
	KubernetesClusterConfig *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig `field:"optional" json:"kubernetesClusterConfig" yaml:"kubernetesClusterConfig"`
	// A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output.
	//
	// If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#staging_bucket GoogleDataprocCluster#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
}

