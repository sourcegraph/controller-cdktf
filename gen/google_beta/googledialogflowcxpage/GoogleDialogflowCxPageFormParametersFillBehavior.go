package googledialogflowcxpage


type GoogleDialogflowCxPageFormParametersFillBehavior struct {
	// initial_prompt_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#initial_prompt_fulfillment GoogleDialogflowCxPage#initial_prompt_fulfillment}
	InitialPromptFulfillment *GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment `field:"optional" json:"initialPromptFulfillment" yaml:"initialPromptFulfillment"`
	// reprompt_event_handlers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#reprompt_event_handlers GoogleDialogflowCxPage#reprompt_event_handlers}
	RepromptEventHandlers interface{} `field:"optional" json:"repromptEventHandlers" yaml:"repromptEventHandlers"`
}

