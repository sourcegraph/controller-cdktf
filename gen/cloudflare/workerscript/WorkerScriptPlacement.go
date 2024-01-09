package workerscript


type WorkerScriptPlacement struct {
	// The placement mode for the Worker. Available values: `smart`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/worker_script#mode WorkerScript#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

