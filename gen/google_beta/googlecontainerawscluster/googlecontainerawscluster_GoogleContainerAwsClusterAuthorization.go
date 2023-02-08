package googlecontainerawscluster


type GoogleContainerAwsClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_cluster#admin_users GoogleContainerAwsCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

