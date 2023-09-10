package rotating

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RotatingConfig struct {
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
	// Base timestamp in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format (see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., `YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rfc3339 Rotating#rfc3339}
	Rfc3339 *string `field:"optional" json:"rfc3339" yaml:"rfc3339"`
	// Number of days to add to the base timestamp to configure the rotation timestamp.
	//
	// When the current time has passed the rotation timestamp, the resource will trigger recreation. At least one of the 'rotation_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rotation_days Rotating#rotation_days}
	RotationDays *float64 `field:"optional" json:"rotationDays" yaml:"rotationDays"`
	// Number of hours to add to the base timestamp to configure the rotation timestamp.
	//
	// When the current time has passed the rotation timestamp, the resource will trigger recreation. At least one of the 'rotation_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rotation_hours Rotating#rotation_hours}
	RotationHours *float64 `field:"optional" json:"rotationHours" yaml:"rotationHours"`
	// Number of minutes to add to the base timestamp to configure the rotation timestamp.
	//
	// When the current time has passed the rotation timestamp, the resource will trigger recreation. At least one of the 'rotation_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rotation_minutes Rotating#rotation_minutes}
	RotationMinutes *float64 `field:"optional" json:"rotationMinutes" yaml:"rotationMinutes"`
	// Number of months to add to the base timestamp to configure the rotation timestamp.
	//
	// When the current time has passed the rotation timestamp, the resource will trigger recreation. At least one of the 'rotation_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rotation_months Rotating#rotation_months}
	RotationMonths *float64 `field:"optional" json:"rotationMonths" yaml:"rotationMonths"`
	// Configure the rotation timestamp with an [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format of the offset timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. At least one of the 'rotation_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rotation_rfc3339 Rotating#rotation_rfc3339}
	RotationRfc3339 *string `field:"optional" json:"rotationRfc3339" yaml:"rotationRfc3339"`
	// Number of years to add to the base timestamp to configure the rotation timestamp.
	//
	// When the current time has passed the rotation timestamp, the resource will trigger recreation. At least one of the 'rotation_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#rotation_years Rotating#rotation_years}
	RotationYears *float64 `field:"optional" json:"rotationYears" yaml:"rotationYears"`
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved.
	//
	// These conditions recreate the resource in addition to other rotation arguments. See [the main provider documentation](../index.md) for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating#triggers Rotating#triggers}
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

