package migrationcenterpreferenceset


type MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences struct {
	// License type to consider when calculating costs for virtual machine insights and recommendations.
	//
	// If unspecified, costs are calculated based on the default licensing plan. Possible values: 'LICENSE_TYPE_UNSPECIFIED', 'LICENSE_TYPE_DEFAULT', 'LICENSE_TYPE_BRING_YOUR_OWN_LICENSE'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/migration_center_preference_set#license_type MigrationCenterPreferenceSet#license_type}
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// machine_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/migration_center_preference_set#machine_preferences MigrationCenterPreferenceSet#machine_preferences}
	MachinePreferences *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences `field:"optional" json:"machinePreferences" yaml:"machinePreferences"`
	// Persistent disk type to use.
	//
	// If unspecified (default), all types are considered, based on available usage data. Possible values: ["PERSISTENT_DISK_TYPE_STANDARD", "PERSISTENT_DISK_TYPE_BALANCED", "PERSISTENT_DISK_TYPE_SSD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/migration_center_preference_set#persistent_disk_type MigrationCenterPreferenceSet#persistent_disk_type}
	PersistentDiskType *string `field:"optional" json:"persistentDiskType" yaml:"persistentDiskType"`
}

