package directsplunk


type DirectSplunkHistoricalDataRetrievalDefaultDuration struct {
	// Must be one of Minute, Hour, or Day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk#unit DirectSplunk#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Must be an integer greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_splunk#value DirectSplunk#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

