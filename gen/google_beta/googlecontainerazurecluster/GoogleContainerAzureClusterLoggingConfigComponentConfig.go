package googlecontainerazurecluster


type GoogleContainerAzureClusterLoggingConfigComponentConfig struct {
	// Components of the logging configuration to be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#enable_components GoogleContainerAzureCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
}

