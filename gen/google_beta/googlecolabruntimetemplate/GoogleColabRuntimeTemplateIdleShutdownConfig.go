package googlecolabruntimetemplate


type GoogleColabRuntimeTemplateIdleShutdownConfig struct {
	// The duration after which the runtime is automatically shut down.
	//
	// An input of 0s disables the idle shutdown feature, and a valid range is [10m, 24h].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_runtime_template#idle_timeout GoogleColabRuntimeTemplate#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

