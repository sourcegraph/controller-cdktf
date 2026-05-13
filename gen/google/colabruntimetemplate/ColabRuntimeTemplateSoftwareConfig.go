package colabruntimetemplate


type ColabRuntimeTemplateSoftwareConfig struct {
	// colab_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_runtime_template#colab_image ColabRuntimeTemplate#colab_image}
	ColabImage *ColabRuntimeTemplateSoftwareConfigColabImage `field:"optional" json:"colabImage" yaml:"colabImage"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_runtime_template#env ColabRuntimeTemplate#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// post_startup_script_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_runtime_template#post_startup_script_config ColabRuntimeTemplate#post_startup_script_config}
	PostStartupScriptConfig *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig `field:"optional" json:"postStartupScriptConfig" yaml:"postStartupScriptConfig"`
}

