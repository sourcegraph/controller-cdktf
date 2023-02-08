package googlecontainerawscluster


type GoogleContainerAwsClusterLoggingConfigComponentConfig struct {
	// Components of the logging configuration to be enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_cluster#enable_components GoogleContainerAwsCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
}

