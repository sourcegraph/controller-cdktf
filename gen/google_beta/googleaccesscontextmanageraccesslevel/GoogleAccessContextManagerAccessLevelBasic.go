package googleaccesscontextmanageraccesslevel


type GoogleAccessContextManagerAccessLevelBasic struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#conditions GoogleAccessContextManagerAccessLevel#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// How the conditions list should be combined to determine if a request is granted this AccessLevel.
	//
	// If AND is used, each Condition in
	// conditions must be satisfied for the AccessLevel to be applied. If
	// OR is used, at least one Condition in conditions must be satisfied
	// for the AccessLevel to be applied. Default value: "AND" Possible values: ["AND", "OR"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#combining_function GoogleAccessContextManagerAccessLevel#combining_function}
	CombiningFunction *string `field:"optional" json:"combiningFunction" yaml:"combiningFunction"`
}

