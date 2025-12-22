package report


type ReportDashboard struct {
	// The ID of the dashboard to be used for the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#id Report#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The query window duration that will be used in the dashboard query in minutes.
	//
	// E.g., if we want the report to contain the last day of data, this needs to be set to 1440.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#query_window_duration_minutes Report#query_window_duration_minutes}
	QueryWindowDurationMinutes *float64 `field:"required" json:"queryWindowDurationMinutes" yaml:"queryWindowDurationMinutes"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#parameters Report#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

