package googleapikeyskey


type GoogleApikeysKeyRestrictionsServerKeyRestrictions struct {
	// A list of the caller IP addresses that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apikeys_key#allowed_ips GoogleApikeysKey#allowed_ips}
	AllowedIps *[]*string `field:"required" json:"allowedIps" yaml:"allowedIps"`
}

