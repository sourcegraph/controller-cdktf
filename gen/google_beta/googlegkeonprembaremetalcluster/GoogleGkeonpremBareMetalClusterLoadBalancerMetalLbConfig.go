package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#address_pools GoogleGkeonpremBareMetalCluster#address_pools}
	AddressPools interface{} `field:"required" json:"addressPools" yaml:"addressPools"`
	// load_balancer_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#load_balancer_node_pool_config GoogleGkeonpremBareMetalCluster#load_balancer_node_pool_config}
	LoadBalancerNodePoolConfig *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfig `field:"optional" json:"loadBalancerNodePoolConfig" yaml:"loadBalancerNodePoolConfig"`
}

