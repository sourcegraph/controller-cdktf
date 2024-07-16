package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterUpgradePolicy struct {
	// Specifies which upgrade policy to use. Possible values: ["SERIAL", "CONCURRENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#policy GoogleGkeonpremBareMetalCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

