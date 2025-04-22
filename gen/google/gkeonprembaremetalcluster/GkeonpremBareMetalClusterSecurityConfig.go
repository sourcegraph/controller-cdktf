package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterSecurityConfig struct {
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gkeonprem_bare_metal_cluster#authorization GkeonpremBareMetalCluster#authorization}
	Authorization *GkeonpremBareMetalClusterSecurityConfigAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
}

