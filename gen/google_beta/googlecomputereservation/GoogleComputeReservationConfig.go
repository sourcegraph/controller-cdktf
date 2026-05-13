package googlecomputereservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeReservationConfig struct {
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
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#name GoogleComputeReservation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#specific_reservation GoogleComputeReservation#specific_reservation}
	SpecificReservation *GoogleComputeReservationSpecificReservation `field:"required" json:"specificReservation" yaml:"specificReservation"`
	// The zone where the reservation is made.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#zone GoogleComputeReservation#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// delete_after_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#delete_after_duration GoogleComputeReservation#delete_after_duration}
	DeleteAfterDuration *GoogleComputeReservationDeleteAfterDuration `field:"optional" json:"deleteAfterDuration" yaml:"deleteAfterDuration"`
	// Absolute time in future when the reservation will be auto-deleted by Compute Engine.
	//
	// Timestamp is represented in RFC3339 text format.
	// Cannot be used with delete_after_duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#delete_at_time GoogleComputeReservation#delete_at_time}
	DeleteAtTime *string `field:"optional" json:"deleteAtTime" yaml:"deleteAtTime"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#description GoogleComputeReservation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates if this group of VMs have emergent maintenance enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#enable_emergent_maintenance GoogleComputeReservation#enable_emergent_maintenance}
	EnableEmergentMaintenance interface{} `field:"optional" json:"enableEmergentMaintenance" yaml:"enableEmergentMaintenance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#project GoogleComputeReservation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// reservation_sharing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#reservation_sharing_policy GoogleComputeReservation#reservation_sharing_policy}
	ReservationSharingPolicy *GoogleComputeReservationReservationSharingPolicy `field:"optional" json:"reservationSharingPolicy" yaml:"reservationSharingPolicy"`
	// share_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#share_settings GoogleComputeReservation#share_settings}
	ShareSettings *GoogleComputeReservationShareSettings `field:"optional" json:"shareSettings" yaml:"shareSettings"`
	// When set to true, only VMs that target this reservation by name can consume this reservation.
	//
	// Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#specific_reservation_required GoogleComputeReservation#specific_reservation_required}
	SpecificReservationRequired interface{} `field:"optional" json:"specificReservationRequired" yaml:"specificReservationRequired"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#timeouts GoogleComputeReservation#timeouts}
	Timeouts *GoogleComputeReservationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

