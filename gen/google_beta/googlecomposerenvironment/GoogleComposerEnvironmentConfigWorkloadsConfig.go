package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWorkloadsConfig struct {
	// scheduler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_composer_environment#scheduler GoogleComposerEnvironment#scheduler}
	Scheduler *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler `field:"optional" json:"scheduler" yaml:"scheduler"`
	// triggerer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_composer_environment#triggerer GoogleComposerEnvironment#triggerer}
	Triggerer *GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer `field:"optional" json:"triggerer" yaml:"triggerer"`
	// web_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_composer_environment#web_server GoogleComposerEnvironment#web_server}
	WebServer *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer `field:"optional" json:"webServer" yaml:"webServer"`
	// worker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_composer_environment#worker GoogleComposerEnvironment#worker}
	Worker *GoogleComposerEnvironmentConfigWorkloadsConfigWorker `field:"optional" json:"worker" yaml:"worker"`
}

