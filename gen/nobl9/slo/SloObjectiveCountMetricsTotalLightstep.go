package slo


type SloObjectiveCountMetricsTotalLightstep struct {
	// Type of data to filter by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#type_of_data Slo#type_of_data}
	TypeOfData *string `field:"required" json:"typeOfData" yaml:"typeOfData"`
	// Optional value to filter by percentiles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#percentile Slo#percentile}
	Percentile *float64 `field:"optional" json:"percentile" yaml:"percentile"`
	// ID of the metrics stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#stream_id Slo#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// UQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#uql Slo#uql}
	Uql *string `field:"optional" json:"uql" yaml:"uql"`
}

