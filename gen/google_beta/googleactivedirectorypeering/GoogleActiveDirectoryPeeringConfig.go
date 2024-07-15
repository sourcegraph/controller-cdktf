package googleactivedirectorypeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleActiveDirectoryPeeringConfig struct {
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
	// The full names of the Google Compute Engine networks to which the instance is connected.
	//
	// Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#authorized_network GoogleActiveDirectoryPeering#authorized_network}
	AuthorizedNetwork *string `field:"required" json:"authorizedNetwork" yaml:"authorizedNetwork"`
	// Full domain resource path for the Managed AD Domain involved in peering.
	//
	// The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#domain_resource GoogleActiveDirectoryPeering#domain_resource}
	DomainResource *string `field:"required" json:"domainResource" yaml:"domainResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#peering_id GoogleActiveDirectoryPeering#peering_id}.
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#id GoogleActiveDirectoryPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels that can contain user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#labels GoogleActiveDirectoryPeering#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#project GoogleActiveDirectoryPeering#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The current state of this Peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#status GoogleActiveDirectoryPeering#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Additional information about the current status of this peering, if available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#status_message GoogleActiveDirectoryPeering#status_message}
	StatusMessage *string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_peering#timeouts GoogleActiveDirectoryPeering#timeouts}
	Timeouts *GoogleActiveDirectoryPeeringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

