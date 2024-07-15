package googlemigrationcenterpreferenceset


type GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesNodeTypes struct {
	// Name of the Sole Tenant node. Consult https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#node_name GoogleMigrationCenterPreferenceSet#node_name}
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
}

