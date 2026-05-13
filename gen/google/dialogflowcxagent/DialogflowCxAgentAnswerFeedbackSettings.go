package dialogflowcxagent


type DialogflowCxAgentAnswerFeedbackSettings struct {
	// If enabled, end users will be able to provide [answer feedback](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.sessions/submitAnswerFeedback#body.AnswerFeedback) to Dialogflow responses. Feature works only if interaction logging is enabled in the Dialogflow agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_agent#enable_answer_feedback DialogflowCxAgent#enable_answer_feedback}
	EnableAnswerFeedback interface{} `field:"optional" json:"enableAnswerFeedback" yaml:"enableAnswerFeedback"`
}

