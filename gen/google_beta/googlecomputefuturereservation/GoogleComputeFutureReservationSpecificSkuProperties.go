package googlecomputefuturereservation


type GoogleComputeFutureReservationSpecificSkuProperties struct {
	// instance_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#instance_properties GoogleComputeFutureReservation#instance_properties}
	InstanceProperties *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties `field:"optional" json:"instanceProperties" yaml:"instanceProperties"`
	// The instance template that will be used to populate the ReservedInstanceProperties of the future reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#source_instance_template GoogleComputeFutureReservation#source_instance_template}
	SourceInstanceTemplate *string `field:"optional" json:"sourceInstanceTemplate" yaml:"sourceInstanceTemplate"`
	// Total number of instances for which capacity assurance is requested at a future time period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#total_count GoogleComputeFutureReservation#total_count}
	TotalCount *string `field:"optional" json:"totalCount" yaml:"totalCount"`
}

