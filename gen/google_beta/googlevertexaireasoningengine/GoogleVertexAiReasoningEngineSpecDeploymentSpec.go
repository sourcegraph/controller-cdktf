package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpec struct {
	// Optional. Concurrency for each container and agent server. Recommended value: 2 * cpu + 1. Defaults to 9.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#container_concurrency GoogleVertexAiReasoningEngine#container_concurrency}
	ContainerConcurrency *float64 `field:"optional" json:"containerConcurrency" yaml:"containerConcurrency"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#env GoogleVertexAiReasoningEngine#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// Optional.
	//
	// The maximum number of application instances that can be
	// launched to handle increased traffic. Defaults to 100.
	// Range: [1, 1000]. If VPC-SC or PSC-I is enabled, the acceptable
	// range is [1, 100].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#max_instances GoogleVertexAiReasoningEngine#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Optional.
	//
	// The minimum number of application instances that will be
	// kept running at all times. Defaults to 1. Range: [0, 10].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#min_instances GoogleVertexAiReasoningEngine#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// psc_interface_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#psc_interface_config GoogleVertexAiReasoningEngine#psc_interface_config}
	PscInterfaceConfig *GoogleVertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig `field:"optional" json:"pscInterfaceConfig" yaml:"pscInterfaceConfig"`
	// Optional. Resource limits for each container. Only 'cpu' and 'memory' keys are supported.
	//
	// Defaults to {"cpu": "4", "memory": "4Gi"}.
	//
	// The only supported values for CPU are '1', '2', '4', '6' and '8'.
	// For more information, go to
	// https://cloud.google.com/run/docs/configuring/cpu.
	//
	// The only supported values for memory are '1Gi', '2Gi', ... '32 Gi'.
	// For more information, go to
	// https://cloud.google.com/run/docs/configuring/memory-limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#resource_limits GoogleVertexAiReasoningEngine#resource_limits}
	ResourceLimits *map[string]*string `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// secret_env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#secret_env GoogleVertexAiReasoningEngine#secret_env}
	SecretEnv interface{} `field:"optional" json:"secretEnv" yaml:"secretEnv"`
}

