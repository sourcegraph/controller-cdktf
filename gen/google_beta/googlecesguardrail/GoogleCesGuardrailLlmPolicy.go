package googlecesguardrail


type GoogleCesGuardrailLlmPolicy struct {
	// Defines when to apply the policy check during the conversation.
	//
	// If set to
	// 'POLICY_SCOPE_UNSPECIFIED', the policy will be applied to the user input.
	// When applying the policy to the agent response, additional latency will
	// be introduced before the agent can respond.
	// Possible values:
	// USER_QUERY
	// AGENT_RESPONSE
	// USER_QUERY_AND_AGENT_RESPONSE Possible values: ["USER_QUERY", "AGENT_RESPONSE", "USER_QUERY_AND_AGENT_RESPONSE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#policy_scope GoogleCesGuardrail#policy_scope}
	PolicyScope *string `field:"required" json:"policyScope" yaml:"policyScope"`
	// Policy prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#prompt GoogleCesGuardrail#prompt}
	Prompt *string `field:"required" json:"prompt" yaml:"prompt"`
	// By default, the LLM policy check is bypassed for short utterances.
	//
	// Enabling this setting applies the policy check to all utterances,
	// including those that would normally be skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#allow_short_utterance GoogleCesGuardrail#allow_short_utterance}
	AllowShortUtterance interface{} `field:"optional" json:"allowShortUtterance" yaml:"allowShortUtterance"`
	// If an error occurs during the policy check, fail open and do not trigger the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#fail_open GoogleCesGuardrail#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// When checking this policy, consider the last 'n' messages in the conversation.
	//
	// When not set a default value of 10 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#max_conversation_messages GoogleCesGuardrail#max_conversation_messages}
	MaxConversationMessages *float64 `field:"optional" json:"maxConversationMessages" yaml:"maxConversationMessages"`
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#model_settings GoogleCesGuardrail#model_settings}
	ModelSettings *GoogleCesGuardrailLlmPolicyModelSettings `field:"optional" json:"modelSettings" yaml:"modelSettings"`
}

