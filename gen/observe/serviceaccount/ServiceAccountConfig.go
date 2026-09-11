package serviceaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServiceAccountConfig struct {
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
	// The label of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/service_account#label ServiceAccount#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// An optional description of the service account's purpose or usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/service_account#description ServiceAccount#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the service account is disabled. Disabled service accounts cannot authenticate. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/service_account#disabled ServiceAccount#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

