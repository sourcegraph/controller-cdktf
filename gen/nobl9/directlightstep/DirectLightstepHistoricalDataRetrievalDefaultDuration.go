package directlightstep


type DirectLightstepHistoricalDataRetrievalDefaultDuration struct {
	// Must be one of Minute, Hour, or Day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_lightstep#unit DirectLightstep#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Must be an integer greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_lightstep#value DirectLightstep#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

