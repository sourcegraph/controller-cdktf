package googlecontainercluster


type GoogleContainerClusterGkeAutoUpgradeConfig struct {
	// The selected auto-upgrade patch type.
	//
	// Accepted values are:
	// * ACCELERATED: Upgrades to the latest available patch version in a given minor and release channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#patch_mode GoogleContainerCluster#patch_mode}
	PatchMode *string `field:"required" json:"patchMode" yaml:"patchMode"`
}

