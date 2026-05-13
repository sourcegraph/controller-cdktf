package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigNewLustre struct {
	// Storage capacity of the instance in gibibytes (GiB). Allowed values are between 18000 and 7632000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#capacity_gb HypercomputeclusterCluster#capacity_gb}
	CapacityGb *string `field:"required" json:"capacityGb" yaml:"capacityGb"`
	// Filesystem name for this instance.
	//
	// This name is used by client-side tools,
	// including when mounting the instance. Must be 8 characters or less and can
	// only contain letters and numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#filesystem HypercomputeclusterCluster#filesystem}
	Filesystem *string `field:"required" json:"filesystem" yaml:"filesystem"`
	// Name of the Managed Lustre instance to create, in the format 'projects/{project}/locations/{location}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#lustre HypercomputeclusterCluster#lustre}
	Lustre *string `field:"required" json:"lustre" yaml:"lustre"`
	// Description of the Managed Lustre instance. Maximum of 2048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#description HypercomputeclusterCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

