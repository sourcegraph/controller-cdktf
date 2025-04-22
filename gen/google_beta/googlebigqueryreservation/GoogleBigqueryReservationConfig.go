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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#name GoogleBigqueryReservation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Minimum slots available to this reservation.
	//
	// A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#slot_capacity GoogleBigqueryReservation#slot_capacity}
	SlotCapacity *float64 `field:"required" json:"slotCapacity" yaml:"slotCapacity"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#autoscale GoogleBigqueryReservation#autoscale}
	Autoscale *GoogleBigqueryReservationAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Maximum number of queries that are allowed to run concurrently in this reservation.
	//
	// This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#concurrency GoogleBigqueryReservation#concurrency}
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#edition GoogleBigqueryReservation#edition}
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#id GoogleBigqueryReservation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project.
	//
	// If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#ignore_idle_slots GoogleBigqueryReservation#ignore_idle_slots}
	IgnoreIdleSlots interface{} `field:"optional" json:"ignoreIdleSlots" yaml:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#location GoogleBigqueryReservation#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#project GoogleBigqueryReservation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The current location of the reservation's secondary replica.
	//
	// This field is only set for
	// reservations using the managed disaster recovery feature. Users can set this in create
	// reservation calls to create a failover reservation or in update reservation calls to convert
	// a non-failover reservation to a failover reservation(or vice versa).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#secondary_location GoogleBigqueryReservation#secondary_location}
	SecondaryLocation *string `field:"optional" json:"secondaryLocation" yaml:"secondaryLocation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_reservation#timeouts GoogleBigqueryReservation#timeouts}
	Timeouts *GoogleBigqueryReservationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

