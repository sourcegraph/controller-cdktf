package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWorkloadsConfig struct {
	// scheduler block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_composer_environment#scheduler GoogleComposerEnvironment#scheduler}
	Scheduler *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler `field:"optional" json:"scheduler" yaml:"scheduler"`
	// web_server block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_composer_environment#web_server GoogleComposerEnvironment#web_server}
	WebServer *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer `field:"optional" json:"webServer" yaml:"webServer"`
	// worker block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_composer_environment#worker GoogleComposerEnvironment#worker}
	Worker *GoogleComposerEnvironmentConfigWorkloadsConfigWorker `field:"optional" json:"worker" yaml:"worker"`
}

