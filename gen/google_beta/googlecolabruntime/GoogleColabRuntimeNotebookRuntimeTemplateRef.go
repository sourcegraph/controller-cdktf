package googlecolabruntime


type GoogleColabRuntimeNotebookRuntimeTemplateRef struct {
	// The resource name of the NotebookRuntimeTemplate based on which a NotebookRuntime will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime#notebook_runtime_template GoogleColabRuntime#notebook_runtime_template}
	NotebookRuntimeTemplate *string `field:"required" json:"notebookRuntimeTemplate" yaml:"notebookRuntimeTemplate"`
}

