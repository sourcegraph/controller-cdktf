package googledataproccluster


type GoogleDataprocClusterClusterConfig struct {
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#autoscaling_config GoogleDataprocCluster#autoscaling_config}
	AutoscalingConfig *GoogleDataprocClusterClusterConfigAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// auxiliary_node_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#auxiliary_node_groups GoogleDataprocCluster#auxiliary_node_groups}
	AuxiliaryNodeGroups interface{} `field:"optional" json:"auxiliaryNodeGroups" yaml:"auxiliaryNodeGroups"`
	// dataproc_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#dataproc_metric_config GoogleDataprocCluster#dataproc_metric_config}
	DataprocMetricConfig *GoogleDataprocClusterClusterConfigDataprocMetricConfig `field:"optional" json:"dataprocMetricConfig" yaml:"dataprocMetricConfig"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#encryption_config GoogleDataprocCluster#encryption_config}
	EncryptionConfig *GoogleDataprocClusterClusterConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#endpoint_config GoogleDataprocCluster#endpoint_config}
	EndpointConfig *GoogleDataprocClusterClusterConfigEndpointConfig `field:"optional" json:"endpointConfig" yaml:"endpointConfig"`
	// gce_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#gce_cluster_config GoogleDataprocCluster#gce_cluster_config}
	GceClusterConfig *GoogleDataprocClusterClusterConfigGceClusterConfig `field:"optional" json:"gceClusterConfig" yaml:"gceClusterConfig"`
	// initialization_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#initialization_action GoogleDataprocCluster#initialization_action}
	InitializationAction interface{} `field:"optional" json:"initializationAction" yaml:"initializationAction"`
	// lifecycle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#lifecycle_config GoogleDataprocCluster#lifecycle_config}
	LifecycleConfig *GoogleDataprocClusterClusterConfigLifecycleConfig `field:"optional" json:"lifecycleConfig" yaml:"lifecycleConfig"`
	// master_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#master_config GoogleDataprocCluster#master_config}
	MasterConfig *GoogleDataprocClusterClusterConfigMasterConfig `field:"optional" json:"masterConfig" yaml:"masterConfig"`
	// metastore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#metastore_config GoogleDataprocCluster#metastore_config}
	MetastoreConfig *GoogleDataprocClusterClusterConfigMetastoreConfig `field:"optional" json:"metastoreConfig" yaml:"metastoreConfig"`
	// preemptible_worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#preemptible_worker_config GoogleDataprocCluster#preemptible_worker_config}
	PreemptibleWorkerConfig *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig `field:"optional" json:"preemptibleWorkerConfig" yaml:"preemptibleWorkerConfig"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#security_config GoogleDataprocCluster#security_config}
	SecurityConfig *GoogleDataprocClusterClusterConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#software_config GoogleDataprocCluster#software_config}
	SoftwareConfig *GoogleDataprocClusterClusterConfigSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// The Cloud Storage staging bucket used to stage files, such as Hadoop jars, between client machines and the cluster.
	//
	// Note: If you don't explicitly specify a staging_bucket then GCP will auto create / assign one for you. However, you are not guaranteed an auto generated bucket which is solely dedicated to your cluster; it may be shared with other clusters in the same region/zone also choosing to use the auto generation option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#staging_bucket GoogleDataprocCluster#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
	// The Cloud Storage temp bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files.
	//
	// Note: If you don't explicitly specify a temp_bucket then GCP will auto create / assign one for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#temp_bucket GoogleDataprocCluster#temp_bucket}
	TempBucket *string `field:"optional" json:"tempBucket" yaml:"tempBucket"`
	// worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#worker_config GoogleDataprocCluster#worker_config}
	WorkerConfig *GoogleDataprocClusterClusterConfigWorkerConfig `field:"optional" json:"workerConfig" yaml:"workerConfig"`
}

