package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/user/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user okta_user}.
type User interface {
	cdktf.TerraformResource
	AdminRoles() *[]*string
	SetAdminRoles(val *[]*string)
	AdminRolesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	SetCity(val *string)
	CityInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostCenter() *string
	SetCostCenter(val *string)
	CostCenterInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CountryCode() *string
	SetCountryCode(val *string)
	CountryCodeInput() *string
	CustomProfileAttributes() *string
	SetCustomProfileAttributes(val *string)
	CustomProfileAttributesInput() *string
	Department() *string
	SetDepartment(val *string)
	DepartmentInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Division() *string
	SetDivision(val *string)
	DivisionInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	EmployeeNumber() *string
	SetEmployeeNumber(val *string)
	EmployeeNumberInput() *string
	ExpirePasswordOnCreate() interface{}
	SetExpirePasswordOnCreate(val interface{})
	ExpirePasswordOnCreateInput() interface{}
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupMemberships() *[]*string
	SetGroupMemberships(val *[]*string)
	GroupMembershipsInput() *[]*string
	HonorificPrefix() *string
	SetHonorificPrefix(val *string)
	HonorificPrefixInput() *string
	HonorificSuffix() *string
	SetHonorificSuffix(val *string)
	HonorificSuffixInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	Login() *string
	SetLogin(val *string)
	LoginInput() *string
	Manager() *string
	SetManager(val *string)
	ManagerId() *string
	SetManagerId(val *string)
	ManagerIdInput() *string
	ManagerInput() *string
	MiddleName() *string
	SetMiddleName(val *string)
	MiddleNameInput() *string
	MobilePhone() *string
	SetMobilePhone(val *string)
	MobilePhoneInput() *string
	NickName() *string
	SetNickName(val *string)
	NickNameInput() *string
	// The tree node.
	Node() constructs.Node
	OldPassword() *string
	SetOldPassword(val *string)
	OldPasswordInput() *string
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordHash() UserPasswordHashOutputReference
	PasswordHashInput() *UserPasswordHash
	PasswordInlineHook() *string
	SetPasswordInlineHook(val *string)
	PasswordInlineHookInput() *string
	PasswordInput() *string
	PostalAddress() *string
	SetPostalAddress(val *string)
	PostalAddressInput() *string
	PreferredLanguage() *string
	SetPreferredLanguage(val *string)
	PreferredLanguageInput() *string
	PrimaryPhone() *string
	SetPrimaryPhone(val *string)
	PrimaryPhoneInput() *string
	ProfileUrl() *string
	SetProfileUrl(val *string)
	ProfileUrlInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RawStatus() *string
	RecoveryAnswer() *string
	SetRecoveryAnswer(val *string)
	RecoveryAnswerInput() *string
	RecoveryQuestion() *string
	SetRecoveryQuestion(val *string)
	RecoveryQuestionInput() *string
	SecondEmail() *string
	SetSecondEmail(val *string)
	SecondEmailInput() *string
	SkipRoles() interface{}
	SetSkipRoles(val interface{})
	SkipRolesInput() interface{}
	State() *string
	SetState(val *string)
	StateInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	StreetAddress() *string
	SetStreetAddress(val *string)
	StreetAddressInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	UserType() *string
	SetUserType(val *string)
	UserTypeInput() *string
	ZipCode() *string
	SetZipCode(val *string)
	ZipCodeInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutPasswordHash(value *UserPasswordHash)
	ResetAdminRoles()
	ResetCity()
	ResetCostCenter()
	ResetCountryCode()
	ResetCustomProfileAttributes()
	ResetDepartment()
	ResetDisplayName()
	ResetDivision()
	ResetEmployeeNumber()
	ResetExpirePasswordOnCreate()
	ResetGroupMemberships()
	ResetHonorificPrefix()
	ResetHonorificSuffix()
	ResetId()
	ResetLocale()
	ResetManager()
	ResetManagerId()
	ResetMiddleName()
	ResetMobilePhone()
	ResetNickName()
	ResetOldPassword()
	ResetOrganization()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPasswordHash()
	ResetPasswordInlineHook()
	ResetPostalAddress()
	ResetPreferredLanguage()
	ResetPrimaryPhone()
	ResetProfileUrl()
	ResetRecoveryAnswer()
	ResetRecoveryQuestion()
	ResetSecondEmail()
	ResetSkipRoles()
	ResetState()
	ResetStatus()
	ResetStreetAddress()
	ResetTimezone()
	ResetTitle()
	ResetUserType()
	ResetZipCode()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) AdminRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AdminRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CostCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CostCenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomProfileAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProfileAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomProfileAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProfileAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DepartmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"departmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Division() *string {
	var returns *string
	_jsii_.Get(
		j,
		"division",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DivisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"divisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ExpirePasswordOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirePasswordOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ExpirePasswordOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirePasswordOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GroupMemberships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMemberships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GroupMembershipsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMembershipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HonorificPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HonorificPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HonorificSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HonorificSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Manager() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ManagerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ManagerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MobilePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MobilePhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NickName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nickName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NickNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nickNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OldPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OldPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordHash() UserPasswordHashOutputReference {
	var returns UserPasswordHashOutputReference
	_jsii_.Get(
		j,
		"passwordHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordHashInput() *UserPasswordHash {
	var returns *UserPasswordHash
	_jsii_.Get(
		j,
		"passwordHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInlineHook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInlineHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInlineHookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInlineHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PostalAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PostalAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PreferredLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrimaryPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrimaryPhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ProfileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ProfileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryAnswer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryAnswerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryQuestion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryQuestionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SecondEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SecondEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SkipRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SkipRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user okta_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-okta.user.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user okta_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.user.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User)SetAdminRoles(val *[]*string) {
	if err := j.validateSetAdminRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminRoles",
		val,
	)
}

func (j *jsiiProxy_User)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_User)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User)SetCostCenter(val *string) {
	if err := j.validateSetCostCenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"costCenter",
		val,
	)
}

