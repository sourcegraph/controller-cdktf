package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestoreFileShares struct {
	// Size of the filestore in GB. Must be between 1024 and 102400, and must meet scalability requirements described at https://cloud.google.com/filestore/docs/service-tiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#capacity_gb GoogleHypercomputeclusterCluster#capacity_gb}
	CapacityGb *string `field:"required" json:"capacityGb" yaml:"capacityGb"`
	// Filestore share location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#file_share GoogleHypercomputeclusterCluster#file_share}
	FileShare *string `field:"required" json:"fileShare" yaml:"fileShare"`
}

