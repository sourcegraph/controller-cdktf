package googleapikeyskey


type GoogleApikeysKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apikeys_key#create GoogleApikeysKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apikeys_key#delete GoogleApikeysKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apikeys_key#update GoogleApikeysKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

