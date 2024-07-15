package googleapikeyskey


type GoogleApikeysKeyRestrictionsServerKeyRestrictions struct {
	// A list of the caller IP addresses that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#allowed_ips GoogleApikeysKey#allowed_ips}
	AllowedIps *[]*string `field:"required" json:"allowedIps" yaml:"allowedIps"`
}

