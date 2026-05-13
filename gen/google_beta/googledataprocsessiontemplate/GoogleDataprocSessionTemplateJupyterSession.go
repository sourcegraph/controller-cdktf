package googledataprocsessiontemplate


type GoogleDataprocSessionTemplateJupyterSession struct {
	// Display name, shown in the Jupyter kernelspec card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#display_name GoogleDataprocSessionTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Kernel to be used with Jupyter interactive session. Possible values: ["PYTHON", "SCALA"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#kernel GoogleDataprocSessionTemplate#kernel}
	Kernel *string `field:"optional" json:"kernel" yaml:"kernel"`
}

