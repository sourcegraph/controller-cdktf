package cesguardrail


type CesGuardrailLlmPolicy struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#policy_scope CesGuardrail#policy_scope}
	PolicyScope *string `field:"required" json:"policyScope" yaml:"policyScope"`
	// Policy prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#prompt CesGuardrail#prompt}
	Prompt *string `field:"required" json:"prompt" yaml:"prompt"`
	// By default, the LLM policy check is bypassed for short utterances.
	//
	// Enabling this setting applies the policy check to all utterances,
	// including those that would normally be skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#allow_short_utterance CesGuardrail#allow_short_utterance}
	AllowShortUtterance interface{} `field:"optional" json:"allowShortUtterance" yaml:"allowShortUtterance"`
	// If an error occurs during the policy check, fail open and do not trigger the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#fail_open CesGuardrail#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// When checking this policy, consider the last 'n' messages in the conversation.
	//
	// When not set a default value of 10 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#max_conversation_messages CesGuardrail#max_conversation_messages}
	MaxConversationMessages *float64 `field:"optional" json:"maxConversationMessages" yaml:"maxConversationMessages"`
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#model_settings CesGuardrail#model_settings}
	ModelSettings *CesGuardrailLlmPolicyModelSettings `field:"optional" json:"modelSettings" yaml:"modelSettings"`
}

