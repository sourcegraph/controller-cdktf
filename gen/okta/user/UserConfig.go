package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// User primary email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#email User#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// User first name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#first_name User#first_name}
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// User last name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#last_name User#last_name}
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// User Okta login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#login User#login}
	Login *string `field:"required" json:"login" yaml:"login"`
	// User city.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#city User#city}
	City *string `field:"optional" json:"city" yaml:"city"`
	// User cost center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#cost_center User#cost_center}
	CostCenter *string `field:"optional" json:"costCenter" yaml:"costCenter"`
	// User country code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#country_code User#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// JSON formatted custom attributes for a user. It must be JSON due to various types Okta allows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#custom_profile_attributes User#custom_profile_attributes}
	CustomProfileAttributes *string `field:"optional" json:"customProfileAttributes" yaml:"customProfileAttributes"`
	// List of custom_profile_attribute keys that should be excluded from being managed by Terraform.
	//
	// This is useful in situations where specific custom fields may contain sensitive information and should be managed outside of Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#custom_profile_attributes_to_ignore User#custom_profile_attributes_to_ignore}
	CustomProfileAttributesToIgnore *[]*string `field:"optional" json:"customProfileAttributesToIgnore" yaml:"customProfileAttributesToIgnore"`
	// User department.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#department User#department}
	Department *string `field:"optional" json:"department" yaml:"department"`
	// User display name, suitable to show end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#display_name User#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// User division.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#division User#division}
	Division *string `field:"optional" json:"division" yaml:"division"`
	// User employee number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#employee_number User#employee_number}
	EmployeeNumber *string `field:"optional" json:"employeeNumber" yaml:"employeeNumber"`
	// If set to `true`, the user will have to change the password at the next login.
	//
	// This property will be used when user is being created and works only when `password` field is set. Default: `false`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#expire_password_on_create User#expire_password_on_create}
	ExpirePasswordOnCreate interface{} `field:"optional" json:"expirePasswordOnCreate" yaml:"expirePasswordOnCreate"`
	// User honorific prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#honorific_prefix User#honorific_prefix}
	HonorificPrefix *string `field:"optional" json:"honorificPrefix" yaml:"honorificPrefix"`
	// User honorific suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#honorific_suffix User#honorific_suffix}
	HonorificSuffix *string `field:"optional" json:"honorificSuffix" yaml:"honorificSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#locale User#locale}
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Manager of User.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#manager User#manager}
	Manager *string `field:"optional" json:"manager" yaml:"manager"`
	// Manager ID of User.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#manager_id User#manager_id}
	ManagerId *string `field:"optional" json:"managerId" yaml:"managerId"`
	// User middle name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#middle_name User#middle_name}
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
	// User mobile phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#mobile_phone User#mobile_phone}
	MobilePhone *string `field:"optional" json:"mobilePhone" yaml:"mobilePhone"`
	// User nickname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#nick_name User#nick_name}
	NickName *string `field:"optional" json:"nickName" yaml:"nickName"`
	// Old User Password.
	//
	// Should be only set in case the password was not changed using the provider. fter successful password change this field should be removed and `password` field should be used for further changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#old_password User#old_password}
	OldPassword *string `field:"optional" json:"oldPassword" yaml:"oldPassword"`
	// User organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#organization User#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// User Password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#password User#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// password_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#password_hash User#password_hash}
	PasswordHash *UserPasswordHash `field:"optional" json:"passwordHash" yaml:"passwordHash"`
	// Specifies that a Password Import Inline Hook should be triggered to handle verification of the user's password the first time the user logs in.
	//
	// This allows an existing password to be imported into Okta directly from some other store. When updating a user with a password hook the user must be in the `STAGED` status. The `password` field should not be specified when using Password Import Inline Hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#password_inline_hook User#password_inline_hook}
	PasswordInlineHook *string `field:"optional" json:"passwordInlineHook" yaml:"passwordInlineHook"`
	// User mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#postal_address User#postal_address}
	PostalAddress *string `field:"optional" json:"postalAddress" yaml:"postalAddress"`
	// User preferred language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#preferred_language User#preferred_language}
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// User primary phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#primary_phone User#primary_phone}
	PrimaryPhone *string `field:"optional" json:"primaryPhone" yaml:"primaryPhone"`
	// User online profile (web page).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#profile_url User#profile_url}
	ProfileUrl *string `field:"optional" json:"profileUrl" yaml:"profileUrl"`
	// User Password Recovery Answer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#recovery_answer User#recovery_answer}
	RecoveryAnswer *string `field:"optional" json:"recoveryAnswer" yaml:"recoveryAnswer"`
	// User Password Recovery Question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#recovery_question User#recovery_question}
	RecoveryQuestion *string `field:"optional" json:"recoveryQuestion" yaml:"recoveryQuestion"`
	// User secondary email address, used for account recovery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#second_email User#second_email}
	SecondEmail *string `field:"optional" json:"secondEmail" yaml:"secondEmail"`
	// Do not populate user roles information (prevents additional API call).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#skip_roles User#skip_roles}
	SkipRoles interface{} `field:"optional" json:"skipRoles" yaml:"skipRoles"`
	// User state or region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#state User#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// User profile property. Valid values are `ACTIVE`, `DEPROVISIONED`, `STAGED`, `SUSPENDED`. Default: `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#status User#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// User street address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#street_address User#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// User default timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#timezone User#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// User title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#title User#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// User employee type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#user_type User#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
	// User zipcode or postal code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user#zip_code User#zip_code}
	ZipCode *string `field:"optional" json:"zipCode" yaml:"zipCode"`
}

