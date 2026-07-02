package schedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ScheduleConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Human readable name synced from external provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#name Schedule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#rotations Schedule#rotations}.
	Rotations any `field:"required" json:"rotations" yaml:"rotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#timezone Schedule#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#holidays_public_config Schedule#holidays_public_config}.
	HolidaysPublicConfig *ScheduleHolidaysPublicConfig `field:"optional" json:"holidaysPublicConfig" yaml:"holidaysPublicConfig"`
	// IDs of teams that own this schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#team_ids Schedule#team_ids}
	TeamIds *[]*string `field:"optional" json:"teamIds" yaml:"teamIds"`
}
