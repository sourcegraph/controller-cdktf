package googleintegrationconnectorsmanagedzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIntegrationConnectorsManagedZoneConfig struct {
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
	// DNS Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#dns GoogleIntegrationConnectorsManagedZone#dns}
	Dns *string `field:"required" json:"dns" yaml:"dns"`
	// Name of Managed Zone needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#name GoogleIntegrationConnectorsManagedZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Target Project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#target_project GoogleIntegrationConnectorsManagedZone#target_project}
	TargetProject *string `field:"required" json:"targetProject" yaml:"targetProject"`
	// The name of the Target Project VPC Network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#target_vpc GoogleIntegrationConnectorsManagedZone#target_vpc}
	TargetVpc *string `field:"required" json:"targetVpc" yaml:"targetVpc"`
	// Description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#description GoogleIntegrationConnectorsManagedZone#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#id GoogleIntegrationConnectorsManagedZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#labels GoogleIntegrationConnectorsManagedZone#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#project GoogleIntegrationConnectorsManagedZone#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_managed_zone#timeouts GoogleIntegrationConnectorsManagedZone#timeouts}
	Timeouts *GoogleIntegrationConnectorsManagedZoneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

