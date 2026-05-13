package discoveryenginewidgetconfig


type DiscoveryEngineWidgetConfigUiSettings struct {
	// data_store_ui_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#data_store_ui_configs DiscoveryEngineWidgetConfig#data_store_ui_configs}
	DataStoreUiConfigs interface{} `field:"optional" json:"dataStoreUiConfigs" yaml:"dataStoreUiConfigs"`
	// The default ordering for search results if specified. Used to set SearchRequest#orderBy on applicable requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#default_search_request_order_by DiscoveryEngineWidgetConfig#default_search_request_order_by}
	DefaultSearchRequestOrderBy *string `field:"optional" json:"defaultSearchRequestOrderBy" yaml:"defaultSearchRequestOrderBy"`
	// If set to true, the widget will not collect user events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#disable_user_events_collection DiscoveryEngineWidgetConfig#disable_user_events_collection}
	DisableUserEventsCollection interface{} `field:"optional" json:"disableUserEventsCollection" yaml:"disableUserEventsCollection"`
	// Whether or not to enable autocomplete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_autocomplete DiscoveryEngineWidgetConfig#enable_autocomplete}
	EnableAutocomplete interface{} `field:"optional" json:"enableAutocomplete" yaml:"enableAutocomplete"`
	// If set to true, the widget will enable the create agent button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_create_agent_button DiscoveryEngineWidgetConfig#enable_create_agent_button}
	EnableCreateAgentButton interface{} `field:"optional" json:"enableCreateAgentButton" yaml:"enableCreateAgentButton"`
	// If set to true, the widget will enable people search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_people_search DiscoveryEngineWidgetConfig#enable_people_search}
	EnablePeopleSearch interface{} `field:"optional" json:"enablePeopleSearch" yaml:"enablePeopleSearch"`
	// Turn on or off collecting the search result quality feedback from end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_quality_feedback DiscoveryEngineWidgetConfig#enable_quality_feedback}
	EnableQualityFeedback interface{} `field:"optional" json:"enableQualityFeedback" yaml:"enableQualityFeedback"`
	// Whether to enable safe search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_safe_search DiscoveryEngineWidgetConfig#enable_safe_search}
	EnableSafeSearch interface{} `field:"optional" json:"enableSafeSearch" yaml:"enableSafeSearch"`
	// Whether to enable search-as-you-type behavior for the search widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_search_as_you_type DiscoveryEngineWidgetConfig#enable_search_as_you_type}
	EnableSearchAsYouType interface{} `field:"optional" json:"enableSearchAsYouType" yaml:"enableSearchAsYouType"`
	// If set to true, the widget will enable visual content summary on applicable search requests. Only used by healthcare search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#enable_visual_content_summary DiscoveryEngineWidgetConfig#enable_visual_content_summary}
	EnableVisualContentSummary interface{} `field:"optional" json:"enableVisualContentSummary" yaml:"enableVisualContentSummary"`
	// generative_answer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#generative_answer_config DiscoveryEngineWidgetConfig#generative_answer_config}
	GenerativeAnswerConfig *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig `field:"optional" json:"generativeAnswerConfig" yaml:"generativeAnswerConfig"`
	// Describes widget (or web app) interaction type Possible values: ["SEARCH_ONLY", "SEARCH_WITH_ANSWER", "SEARCH_WITH_FOLLOW_UPS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#interaction_type DiscoveryEngineWidgetConfig#interaction_type}
	InteractionType *string `field:"optional" json:"interactionType" yaml:"interactionType"`
	// Controls whether result extract is display and how (snippet or extractive answer).
	//
	// Default to no result if unspecified. Possible values: ["SNIPPET", "EXTRACTIVE_ANSWER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#result_description_type DiscoveryEngineWidgetConfig#result_description_type}
	ResultDescriptionType *string `field:"optional" json:"resultDescriptionType" yaml:"resultDescriptionType"`
}

