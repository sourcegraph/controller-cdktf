package policypassword

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policypassword/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_password okta_policy_password}.
type PolicyPassword interface {
	cdktf.TerraformResource
	AuthProvider() *string
	SetAuthProvider(val *string)
	AuthProviderInput() *string
	CallRecovery() *string
	SetCallRecovery(val *string)
	CallRecoveryInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailRecovery() *string
	SetEmailRecovery(val *string)
	EmailRecoveryInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsIncluded() *[]*string
	SetGroupsIncluded(val *[]*string)
	GroupsIncludedInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PasswordAutoUnlockMinutes() *float64
	SetPasswordAutoUnlockMinutes(val *float64)
	PasswordAutoUnlockMinutesInput() *float64
	PasswordDictionaryLookup() interface{}
	SetPasswordDictionaryLookup(val interface{})
	PasswordDictionaryLookupInput() interface{}
	PasswordExcludeFirstName() interface{}
	SetPasswordExcludeFirstName(val interface{})
	PasswordExcludeFirstNameInput() interface{}
	PasswordExcludeLastName() interface{}
	SetPasswordExcludeLastName(val interface{})
	PasswordExcludeLastNameInput() interface{}
	PasswordExcludeUsername() interface{}
	SetPasswordExcludeUsername(val interface{})
	PasswordExcludeUsernameInput() interface{}
	PasswordExpireWarnDays() *float64
	SetPasswordExpireWarnDays(val *float64)
	PasswordExpireWarnDaysInput() *float64
	PasswordHistoryCount() *float64
	SetPasswordHistoryCount(val *float64)
	PasswordHistoryCountInput() *float64
	PasswordLockoutNotificationChannels() *[]*string
	SetPasswordLockoutNotificationChannels(val *[]*string)
	PasswordLockoutNotificationChannelsInput() *[]*string
	PasswordMaxAgeDays() *float64
	SetPasswordMaxAgeDays(val *float64)
	PasswordMaxAgeDaysInput() *float64
	PasswordMaxLockoutAttempts() *float64
	SetPasswordMaxLockoutAttempts(val *float64)
	PasswordMaxLockoutAttemptsInput() *float64
	PasswordMinAgeMinutes() *float64
	SetPasswordMinAgeMinutes(val *float64)
	PasswordMinAgeMinutesInput() *float64
	PasswordMinLength() *float64
	SetPasswordMinLength(val *float64)
	PasswordMinLengthInput() *float64
	PasswordMinLowercase() *float64
	SetPasswordMinLowercase(val *float64)
	PasswordMinLowercaseInput() *float64
	PasswordMinNumber() *float64
	SetPasswordMinNumber(val *float64)
	PasswordMinNumberInput() *float64
	PasswordMinSymbol() *float64
	SetPasswordMinSymbol(val *float64)
	PasswordMinSymbolInput() *float64
	PasswordMinUppercase() *float64
	SetPasswordMinUppercase(val *float64)
	PasswordMinUppercaseInput() *float64
	PasswordShowLockoutFailures() interface{}
	SetPasswordShowLockoutFailures(val interface{})
	PasswordShowLockoutFailuresInput() interface{}
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuestionMinLength() *float64
	SetQuestionMinLength(val *float64)
	QuestionMinLengthInput() *float64
	QuestionRecovery() *string
	SetQuestionRecovery(val *string)
	QuestionRecoveryInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecoveryEmailToken() *float64
	SetRecoveryEmailToken(val *float64)
	RecoveryEmailTokenInput() *float64
	SkipUnlock() interface{}
	SetSkipUnlock(val interface{})
	SkipUnlockInput() interface{}
	SmsRecovery() *string
	SetSmsRecovery(val *string)
	SmsRecoveryInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAuthProvider()
	ResetCallRecovery()
	ResetDescription()
	ResetEmailRecovery()
	ResetGroupsIncluded()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordAutoUnlockMinutes()
	ResetPasswordDictionaryLookup()
	ResetPasswordExcludeFirstName()
	ResetPasswordExcludeLastName()
	ResetPasswordExcludeUsername()
	ResetPasswordExpireWarnDays()
	ResetPasswordHistoryCount()
	ResetPasswordLockoutNotificationChannels()
	ResetPasswordMaxAgeDays()
	ResetPasswordMaxLockoutAttempts()
	ResetPasswordMinAgeMinutes()
	ResetPasswordMinLength()
	ResetPasswordMinLowercase()
	ResetPasswordMinNumber()
	ResetPasswordMinSymbol()
	ResetPasswordMinUppercase()
	ResetPasswordShowLockoutFailures()
	ResetPriority()
	ResetQuestionMinLength()
	ResetQuestionRecovery()
	ResetRecoveryEmailToken()
	ResetSkipUnlock()
	ResetSmsRecovery()
	ResetStatus()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PolicyPassword
type jsiiProxy_PolicyPassword struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyPassword) AuthProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) AuthProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) CallRecovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) CallRecoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) EmailRecovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) EmailRecoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) GroupsIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) GroupsIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordAutoUnlockMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordAutoUnlockMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordAutoUnlockMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordAutoUnlockMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordDictionaryLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordDictionaryLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordDictionaryLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordDictionaryLookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExcludeFirstName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordExcludeFirstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExcludeFirstNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordExcludeFirstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExcludeLastName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordExcludeLastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExcludeLastNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordExcludeLastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExcludeUsername() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordExcludeUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExcludeUsernameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordExcludeUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExpireWarnDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordExpireWarnDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordExpireWarnDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordExpireWarnDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordHistoryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistoryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordHistoryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistoryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordLockoutNotificationChannels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passwordLockoutNotificationChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordLockoutNotificationChannelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passwordLockoutNotificationChannelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMaxAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMaxAgeDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxAgeDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMaxLockoutAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxLockoutAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMaxLockoutAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxLockoutAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinAgeMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinAgeMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinAgeMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinAgeMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinLowercase() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinLowercaseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinSymbol() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinSymbolInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinSymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinUppercase() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordMinUppercaseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordShowLockoutFailures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordShowLockoutFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PasswordShowLockoutFailuresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordShowLockoutFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) QuestionMinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"questionMinLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) QuestionMinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"questionMinLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) QuestionRecovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) QuestionRecoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) RecoveryEmailToken() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryEmailToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) RecoveryEmailTokenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryEmailTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) SkipUnlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) SkipUnlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUnlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) SmsRecovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) SmsRecoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyPassword) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_password okta_policy_password} Resource.
func NewPolicyPassword(scope constructs.Construct, id *string, config *PolicyPasswordConfig) PolicyPassword {
	_init_.Initialize()

	if err := validateNewPolicyPasswordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyPassword{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_password okta_policy_password} Resource.
func NewPolicyPassword_Override(p PolicyPassword, scope constructs.Construct, id *string, config *PolicyPasswordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyPassword)SetAuthProvider(val *string) {
	if err := j.validateSetAuthProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authProvider",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetCallRecovery(val *string) {
	if err := j.validateSetCallRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callRecovery",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetEmailRecovery(val *string) {
	if err := j.validateSetEmailRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailRecovery",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetGroupsIncluded(val *[]*string) {
	if err := j.validateSetGroupsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsIncluded",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordAutoUnlockMinutes(val *float64) {
	if err := j.validateSetPasswordAutoUnlockMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordAutoUnlockMinutes",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordDictionaryLookup(val interface{}) {
	if err := j.validateSetPasswordDictionaryLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordDictionaryLookup",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordExcludeFirstName(val interface{}) {
	if err := j.validateSetPasswordExcludeFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordExcludeFirstName",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordExcludeLastName(val interface{}) {
	if err := j.validateSetPasswordExcludeLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordExcludeLastName",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordExcludeUsername(val interface{}) {
	if err := j.validateSetPasswordExcludeUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordExcludeUsername",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordExpireWarnDays(val *float64) {
	if err := j.validateSetPasswordExpireWarnDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordExpireWarnDays",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordHistoryCount(val *float64) {
	if err := j.validateSetPasswordHistoryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordHistoryCount",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordLockoutNotificationChannels(val *[]*string) {
	if err := j.validateSetPasswordLockoutNotificationChannelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLockoutNotificationChannels",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMaxAgeDays(val *float64) {
	if err := j.validateSetPasswordMaxAgeDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMaxAgeDays",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMaxLockoutAttempts(val *float64) {
	if err := j.validateSetPasswordMaxLockoutAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMaxLockoutAttempts",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMinAgeMinutes(val *float64) {
	if err := j.validateSetPasswordMinAgeMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMinAgeMinutes",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMinLength(val *float64) {
	if err := j.validateSetPasswordMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMinLength",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMinLowercase(val *float64) {
	if err := j.validateSetPasswordMinLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMinLowercase",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMinNumber(val *float64) {
	if err := j.validateSetPasswordMinNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMinNumber",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMinSymbol(val *float64) {
	if err := j.validateSetPasswordMinSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMinSymbol",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordMinUppercase(val *float64) {
	if err := j.validateSetPasswordMinUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordMinUppercase",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPasswordShowLockoutFailures(val interface{}) {
	if err := j.validateSetPasswordShowLockoutFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordShowLockoutFailures",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetQuestionMinLength(val *float64) {
	if err := j.validateSetQuestionMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"questionMinLength",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetQuestionRecovery(val *string) {
	if err := j.validateSetQuestionRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"questionRecovery",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetRecoveryEmailToken(val *float64) {
	if err := j.validateSetRecoveryEmailTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryEmailToken",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetSkipUnlock(val interface{}) {
	if err := j.validateSetSkipUnlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUnlock",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetSmsRecovery(val *string) {
	if err := j.validateSetSmsRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsRecovery",
		val,
	)
}

func (j *jsiiProxy_PolicyPassword)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
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
func PolicyPassword_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyPassword_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyPassword_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyPassword_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyPassword_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyPassword_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyPassword_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyPassword.PolicyPassword",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyPassword) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyPassword) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyPassword) ResetAuthProvider() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthProvider",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetCallRecovery() {
	_jsii_.InvokeVoid(
		p,
		"resetCallRecovery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetEmailRecovery() {
	_jsii_.InvokeVoid(
		p,
		"resetEmailRecovery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetGroupsIncluded() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupsIncluded",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordAutoUnlockMinutes() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordAutoUnlockMinutes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordDictionaryLookup() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordDictionaryLookup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordExcludeFirstName() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordExcludeFirstName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordExcludeLastName() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordExcludeLastName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordExcludeUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordExcludeUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordExpireWarnDays() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordExpireWarnDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordHistoryCount() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordHistoryCount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordLockoutNotificationChannels() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordLockoutNotificationChannels",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMaxAgeDays() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMaxAgeDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMaxLockoutAttempts() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMaxLockoutAttempts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMinAgeMinutes() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMinAgeMinutes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMinLength() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMinLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMinLowercase() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMinLowercase",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMinNumber() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMinNumber",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMinSymbol() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMinSymbol",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordMinUppercase() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordMinUppercase",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPasswordShowLockoutFailures() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordShowLockoutFailures",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetPriority() {
	_jsii_.InvokeVoid(
		p,
		"resetPriority",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetQuestionMinLength() {
	_jsii_.InvokeVoid(
		p,
		"resetQuestionMinLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetQuestionRecovery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuestionRecovery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetRecoveryEmailToken() {
	_jsii_.InvokeVoid(
		p,
		"resetRecoveryEmailToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetSkipUnlock() {
	_jsii_.InvokeVoid(
		p,
		"resetSkipUnlock",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetSmsRecovery() {
	_jsii_.InvokeVoid(
		p,
		"resetSmsRecovery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) ResetStatus() {
	_jsii_.InvokeVoid(
		p,
		"resetStatus",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyPassword) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyPassword) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

