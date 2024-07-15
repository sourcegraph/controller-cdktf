package googlecontainerazurecluster


type GoogleContainerAzureClusterAuthorizationAdminGroups struct {
	// The name of the group, e.g. `my-group@domain.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_cluster#group GoogleContainerAzureCluster#group}
	Group *string `field:"required" json:"group" yaml:"group"`
}

