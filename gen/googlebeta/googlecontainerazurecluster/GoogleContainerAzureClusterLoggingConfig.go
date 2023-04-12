package googlecontainerazurecluster


type GoogleContainerAzureClusterLoggingConfig struct {
	// component_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_cluster#component_config GoogleContainerAzureCluster#component_config}
	ComponentConfig *GoogleContainerAzureClusterLoggingConfigComponentConfig `field:"optional" json:"componentConfig" yaml:"componentConfig"`
}

