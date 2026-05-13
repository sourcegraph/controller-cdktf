package googlefirebaseailogicconfig


type GoogleFirebaseAiLogicConfigTrafficFilter struct {
	// Only allows users to use AI Logic via prompt templates for this project.
	//
	// If true, only calls using server templates are permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#template_only GoogleFirebaseAiLogicConfig#template_only}
	TemplateOnly interface{} `field:"optional" json:"templateOnly" yaml:"templateOnly"`
}

