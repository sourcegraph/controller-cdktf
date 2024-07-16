package googlebigtableappprofile


type GoogleBigtableAppProfileStandardIsolation struct {
	// The priority of requests sent using this app profile. Possible values: ["PRIORITY_LOW", "PRIORITY_MEDIUM", "PRIORITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_app_profile#priority GoogleBigtableAppProfile#priority}
	Priority *string `field:"required" json:"priority" yaml:"priority"`
}

