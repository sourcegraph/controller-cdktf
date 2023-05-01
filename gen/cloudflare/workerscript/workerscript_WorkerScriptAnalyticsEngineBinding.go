package workerscript


type WorkerScriptAnalyticsEngineBinding struct {
	// The name of the Analytics Engine dataset to write to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#dataset WorkerScript#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

