package escalationpath

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EscalationPathConfig struct {
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
	// The name of this escalation path, for the user's reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#name EscalationPath#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The nodes that form the levels and branches of this escalation path.
	//
	// -->**Note** Although the `if_else` block is recursive, currently a maximum of 3 levels are supported. Attempting to configure more than 3 levels of nesting will result in a schema error.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#path EscalationPath#path}
	Path interface{} `field:"required" json:"path" yaml:"path"`
	// IDs of the teams that own this escalation path.
	//
	// This will automatically sync escalation paths with the right teams in Catalog. If you have an escalation paths attribute on your Teams, this attribute is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#team_ids EscalationPath#team_ids}
	TeamIds *[]*string `field:"optional" json:"teamIds" yaml:"teamIds"`
	// The working hours for this escalation path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/escalation_path#working_hours EscalationPath#working_hours}
	WorkingHours interface{} `field:"optional" json:"workingHours" yaml:"workingHours"`
}

