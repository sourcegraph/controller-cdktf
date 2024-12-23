package googlecontainercluster


type GoogleContainerClusterReleaseChannel struct {
	// The selected release channel.
	//
	// Accepted values are:
	// * UNSPECIFIED: Not set.
	// * RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
	// * REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
	// * STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#channel GoogleContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

