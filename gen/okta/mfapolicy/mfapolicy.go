package mfapolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/mfapolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/mfa_policy okta_mfa_policy}.
type MfaPolicy interface {
	cdktf.TerraformResource
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
	Duo() *map[string]*string
	SetDuo(val *map[string]*string)
	DuoInput() *map[string]*string
	ExternalIdp() *map[string]*string
	SetExternalIdp(val *map[string]*string)
	ExternalIdpInput() *map[string]*string
	FidoU2F() *map[string]*string
	SetFidoU2F(val *map[string]*string)
	FidoU2FInput() *map[string]*string
	FidoWebauthn() *map[string]*string
	SetFidoWebauthn(val *map[string]*string)
	FidoWebauthnInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleOtp() *map[string]*string
	SetGoogleOtp(val *map[string]*string)
	GoogleOtpInput() *map[string]*string
	GroupsIncluded() *[]*string
	SetGroupsIncluded(val *[]*string)
	GroupsIncludedInput() *[]*string
	Hotp() *map[string]*string
	SetHotp(val *map[string]*string)
	HotpInput() *map[string]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsOie() interface{}
	SetIsOie(val interface{})
	IsOieInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OktaCall() *map[string]*string
	SetOktaCall(val *map[string]*string)
	OktaCallInput() *map[string]*string
	OktaEmail() *map[string]*string
	SetOktaEmail(val *map[string]*string)
	OktaEmailInput() *map[string]*string
	OktaOtp() *map[string]*string
	SetOktaOtp(val *map[string]*string)
	OktaOtpInput() *map[string]*string
	OktaPassword() *map[string]*string
	SetOktaPassword(val *map[string]*string)
	OktaPasswordInput() *map[string]*string
	OktaPush() *map[string]*string
	SetOktaPush(val *map[string]*string)
	OktaPushInput() *map[string]*string
	OktaQuestion() *map[string]*string
	SetOktaQuestion(val *map[string]*string)
	OktaQuestionInput() *map[string]*string
	OktaSms() *map[string]*string
	SetOktaSms(val *map[string]*string)
	OktaSmsInput() *map[string]*string
	OktaVerify() *map[string]*string
	SetOktaVerify(val *map[string]*string)
	OktaVerifyInput() *map[string]*string
	OnpremMfa() *map[string]*string
	SetOnpremMfa(val *map[string]*string)
	OnpremMfaInput() *map[string]*string
	PhoneNumber() *map[string]*string
	SetPhoneNumber(val *map[string]*string)
	PhoneNumberInput() *map[string]*string
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
	// Experimental.
	RawOverrides() interface{}
	RsaToken() *map[string]*string
	SetRsaToken(val *map[string]*string)
	RsaTokenInput() *map[string]*string
	SecurityQuestion() *map[string]*string
	SetSecurityQuestion(val *map[string]*string)
	SecurityQuestionInput() *map[string]*string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SymantecVip() *map[string]*string
	SetSymantecVip(val *map[string]*string)
	SymantecVipInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Webauthn() *map[string]*string
	SetWebauthn(val *map[string]*string)
	WebauthnInput() *map[string]*string
	YubikeyToken() *map[string]*string
	SetYubikeyToken(val *map[string]*string)
	YubikeyTokenInput() *map[string]*string
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
	ResetDescription()
	ResetDuo()
	ResetExternalIdp()
	ResetFidoU2F()
	ResetFidoWebauthn()
	ResetGoogleOtp()
	ResetGroupsIncluded()
	ResetHotp()
	ResetId()
	ResetIsOie()
	ResetOktaCall()
	ResetOktaEmail()
	ResetOktaOtp()
	ResetOktaPassword()
	ResetOktaPush()
	ResetOktaQuestion()
	ResetOktaSms()
	ResetOktaVerify()
	ResetOnpremMfa()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhoneNumber()
	ResetPriority()
	ResetRsaToken()
	ResetSecurityQuestion()
	ResetStatus()
	ResetSymantecVip()
	ResetWebauthn()
	ResetYubikeyToken()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MfaPolicy
type jsiiProxy_MfaPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MfaPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Duo() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"duo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) DuoInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"duoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) ExternalIdp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"externalIdp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) ExternalIdpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"externalIdpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) FidoU2F() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoU2F",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) FidoU2FInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoU2FInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) FidoWebauthn() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoWebauthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) FidoWebauthnInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoWebauthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) GoogleOtp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleOtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) GoogleOtpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleOtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) GroupsIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) GroupsIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Hotp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hotp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) HotpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hotpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) IsOie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) IsOieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaCall() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaCallInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaEmail() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaEmailInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaOtp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaOtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaOtpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaOtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaPassword() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaPasswordInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaPush() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaPushInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaQuestion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaQuestionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaSms() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaSms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaSmsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaSmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaVerify() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OktaVerifyInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OnpremMfa() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"onpremMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) OnpremMfaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"onpremMfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) PhoneNumber() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) PhoneNumberInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) RsaToken() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rsaToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) RsaTokenInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rsaTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) SecurityQuestion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"securityQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) SecurityQuestionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"securityQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) SymantecVip() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"symantecVip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) SymantecVipInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"symantecVipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) Webauthn() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webauthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) WebauthnInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webauthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) YubikeyToken() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"yubikeyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaPolicy) YubikeyTokenInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"yubikeyTokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/mfa_policy okta_mfa_policy} Resource.
func NewMfaPolicy(scope constructs.Construct, id *string, config *MfaPolicyConfig) MfaPolicy {
	_init_.Initialize()

	if err := validateNewMfaPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MfaPolicy{}

	_jsii_.Create(
		"@cdktf/provider-okta.mfaPolicy.MfaPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/mfa_policy okta_mfa_policy} Resource.
func NewMfaPolicy_Override(m MfaPolicy, scope constructs.Construct, id *string, config *MfaPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.mfaPolicy.MfaPolicy",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MfaPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetDuo(val *map[string]*string) {
	if err := j.validateSetDuoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duo",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetExternalIdp(val *map[string]*string) {
	if err := j.validateSetExternalIdpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIdp",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetFidoU2F(val *map[string]*string) {
	if err := j.validateSetFidoU2FParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidoU2F",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetFidoWebauthn(val *map[string]*string) {
	if err := j.validateSetFidoWebauthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidoWebauthn",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetGoogleOtp(val *map[string]*string) {
	if err := j.validateSetGoogleOtpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleOtp",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetGroupsIncluded(val *[]*string) {
	if err := j.validateSetGroupsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsIncluded",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetHotp(val *map[string]*string) {
	if err := j.validateSetHotpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotp",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetIsOie(val interface{}) {
	if err := j.validateSetIsOieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOie",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaCall(val *map[string]*string) {
	if err := j.validateSetOktaCallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaCall",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaEmail(val *map[string]*string) {
	if err := j.validateSetOktaEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaEmail",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaOtp(val *map[string]*string) {
	if err := j.validateSetOktaOtpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaOtp",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaPassword(val *map[string]*string) {
	if err := j.validateSetOktaPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaPassword",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaPush(val *map[string]*string) {
	if err := j.validateSetOktaPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaPush",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaQuestion(val *map[string]*string) {
	if err := j.validateSetOktaQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaQuestion",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaSms(val *map[string]*string) {
	if err := j.validateSetOktaSmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaSms",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOktaVerify(val *map[string]*string) {
	if err := j.validateSetOktaVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaVerify",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetOnpremMfa(val *map[string]*string) {
	if err := j.validateSetOnpremMfaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onpremMfa",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetPhoneNumber(val *map[string]*string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetRsaToken(val *map[string]*string) {
	if err := j.validateSetRsaTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaToken",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetSecurityQuestion(val *map[string]*string) {
	if err := j.validateSetSecurityQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityQuestion",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetSymantecVip(val *map[string]*string) {
	if err := j.validateSetSymantecVipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symantecVip",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetWebauthn(val *map[string]*string) {
	if err := j.validateSetWebauthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webauthn",
		val,
	)
}

func (j *jsiiProxy_MfaPolicy)SetYubikeyToken(val *map[string]*string) {
	if err := j.validateSetYubikeyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yubikeyToken",
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
func MfaPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMfaPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.mfaPolicy.MfaPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MfaPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMfaPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.mfaPolicy.MfaPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MfaPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMfaPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.mfaPolicy.MfaPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MfaPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.mfaPolicy.MfaPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MfaPolicy) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MfaPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MfaPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetDuo() {
	_jsii_.InvokeVoid(
		m,
		"resetDuo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetExternalIdp() {
	_jsii_.InvokeVoid(
		m,
		"resetExternalIdp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetFidoU2F() {
	_jsii_.InvokeVoid(
		m,
		"resetFidoU2F",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetFidoWebauthn() {
	_jsii_.InvokeVoid(
		m,
		"resetFidoWebauthn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetGoogleOtp() {
	_jsii_.InvokeVoid(
		m,
		"resetGoogleOtp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetGroupsIncluded() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupsIncluded",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetHotp() {
	_jsii_.InvokeVoid(
		m,
		"resetHotp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetIsOie() {
	_jsii_.InvokeVoid(
		m,
		"resetIsOie",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaCall() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaCall",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaEmail() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaEmail",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaOtp() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaOtp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaPush() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaPush",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaQuestion() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaQuestion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaSms() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaSms",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOktaVerify() {
	_jsii_.InvokeVoid(
		m,
		"resetOktaVerify",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOnpremMfa() {
	_jsii_.InvokeVoid(
		m,
		"resetOnpremMfa",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		m,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetPriority() {
	_jsii_.InvokeVoid(
		m,
		"resetPriority",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetRsaToken() {
	_jsii_.InvokeVoid(
		m,
		"resetRsaToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetSecurityQuestion() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityQuestion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetSymantecVip() {
	_jsii_.InvokeVoid(
		m,
		"resetSymantecVip",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetWebauthn() {
	_jsii_.InvokeVoid(
		m,
		"resetWebauthn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) ResetYubikeyToken() {
	_jsii_.InvokeVoid(
		m,
		"resetYubikeyToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

