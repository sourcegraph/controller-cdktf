package hypercomputeclustercluster


type HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk struct {
	// Size of the disk in gigabytes. Must be at least 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#size_gb HypercomputeclusterCluster#size_gb}
	SizeGb *string `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// [Persistent disk type](https://cloud.google.com/compute/docs/disks#disk-types), in the format 'projects/{project}/zones/{zone}/diskTypes/{disk_type}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#type HypercomputeclusterCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

