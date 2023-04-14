package healthcheck


type HealthcheckHeader struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#header Healthcheck#header}
	Header *string `field:"required" json:"header" yaml:"header"`
	// A list of string values for the header.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#values Healthcheck#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

