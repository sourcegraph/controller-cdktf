package googlecomputefuturereservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeFutureReservationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the las
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#name GoogleComputeFutureReservation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#time_window GoogleComputeFutureReservation#time_window}
	TimeWindow *GoogleComputeFutureReservationTimeWindow `field:"required" json:"timeWindow" yaml:"timeWindow"`
	// aggregate_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#aggregate_reservation GoogleComputeFutureReservation#aggregate_reservation}
	AggregateReservation *GoogleComputeFutureReservationAggregateReservation `field:"optional" json:"aggregateReservation" yaml:"aggregateReservation"`
	// Future timestamp when the FR auto-created reservations will be deleted by Compute Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#auto_created_reservations_delete_time GoogleComputeFutureReservation#auto_created_reservations_delete_time}
	AutoCreatedReservationsDeleteTime *string `field:"optional" json:"autoCreatedReservationsDeleteTime" yaml:"autoCreatedReservationsDeleteTime"`
	// auto_created_reservations_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#auto_created_reservations_duration GoogleComputeFutureReservation#auto_created_reservations_duration}
	AutoCreatedReservationsDuration *GoogleComputeFutureReservationAutoCreatedReservationsDuration `field:"optional" json:"autoCreatedReservationsDuration" yaml:"autoCreatedReservationsDuration"`
	// Setting for enabling or disabling automatic deletion for auto-created reservation.
	//
	// If set to true, auto-created reservations will be deleted at Future Reservation's end time (default) or at user's defined timestamp if any of the [autoCreatedReservationsDeleteTime, autoCreatedReservationsDuration] values is specified. For keeping auto-created reservation indefinitely, this value should be set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#auto_delete_auto_created_reservations GoogleComputeFutureReservation#auto_delete_auto_created_reservations}
	AutoDeleteAutoCreatedReservations interface{} `field:"optional" json:"autoDeleteAutoCreatedReservations" yaml:"autoDeleteAutoCreatedReservations"`
	// commitment_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#commitment_info GoogleComputeFutureReservation#commitment_info}
	CommitmentInfo *GoogleComputeFutureReservationCommitmentInfo `field:"optional" json:"commitmentInfo" yaml:"commitmentInfo"`
	// Type of the deployment requested as part of future reservation. Possible values: ["DENSE", "FLEXIBLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#deployment_type GoogleComputeFutureReservation#deployment_type}
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#description GoogleComputeFutureReservation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#id GoogleComputeFutureReservation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name prefix for the reservations to be created at the time of delivery.
	//
	// The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#name_prefix GoogleComputeFutureReservation#name_prefix}
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Planning state before being submitted for evaluation Possible values: ["DRAFT", "SUBMITTED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#planning_status GoogleComputeFutureReservation#planning_status}
	PlanningStatus *string `field:"optional" json:"planningStatus" yaml:"planningStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#project GoogleComputeFutureReservation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The reservation mode which determines reservation-termination behavior and expected pricing. Possible values: ["CALENDAR", "DEFAULT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#reservation_mode GoogleComputeFutureReservation#reservation_mode}
	ReservationMode *string `field:"optional" json:"reservationMode" yaml:"reservationMode"`
	// Name of reservations where the capacity is provisioned at the time of delivery of future reservations.
	//
	// If the reservation with the given name does not exist already, it is created automatically at the time of Approval with INACTIVE state till specified start-time. Either provide the reservationName or a namePrefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#reservation_name GoogleComputeFutureReservation#reservation_name}
	ReservationName *string `field:"optional" json:"reservationName" yaml:"reservationName"`
	// Maintenance information for this reservation Possible values: ["GROUPED", "INDEPENDENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#scheduling_type GoogleComputeFutureReservation#scheduling_type}
	SchedulingType *string `field:"optional" json:"schedulingType" yaml:"schedulingType"`
	// share_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#share_settings GoogleComputeFutureReservation#share_settings}
	ShareSettings *GoogleComputeFutureReservationShareSettings `field:"optional" json:"shareSettings" yaml:"shareSettings"`
	// Indicates whether the auto-created reservation can be consumed by VMs with affinity for "any" reservation.
	//
	// If the field is set, then only VMs that target the reservation by name can consume from the delivered reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#specific_reservation_required GoogleComputeFutureReservation#specific_reservation_required}
	SpecificReservationRequired interface{} `field:"optional" json:"specificReservationRequired" yaml:"specificReservationRequired"`
	// specific_sku_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#specific_sku_properties GoogleComputeFutureReservation#specific_sku_properties}
	SpecificSkuProperties *GoogleComputeFutureReservationSpecificSkuProperties `field:"optional" json:"specificSkuProperties" yaml:"specificSkuProperties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#timeouts GoogleComputeFutureReservation#timeouts}
	Timeouts *GoogleComputeFutureReservationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

