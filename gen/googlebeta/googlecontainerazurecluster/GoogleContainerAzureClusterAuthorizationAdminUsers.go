package googlecontainerazurecluster


type GoogleContainerAzureClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_cluster#username GoogleContainerAzureCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

