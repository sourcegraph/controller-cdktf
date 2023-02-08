package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfig struct {
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#boot_disk_kms_key GoogleContainerCluster#boot_disk_kms_key}
	BootDiskKmsKey *string `field:"optional" json:"bootDiskKmsKey" yaml:"bootDiskKmsKey"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#disk_size_gb GoogleContainerCluster#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#disk_type GoogleContainerCluster#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// ephemeral_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#ephemeral_storage_config GoogleContainerCluster#ephemeral_storage_config}
	EphemeralStorageConfig *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfig `field:"optional" json:"ephemeralStorageConfig" yaml:"ephemeralStorageConfig"`
	// gcfs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#gcfs_config GoogleContainerCluster#gcfs_config}
	GcfsConfig *GoogleContainerClusterNodePoolNodeConfigGcfsConfig `field:"optional" json:"gcfsConfig" yaml:"gcfsConfig"`
	// List of the type and count of accelerator cards attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#guest_accelerator GoogleContainerCluster#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// gvnic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#gvnic GoogleContainerCluster#gvnic}
	Gvnic *GoogleContainerClusterNodePoolNodeConfigGvnic `field:"optional" json:"gvnic" yaml:"gvnic"`
	// The image type to use for this node.
	//
	// Note that for a given image type, the latest version of it will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#image_type GoogleContainerCluster#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#kubelet_config GoogleContainerCluster#kubelet_config}
	KubeletConfig *GoogleContainerClusterNodePoolNodeConfigKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s) that Kubernetes may apply to the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#labels GoogleContainerCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linux_node_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#linux_node_config GoogleContainerCluster#linux_node_config}
	LinuxNodeConfig *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig `field:"optional" json:"linuxNodeConfig" yaml:"linuxNodeConfig"`
	// The number of local SSD disks to be attached to the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#local_ssd_count GoogleContainerCluster#local_ssd_count}
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// The name of a Google Compute Engine machine type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#machine_type GoogleContainerCluster#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The metadata key/value pairs assigned to instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#metadata GoogleContainerCluster#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Minimum CPU platform to be used by this instance.
	//
	// The instance may be scheduled on the specified or newer CPU platform.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#min_cpu_platform GoogleContainerCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Setting this field will assign instances of this pool to run on the specified node group.
	//
	// This is useful for running workloads on sole tenant nodes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#node_group GoogleContainerCluster#node_group}
	NodeGroup *string `field:"optional" json:"nodeGroup" yaml:"nodeGroup"`
	// The set of Google API scopes to be made available on all of the node VMs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#oauth_scopes GoogleContainerCluster#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// Whether the nodes are created as preemptible VM instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#preemptible GoogleContainerCluster#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#reservation_affinity GoogleContainerCluster#reservation_affinity}
	ReservationAffinity *GoogleContainerClusterNodePoolNodeConfigReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// sandbox_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#sandbox_config GoogleContainerCluster#sandbox_config}
	SandboxConfig *GoogleContainerClusterNodePoolNodeConfigSandboxConfig `field:"optional" json:"sandboxConfig" yaml:"sandboxConfig"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#service_account GoogleContainerCluster#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#shielded_instance_config GoogleContainerCluster#shielded_instance_config}
	ShieldedInstanceConfig *GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Whether the nodes are created as spot VM instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#spot GoogleContainerCluster#spot}
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
	// The list of instance tags applied to all nodes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#tags GoogleContainerCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// List of Kubernetes taints to be applied to each node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#taint GoogleContainerCluster#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// workload_metadata_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#workload_metadata_config GoogleContainerCluster#workload_metadata_config}
	WorkloadMetadataConfig *GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfig `field:"optional" json:"workloadMetadataConfig" yaml:"workloadMetadataConfig"`
}

