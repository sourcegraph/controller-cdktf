package dataincidentschedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataIncidentScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Unique internal ID of the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/schedule#id DataIncidentSchedule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Human readable name synced from external provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/schedule#name DataIncidentSchedule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

