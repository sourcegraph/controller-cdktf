package googlecomputereservation


type GoogleComputeReservationReservationSharingPolicy struct {
	// Sharing config for all Google Cloud services. Possible values: ["ALLOW_ALL", "DISALLOW_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#service_share_type GoogleComputeReservation#service_share_type}
	ServiceShareType *string `field:"optional" json:"serviceShareType" yaml:"serviceShareType"`
}

