package googlecontainerawscluster


type GoogleContainerAwsClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_cluster#admin_users GoogleContainerAwsCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
	// admin_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_cluster#admin_groups GoogleContainerAwsCluster#admin_groups}
	AdminGroups interface{} `field:"optional" json:"adminGroups" yaml:"adminGroups"`
}

