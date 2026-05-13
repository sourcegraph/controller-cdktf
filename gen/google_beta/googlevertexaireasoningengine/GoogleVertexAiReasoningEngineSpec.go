package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpec struct {
	// Optional. The OSS agent framework used to develop the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#agent_framework GoogleVertexAiReasoningEngine#agent_framework}
	AgentFramework *string `field:"optional" json:"agentFramework" yaml:"agentFramework"`
	// Optional. Declarations for object class methods in OpenAPI specification format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#class_methods GoogleVertexAiReasoningEngine#class_methods}
	ClassMethods *string `field:"optional" json:"classMethods" yaml:"classMethods"`
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#container_spec GoogleVertexAiReasoningEngine#container_spec}
	ContainerSpec *GoogleVertexAiReasoningEngineSpecContainerSpec `field:"optional" json:"containerSpec" yaml:"containerSpec"`
	// deployment_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#deployment_spec GoogleVertexAiReasoningEngine#deployment_spec}
	DeploymentSpec *GoogleVertexAiReasoningEngineSpecDeploymentSpec `field:"optional" json:"deploymentSpec" yaml:"deploymentSpec"`
	// Optional.
	//
	// The identity type to use for the Reasoning Engine.
	// If not specified, the 'service_account' field will be used if set,
	// otherwise the default Vertex AI Reasoning Engine Service Agent in the project will be used.
	// Possible values:
	// * 'SERVICE_ACCOUNT': Use a custom service account if the 'service_account' field is set, otherwise use the default Vertex AI Reasoning Engine Service Agent in the project.
	// * 'AGENT_IDENTITY': Use Agent Identity. The 'service_account' field must not be set. Possible values: ["SERVICE_ACCOUNT", "AGENT_IDENTITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#identity_type GoogleVertexAiReasoningEngine#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// package_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#package_spec GoogleVertexAiReasoningEngine#package_spec}
	PackageSpec *GoogleVertexAiReasoningEngineSpecPackageSpec `field:"optional" json:"packageSpec" yaml:"packageSpec"`
	// Optional.
	//
	// The service account that the Reasoning Engine artifact runs
	// as. It should have "roles/storage.objectViewer" for reading the user
	// project's Cloud Storage and "roles/aiplatform.user" for using Vertex
	// extensions. If not specified, the Vertex AI Reasoning Engine service
	// Agent in the project will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#service_account GoogleVertexAiReasoningEngine#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// source_code_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#source_code_spec GoogleVertexAiReasoningEngine#source_code_spec}
	SourceCodeSpec *GoogleVertexAiReasoningEngineSpecSourceCodeSpec `field:"optional" json:"sourceCodeSpec" yaml:"sourceCodeSpec"`
}

