package discoveryenginewidgetconfig


type DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig struct {
	// Whether generated answer contains suggested related questions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#disable_related_questions DiscoveryEngineWidgetConfig#disable_related_questions}
	DisableRelatedQuestions interface{} `field:"optional" json:"disableRelatedQuestions" yaml:"disableRelatedQuestions"`
	// Specifies whether to filter out queries that are adversarial.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#ignore_adversarial_query DiscoveryEngineWidgetConfig#ignore_adversarial_query}
	IgnoreAdversarialQuery interface{} `field:"optional" json:"ignoreAdversarialQuery" yaml:"ignoreAdversarialQuery"`
	// Specifies whether to filter out queries that are not relevant to the content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#ignore_low_relevant_content DiscoveryEngineWidgetConfig#ignore_low_relevant_content}
	IgnoreLowRelevantContent interface{} `field:"optional" json:"ignoreLowRelevantContent" yaml:"ignoreLowRelevantContent"`
	// Specifies whether to filter out queries that are not answer-seeking.
	//
	// The default value is 'false'. No answer is returned if the search query
	// is classified as a non-answer seeking query.
	// If this field is set to 'true', we skip generating answers for
	// non-answer seeking queries and return fallback messages instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#ignore_non_answer_seeking_query DiscoveryEngineWidgetConfig#ignore_non_answer_seeking_query}
	IgnoreNonAnswerSeekingQuery interface{} `field:"optional" json:"ignoreNonAnswerSeekingQuery" yaml:"ignoreNonAnswerSeekingQuery"`
	// Source of image returned in the answer. Possible values: ["ALL_AVAILABLE_SOURCES", "CORPUS_IMAGE_ONLY", "FIGURE_GENERATION_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#image_source DiscoveryEngineWidgetConfig#image_source}
	ImageSource *string `field:"optional" json:"imageSource" yaml:"imageSource"`
	// Language code for Summary. Use language tags defined by [BCP47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). Note: This is an experimental feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#language_code DiscoveryEngineWidgetConfig#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Max rephrase steps.
	//
	// The max number is 5 steps. If not set or
	// set to < 1, it will be set to 1 by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#max_rephrase_steps DiscoveryEngineWidgetConfig#max_rephrase_steps}
	MaxRephraseSteps *float64 `field:"optional" json:"maxRephraseSteps" yaml:"maxRephraseSteps"`
	// Text at the beginning of the prompt that instructs the model that generates the answer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#model_prompt_preamble DiscoveryEngineWidgetConfig#model_prompt_preamble}
	ModelPromptPreamble *string `field:"optional" json:"modelPromptPreamble" yaml:"modelPromptPreamble"`
	// The model version used to generate the answer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#model_version DiscoveryEngineWidgetConfig#model_version}
	ModelVersion *string `field:"optional" json:"modelVersion" yaml:"modelVersion"`
	// The number of top results to generate the answer from. Up to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#result_count DiscoveryEngineWidgetConfig#result_count}
	ResultCount *float64 `field:"optional" json:"resultCount" yaml:"resultCount"`
}

