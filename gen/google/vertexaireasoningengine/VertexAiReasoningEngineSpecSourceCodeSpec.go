package vertexaireasoningengine


type VertexAiReasoningEngineSpecSourceCodeSpec struct {
	// developer_connect_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#developer_connect_source VertexAiReasoningEngine#developer_connect_source}
	DeveloperConnectSource *VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource `field:"optional" json:"developerConnectSource" yaml:"developerConnectSource"`
	// image_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#image_spec VertexAiReasoningEngine#image_spec}
	ImageSpec *VertexAiReasoningEngineSpecSourceCodeSpecImageSpec `field:"optional" json:"imageSpec" yaml:"imageSpec"`
	// inline_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#inline_source VertexAiReasoningEngine#inline_source}
	InlineSource *VertexAiReasoningEngineSpecSourceCodeSpecInlineSource `field:"optional" json:"inlineSource" yaml:"inlineSource"`
	// python_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#python_spec VertexAiReasoningEngine#python_spec}
	PythonSpec *VertexAiReasoningEngineSpecSourceCodeSpecPythonSpec `field:"optional" json:"pythonSpec" yaml:"pythonSpec"`
}

