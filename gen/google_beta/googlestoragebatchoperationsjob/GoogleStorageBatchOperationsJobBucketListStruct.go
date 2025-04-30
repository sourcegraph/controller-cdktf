package googlestoragebatchoperationsjob


type GoogleStorageBatchOperationsJobBucketListStruct struct {
	// buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#buckets GoogleStorageBatchOperationsJob#buckets}
	Buckets *GoogleStorageBatchOperationsJobBucketListBuckets `field:"required" json:"buckets" yaml:"buckets"`
}

