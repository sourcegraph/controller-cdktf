package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity struct {
	// Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#consume_reservation_type GoogleDataprocWorkflowTemplate#consume_reservation_type}
	ConsumeReservationType *string `field:"optional" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Optional. Corresponds to the label key of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#key GoogleDataprocWorkflowTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Optional. Corresponds to the label values of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#values GoogleDataprocWorkflowTemplate#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

