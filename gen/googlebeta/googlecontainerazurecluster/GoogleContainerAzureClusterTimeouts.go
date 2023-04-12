package googlecontainerazurecluster


type GoogleContainerAzureClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_cluster#create GoogleContainerAzureCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_cluster#delete GoogleContainerAzureCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_cluster#update GoogleContainerAzureCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