func (j *jsiiProxy_User)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_User)SetCustomProfileAttributes(val *string) {
	if err := j.validateSetCustomProfileAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customProfileAttributes",
		val,
	)
}

func (j *jsiiProxy_User)SetDepartment(val *string) {
	if err := j.validateSetDepartmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"department",
		val,
	)
}

func (j *jsiiProxy_User)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_User)SetDivision(val *string) {
	if err := j.validateSetDivisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"division",
		val,
	)
}

func (j *jsiiProxy_User)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_User)SetEmployeeNumber(val *string) {
	if err := j.validateSetEmployeeNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeNumber",
		val,
	)
}

func (j *jsiiProxy_User)SetExpirePasswordOnCreate(val interface{}) {
	if err := j.validateSetExpirePasswordOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirePasswordOnCreate",
		val,
	)
}

func (j *jsiiProxy_User)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_User)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User)SetGroupMemberships(val *[]*string) {
	if err := j.validateSetGroupMembershipsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMemberships",
		val,
	)
}

func (j *jsiiProxy_User)SetHonorificPrefix(val *string) {
	if err := j.validateSetHonorificPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"honorificPrefix",
		val,
	)
}

func (j *jsiiProxy_User)SetHonorificSuffix(val *string) {
	if err := j.validateSetHonorificSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"honorificSuffix",
		val,
	)
}

func (j *jsiiProxy_User)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_User)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_User)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_User)SetLogin(val *string) {
	if err := j.validateSetLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_User)SetManager(val *string) {
	if err := j.validateSetManagerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manager",
		val,
	)
}

func (j *jsiiProxy_User)SetManagerId(val *string) {
	if err := j.validateSetManagerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managerId",
		val,
	)
}

func (j *jsiiProxy_User)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_User)SetMobilePhone(val *string) {
	if err := j.validateSetMobilePhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhone",
		val,
	)
}

func (j *jsiiProxy_User)SetNickName(val *string) {
	if err := j.validateSetNickNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nickName",
		val,
	)
}

func (j *jsiiProxy_User)SetOldPassword(val *string) {
	if err := j.validateSetOldPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oldPassword",
		val,
	)
}

func (j *jsiiProxy_User)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_User)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_User)SetPasswordInlineHook(val *string) {
	if err := j.validateSetPasswordInlineHookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordInlineHook",
		val,
	)
}

func (j *jsiiProxy_User)SetPostalAddress(val *string) {
	if err := j.validateSetPostalAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalAddress",
		val,
	)
}

func (j *jsiiProxy_User)SetPreferredLanguage(val *string) {
	if err := j.validateSetPreferredLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredLanguage",
		val,
	)
}

func (j *jsiiProxy_User)SetPrimaryPhone(val *string) {
	if err := j.validateSetPrimaryPhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryPhone",
		val,
	)
}

func (j *jsiiProxy_User)SetProfileUrl(val *string) {
	if err := j.validateSetProfileUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileUrl",
		val,
	)
}

