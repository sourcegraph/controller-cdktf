package monitor


type MonitorRuleFacet struct {
	// Comparison function to use when comparing the field against the desired value(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#facet_function Monitor#facet_function}
	FacetFunction *string `field:"required" json:"facetFunction" yaml:"facetFunction"`
	// The values to compare the field against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#facet_values Monitor#facet_values}
	FacetValues *[]*string `field:"required" json:"facetValues" yaml:"facetValues"`
	// Time window to evaluate time_function over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#lookback_time Monitor#lookback_time}
	LookbackTime *string `field:"required" json:"lookbackTime" yaml:"lookbackTime"`
	// Temporal condition to evaluate the matches against (e.g. "at least once in window").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#time_function Monitor#time_function}
	TimeFunction *string `field:"required" json:"timeFunction" yaml:"timeFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#time_value Monitor#time_value}.
	TimeValue *float64 `field:"optional" json:"timeValue" yaml:"timeValue"`
}

