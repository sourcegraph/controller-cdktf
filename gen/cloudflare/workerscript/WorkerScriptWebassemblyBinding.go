package workerscript


type WorkerScriptWebassemblyBinding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#module WorkerScript#module}.
	Module *string `field:"required" json:"module" yaml:"module"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#name WorkerScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

