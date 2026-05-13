package googlecontactcenterinsightsanalysisrule


type GoogleContactCenterInsightsAnalysisRuleAnnotatorSelector struct {
	// The issue model to run.
	//
	// If not provided, the most recently deployed topic
	// model will be used. The provided issue model will only be used for
	// inference if the issue model is deployed and if run_issue_model_annotator
	// is set to true. If more than one issue model is provided, only the first
	// provided issue model will be used for inference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#issue_models GoogleContactCenterInsightsAnalysisRule#issue_models}
	IssueModels *[]*string `field:"optional" json:"issueModels" yaml:"issueModels"`
	// The list of phrase matchers to run.
	//
	// If not provided, all active phrase
	// matchers will be used. If inactive phrase matchers are provided, they will
	// not be used. Phrase matchers will be run only if
	// run_phrase_matcher_annotator is set to true. Format:
	// projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#phrase_matchers GoogleContactCenterInsightsAnalysisRule#phrase_matchers}
	PhraseMatchers *[]*string `field:"optional" json:"phraseMatchers" yaml:"phraseMatchers"`
	// qa_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#qa_config GoogleContactCenterInsightsAnalysisRule#qa_config}
	QaConfig *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig `field:"optional" json:"qaConfig" yaml:"qaConfig"`
	// Whether to run the entity annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_entity_annotator GoogleContactCenterInsightsAnalysisRule#run_entity_annotator}
	RunEntityAnnotator interface{} `field:"optional" json:"runEntityAnnotator" yaml:"runEntityAnnotator"`
	// Whether to run the intent annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_intent_annotator GoogleContactCenterInsightsAnalysisRule#run_intent_annotator}
	RunIntentAnnotator interface{} `field:"optional" json:"runIntentAnnotator" yaml:"runIntentAnnotator"`
	// Whether to run the interruption annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_interruption_annotator GoogleContactCenterInsightsAnalysisRule#run_interruption_annotator}
	RunInterruptionAnnotator interface{} `field:"optional" json:"runInterruptionAnnotator" yaml:"runInterruptionAnnotator"`
	// Whether to run the issue model annotator. A model should have already been deployed for this to take effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_issue_model_annotator GoogleContactCenterInsightsAnalysisRule#run_issue_model_annotator}
	RunIssueModelAnnotator interface{} `field:"optional" json:"runIssueModelAnnotator" yaml:"runIssueModelAnnotator"`
	// Whether to run the active phrase matcher annotator(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_phrase_matcher_annotator GoogleContactCenterInsightsAnalysisRule#run_phrase_matcher_annotator}
	RunPhraseMatcherAnnotator interface{} `field:"optional" json:"runPhraseMatcherAnnotator" yaml:"runPhraseMatcherAnnotator"`
	// Whether to run the QA annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_qa_annotator GoogleContactCenterInsightsAnalysisRule#run_qa_annotator}
	RunQaAnnotator interface{} `field:"optional" json:"runQaAnnotator" yaml:"runQaAnnotator"`
	// Whether to run the sentiment annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_sentiment_annotator GoogleContactCenterInsightsAnalysisRule#run_sentiment_annotator}
	RunSentimentAnnotator interface{} `field:"optional" json:"runSentimentAnnotator" yaml:"runSentimentAnnotator"`
	// Whether to run the silence annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_silence_annotator GoogleContactCenterInsightsAnalysisRule#run_silence_annotator}
	RunSilenceAnnotator interface{} `field:"optional" json:"runSilenceAnnotator" yaml:"runSilenceAnnotator"`
	// Whether to run the summarization annotator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#run_summarization_annotator GoogleContactCenterInsightsAnalysisRule#run_summarization_annotator}
	RunSummarizationAnnotator interface{} `field:"optional" json:"runSummarizationAnnotator" yaml:"runSummarizationAnnotator"`
	// summarization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_analysis_rule#summarization_config GoogleContactCenterInsightsAnalysisRule#summarization_config}
	SummarizationConfig *GoogleContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig `field:"optional" json:"summarizationConfig" yaml:"summarizationConfig"`
}

