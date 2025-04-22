package googlecolabruntimetemplate


type GoogleColabRuntimeTemplateSoftwareConfig struct {
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime_template#env GoogleColabRuntimeTemplate#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// post_startup_script_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime_template#post_startup_script_config GoogleColabRuntimeTemplate#post_startup_script_config}
	PostStartupScriptConfig *GoogleColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig `field:"optional" json:"postStartupScriptConfig" yaml:"postStartupScriptConfig"`
}

