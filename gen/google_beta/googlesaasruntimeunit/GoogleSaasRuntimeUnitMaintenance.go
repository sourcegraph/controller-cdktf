package googlesaasruntimeunit


type GoogleSaasRuntimeUnitMaintenance struct {
	// If present, it fixes the release on the unit until the given time;
	//
	// i.e.
	// changes to the release field will be rejected. Rollouts should and will
	// also respect this by not requesting an upgrade in the first place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#pinned_until_time GoogleSaasRuntimeUnit#pinned_until_time}
	PinnedUntilTime *string `field:"optional" json:"pinnedUntilTime" yaml:"pinnedUntilTime"`
}