func (j *jsiiProxy_User)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User)SetRecoveryAnswer(val *string) {
	if err := j.validateSetRecoveryAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryAnswer",
		val,
	)
}

func (j *jsiiProxy_User)SetRecoveryQuestion(val *string) {
	if err := j.validateSetRecoveryQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryQuestion",
		val,
	)
}

func (j *jsiiProxy_User)SetSecondEmail(val *string) {
	if err := j.validateSetSecondEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondEmail",
		val,
	)
}

func (j *jsiiProxy_User)SetSkipRoles(val interface{}) {
	if err := j.validateSetSkipRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipRoles",
		val,
	)
}

func (j *jsiiProxy_User)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_User)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_User)SetStreetAddress(val *string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_User)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_User)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_User)SetUserType(val *string) {
	if err := j.validateSetUserTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userType",
		val,
	)
}

func (j *jsiiProxy_User)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.user.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.user.User",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.user.User",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.user.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) PutPasswordHash(value *UserPasswordHash) {
	if err := u.validatePutPasswordHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putPasswordHash",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) ResetAdminRoles() {
	_jsii_.InvokeVoid(
		u,
		"resetAdminRoles",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCity() {
	_jsii_.InvokeVoid(
		u,
		"resetCity",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCostCenter() {
	_jsii_.InvokeVoid(
		u,
		"resetCostCenter",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCountryCode() {
	_jsii_.InvokeVoid(
		u,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCustomProfileAttributes() {
	_jsii_.InvokeVoid(
		u,
		"resetCustomProfileAttributes",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDepartment() {
	_jsii_.InvokeVoid(
		u,
		"resetDepartment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisplayName() {
	_jsii_.InvokeVoid(
		u,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDivision() {
	_jsii_.InvokeVoid(
		u,
		"resetDivision",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmployeeNumber() {
	_jsii_.InvokeVoid(
		u,
		"resetEmployeeNumber",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetExpirePasswordOnCreate() {
	_jsii_.InvokeVoid(
		u,
		"resetExpirePasswordOnCreate",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetGroupMemberships() {
	_jsii_.InvokeVoid(
		u,
		"resetGroupMemberships",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHonorificPrefix() {
	_jsii_.InvokeVoid(
		u,
		"resetHonorificPrefix",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHonorificSuffix() {
	_jsii_.InvokeVoid(
		u,
		"resetHonorificSuffix",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLocale() {
	_jsii_.InvokeVoid(
		u,
		"resetLocale",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetManager() {
	_jsii_.InvokeVoid(
		u,
		"resetManager",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetManagerId() {
	_jsii_.InvokeVoid(
		u,
		"resetManagerId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMiddleName() {
	_jsii_.InvokeVoid(
		u,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMobilePhone() {
	_jsii_.InvokeVoid(
		u,
		"resetMobilePhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetNickName() {
	_jsii_.InvokeVoid(
		u,
		"resetNickName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOldPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetOldPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOrganization() {
	_jsii_.InvokeVoid(
		u,
		"resetOrganization",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPasswordHash() {
	_jsii_.InvokeVoid(
		u,
		"resetPasswordHash",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPasswordInlineHook() {
	_jsii_.InvokeVoid(
		u,
		"resetPasswordInlineHook",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPostalAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetPostalAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPreferredLanguage() {
	_jsii_.InvokeVoid(
		u,
		"resetPreferredLanguage",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPrimaryPhone() {
	_jsii_.InvokeVoid(
		u,
		"resetPrimaryPhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetProfileUrl() {
	_jsii_.InvokeVoid(
		u,
		"resetProfileUrl",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRecoveryAnswer() {
	_jsii_.InvokeVoid(
		u,
		"resetRecoveryAnswer",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRecoveryQuestion() {
	_jsii_.InvokeVoid(
		u,
		"resetRecoveryQuestion",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSecondEmail() {
	_jsii_.InvokeVoid(
		u,
		"resetSecondEmail",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSkipRoles() {
	_jsii_.InvokeVoid(
		u,
		"resetSkipRoles",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetState() {
	_jsii_.InvokeVoid(
		u,
		"resetState",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStatus() {
	_jsii_.InvokeVoid(
		u,
		"resetStatus",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimezone() {
	_jsii_.InvokeVoid(
		u,
		"resetTimezone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTitle() {
	_jsii_.InvokeVoid(
		u,
		"resetTitle",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetUserType() {
	_jsii_.InvokeVoid(
		u,
		"resetUserType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetZipCode() {
	_jsii_.InvokeVoid(
		u,
		"resetZipCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

