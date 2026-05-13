package googlenetworkservicesmulticastdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesMulticastDomainConfig struct {
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
	// The resource name of the multicast admin VPC network. Use the following format: 'projects/{project}/locations/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#admin_network GoogleNetworkServicesMulticastDomain#admin_network}
	AdminNetwork *string `field:"required" json:"adminNetwork" yaml:"adminNetwork"`
	// connection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#connection_config GoogleNetworkServicesMulticastDomain#connection_config}
	ConnectionConfig *GoogleNetworkServicesMulticastDomainConnectionConfig `field:"required" json:"connectionConfig" yaml:"connectionConfig"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#location GoogleNetworkServicesMulticastDomain#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A unique name for the multicast domain.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#multicast_domain_id GoogleNetworkServicesMulticastDomain#multicast_domain_id}
	MulticastDomainId *string `field:"required" json:"multicastDomainId" yaml:"multicastDomainId"`
	// An optional text description of the multicast domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#description GoogleNetworkServicesMulticastDomain#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#id GoogleNetworkServicesMulticastDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#labels GoogleNetworkServicesMulticastDomain#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The multicast domain group this domain should be associated with. Use the following format: 'projects/{project}/locations/global/multicastDomainGroups/{multicast_domain_group}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#multicast_domain_group GoogleNetworkServicesMulticastDomain#multicast_domain_group}
	MulticastDomainGroup *string `field:"optional" json:"multicastDomainGroup" yaml:"multicastDomainGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#project GoogleNetworkServicesMulticastDomain#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#timeouts GoogleNetworkServicesMulticastDomain#timeouts}
	Timeouts *GoogleNetworkServicesMulticastDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ull_multicast_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#ull_multicast_domain GoogleNetworkServicesMulticastDomain#ull_multicast_domain}
	UllMulticastDomain *GoogleNetworkServicesMulticastDomainUllMulticastDomain `field:"optional" json:"ullMulticastDomain" yaml:"ullMulticastDomain"`
}

