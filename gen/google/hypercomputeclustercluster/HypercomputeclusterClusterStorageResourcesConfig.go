package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfig struct {
	// existing_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#existing_bucket HypercomputeclusterCluster#existing_bucket}
	ExistingBucket *HypercomputeclusterClusterStorageResourcesConfigExistingBucket `field:"optional" json:"existingBucket" yaml:"existingBucket"`
	// existing_filestore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#existing_filestore HypercomputeclusterCluster#existing_filestore}
	ExistingFilestore *HypercomputeclusterClusterStorageResourcesConfigExistingFilestore `field:"optional" json:"existingFilestore" yaml:"existingFilestore"`
	// existing_lustre block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#existing_lustre HypercomputeclusterCluster#existing_lustre}
	ExistingLustre *HypercomputeclusterClusterStorageResourcesConfigExistingLustre `field:"optional" json:"existingLustre" yaml:"existingLustre"`
	// new_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_bucket HypercomputeclusterCluster#new_bucket}
	NewBucket *HypercomputeclusterClusterStorageResourcesConfigNewBucket `field:"optional" json:"newBucket" yaml:"newBucket"`
	// new_filestore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_filestore HypercomputeclusterCluster#new_filestore}
	NewFilestore *HypercomputeclusterClusterStorageResourcesConfigNewFilestore `field:"optional" json:"newFilestore" yaml:"newFilestore"`
	// new_lustre block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#new_lustre HypercomputeclusterCluster#new_lustre}
	NewLustre *HypercomputeclusterClusterStorageResourcesConfigNewLustre `field:"optional" json:"newLustre" yaml:"newLustre"`
}

