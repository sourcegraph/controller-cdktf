package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore struct {
	// Name of the Filestore instance to import, in the format 'projects/{project}/locations/{location}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#filestore GoogleHypercomputeclusterCluster#filestore}
	Filestore *string `field:"required" json:"filestore" yaml:"filestore"`
}

