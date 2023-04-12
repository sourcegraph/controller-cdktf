package googlecontainercluster


type GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaults struct {
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#boot_disk_kms_key GoogleContainerCluster#boot_disk_kms_key}
	BootDiskKmsKey *string `field:"optional" json:"bootDiskKmsKey" yaml:"bootDiskKmsKey"`
	// The default image type used by NAP once a new node pool is being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#image_type GoogleContainerCluster#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// Minimum CPU platform to be used by this instance.
	//
	// The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#min_cpu_platform GoogleContainerCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Scopes that are used by NAP when creating node pools.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#oauth_scopes GoogleContainerCluster#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#service_account GoogleContainerCluster#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

