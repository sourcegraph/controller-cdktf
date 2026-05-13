package googlecontactcenterinsightsassessmentrule


type GoogleContactCenterInsightsAssessmentRuleSampleRule struct {
	// To specify the filter for the conversions that should apply this sample rule.
	//
	// An empty filter means this sample rule applies to all conversations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#conversation_filter GoogleContactCenterInsightsAssessmentRule#conversation_filter}
	ConversationFilter *string `field:"optional" json:"conversationFilter" yaml:"conversationFilter"`
	// Group by dimension to sample the conversation.
	//
	// If no dimension is
	// provided, the sampling will be applied to the project level.
	// Current supported dimensions is 'quality_metadata.agent_info.agent_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#dimension GoogleContactCenterInsightsAssessmentRule#dimension}
	Dimension *string `field:"optional" json:"dimension" yaml:"dimension"`
	// Percentage of conversations that we should sample  based on the dimension between [0, 100].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#sample_percentage GoogleContactCenterInsightsAssessmentRule#sample_percentage}
	SamplePercentage *float64 `field:"optional" json:"samplePercentage" yaml:"samplePercentage"`
	// Number of the conversations that we should sample based on the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_assessment_rule#sample_row GoogleContactCenterInsightsAssessmentRule#sample_row}
	SampleRow *float64 `field:"optional" json:"sampleRow" yaml:"sampleRow"`
}

