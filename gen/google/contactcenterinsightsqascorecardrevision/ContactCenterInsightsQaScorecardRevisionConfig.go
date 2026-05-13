package contactcenterinsightsqascorecardrevision

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContactCenterInsightsQaScorecardRevisionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_scorecard_revision#location ContactCenterInsightsQaScorecardRevision#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_scorecard_revision#qa_scorecard ContactCenterInsightsQaScorecardRevision#qa_scorecard}
	QaScorecard *string `field:"required" json:"qaScorecard" yaml:"qaScorecard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_scorecard_revision#id ContactCenterInsightsQaScorecardRevision#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_scorecard_revision#project ContactCenterInsightsQaScorecardRevision#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A unique ID for the new QaScorecardRevision.
	//
	// This ID will become the final
	// component of the QaScorecardRevision's resource name.
	// If no ID is specified this resource will get the latest revision on the given scorecard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_scorecard_revision#qa_scorecard_revision_id ContactCenterInsightsQaScorecardRevision#qa_scorecard_revision_id}
	QaScorecardRevisionId *string `field:"optional" json:"qaScorecardRevisionId" yaml:"qaScorecardRevisionId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_scorecard_revision#timeouts ContactCenterInsightsQaScorecardRevision#timeouts}
	Timeouts *ContactCenterInsightsQaScorecardRevisionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

