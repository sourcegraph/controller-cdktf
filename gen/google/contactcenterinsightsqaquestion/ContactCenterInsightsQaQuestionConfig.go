package contactcenterinsightsqaquestion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContactCenterInsightsQaQuestionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#location ContactCenterInsightsQaQuestion#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#qa_scorecard ContactCenterInsightsQaQuestion#qa_scorecard}
	QaScorecard *string `field:"required" json:"qaScorecard" yaml:"qaScorecard"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#revision ContactCenterInsightsQaQuestion#revision}
	Revision *string `field:"required" json:"revision" yaml:"revision"`
	// Short, descriptive string, used in the UI where it's not practical to display the full question body. E.g., "Greeting".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#abbreviation ContactCenterInsightsQaQuestion#abbreviation}
	Abbreviation *string `field:"optional" json:"abbreviation" yaml:"abbreviation"`
	// answer_choices block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#answer_choices ContactCenterInsightsQaQuestion#answer_choices}
	AnswerChoices interface{} `field:"optional" json:"answerChoices" yaml:"answerChoices"`
	// Instructions describing how to determine the answer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#answer_instructions ContactCenterInsightsQaQuestion#answer_instructions}
	AnswerInstructions *string `field:"optional" json:"answerInstructions" yaml:"answerInstructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#id ContactCenterInsightsQaQuestion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#metrics ContactCenterInsightsQaQuestion#metrics}
	Metrics *ContactCenterInsightsQaQuestionMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// Defines the order of the question within its parent scorecard revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#order ContactCenterInsightsQaQuestion#order}
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// predefined_question_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#predefined_question_config ContactCenterInsightsQaQuestion#predefined_question_config}
	PredefinedQuestionConfig *ContactCenterInsightsQaQuestionPredefinedQuestionConfig `field:"optional" json:"predefinedQuestionConfig" yaml:"predefinedQuestionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#project ContactCenterInsightsQaQuestion#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// qa_question_data_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#qa_question_data_options ContactCenterInsightsQaQuestion#qa_question_data_options}
	QaQuestionDataOptions *ContactCenterInsightsQaQuestionQaQuestionDataOptions `field:"optional" json:"qaQuestionDataOptions" yaml:"qaQuestionDataOptions"`
	// Question text. E.g., "Did the agent greet the customer?".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#question_body ContactCenterInsightsQaQuestion#question_body}
	QuestionBody *string `field:"optional" json:"questionBody" yaml:"questionBody"`
	// The type of question. Possible values: CUSTOMIZABLE PREDEFINED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#question_type ContactCenterInsightsQaQuestion#question_type}
	QuestionType *string `field:"optional" json:"questionType" yaml:"questionType"`
	// Questions are tagged for categorization and scoring.
	//
	// Tags can either be:
	// - Default Tags: These are predefined categories. They are identified by
	// their string value (e.g., "BUSINESS", "COMPLIANCE", and "CUSTOMER").
	// - Custom Tags: These are user-defined categories. They are identified by
	// their full resource name (e.g.,
	// projects/{project}/locations/{location}/qaQuestionTags/{qa_question_tag}).
	// Both default and custom tags are used to group questions and to influence
	// the scoring of each question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#tags ContactCenterInsightsQaQuestion#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#timeouts ContactCenterInsightsQaQuestion#timeouts}
	Timeouts *ContactCenterInsightsQaQuestionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tuning_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question#tuning_metadata ContactCenterInsightsQaQuestion#tuning_metadata}
	TuningMetadata *ContactCenterInsightsQaQuestionTuningMetadata `field:"optional" json:"tuningMetadata" yaml:"tuningMetadata"`
}

