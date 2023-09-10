package offset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OffsetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#base_rfc3339 Offset#base_rfc3339}
	BaseRfc3339 *string `field:"optional" json:"baseRfc3339" yaml:"baseRfc3339"`
	// Number of days to offset the base timestamp. At least one of the 'offset_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#offset_days Offset#offset_days}
	OffsetDays *float64 `field:"optional" json:"offsetDays" yaml:"offsetDays"`
	// Number of hours to offset the base timestamp. At least one of the 'offset_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#offset_hours Offset#offset_hours}
	OffsetHours *float64 `field:"optional" json:"offsetHours" yaml:"offsetHours"`
	// Number of minutes to offset the base timestamp. At least one of the 'offset_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#offset_minutes Offset#offset_minutes}
	OffsetMinutes *float64 `field:"optional" json:"offsetMinutes" yaml:"offsetMinutes"`
	// Number of months to offset the base timestamp. At least one of the 'offset_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#offset_months Offset#offset_months}
	OffsetMonths *float64 `field:"optional" json:"offsetMonths" yaml:"offsetMonths"`
	// Number of seconds to offset the base timestamp. At least one of the 'offset_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#offset_seconds Offset#offset_seconds}
	OffsetSeconds *float64 `field:"optional" json:"offsetSeconds" yaml:"offsetSeconds"`
	// Number of years to offset the base timestamp. At least one of the 'offset_' arguments must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#offset_years Offset#offset_years}
	OffsetYears *float64 `field:"optional" json:"offsetYears" yaml:"offsetYears"`
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved.
	//
	// See [the main provider documentation](../index.md) for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset#triggers Offset#triggers}
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

