package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsUserInput struct {
	// Audio data from the end user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#audio GoogleCesEvaluation#audio}
	Audio *string `field:"optional" json:"audio" yaml:"audio"`
	// blob block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#blob GoogleCesEvaluation#blob}
	Blob *GoogleCesEvaluationGoldenTurnsStepsUserInputBlob `field:"optional" json:"blob" yaml:"blob"`
	// DTMF digits from the end user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#dtmf GoogleCesEvaluation#dtmf}
	Dtmf *string `field:"optional" json:"dtmf" yaml:"dtmf"`
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#event GoogleCesEvaluation#event}
	Event *GoogleCesEvaluationGoldenTurnsStepsUserInputEvent `field:"optional" json:"event" yaml:"event"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#image GoogleCesEvaluation#image}
	Image *GoogleCesEvaluationGoldenTurnsStepsUserInputImage `field:"optional" json:"image" yaml:"image"`
	// Natural language query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#text GoogleCesEvaluation#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// tool_responses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tool_responses GoogleCesEvaluation#tool_responses}
	ToolResponses *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponses `field:"optional" json:"toolResponses" yaml:"toolResponses"`
	// Map of variables to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#variables GoogleCesEvaluation#variables}
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
	// Whether the session should continue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#will_continue GoogleCesEvaluation#will_continue}
	WillContinue interface{} `field:"optional" json:"willContinue" yaml:"willContinue"`
}

