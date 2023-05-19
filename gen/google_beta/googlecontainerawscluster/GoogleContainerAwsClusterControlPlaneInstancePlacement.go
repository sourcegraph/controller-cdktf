package googlecontainerawscluster


type GoogleContainerAwsClusterControlPlaneInstancePlacement struct {
	// The tenancy for the instance. Possible values: TENANCY_UNSPECIFIED, DEFAULT, DEDICATED, HOST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_aws_cluster#tenancy GoogleContainerAwsCluster#tenancy}
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

