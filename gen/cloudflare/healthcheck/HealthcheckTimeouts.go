package healthcheck


type HealthcheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#create Healthcheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

