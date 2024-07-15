package googlecomputeinstancesettings


type GoogleComputeInstanceSettingsMetadata struct {
	// A metadata key/value items map. The total size of all keys and values must be less than 512KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_settings#items GoogleComputeInstanceSettings#items}
	Items *map[string]*string `field:"optional" json:"items" yaml:"items"`
}

