package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigExistingLustre struct {
	// Name of the Managed Lustre instance to import, in the format 'projects/{project}/locations/{location}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#lustre HypercomputeclusterCluster#lustre}
	Lustre *string `field:"required" json:"lustre" yaml:"lustre"`
}

