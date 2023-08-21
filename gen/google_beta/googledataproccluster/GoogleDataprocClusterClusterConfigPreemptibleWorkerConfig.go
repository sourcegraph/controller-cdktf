package googledataproccluster


type GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig struct {
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataproc_cluster#disk_config GoogleDataprocCluster#disk_config}
	DiskConfig *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// Specifies the number of preemptible nodes to create. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataproc_cluster#num_instances GoogleDataprocCluster#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
	// Specifies the preemptibility of the secondary nodes. Defaults to PREEMPTIBLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataproc_cluster#preemptibility GoogleDataprocCluster#preemptibility}
	Preemptibility *string `field:"optional" json:"preemptibility" yaml:"preemptibility"`
}

