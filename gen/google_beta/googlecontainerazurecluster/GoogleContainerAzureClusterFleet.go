package googlecontainerazurecluster


type GoogleContainerAzureClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#project GoogleContainerAzureCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

