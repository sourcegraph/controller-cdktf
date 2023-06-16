package googlecontainernodepool


type GoogleContainerNodePoolNodeConfig struct {
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#advanced_machine_features GoogleContainerNodePool#advanced_machine_features}
	AdvancedMachineFeatures *GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#boot_disk_kms_key GoogleContainerNodePool#boot_disk_kms_key}
	BootDiskKmsKey *string `field:"optional" json:"bootDiskKmsKey" yaml:"bootDiskKmsKey"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#disk_size_gb GoogleContainerNodePool#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#disk_type GoogleContainerNodePool#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// ephemeral_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#ephemeral_storage_config GoogleContainerNodePool#ephemeral_storage_config}
	EphemeralStorageConfig *GoogleContainerNodePoolNodeConfigEphemeralStorageConfig `field:"optional" json:"ephemeralStorageConfig" yaml:"ephemeralStorageConfig"`
	// ephemeral_storage_local_ssd_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#ephemeral_storage_local_ssd_config GoogleContainerNodePool#ephemeral_storage_local_ssd_config}
	EphemeralStorageLocalSsdConfig *GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig `field:"optional" json:"ephemeralStorageLocalSsdConfig" yaml:"ephemeralStorageLocalSsdConfig"`
	// gcfs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#gcfs_config GoogleContainerNodePool#gcfs_config}
	GcfsConfig *GoogleContainerNodePoolNodeConfigGcfsConfig `field:"optional" json:"gcfsConfig" yaml:"gcfsConfig"`
	// List of the type and count of accelerator cards attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#guest_accelerator GoogleContainerNodePool#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// gvnic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#gvnic GoogleContainerNodePool#gvnic}
	Gvnic *GoogleContainerNodePoolNodeConfigGvnic `field:"optional" json:"gvnic" yaml:"gvnic"`
	// The image type to use for this node.
	//
	// Note that for a given image type, the latest version of it will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#image_type GoogleContainerNodePool#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#kubelet_config GoogleContainerNodePool#kubelet_config}
	KubeletConfig *GoogleContainerNodePoolNodeConfigKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s) that Kubernetes may apply to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#labels GoogleContainerNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linux_node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#linux_node_config GoogleContainerNodePool#linux_node_config}
	LinuxNodeConfig *GoogleContainerNodePoolNodeConfigLinuxNodeConfig `field:"optional" json:"linuxNodeConfig" yaml:"linuxNodeConfig"`
	// local_nvme_ssd_block_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#local_nvme_ssd_block_config GoogleContainerNodePool#local_nvme_ssd_block_config}
	LocalNvmeSsdBlockConfig *GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig `field:"optional" json:"localNvmeSsdBlockConfig" yaml:"localNvmeSsdBlockConfig"`
	// The number of local SSD disks to be attached to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#local_ssd_count GoogleContainerNodePool#local_ssd_count}
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// Type of logging agent that is used as the default value for node pools in the cluster.
	//
	// Valid values include DEFAULT and MAX_THROUGHPUT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#logging_variant GoogleContainerNodePool#logging_variant}
	LoggingVariant *string `field:"optional" json:"loggingVariant" yaml:"loggingVariant"`
	// The name of a Google Compute Engine machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#machine_type GoogleContainerNodePool#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The metadata key/value pairs assigned to instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#metadata GoogleContainerNodePool#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Minimum CPU platform to be used by this instance.
	//
	// The instance may be scheduled on the specified or newer CPU platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#min_cpu_platform GoogleContainerNodePool#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Setting this field will assign instances of this pool to run on the specified node group.
	//
	// This is useful for running workloads on sole tenant nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#node_group GoogleContainerNodePool#node_group}
	NodeGroup *string `field:"optional" json:"nodeGroup" yaml:"nodeGroup"`
	// The set of Google API scopes to be made available on all of the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#oauth_scopes GoogleContainerNodePool#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// Whether the nodes are created as preemptible VM instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#preemptible GoogleContainerNodePool#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#reservation_affinity GoogleContainerNodePool#reservation_affinity}
	ReservationAffinity *GoogleContainerNodePoolNodeConfigReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The GCE resource labels (a map of key/value pairs) to be applied to the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#resource_labels GoogleContainerNodePool#resource_labels}
	ResourceLabels *map[string]*string `field:"optional" json:"resourceLabels" yaml:"resourceLabels"`
	// sandbox_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#sandbox_config GoogleContainerNodePool#sandbox_config}
	SandboxConfig *GoogleContainerNodePoolNodeConfigSandboxConfig `field:"optional" json:"sandboxConfig" yaml:"sandboxConfig"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#service_account GoogleContainerNodePool#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#shielded_instance_config GoogleContainerNodePool#shielded_instance_config}
	ShieldedInstanceConfig *GoogleContainerNodePoolNodeConfigShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Whether the nodes are created as spot VM instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#spot GoogleContainerNodePool#spot}
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
	// The list of instance tags applied to all nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#tags GoogleContainerNodePool#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// List of Kubernetes taints to be applied to each node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#taint GoogleContainerNodePool#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// workload_metadata_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_node_pool#workload_metadata_config GoogleContainerNodePool#workload_metadata_config}
	WorkloadMetadataConfig *GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig `field:"optional" json:"workloadMetadataConfig" yaml:"workloadMetadataConfig"`
}

