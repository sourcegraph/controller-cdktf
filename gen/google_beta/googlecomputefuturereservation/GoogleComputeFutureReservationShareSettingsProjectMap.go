package googlecomputefuturereservation


type GoogleComputeFutureReservationShareSettingsProjectMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#id GoogleComputeFutureReservation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The project ID, should be same as the key of this project config in the parent map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#project_id GoogleComputeFutureReservation#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

