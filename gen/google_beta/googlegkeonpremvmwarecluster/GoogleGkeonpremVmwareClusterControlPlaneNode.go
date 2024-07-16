package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterControlPlaneNode struct {
	// auto_resize_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#auto_resize_config GoogleGkeonpremVmwareCluster#auto_resize_config}
	AutoResizeConfig *GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig `field:"optional" json:"autoResizeConfig" yaml:"autoResizeConfig"`
	// The number of CPUs for each admin cluster node that serve as control planes for this VMware User Cluster.
	//
	// (default: 4 CPUs)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#cpus GoogleGkeonpremVmwareCluster#cpus}
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// The megabytes of memory for each admin cluster node that serves as a control plane for this VMware User Cluster (default: 8192 MB memory).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#memory GoogleGkeonpremVmwareCluster#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The number of control plane nodes for this VMware User Cluster. (default: 1 replica).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#replicas GoogleGkeonpremVmwareCluster#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

