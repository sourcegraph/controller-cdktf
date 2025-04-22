package googlebeyondcorpapplication


type GoogleBeyondcorpApplicationEndpointMatchers struct {
	// Required. Hostname of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application#hostname GoogleBeyondcorpApplication#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Optional. Ports of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application#ports GoogleBeyondcorpApplication#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

