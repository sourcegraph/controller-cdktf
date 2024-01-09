package slo


type SloAttachment struct {
	// URL to the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#url Slo#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Name displayed for the attachment. Max. length: 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#display_name Slo#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

