package googlecontainerawscluster


type GoogleContainerAwsClusterLoggingConfigComponentConfig struct {
	// Components of the logging configuration to be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_container_aws_cluster#enable_components GoogleContainerAwsCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
}

