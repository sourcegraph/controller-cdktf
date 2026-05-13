package googlenetworksecuritysacattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecuritySacAttachmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The location of the SACAttachment resource. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#location GoogleNetworkSecuritySacAttachment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Identifier. Resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#name GoogleNetworkSecuritySacAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// NCC Gateway associated with the attachment. This can be input as an ID or a full resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#ncc_gateway GoogleNetworkSecuritySacAttachment#ncc_gateway}
	NccGateway *string `field:"required" json:"nccGateway" yaml:"nccGateway"`
	// SAC Realm which owns the attachment. This can be input as an ID or a full resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#sac_realm GoogleNetworkSecuritySacAttachment#sac_realm}
	SacRealm *string `field:"required" json:"sacRealm" yaml:"sacRealm"`
	// Case-insensitive ISO-3166 alpha-2 country code used for localization. Only valid for Symantec attachments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#country GoogleNetworkSecuritySacAttachment#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#id GoogleNetworkSecuritySacAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://docs.cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#labels GoogleNetworkSecuritySacAttachment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#project GoogleNetworkSecuritySacAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// symantec_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#symantec_options GoogleNetworkSecuritySacAttachment#symantec_options}
	SymantecOptions *GoogleNetworkSecuritySacAttachmentSymantecOptions `field:"optional" json:"symantecOptions" yaml:"symantecOptions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#timeouts GoogleNetworkSecuritySacAttachment#timeouts}
	Timeouts *GoogleNetworkSecuritySacAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Case-sensitive tzinfo identifier used for localization. Only valid for Symantec attachments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_attachment#time_zone GoogleNetworkSecuritySacAttachment#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

