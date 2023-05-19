package workerscript


type WorkerScriptR2BucketBinding struct {
	// The name of the Bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#bucket_name WorkerScript#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

