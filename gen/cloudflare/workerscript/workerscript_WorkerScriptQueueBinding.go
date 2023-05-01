package workerscript


type WorkerScriptQueueBinding struct {
	// The name of the global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#binding WorkerScript#binding}
	Binding *string `field:"required" json:"binding" yaml:"binding"`
	// Name of the queue you want to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#queue WorkerScript#queue}
	Queue *string `field:"required" json:"queue" yaml:"queue"`
}

