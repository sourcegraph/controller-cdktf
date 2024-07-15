package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterSecurityConfig struct {
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#authorization GoogleGkeonpremBareMetalCluster#authorization}
	Authorization *GoogleGkeonpremBareMetalClusterSecurityConfigAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
}

