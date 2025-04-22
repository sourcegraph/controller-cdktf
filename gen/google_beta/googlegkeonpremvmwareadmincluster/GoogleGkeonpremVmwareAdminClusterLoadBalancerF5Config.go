package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config struct {
	// The load balancer's IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#address GoogleGkeonpremVmwareAdminCluster#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// he preexisting partition to be used by the load balancer.
	//
	// T
	// his partition is usually created for the admin cluster for example:
	// 'my-f5-admin-partition'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#partition GoogleGkeonpremVmwareAdminCluster#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The pool name. Only necessary, if using SNAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#snat_pool GoogleGkeonpremVmwareAdminCluster#snat_pool}
	SnatPool *string `field:"optional" json:"snatPool" yaml:"snatPool"`
}

