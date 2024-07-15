package googleapikeyskey


type GoogleApikeysKeyRestrictionsIosKeyRestrictions struct {
	// A list of bundle IDs that are allowed when making API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#allowed_bundle_ids GoogleApikeysKey#allowed_bundle_ids}
	AllowedBundleIds *[]*string `field:"required" json:"allowedBundleIds" yaml:"allowedBundleIds"`
}

