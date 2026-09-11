package severity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SeverityConfig struct {
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
	// Description of the severity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/severity#description Severity#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Human readable name of the severity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/severity#name Severity#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Rank to help sort severities (lower numbers are less severe).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/severity#rank Severity#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

