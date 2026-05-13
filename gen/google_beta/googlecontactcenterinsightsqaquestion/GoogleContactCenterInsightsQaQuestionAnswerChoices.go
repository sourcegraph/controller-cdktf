package googlecontactcenterinsightsqaquestion


type GoogleContactCenterInsightsQaQuestionAnswerChoices struct {
	// Boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#bool_value GoogleContactCenterInsightsQaQuestion#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// A short string used as an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#key GoogleContactCenterInsightsQaQuestion#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value of "Not Applicable (N/A)".
	//
	// If provided, this field may only
	// be set to 'true'. If a question receives this answer, it will be
	// excluded from any score calculations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#na_value GoogleContactCenterInsightsQaQuestion#na_value}
	NaValue interface{} `field:"optional" json:"naValue" yaml:"naValue"`
	// Numerical value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#num_value GoogleContactCenterInsightsQaQuestion#num_value}
	NumValue *float64 `field:"optional" json:"numValue" yaml:"numValue"`
	// Numerical score of the answer, used for generating the overall score of a QaScorecardResult.
	//
	// If the answer uses na_value, this field is unused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#score GoogleContactCenterInsightsQaQuestion#score}
	Score *float64 `field:"optional" json:"score" yaml:"score"`
	// String value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#str_value GoogleContactCenterInsightsQaQuestion#str_value}
	StrValue *string `field:"optional" json:"strValue" yaml:"strValue"`
}

