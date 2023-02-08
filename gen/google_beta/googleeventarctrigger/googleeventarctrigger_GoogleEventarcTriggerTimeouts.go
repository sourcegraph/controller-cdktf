package googleeventarctrigger


type GoogleEventarcTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#create GoogleEventarcTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#delete GoogleEventarcTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#update GoogleEventarcTrigger#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

