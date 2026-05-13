package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigExistingFilestore struct {
	// Name of the Filestore instance to import, in the format 'projects/{project}/locations/{location}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#filestore HypercomputeclusterCluster#filestore}
	Filestore *string `field:"required" json:"filestore" yaml:"filestore"`
}

