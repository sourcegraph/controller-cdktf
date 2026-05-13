package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfig struct {
	// existing_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#existing_bucket GoogleHypercomputeclusterCluster#existing_bucket}
	ExistingBucket *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucket `field:"optional" json:"existingBucket" yaml:"existingBucket"`
	// existing_filestore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#existing_filestore GoogleHypercomputeclusterCluster#existing_filestore}
	ExistingFilestore *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore `field:"optional" json:"existingFilestore" yaml:"existingFilestore"`
	// existing_lustre block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#existing_lustre GoogleHypercomputeclusterCluster#existing_lustre}
	ExistingLustre *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre `field:"optional" json:"existingLustre" yaml:"existingLustre"`
	// new_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_bucket GoogleHypercomputeclusterCluster#new_bucket}
	NewBucket *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket `field:"optional" json:"newBucket" yaml:"newBucket"`
	// new_filestore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_filestore GoogleHypercomputeclusterCluster#new_filestore}
	NewFilestore *GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestore `field:"optional" json:"newFilestore" yaml:"newFilestore"`
	// new_lustre block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#new_lustre GoogleHypercomputeclusterCluster#new_lustre}
	NewLustre *GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustre `field:"optional" json:"newLustre" yaml:"newLustre"`
}

