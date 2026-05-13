package googlecolabnotebookexecution


type GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec struct {
	// Enable public internet access for the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#enable_internet_access GoogleColabNotebookExecution#enable_internet_access}
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// The name of the VPC that this runtime is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#network GoogleColabNotebookExecution#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnetwork that this runtime is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#subnetwork GoogleColabNotebookExecution#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

