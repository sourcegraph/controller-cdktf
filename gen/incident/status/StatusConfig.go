package status

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatusConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// What category of status it is.
	//
	// All statuses apart from live (renamed in the app to Active) and learning (renamed in the app to Post-incident) are managed by incident.io and cannot be configured. Possible values are: `triage`, `declined`, `merged`, `canceled`, `live`, `learning`, `closed`, `paused`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/status#category Status#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// Rich text description of the incident status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/status#description Status#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Unique name of this status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/status#name Status#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}
