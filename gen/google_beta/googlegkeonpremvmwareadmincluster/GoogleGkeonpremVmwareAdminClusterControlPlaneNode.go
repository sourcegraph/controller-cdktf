package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterControlPlaneNode struct {
	// The number of vCPUs for the control-plane node of the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#cpus GoogleGkeonpremVmwareAdminCluster#cpus}
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// The number of mebibytes of memory for the control-plane node of the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#memory GoogleGkeonpremVmwareAdminCluster#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The number of control plane nodes for this VMware admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#replicas GoogleGkeonpremVmwareAdminCluster#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

