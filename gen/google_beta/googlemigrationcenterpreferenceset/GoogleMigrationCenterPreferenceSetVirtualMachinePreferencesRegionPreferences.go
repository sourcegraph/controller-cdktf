package googlemigrationcenterpreferenceset


type GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences struct {
	// A list of preferred regions, ordered by the most preferred region first.
	//
	// Set only valid Google Cloud region names. See https://cloud.google.com/compute/docs/regions-zones for available regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#preferred_regions GoogleMigrationCenterPreferenceSet#preferred_regions}
	PreferredRegions *[]*string `field:"optional" json:"preferredRegions" yaml:"preferredRegions"`
}

