package datagooglecomputereservationsubblock

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeReservationSubBlockConfig struct {
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
	// The name of the reservation sub-block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_compute_reservation_sub_block#name DataGoogleComputeReservationSubBlock#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the parent reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_compute_reservation_sub_block#reservation DataGoogleComputeReservationSubBlock#reservation}
	Reservation *string `field:"required" json:"reservation" yaml:"reservation"`
	// The name of the parent reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_compute_reservation_sub_block#reservation_block DataGoogleComputeReservationSubBlock#reservation_block}
	ReservationBlock *string `field:"required" json:"reservationBlock" yaml:"reservationBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_compute_reservation_sub_block#id DataGoogleComputeReservationSubBlock#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project in which the resource belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_compute_reservation_sub_block#project DataGoogleComputeReservationSubBlock#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The zone where the reservation sub-block resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_compute_reservation_sub_block#zone DataGoogleComputeReservationSubBlock#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

