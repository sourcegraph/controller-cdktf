package googlecomputereservation


type GoogleComputeReservationShareSettings struct {
	// project_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#project_map GoogleComputeReservation#project_map}
	ProjectMap interface{} `field:"optional" json:"projectMap" yaml:"projectMap"`
	// Type of sharing for this shared-reservation Possible values: ["LOCAL", "SPECIFIC_PROJECTS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#share_type GoogleComputeReservation#share_type}
	ShareType *string `field:"optional" json:"shareType" yaml:"shareType"`
}

