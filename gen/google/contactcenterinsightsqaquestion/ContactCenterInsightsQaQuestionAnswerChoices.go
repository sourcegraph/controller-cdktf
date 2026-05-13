package contactcenterinsightsqaquestion


type ContactCenterInsightsQaQuestionAnswerChoices struct {
	// Boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#bool_value ContactCenterInsightsQaQuestion#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// A short string used as an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#key ContactCenterInsightsQaQuestion#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value of "Not Applicable (N/A)".
	//
	// If provided, this field may only
	// be set to 'true'. If a question receives this answer, it will be
	// excluded from any score calculations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#na_value ContactCenterInsightsQaQuestion#na_value}
	NaValue interface{} `field:"optional" json:"naValue" yaml:"naValue"`
	// Numerical value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#num_value ContactCenterInsightsQaQuestion#num_value}
	NumValue *float64 `field:"optional" json:"numValue" yaml:"numValue"`
	// Numerical score of the answer, used for generating the overall score of a QaScorecardResult.
	//
	// If the answer uses na_value, this field is unused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#score ContactCenterInsightsQaQuestion#score}
	Score *float64 `field:"optional" json:"score" yaml:"score"`
	// String value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#str_value ContactCenterInsightsQaQuestion#str_value}
	StrValue *string `field:"optional" json:"strValue" yaml:"strValue"`
}

