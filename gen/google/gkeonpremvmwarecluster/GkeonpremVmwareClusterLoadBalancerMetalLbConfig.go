package gkeonpremvmwarecluster


type GkeonpremVmwareClusterLoadBalancerMetalLbConfig struct {
	// address_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gkeonprem_vmware_cluster#address_pools GkeonpremVmwareCluster#address_pools}
	AddressPools interface{} `field:"required" json:"addressPools" yaml:"addressPools"`
}

