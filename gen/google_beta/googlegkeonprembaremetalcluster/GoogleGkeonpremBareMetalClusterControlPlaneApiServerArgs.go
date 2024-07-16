package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterControlPlaneApiServerArgs struct {
	// The argument name as it appears on the API Server command line please make sure to remove the leading dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#argument GoogleGkeonpremBareMetalCluster#argument}
	Argument *string `field:"required" json:"argument" yaml:"argument"`
	// The value of the arg as it will be passed to the API Server command line.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#value GoogleGkeonpremBareMetalCluster#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

