package cestool


type CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecs struct {
	// An expression which specifies a boost condition.
	//
	// The syntax is the same
	// as filter expression syntax. Currently, the only supported condition is
	// a list of BCP-47 lang codes.
	// Example: To boost suggestions in languages en or fr:
	// (lang_code: ANY("en", "fr"))
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#condition CesTool#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Strength of the boost, which should be in [-1, 1].
	//
	// Negative boost means
	// demotion. Default is 0.0.
	// Setting to 1.0 gives the suggestions a big promotion. However, it does
	// not necessarily mean that the top result will be a boosted suggestion.
	// Setting to -1.0 gives the suggestions a big demotion. However, other
	// suggestions that are relevant might still be shown.
	// Setting to 0.0 means no boost applied. The boosting condition is
	// ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#boost CesTool#boost}
	Boost *float64 `field:"optional" json:"boost" yaml:"boost"`
	// boost_control_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#boost_control_spec CesTool#boost_control_spec}
	BoostControlSpec *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec `field:"optional" json:"boostControlSpec" yaml:"boostControlSpec"`
}

