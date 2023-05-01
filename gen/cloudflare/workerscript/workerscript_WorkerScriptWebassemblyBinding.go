package workerscript


type WorkerScriptWebassemblyBinding struct {
	// The base64 encoded wasm module you want to store.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#module WorkerScript#module}
	Module *string `field:"required" json:"module" yaml:"module"`
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

