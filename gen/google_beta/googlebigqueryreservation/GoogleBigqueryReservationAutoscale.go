package googlebigqueryreservation


type GoogleBigqueryReservationAutoscale struct {
	// Number of slots to be scaled when needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_reservation#max_slots GoogleBigqueryReservation#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
}

