package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre struct {
	// Name of the Managed Lustre instance to import, in the format 'projects/{project}/locations/{location}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#lustre GoogleHypercomputeclusterCluster#lustre}
	Lustre *string `field:"required" json:"lustre" yaml:"lustre"`
}

