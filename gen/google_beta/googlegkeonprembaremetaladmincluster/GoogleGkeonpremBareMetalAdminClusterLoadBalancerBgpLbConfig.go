package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#address_pools GoogleGkeonpremBareMetalAdminCluster#address_pools}
	AddressPools interface{} `field:"optional" json:"addressPools" yaml:"addressPools"`
	// BGP autonomous system number (ASN) of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#asn GoogleGkeonpremBareMetalAdminCluster#asn}
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// bgp_peer_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#bgp_peer_configs GoogleGkeonpremBareMetalAdminCluster#bgp_peer_configs}
	BgpPeerConfigs interface{} `field:"optional" json:"bgpPeerConfigs" yaml:"bgpPeerConfigs"`
	// load_balancer_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#load_balancer_node_pool_config GoogleGkeonpremBareMetalAdminCluster#load_balancer_node_pool_config}
	LoadBalancerNodePoolConfig *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig `field:"optional" json:"loadBalancerNodePoolConfig" yaml:"loadBalancerNodePoolConfig"`
}

