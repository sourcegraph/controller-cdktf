package googlecomputefuturereservation


type GoogleComputeFutureReservationShareSettings struct {
	// project_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#project_map GoogleComputeFutureReservation#project_map}
	ProjectMap interface{} `field:"optional" json:"projectMap" yaml:"projectMap"`
	// list of Project names to specify consumer projects for this shared-reservation.
	//
	// This is only valid when shareType's value is SPECIFIC_PROJECTS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#projects GoogleComputeFutureReservation#projects}
	Projects *[]*string `field:"optional" json:"projects" yaml:"projects"`
	// Type of sharing for this future reservation. Possible values: ["LOCAL", "SPECIFIC_PROJECTS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#share_type GoogleComputeFutureReservation#share_type}
	ShareType *string `field:"optional" json:"shareType" yaml:"shareType"`
}

