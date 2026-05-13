package storagebucketobject


type StorageBucketObjectContexts struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_bucket_object#custom StorageBucketObject#custom}
	Custom interface{} `field:"required" json:"custom" yaml:"custom"`
}

