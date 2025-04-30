package googlelookerinstance


type GoogleLookerInstancePscConfigServiceAttachments struct {
	// Fully qualified domain name that will be used in the private DNS record created for the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_looker_instance#local_fqdn GoogleLookerInstance#local_fqdn}
	LocalFqdn *string `field:"optional" json:"localFqdn" yaml:"localFqdn"`
	// URI of the service attachment to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_looker_instance#target_service_attachment_uri GoogleLookerInstance#target_service_attachment_uri}
	TargetServiceAttachmentUri *string `field:"optional" json:"targetServiceAttachmentUri" yaml:"targetServiceAttachmentUri"`
}

