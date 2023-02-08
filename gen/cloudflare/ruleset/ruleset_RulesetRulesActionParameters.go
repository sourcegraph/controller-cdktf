package ruleset


type RulesetRulesActionParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cookie_fields Ruleset#cookie_fields}.
	CookieFields *[]*string `field:"optional" json:"cookieFields" yaml:"cookieFields"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#headers Ruleset#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#host_header Ruleset#host_header}.
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#id Ruleset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#increment Ruleset#increment}.
	Increment *float64 `field:"optional" json:"increment" yaml:"increment"`
	// matched_data block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#matched_data Ruleset#matched_data}
	MatchedData *RulesetRulesActionParametersMatchedData `field:"optional" json:"matchedData" yaml:"matchedData"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#origin Ruleset#origin}
	Origin *RulesetRulesActionParametersOrigin `field:"optional" json:"origin" yaml:"origin"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#overrides Ruleset#overrides}
	Overrides *RulesetRulesActionParametersOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#phases Ruleset#phases}.
	Phases *[]*string `field:"optional" json:"phases" yaml:"phases"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#products Ruleset#products}.
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#request_fields Ruleset#request_fields}.
	RequestFields *[]*string `field:"optional" json:"requestFields" yaml:"requestFields"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#response Ruleset#response}
	Response interface{} `field:"optional" json:"response" yaml:"response"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#response_fields Ruleset#response_fields}.
	ResponseFields *[]*string `field:"optional" json:"responseFields" yaml:"responseFields"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#rules Ruleset#rules}.
	Rules *map[string]*string `field:"optional" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#ruleset Ruleset#ruleset}.
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#rulesets Ruleset#rulesets}.
	Rulesets *[]*string `field:"optional" json:"rulesets" yaml:"rulesets"`
	// uri block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#uri Ruleset#uri}
	Uri *RulesetRulesActionParametersUri `field:"optional" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#version Ruleset#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

