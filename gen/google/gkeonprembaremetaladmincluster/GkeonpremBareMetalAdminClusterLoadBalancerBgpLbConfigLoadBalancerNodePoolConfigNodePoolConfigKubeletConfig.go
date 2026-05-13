package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigKubeletConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#registry_burst GkeonpremBareMetalAdminCluster#registry_burst}.
	RegistryBurst *float64 `field:"optional" json:"registryBurst" yaml:"registryBurst"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#registry_pull_qps GkeonpremBareMetalAdminCluster#registry_pull_qps}.
	RegistryPullQps *float64 `field:"optional" json:"registryPullQps" yaml:"registryPullQps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#serialize_image_pulls_disabled GkeonpremBareMetalAdminCluster#serialize_image_pulls_disabled}.
	SerializeImagePullsDisabled interface{} `field:"optional" json:"serializeImagePullsDisabled" yaml:"serializeImagePullsDisabled"`
}

