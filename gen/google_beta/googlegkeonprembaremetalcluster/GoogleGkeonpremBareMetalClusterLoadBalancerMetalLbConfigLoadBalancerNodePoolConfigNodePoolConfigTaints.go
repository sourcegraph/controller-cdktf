package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaints struct {
	// Specifies the nodes operating system (default: LINUX). Possible values: ["EFFECT_UNSPECIFIED", "PREFER_NO_SCHEDULE", "NO_EXECUTE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#effect GoogleGkeonpremBareMetalCluster#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Key associated with the effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#key GoogleGkeonpremBareMetalCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value associated with the effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#value GoogleGkeonpremBareMetalCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

