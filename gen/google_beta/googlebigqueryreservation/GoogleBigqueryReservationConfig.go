package googlebigqueryreservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryReservationConfig struct {
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
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#name GoogleBigqueryReservation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Minimum slots available to this reservation.
	//
	// A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#slot_capacity GoogleBigqueryReservation#slot_capacity}
	SlotCapacity *float64 `field:"required" json:"slotCapacity" yaml:"slotCapacity"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#autoscale GoogleBigqueryReservation#autoscale}
	Autoscale *GoogleBigqueryReservationAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Maximum number of queries that are allowed to run concurrently in this reservation.
	//
	// This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#concurrency GoogleBigqueryReservation#concurrency}
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#edition GoogleBigqueryReservation#edition}
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#id GoogleBigqueryReservation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project.
	//
	// If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#ignore_idle_slots GoogleBigqueryReservation#ignore_idle_slots}
	IgnoreIdleSlots interface{} `field:"optional" json:"ignoreIdleSlots" yaml:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#location GoogleBigqueryReservation#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The overall max slots for the reservation, covering slotCapacity (baseline), idle slots (if ignoreIdleSlots is false) and scaled slots.
	//
	// If present, the reservation won't use
	// more than the specified number of slots, even if there is demand and supply (from idle
	// slots). NOTE: capping a reservation's idle slot usage is best effort and its usage may
	// exceed the maxSlots value. However, in terms of autoscale.current_slots (which accounts
	// for the additional added slots), it will never exceed the maxSlots - baseline.
	//
	// This field must be set together with the scalingMode enum value, otherwise the request
	// will be rejected with error code google.rpc.Code.INVALID_ARGUMENT.
	//
	// If the maxSlots and scalingMode are set, the autoscale or autoscale.max_slots field
	// must be unset. Otherwise the request will be rejected with error code
	// google.rpc.Code.INVALID_ARGUMENT. However, the autoscale field may still be in the
	// output. The autopscale.max_slots will always show as 0 and the autoscaler.current_slots
	// will represent the current slots from autoscaler excluding idle slots. For example,
	// if the maxSlots is 1000 and scalingMode is AUTOSCALE_ONLY, then in the output, the
	// autoscaler.max_slots will be 0 and the autoscaler.current_slots may be any value
	// between 0 and 1000.
	//
	// If the maxSlots is 1000, scalingMode is ALL_SLOTS, the baseline is 100 and idle slots
	// usage is 200, then in the output, the autoscaler.max_slots will be 0 and the
	// autoscaler.current_slots will not be higher than 700.
	//
	// If the maxSlots is 1000, scalingMode is IDLE_SLOTS_ONLY, then in the output, the
	// autoscaler field will be null.
	//
	// If the maxSlots and scalingMode are set, then the ignoreIdleSlots field must be
	// aligned with the scalingMode enum value.(See details in ScalingMode comments).
	// Otherwise the request will be rejected with error code google.rpc.Code.INVALID_ARGUMENT.
	//
	// Please note, the maxSlots is for user to manage the part of slots greater than the
	// baseline. Therefore, we don't allow users to set maxSlots smaller or equal to the
	// baseline as it will not be meaningful. If the field is present and
	// slotCapacity>=maxSlots, requests will be rejected with error code
	// google.rpc.Code.INVALID_ARGUMENT.
	//
	// Please note that if maxSlots is set to 0, we will treat it as unset. Customers can set
	// maxSlots to 0 and set scalingMode to SCALING_MODE_UNSPECIFIED to disable the maxSlots
	// feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#max_slots GoogleBigqueryReservation#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#project GoogleBigqueryReservation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The reservation group that this reservation belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#reservation_group GoogleBigqueryReservation#reservation_group}
	ReservationGroup *string `field:"optional" json:"reservationGroup" yaml:"reservationGroup"`
	// The scaling mode for the reservation.
	//
	// If the field is present but maxSlots is not present,
	// requests will be rejected with error code google.rpc.Code.INVALID_ARGUMENT.
	//
	// Enum values:
	//
	// 'SCALING_MODE_UNSPECIFIED': Default value of ScalingMode.
	//
	// 'AUTOSCALE_ONLY': The reservation will scale up only using slots from autoscaling. It will
	// not use any idle slots even if there may be some available. The upper limit that autoscaling
	// can scale up to will be maxSlots - baseline. For example, if maxSlots is 1000, baseline is 200
	// and customer sets ScalingMode to AUTOSCALE_ONLY, then autoscalerg will scale up to 800 slots
	// and no idle slots will be used. Please note, in this mode, the ignoreIdleSlots field must be
	// set to true. Otherwise the request will be rejected with error code
	// google.rpc.Code.INVALID_ARGUMENT.
	//
	// 'IDLE_SLOTS_ONLY': The reservation will scale up using only idle slots contributed by other
	// reservations or from unassigned commitments. If no idle slots are available it will not scale
	// up further. If the idle slots which it is using are reclaimed by the contributing reservation(s)
	// it may be forced to scale down. The max idle slots the reservation can be maxSlots - baseline
	// capacity. For example, if maxSlots is 1000, baseline is 200 and customer sets ScalingMode to
	// IDLE_SLOTS_ONLY, 1. if there are 1000 idle slots available in other reservations, the
	// reservation will scale up to 1000 slots with 200 baseline and 800 idle slots. 2. if there are
	// 500 idle slots available in other reservations, the reservation will scale up to 700 slots with
	// 200 baseline and 300 idle slots. Please note, in this mode, the reservation might not be able to
	// scale up to maxSlots. Please note, in this mode, the ignoreIdleSlots field must be set to false.
	// Otherwise the request will be rejected with error code google.rpc.Code.INVALID_ARGUMENT
	//
	// 'ALL_SLOTS': The reservation will scale up using all slots available to it. It will use idle slots
	// contributed by other reservations or from unassigned commitments first. If no idle slots are
	// available it will scale up using autoscaling. For example, if maxSlots is 1000, baseline is 200
	// and customer sets ScalingMode to ALL_SLOTS, 1. if there are 800 idle slots available in other
	// reservations, the reservation will scale up to 1000 slots with 200 baseline and 800 idle slots. 2.
	// if there are 500 idle slots available in other reservations, the reservation will scale up to 1000
	// slots with 200 baseline, 500 idle slots and 300 autoscaling slots. 3. if there are no idle slots
	// available in other reservations, it will scale up to 1000 slots with 200 baseline and 800
	// autoscaling slots. Please note, in this mode, the ignoreIdleSlots field must be set to false.
	// Otherwise the request will be rejected with error code google.rpc.Code.INVALID_ARGUMENT. Possible values: ["SCALING_MODE_UNSPECIFIED", "AUTOSCALE_ONLY", "IDLE_SLOTS_ONLY", "ALL_SLOTS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#scaling_mode GoogleBigqueryReservation#scaling_mode}
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
	// The current location of the reservation's secondary replica.
	//
	// This field is only set for
	// reservations using the managed disaster recovery feature. Users can set this in create
	// reservation calls to create a failover reservation or in update reservation calls to convert
	// a non-failover reservation to a failover reservation(or vice versa).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#secondary_location GoogleBigqueryReservation#secondary_location}
	SecondaryLocation *string `field:"optional" json:"secondaryLocation" yaml:"secondaryLocation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_reservation#timeouts GoogleBigqueryReservation#timeouts}
	Timeouts *GoogleBigqueryReservationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

