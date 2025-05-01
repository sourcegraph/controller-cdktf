package googlestoragebatchoperationsjob


type GoogleStorageBatchOperationsJobBucketListBuckets struct {
	// Bucket name for the objects to be transformed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#bucket GoogleStorageBatchOperationsJob#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// manifest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#manifest GoogleStorageBatchOperationsJob#manifest}
	Manifest *GoogleStorageBatchOperationsJobBucketListBucketsManifest `field:"optional" json:"manifest" yaml:"manifest"`
	// prefix_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#prefix_list GoogleStorageBatchOperationsJob#prefix_list}
	PrefixList *GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStruct `field:"optional" json:"prefixList" yaml:"prefixList"`
}

