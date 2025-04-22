package googledataprocbatch


type GoogleDataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig struct {
	// Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#dataproc_cluster GoogleDataprocBatch#dataproc_cluster}
	DataprocCluster *string `field:"optional" json:"dataprocCluster" yaml:"dataprocCluster"`
}

