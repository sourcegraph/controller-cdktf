package password

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PasswordConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The length of the string desired.
	//
	// The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#length Password#length}
	Length *float64 `field:"required" json:"length" yaml:"length"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource.
	//
	// See [the main provider documentation](../index.html) for more information.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#keepers Password#keepers}
	Keepers *map[string]*string `field:"optional" json:"keepers" yaml:"keepers"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#lower Password#lower}
	Lower interface{} `field:"optional" json:"lower" yaml:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#min_lower Password#min_lower}
	MinLower *float64 `field:"optional" json:"minLower" yaml:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#min_numeric Password#min_numeric}
	MinNumeric *float64 `field:"optional" json:"minNumeric" yaml:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#min_special Password#min_special}
	MinSpecial *float64 `field:"optional" json:"minSpecial" yaml:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#min_upper Password#min_upper}
	MinUpper *float64 `field:"optional" json:"minUpper" yaml:"minUpper"`
	// Include numeric characters in the result. Default value is `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#number Password#number}
	Number interface{} `field:"optional" json:"number" yaml:"number"`
	// Supply your own list of special characters to use for string generation.
	//
	// This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#override_special Password#override_special}
	OverrideSpecial *string `field:"optional" json:"overrideSpecial" yaml:"overrideSpecial"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#special Password#special}
	Special interface{} `field:"optional" json:"special" yaml:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/password#upper Password#upper}
	Upper interface{} `field:"optional" json:"upper" yaml:"upper"`
}

