package lookerinstance

type LookerInstancePscConfig struct {
	// List of VPCs that are allowed ingress into the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/looker_instance#allowed_vpcs LookerInstance#allowed_vpcs}
	AllowedVpcs *[]*string `field:"optional" json:"allowedVpcs" yaml:"allowedVpcs"`
	// service_attachments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/looker_instance#service_attachments LookerInstance#service_attachments}
	ServiceAttachments any `field:"optional" json:"serviceAttachments" yaml:"serviceAttachments"`
}
