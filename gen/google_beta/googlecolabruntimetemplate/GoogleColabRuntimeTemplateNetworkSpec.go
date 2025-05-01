package googlecolabruntimetemplate


type GoogleColabRuntimeTemplateNetworkSpec struct {
	// Enable public internet access for the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template#enable_internet_access GoogleColabRuntimeTemplate#enable_internet_access}
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// The name of the VPC that this runtime is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template#network GoogleColabRuntimeTemplate#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnetwork that this runtime is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template#subnetwork GoogleColabRuntimeTemplate#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

