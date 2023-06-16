package workerscript


type WorkerScriptServiceBinding struct {
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#service WorkerScript#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// The name of the Worker environment to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#environment WorkerScript#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
}

