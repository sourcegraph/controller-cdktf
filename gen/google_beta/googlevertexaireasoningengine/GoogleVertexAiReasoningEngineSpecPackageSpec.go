package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecPackageSpec struct {
	// Optional. The Cloud Storage URI of the dependency files in tar.gz format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#dependency_files_gcs_uri GoogleVertexAiReasoningEngine#dependency_files_gcs_uri}
	DependencyFilesGcsUri *string `field:"optional" json:"dependencyFilesGcsUri" yaml:"dependencyFilesGcsUri"`
	// Optional. The Cloud Storage URI of the pickled python object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#pickle_object_gcs_uri GoogleVertexAiReasoningEngine#pickle_object_gcs_uri}
	PickleObjectGcsUri *string `field:"optional" json:"pickleObjectGcsUri" yaml:"pickleObjectGcsUri"`
	// Optional. The Python version. Currently support 3.8, 3.9, 3.10, 3.11, 3.12, 3.13. If not specified, default value is 3.10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#python_version GoogleVertexAiReasoningEngine#python_version}
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Optional. The Cloud Storage URI of the requirements.txtfile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#requirements_gcs_uri GoogleVertexAiReasoningEngine#requirements_gcs_uri}
	RequirementsGcsUri *string `field:"optional" json:"requirementsGcsUri" yaml:"requirementsGcsUri"`
}

