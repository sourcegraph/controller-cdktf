package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpec struct {
	// developer_connect_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#developer_connect_source GoogleVertexAiReasoningEngine#developer_connect_source}
	DeveloperConnectSource *GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource `field:"optional" json:"developerConnectSource" yaml:"developerConnectSource"`
	// image_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#image_spec GoogleVertexAiReasoningEngine#image_spec}
	ImageSpec *GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec `field:"optional" json:"imageSpec" yaml:"imageSpec"`
	// inline_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#inline_source GoogleVertexAiReasoningEngine#inline_source}
	InlineSource *GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource `field:"optional" json:"inlineSource" yaml:"inlineSource"`
	// python_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#python_spec GoogleVertexAiReasoningEngine#python_spec}
	PythonSpec *GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec `field:"optional" json:"pythonSpec" yaml:"pythonSpec"`
}

