package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#address_pools GkeonpremBareMetalAdminCluster#address_pools}
	AddressPools interface{} `field:"optional" json:"addressPools" yaml:"addressPools"`
	// BGP autonomous system number (ASN) of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#asn GkeonpremBareMetalAdminCluster#asn}
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// bgp_peer_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#bgp_peer_configs GkeonpremBareMetalAdminCluster#bgp_peer_configs}
	BgpPeerConfigs interface{} `field:"optional" json:"bgpPeerConfigs" yaml:"bgpPeerConfigs"`
	// load_balancer_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#load_balancer_node_pool_config GkeonpremBareMetalAdminCluster#load_balancer_node_pool_config}
	LoadBalancerNodePoolConfig *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig `field:"optional" json:"loadBalancerNodePoolConfig" yaml:"loadBalancerNodePoolConfig"`
}

