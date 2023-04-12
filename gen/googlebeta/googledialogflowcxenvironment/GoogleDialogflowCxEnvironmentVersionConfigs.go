package googledialogflowcxenvironment


type GoogleDialogflowCxEnvironmentVersionConfigs struct {
	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_cx_environment#version GoogleDialogflowCxEnvironment#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

