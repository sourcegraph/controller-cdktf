package googlecomputeresourcepolicy


type GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name.
	//
	// The chain name must be 1-63 characters long and comply
	// with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#chain_name GoogleComputeResourcePolicy#chain_name}
	ChainName *string `field:"optional" json:"chainName" yaml:"chainName"`
	// Whether to perform a 'guest aware' snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#guest_flush GoogleComputeResourcePolicy#guest_flush}
	GuestFlush interface{} `field:"optional" json:"guestFlush" yaml:"guestFlush"`
	// A set of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#labels GoogleComputeResourcePolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Cloud Storage bucket location to store the auto snapshot (regional or multi-regional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#storage_locations GoogleComputeResourcePolicy#storage_locations}
	StorageLocations *[]*string `field:"optional" json:"storageLocations" yaml:"storageLocations"`
}

