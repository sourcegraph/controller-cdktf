package googlecontainerazurecluster


type GoogleContainerAzureClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#admin_users GoogleContainerAzureCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
	// admin_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#admin_groups GoogleContainerAzureCluster#admin_groups}
	AdminGroups interface{} `field:"optional" json:"adminGroups" yaml:"adminGroups"`
}

