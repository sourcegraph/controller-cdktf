package googlelookerinstance


type GoogleLookerInstancePscConfig struct {
	// List of VPCs that are allowed ingress into the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_looker_instance#allowed_vpcs GoogleLookerInstance#allowed_vpcs}
	AllowedVpcs *[]*string `field:"optional" json:"allowedVpcs" yaml:"allowedVpcs"`
	// service_attachments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_looker_instance#service_attachments GoogleLookerInstance#service_attachments}
	ServiceAttachments interface{} `field:"optional" json:"serviceAttachments" yaml:"serviceAttachments"`
}

