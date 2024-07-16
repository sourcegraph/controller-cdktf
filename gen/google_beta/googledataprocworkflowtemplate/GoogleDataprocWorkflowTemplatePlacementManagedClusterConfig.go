package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig struct {
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#autoscaling_config GoogleDataprocWorkflowTemplate#autoscaling_config}
	AutoscalingConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#encryption_config GoogleDataprocWorkflowTemplate#encryption_config}
	EncryptionConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#endpoint_config GoogleDataprocWorkflowTemplate#endpoint_config}
	EndpointConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig `field:"optional" json:"endpointConfig" yaml:"endpointConfig"`
	// gce_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#gce_cluster_config GoogleDataprocWorkflowTemplate#gce_cluster_config}
	GceClusterConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig `field:"optional" json:"gceClusterConfig" yaml:"gceClusterConfig"`
	// gke_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#gke_cluster_config GoogleDataprocWorkflowTemplate#gke_cluster_config}
	GkeClusterConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig `field:"optional" json:"gkeClusterConfig" yaml:"gkeClusterConfig"`
	// initialization_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#initialization_actions GoogleDataprocWorkflowTemplate#initialization_actions}
	InitializationActions interface{} `field:"optional" json:"initializationActions" yaml:"initializationActions"`
	// lifecycle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#lifecycle_config GoogleDataprocWorkflowTemplate#lifecycle_config}
	LifecycleConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig `field:"optional" json:"lifecycleConfig" yaml:"lifecycleConfig"`
	// master_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#master_config GoogleDataprocWorkflowTemplate#master_config}
	MasterConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig `field:"optional" json:"masterConfig" yaml:"masterConfig"`
	// metastore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#metastore_config GoogleDataprocWorkflowTemplate#metastore_config}
	MetastoreConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig `field:"optional" json:"metastoreConfig" yaml:"metastoreConfig"`
	// secondary_worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#secondary_worker_config GoogleDataprocWorkflowTemplate#secondary_worker_config}
	SecondaryWorkerConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig `field:"optional" json:"secondaryWorkerConfig" yaml:"secondaryWorkerConfig"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#security_config GoogleDataprocWorkflowTemplate#security_config}
	SecurityConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#software_config GoogleDataprocWorkflowTemplate#software_config}
	SoftwareConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// Optional.
	//
	// A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging bucket](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#staging_bucket GoogleDataprocWorkflowTemplate#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
	// Optional.
	//
	// A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket. **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#temp_bucket GoogleDataprocWorkflowTemplate#temp_bucket}
	TempBucket *string `field:"optional" json:"tempBucket" yaml:"tempBucket"`
	// worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#worker_config GoogleDataprocWorkflowTemplate#worker_config}
	WorkerConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig `field:"optional" json:"workerConfig" yaml:"workerConfig"`
}

