package googlecontactcenterinsightsautolabelingrule


type GoogleContactCenterInsightsAutoLabelingRuleConditions struct {
	// A optional CEL expression to be evaluated as a boolean value.
	//
	// Once evaluated as true, then we will proceed with the value evaluation.
	// An empty condition will be auto evaluated as true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_auto_labeling_rule#condition GoogleContactCenterInsightsAutoLabelingRule#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// CEL expression to be evaluated as the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_auto_labeling_rule#value GoogleContactCenterInsightsAutoLabelingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

