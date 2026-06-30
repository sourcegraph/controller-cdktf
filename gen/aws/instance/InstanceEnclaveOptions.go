package instance

type InstanceEnclaveOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/instance#enabled Instance#enabled}.
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
}
