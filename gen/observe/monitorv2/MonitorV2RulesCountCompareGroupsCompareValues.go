package monitorv2


type MonitorV2RulesCountCompareGroupsCompareValues struct {
	// the type of comparison (greater, less, equal, etc.).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#compare_fn MonitorV2#compare_fn}
	CompareFn *string `field:"required" json:"compareFn" yaml:"compareFn"`
	// list of size <=1 consisting of a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_bool MonitorV2#value_bool}
	ValueBool interface{} `field:"optional" json:"valueBool" yaml:"valueBool"`
	// list of size <=1 consisting of a duration value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_duration MonitorV2#value_duration}
	ValueDuration *[]*string `field:"optional" json:"valueDuration" yaml:"valueDuration"`
	// list of size <=1 consisting of a float value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_float64 MonitorV2#value_float64}
	ValueFloat64 *[]*float64 `field:"optional" json:"valueFloat64" yaml:"valueFloat64"`
	// list of size <=1 consisting of an integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_int64 MonitorV2#value_int64}
	ValueInt64 *[]*float64 `field:"optional" json:"valueInt64" yaml:"valueInt64"`
	// list of size <=1 consisting of a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_string MonitorV2#value_string}
	ValueString *[]*string `field:"optional" json:"valueString" yaml:"valueString"`
	// list of size <=1 consisting of a timestamp value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value_timestamp MonitorV2#value_timestamp}
	ValueTimestamp *[]*string `field:"optional" json:"valueTimestamp" yaml:"valueTimestamp"`
}

