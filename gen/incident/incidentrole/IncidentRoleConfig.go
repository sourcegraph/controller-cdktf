package incidentrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IncidentRoleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Describes the purpose of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/incident_role#description IncidentRole#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Provided to whoever is nominated for the role. Note that this will be empty for the 'reporter' role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/incident_role#instructions IncidentRole#instructions}
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// Human readable name of the incident role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/incident_role#name IncidentRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Short human readable name for Slack. Note that this will be empty for the 'reporter' role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/incident_role#shortform IncidentRole#shortform}
	Shortform *string `field:"required" json:"shortform" yaml:"shortform"`
}

