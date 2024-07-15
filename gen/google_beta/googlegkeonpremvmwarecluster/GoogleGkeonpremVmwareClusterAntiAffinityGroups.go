package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterAntiAffinityGroups struct {
	// Spread nodes across at least three physical hosts (requires at least three hosts). Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#aag_config_disabled GoogleGkeonpremVmwareCluster#aag_config_disabled}
	AagConfigDisabled interface{} `field:"required" json:"aagConfigDisabled" yaml:"aagConfigDisabled"`
}

