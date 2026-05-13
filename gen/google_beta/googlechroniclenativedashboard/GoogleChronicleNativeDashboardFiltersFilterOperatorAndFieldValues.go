package googlechroniclenativedashboard


type GoogleChronicleNativeDashboardFiltersFilterOperatorAndFieldValues struct {
	// The values for the modifier. All operators should have a single value other than 'IN' and 'BETWEEN'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#field_values GoogleChronicleNativeDashboard#field_values}
	FieldValues *[]*string `field:"optional" json:"fieldValues" yaml:"fieldValues"`
	// The operator to apply to the field.
	//
	// Possible values: ["EQUAL", "NOT_EQUAL", "IN", "GREATER_THAN", "GREATER_THAN_OR_EQUAL_TO", "LESS_THAN", "LESS_THAN_OR_EQUAL_TO", "BETWEEN", "PAST", "IS_NULL", "IS_NOT_NULL", "STARTS_WITH", "ENDS_WITH", "DOES_NOT_STARTS_WITH", "DOES_NOT_ENDS_WITH", "NOT_IN", "CONTAINS", "DOES_NOT_CONTAIN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#filter_operator GoogleChronicleNativeDashboard#filter_operator}
	FilterOperator *string `field:"optional" json:"filterOperator" yaml:"filterOperator"`
}

