package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterStorage struct {
	// Whether or not to deploy vSphere CSI components in the VMware User Cluster. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#vsphere_csi_disabled GoogleGkeonpremVmwareCluster#vsphere_csi_disabled}
	VsphereCsiDisabled interface{} `field:"required" json:"vsphereCsiDisabled" yaml:"vsphereCsiDisabled"`
}

