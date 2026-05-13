package googledialogflowcxgenerator


type GoogleDialogflowCxGeneratorModelParameter struct {
	// The maximum number of tokens to generate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#max_decode_steps GoogleDialogflowCxGenerator#max_decode_steps}
	MaxDecodeSteps *float64 `field:"optional" json:"maxDecodeSteps" yaml:"maxDecodeSteps"`
	// The temperature used for sampling.
	//
	// Temperature sampling occurs after both topP and topK have been applied.
	// Valid range: [0.0, 1.0] Low temperature = less random. High temperature = more random.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#temperature GoogleDialogflowCxGenerator#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// If set, the sampling process in each step is limited to the topK tokens with highest probabilities.
	//
	// Valid range: [1, 40] or 1000+. Small topK = less random. Large topK = more random.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#top_k GoogleDialogflowCxGenerator#top_k}
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// If set, only the tokens comprising the top topP probability mass are considered.
	//
	// If both topP and topK are set, topP will be used for further refining candidates selected with topK.
	// Valid range: (0.0, 1.0]. Small topP = less random. Large topP = more random.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#top_p GoogleDialogflowCxGenerator#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

