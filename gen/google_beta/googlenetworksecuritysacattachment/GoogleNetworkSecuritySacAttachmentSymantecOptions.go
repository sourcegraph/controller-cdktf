package googlenetworksecuritysacattachment


type GoogleNetworkSecuritySacAttachmentSymantecOptions struct {
	// Name to be used when creating a location on the customer's behalf in Symantec's Location API.
	//
	// Not to be confused with Google Cloud locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#symantec_location_name GoogleNetworkSecuritySacAttachment#symantec_location_name}
	SymantecLocationName *string `field:"optional" json:"symantecLocationName" yaml:"symantecLocationName"`
	// Symantec data center identifier that this attachment will connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#symantec_site GoogleNetworkSecuritySacAttachment#symantec_site}
	SymantecSite *string `field:"optional" json:"symantecSite" yaml:"symantecSite"`
}

