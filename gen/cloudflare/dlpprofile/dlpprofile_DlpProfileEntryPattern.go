package dlpprofile


type DlpProfileEntryPattern struct {
	// The regex that defines the pattern.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/dlp_profile#regex DlpProfile#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// The validation algorithm to apply with this pattern.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/dlp_profile#validation DlpProfile#validation}
	Validation *string `field:"optional" json:"validation" yaml:"validation"`
}

