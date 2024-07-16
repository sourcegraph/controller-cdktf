package googlednsmanagedzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDnsManagedZoneConfig struct {
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
	// The DNS name of this managed zone, for instance "example.com.".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#dns_name GoogleDnsManagedZone#dns_name}
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// User assigned name for this resource. Must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#name GoogleDnsManagedZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#cloud_logging_config GoogleDnsManagedZone#cloud_logging_config}
	CloudLoggingConfig *GoogleDnsManagedZoneCloudLoggingConfig `field:"optional" json:"cloudLoggingConfig" yaml:"cloudLoggingConfig"`
	// A textual description field. Defaults to 'Managed by Terraform'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#description GoogleDnsManagedZone#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dnssec_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#dnssec_config GoogleDnsManagedZone#dnssec_config}
	DnssecConfig *GoogleDnsManagedZoneDnssecConfig `field:"optional" json:"dnssecConfig" yaml:"dnssecConfig"`
	// Set this true to delete all records in the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#force_destroy GoogleDnsManagedZone#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// forwarding_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#forwarding_config GoogleDnsManagedZone#forwarding_config}
	ForwardingConfig *GoogleDnsManagedZoneForwardingConfig `field:"optional" json:"forwardingConfig" yaml:"forwardingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#id GoogleDnsManagedZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to this ManagedZone.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#labels GoogleDnsManagedZone#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#peering_config GoogleDnsManagedZone#peering_config}
	PeeringConfig *GoogleDnsManagedZonePeeringConfig `field:"optional" json:"peeringConfig" yaml:"peeringConfig"`
	// private_visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#private_visibility_config GoogleDnsManagedZone#private_visibility_config}
	PrivateVisibilityConfig *GoogleDnsManagedZonePrivateVisibilityConfig `field:"optional" json:"privateVisibilityConfig" yaml:"privateVisibilityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#project GoogleDnsManagedZone#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Specifies if this is a managed reverse lookup zone.
	//
	// If true, Cloud DNS will resolve reverse
	// lookup queries using automatically configured records for VPC resources. This only applies
	// to networks listed under 'private_visibility_config'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#reverse_lookup GoogleDnsManagedZone#reverse_lookup}
	ReverseLookup interface{} `field:"optional" json:"reverseLookup" yaml:"reverseLookup"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#service_directory_config GoogleDnsManagedZone#service_directory_config}
	ServiceDirectoryConfig *GoogleDnsManagedZoneServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#timeouts GoogleDnsManagedZone#timeouts}
	Timeouts *GoogleDnsManagedZoneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	//
	// Default value: "public" Possible values: ["private", "public"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#visibility GoogleDnsManagedZone#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

