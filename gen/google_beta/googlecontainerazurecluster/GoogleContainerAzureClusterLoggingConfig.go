package googlecontainerazurecluster


type GoogleContainerAzureClusterLoggingConfig struct {
	// component_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#component_config GoogleContainerAzureCluster#component_config}
	ComponentConfig *GoogleContainerAzureClusterLoggingConfigComponentConfig `field:"optional" json:"componentConfig" yaml:"componentConfig"`
}

