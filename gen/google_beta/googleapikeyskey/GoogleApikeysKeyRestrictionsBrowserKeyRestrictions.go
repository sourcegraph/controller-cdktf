package googleapikeyskey


type GoogleApikeysKeyRestrictionsBrowserKeyRestrictions struct {
	// A list of regular expressions for the referrer URLs that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#allowed_referrers GoogleApikeysKey#allowed_referrers}
	AllowedReferrers *[]*string `field:"required" json:"allowedReferrers" yaml:"allowedReferrers"`
}

