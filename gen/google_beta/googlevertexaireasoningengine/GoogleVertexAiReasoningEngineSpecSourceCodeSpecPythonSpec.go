package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec struct {
	// Optional.
	//
	// The Python module to load as the entrypoint,
	// specified as a fully qualified module name. For example:
	// path.to.agent. If not specified, defaults to "agent".
	// The project root will be added to Python sys.path, allowing
	// imports to be specified relative to the root.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#entrypoint_module GoogleVertexAiReasoningEngine#entrypoint_module}
	EntrypointModule *string `field:"optional" json:"entrypointModule" yaml:"entrypointModule"`
	// Optional. The name of the callable object within the entrypointModule to use as the application If not specified, defaults to "root_agent".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#entrypoint_object GoogleVertexAiReasoningEngine#entrypoint_object}
	EntrypointObject *string `field:"optional" json:"entrypointObject" yaml:"entrypointObject"`
	// Optional. The path to the requirements file, relative to the source root. If not specified, defaults to "requirements.txt".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#requirements_file GoogleVertexAiReasoningEngine#requirements_file}
	RequirementsFile *string `field:"optional" json:"requirementsFile" yaml:"requirementsFile"`
	// Optional.
	//
	// The version of Python to use. Support version
	// includes 3.9, 3.10, 3.11, 3.12, 3.13. If not specified,
	// default value is 3.10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#version GoogleVertexAiReasoningEngine#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

