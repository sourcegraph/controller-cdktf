package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigNewBucket struct {
	// Name of the Cloud Storage bucket to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#bucket HypercomputeclusterCluster#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// autoclass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#autoclass HypercomputeclusterCluster#autoclass}
	Autoclass *HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass `field:"optional" json:"autoclass" yaml:"autoclass"`
	// hierarchical_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#hierarchical_namespace HypercomputeclusterCluster#hierarchical_namespace}
	HierarchicalNamespace *HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace `field:"optional" json:"hierarchicalNamespace" yaml:"hierarchicalNamespace"`
	// If set, uses the provided storage class as the bucket's default storage class. Possible values: STANDARD NEARLINE COLDLINE ARCHIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#storage_class HypercomputeclusterCluster#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

