package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#address_pools GoogleGkeonpremBareMetalCluster#address_pools}
	AddressPools interface{} `field:"required" json:"addressPools" yaml:"addressPools"`
	// BGP autonomous system number (ASN) of the cluster. This field can be updated after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#asn GoogleGkeonpremBareMetalCluster#asn}
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// bgp_peer_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#bgp_peer_configs GoogleGkeonpremBareMetalCluster#bgp_peer_configs}
	BgpPeerConfigs interface{} `field:"required" json:"bgpPeerConfigs" yaml:"bgpPeerConfigs"`
	// load_balancer_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#load_balancer_node_pool_config GoogleGkeonpremBareMetalCluster#load_balancer_node_pool_config}
	LoadBalancerNodePoolConfig *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig `field:"optional" json:"loadBalancerNodePoolConfig" yaml:"loadBalancerNodePoolConfig"`
}

