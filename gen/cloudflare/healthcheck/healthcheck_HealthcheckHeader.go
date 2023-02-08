package healthcheck


type HealthcheckHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#header Healthcheck#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#values Healthcheck#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

