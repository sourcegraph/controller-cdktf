package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket struct {
	// Name of the Cloud Storage bucket to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#bucket GoogleHypercomputeclusterCluster#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// autoclass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#autoclass GoogleHypercomputeclusterCluster#autoclass}
	Autoclass *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass `field:"optional" json:"autoclass" yaml:"autoclass"`
	// hierarchical_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#hierarchical_namespace GoogleHypercomputeclusterCluster#hierarchical_namespace}
	HierarchicalNamespace *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace `field:"optional" json:"hierarchicalNamespace" yaml:"hierarchicalNamespace"`
	// If set, uses the provided storage class as the bucket's default storage class. Possible values: STANDARD NEARLINE COLDLINE ARCHIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#storage_class GoogleHypercomputeclusterCluster#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

