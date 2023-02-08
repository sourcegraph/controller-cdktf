package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions struct {
	// The scope of automatic upgrades to restrict in the exclusion window.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#scope GoogleContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

