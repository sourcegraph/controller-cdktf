package googlecontainerawscluster


type GoogleContainerAwsClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_cluster#project GoogleContainerAwsCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

