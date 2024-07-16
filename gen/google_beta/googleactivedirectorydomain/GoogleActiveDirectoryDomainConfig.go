package googleactivedirectorydomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleActiveDirectoryDomainConfig struct {
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
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions of https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#domain_name GoogleActiveDirectoryDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Locations where domain needs to be provisioned.
	//
	// [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#locations GoogleActiveDirectoryDomain#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// The CIDR range of internal addresses that are reserved for this domain.
	//
	// Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#reserved_ip_range GoogleActiveDirectoryDomain#reserved_ip_range}
	ReservedIpRange *string `field:"required" json:"reservedIpRange" yaml:"reservedIpRange"`
	// The name of delegated administrator account used to perform Active Directory operations. If not specified, setupadmin will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#admin GoogleActiveDirectoryDomain#admin}
	Admin *string `field:"optional" json:"admin" yaml:"admin"`
	// The full names of the Google Compute Engine networks the domain instance is connected to.
	//
	// The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#authorized_networks GoogleActiveDirectoryDomain#authorized_networks}
	AuthorizedNetworks *[]*string `field:"optional" json:"authorizedNetworks" yaml:"authorizedNetworks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#id GoogleActiveDirectoryDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels that can contain user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#labels GoogleActiveDirectoryDomain#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#project GoogleActiveDirectoryDomain#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain#timeouts GoogleActiveDirectoryDomain#timeouts}
	Timeouts *GoogleActiveDirectoryDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

