package containerazurecluster

type ContainerAzureClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_azure_cluster#admin_users ContainerAzureCluster#admin_users}
	AdminUsers any `field:"required" json:"adminUsers" yaml:"adminUsers"`
	// admin_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_azure_cluster#admin_groups ContainerAzureCluster#admin_groups}
	AdminGroups any `field:"optional" json:"adminGroups" yaml:"adminGroups"`
}
