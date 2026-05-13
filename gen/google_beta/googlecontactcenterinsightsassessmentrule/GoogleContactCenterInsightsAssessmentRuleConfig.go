package googlecontactcenterinsightsassessmentrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleContactCenterInsightsAssessmentRuleConfig struct {
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
	// Location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#location GoogleContactCenterInsightsAssessmentRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// If true, apply this rule to conversations. Otherwise, this rule is inactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#active GoogleContactCenterInsightsAssessmentRule#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// A unique ID for the new AssessmentRule.
	//
	// This ID will become the final
	// component of the AssessmentRule's resource name. If no ID is specified,
	// a server-generated ID will be used.
	//
	// This value should be 4-64 characters and must match the regular
	// expression '^[A-Za-z0-9]{4,64}$'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#assessment_rule_id GoogleContactCenterInsightsAssessmentRule#assessment_rule_id}
	AssessmentRuleId *string `field:"optional" json:"assessmentRuleId" yaml:"assessmentRuleId"`
	// Display Name of the assessment rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#display_name GoogleContactCenterInsightsAssessmentRule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#id GoogleContactCenterInsightsAssessmentRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#project GoogleContactCenterInsightsAssessmentRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// sample_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#sample_rule GoogleContactCenterInsightsAssessmentRule#sample_rule}
	SampleRule *GoogleContactCenterInsightsAssessmentRuleSampleRule `field:"optional" json:"sampleRule" yaml:"sampleRule"`
	// schedule_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#schedule_info GoogleContactCenterInsightsAssessmentRule#schedule_info}
	ScheduleInfo *GoogleContactCenterInsightsAssessmentRuleScheduleInfo `field:"optional" json:"scheduleInfo" yaml:"scheduleInfo"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#timeouts GoogleContactCenterInsightsAssessmentRule#timeouts}
	Timeouts *GoogleContactCenterInsightsAssessmentRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

