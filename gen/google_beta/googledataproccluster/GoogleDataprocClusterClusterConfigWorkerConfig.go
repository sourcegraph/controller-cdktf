package googledataproccluster


type GoogleDataprocClusterClusterConfigWorkerConfig struct {
	// accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#accelerators GoogleDataprocCluster#accelerators}
	Accelerators interface{} `field:"optional" json:"accelerators" yaml:"accelerators"`
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#disk_config GoogleDataprocCluster#disk_config}
	DiskConfig *GoogleDataprocClusterClusterConfigWorkerConfigDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// The URI for the image to use for this master/worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#image_uri GoogleDataprocCluster#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// The name of a Google Compute Engine machine type to create for the master/worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#machine_type GoogleDataprocCluster#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The name of a minimum generation of CPU family for the master/worker.
	//
	// If not specified, GCP will default to a predetermined computed value for each zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#min_cpu_platform GoogleDataprocCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// The minimum number of primary worker instances to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#min_num_instances GoogleDataprocCluster#min_num_instances}
	MinNumInstances *float64 `field:"optional" json:"minNumInstances" yaml:"minNumInstances"`
	// Specifies the number of worker nodes to create. If not specified, GCP will default to a predetermined computed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#num_instances GoogleDataprocCluster#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
}

