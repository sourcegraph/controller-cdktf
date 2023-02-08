package accesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#application_id AccessPolicy#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#decision AccessPolicy#decision}.
	Decision *string `field:"required" json:"decision" yaml:"decision"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#include AccessPolicy#include}
	Include interface{} `field:"required" json:"include" yaml:"include"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#name AccessPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#precedence AccessPolicy#precedence}.
	Precedence *float64 `field:"required" json:"precedence" yaml:"precedence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#account_id AccessPolicy#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// approval_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#approval_group AccessPolicy#approval_group}
	ApprovalGroup interface{} `field:"optional" json:"approvalGroup" yaml:"approvalGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#approval_required AccessPolicy#approval_required}.
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#exclude AccessPolicy#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#id AccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#purpose_justification_prompt AccessPolicy#purpose_justification_prompt}.
	PurposeJustificationPrompt *string `field:"optional" json:"purposeJustificationPrompt" yaml:"purposeJustificationPrompt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#purpose_justification_required AccessPolicy#purpose_justification_required}.
	PurposeJustificationRequired interface{} `field:"optional" json:"purposeJustificationRequired" yaml:"purposeJustificationRequired"`
	// require block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#require AccessPolicy#require}
	Require interface{} `field:"optional" json:"require" yaml:"require"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#zone_id AccessPolicy#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

