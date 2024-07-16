package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterSecurityConfigAuthorizationAdminUsers struct {
	// The name of the user, e.g. 'my-gcp-id@gmail.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#username GoogleGkeonpremBareMetalCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

