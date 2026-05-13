package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#registry_burst GoogleGkeonpremBareMetalAdminCluster#registry_burst}.
	RegistryBurst *float64 `field:"optional" json:"registryBurst" yaml:"registryBurst"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#registry_pull_qps GoogleGkeonpremBareMetalAdminCluster#registry_pull_qps}.
	RegistryPullQps *float64 `field:"optional" json:"registryPullQps" yaml:"registryPullQps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#serialize_image_pulls_disabled GoogleGkeonpremBareMetalAdminCluster#serialize_image_pulls_disabled}.
	SerializeImagePullsDisabled interface{} `field:"optional" json:"serializeImagePullsDisabled" yaml:"serializeImagePullsDisabled"`
}

