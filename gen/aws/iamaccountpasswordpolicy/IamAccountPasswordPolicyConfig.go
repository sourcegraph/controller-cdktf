package iamaccountpasswordpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamAccountPasswordPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#allow_users_to_change_password IamAccountPasswordPolicy#allow_users_to_change_password}.
	AllowUsersToChangePassword any `field:"optional" json:"allowUsersToChangePassword" yaml:"allowUsersToChangePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#hard_expiry IamAccountPasswordPolicy#hard_expiry}.
	HardExpiry any `field:"optional" json:"hardExpiry" yaml:"hardExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#id IamAccountPasswordPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#max_password_age IamAccountPasswordPolicy#max_password_age}.
	MaxPasswordAge *float64 `field:"optional" json:"maxPasswordAge" yaml:"maxPasswordAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#minimum_password_length IamAccountPasswordPolicy#minimum_password_length}.
	MinimumPasswordLength *float64 `field:"optional" json:"minimumPasswordLength" yaml:"minimumPasswordLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#password_reuse_prevention IamAccountPasswordPolicy#password_reuse_prevention}.
	PasswordReusePrevention *float64 `field:"optional" json:"passwordReusePrevention" yaml:"passwordReusePrevention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#require_lowercase_characters IamAccountPasswordPolicy#require_lowercase_characters}.
	RequireLowercaseCharacters any `field:"optional" json:"requireLowercaseCharacters" yaml:"requireLowercaseCharacters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#require_numbers IamAccountPasswordPolicy#require_numbers}.
	RequireNumbers any `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#require_symbols IamAccountPasswordPolicy#require_symbols}.
	RequireSymbols any `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iam_account_password_policy#require_uppercase_characters IamAccountPasswordPolicy#require_uppercase_characters}.
	RequireUppercaseCharacters any `field:"optional" json:"requireUppercaseCharacters" yaml:"requireUppercaseCharacters"`
}
