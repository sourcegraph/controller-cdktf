package directazuremonitor


type DirectAzureMonitorHistoricalDataRetrievalDefaultDuration struct {
	// Must be one of Minute, Hour, or Day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_azure_monitor#unit DirectAzureMonitor#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Must be an integer greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_azure_monitor#value DirectAzureMonitor#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

