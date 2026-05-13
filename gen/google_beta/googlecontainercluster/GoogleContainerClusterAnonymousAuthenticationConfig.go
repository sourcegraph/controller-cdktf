package googlecontainercluster


type GoogleContainerClusterAnonymousAuthenticationConfig struct {
	// Setting this to LIMITED will restrict authentication of anonymous users to health check endpoints only.
	//
	// Accepted values are:
	// * ENABLED: Authentication of anonymous users is enabled for all endpoints.
	// * LIMITED: Anonymous access is only allowed for health check endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#mode GoogleContainerCluster#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

