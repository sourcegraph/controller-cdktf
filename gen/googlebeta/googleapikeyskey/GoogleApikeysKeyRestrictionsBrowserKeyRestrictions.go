package googleapikeyskey


type GoogleApikeysKeyRestrictionsBrowserKeyRestrictions struct {
	// A list of regular expressions for the referrer URLs that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apikeys_key#allowed_referrers GoogleApikeysKey#allowed_referrers}
	AllowedReferrers *[]*string `field:"required" json:"allowedReferrers" yaml:"allowedReferrers"`
}

