package googlecontactcenterinsightsqascorecard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleContactCenterInsightsQaScorecardConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#location GoogleContactCenterInsightsQaScorecard#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A unique ID for the new QaScorecard.
	//
	// This ID will become the final
	// component of the QaScorecard's resource name. If no ID is specified, a
	// server-generated ID will be used.
	//
	// This value should be 4-64 characters and must match the regular
	// expression '^[a-z0-9-]{4,64}$'. Valid characters are 'a-z-'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#qa_scorecard_id GoogleContactCenterInsightsQaScorecard#qa_scorecard_id}
	QaScorecardId *string `field:"required" json:"qaScorecardId" yaml:"qaScorecardId"`
	// A text description explaining the intent of the scorecard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#description GoogleContactCenterInsightsQaScorecard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The user-specified display name of the scorecard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#display_name GoogleContactCenterInsightsQaScorecard#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#id GoogleContactCenterInsightsQaScorecard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the scorecard is the default one for the project.
	//
	// A default scorecard cannot be deleted and will always appear first in
	// scorecard selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#is_default GoogleContactCenterInsightsQaScorecard#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#project GoogleContactCenterInsightsQaScorecard#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Possible values: QA_SCORECARD_SOURCE_CUSTOMER_DEFINED QA_SCORECARD_SOURCE_DISCOVERY_ENGINE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#source GoogleContactCenterInsightsQaScorecard#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_scorecard#timeouts GoogleContactCenterInsightsQaScorecard#timeouts}
	Timeouts *GoogleContactCenterInsightsQaScorecardTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

