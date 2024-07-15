package googledialogflowcxflow


type GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillment struct {
	// conditional_cases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#conditional_cases GoogleDialogflowCxFlow#conditional_cases}
	ConditionalCases interface{} `field:"optional" json:"conditionalCases" yaml:"conditionalCases"`
	// messages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#messages GoogleDialogflowCxFlow#messages}
	Messages interface{} `field:"optional" json:"messages" yaml:"messages"`
	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs.
	//
	// If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#return_partial_responses GoogleDialogflowCxFlow#return_partial_responses}
	ReturnPartialResponses interface{} `field:"optional" json:"returnPartialResponses" yaml:"returnPartialResponses"`
	// set_parameter_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#set_parameter_actions GoogleDialogflowCxFlow#set_parameter_actions}
	SetParameterActions interface{} `field:"optional" json:"setParameterActions" yaml:"setParameterActions"`
	// The tag used by the webhook to identify which fulfillment is being called.
	//
	// This field is required if webhook is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#tag GoogleDialogflowCxFlow#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#webhook GoogleDialogflowCxFlow#webhook}
	Webhook *string `field:"optional" json:"webhook" yaml:"webhook"`
}

