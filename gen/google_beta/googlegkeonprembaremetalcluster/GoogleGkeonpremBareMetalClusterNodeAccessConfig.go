package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterNodeAccessConfig struct {
	// LoginUser is the user name used to access node machines. It defaults to "root" if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#login_user GoogleGkeonpremBareMetalCluster#login_user}
	LoginUser *string `field:"optional" json:"loginUser" yaml:"loginUser"`
}

