package googlegkeonpremvmwarenodepool


type GoogleGkeonpremVmwareNodePoolConfigVsphereConfig struct {
	// The name of the vCenter datastore. Inherited from the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_node_pool#datastore GoogleGkeonpremVmwareNodePool#datastore}
	Datastore *string `field:"optional" json:"datastore" yaml:"datastore"`
	// Vsphere host groups to apply to all VMs in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_node_pool#host_groups GoogleGkeonpremVmwareNodePool#host_groups}
	HostGroups *[]*string `field:"optional" json:"hostGroups" yaml:"hostGroups"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_node_pool#tags GoogleGkeonpremVmwareNodePool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

