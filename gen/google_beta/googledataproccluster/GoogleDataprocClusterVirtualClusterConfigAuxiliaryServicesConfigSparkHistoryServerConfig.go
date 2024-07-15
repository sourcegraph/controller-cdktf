package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig struct {
	// Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#dataproc_cluster GoogleDataprocCluster#dataproc_cluster}
	DataprocCluster *string `field:"optional" json:"dataprocCluster" yaml:"dataprocCluster"`
}

