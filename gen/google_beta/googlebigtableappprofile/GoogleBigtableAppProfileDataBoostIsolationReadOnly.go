package googlebigtableappprofile


type GoogleBigtableAppProfileDataBoostIsolationReadOnly struct {
	// The Compute Billing Owner for this Data Boost App Profile. Possible values: ["HOST_PAYS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_bigtable_app_profile#compute_billing_owner GoogleBigtableAppProfile#compute_billing_owner}
	ComputeBillingOwner *string `field:"required" json:"computeBillingOwner" yaml:"computeBillingOwner"`
}

