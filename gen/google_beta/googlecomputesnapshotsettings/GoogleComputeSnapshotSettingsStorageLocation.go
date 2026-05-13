package googlecomputesnapshotsettings


type GoogleComputeSnapshotSettingsStorageLocation struct {
	// The chosen location policy Possible values: ["NEAREST_MULTI_REGION", "LOCAL_REGION", "SPECIFIC_LOCATIONS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_snapshot_settings#policy GoogleComputeSnapshotSettings#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_snapshot_settings#locations GoogleComputeSnapshotSettings#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
}

