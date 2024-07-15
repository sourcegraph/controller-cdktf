package googlediscoveryenginechatengine


type GoogleDiscoveryEngineChatEngineCommonConfig struct {
	// The name of the company, business or entity that is associated with the engine.
	//
	// Setting this may help improve LLM related features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_discovery_engine_chat_engine#company_name GoogleDiscoveryEngineChatEngine#company_name}
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
}

