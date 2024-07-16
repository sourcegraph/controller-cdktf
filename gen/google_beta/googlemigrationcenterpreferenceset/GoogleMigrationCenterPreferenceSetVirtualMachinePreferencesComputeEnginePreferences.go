package googlemigrationcenterpreferenceset


type GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences struct {
	// License type to consider when calculating costs for virtual machine insights and recommendations.
	//
	// If unspecified, costs are calculated based on the default licensing plan.
	//  Possible values:
	//  LICENSE_TYPE_UNSPECIFIED
	// LICENSE_TYPE_DEFAULT
	// LICENSE_TYPE_BRING_YOUR_OWN_LICENSE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#license_type GoogleMigrationCenterPreferenceSet#license_type}
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// machine_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#machine_preferences GoogleMigrationCenterPreferenceSet#machine_preferences}
	MachinePreferences *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences `field:"optional" json:"machinePreferences" yaml:"machinePreferences"`
}

