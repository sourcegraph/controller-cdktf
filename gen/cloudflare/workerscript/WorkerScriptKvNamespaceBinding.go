package workerscript


type WorkerScriptKvNamespaceBinding struct {
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the KV namespace you want to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#namespace_id WorkerScript#namespace_id}
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
}

