package policymfa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policymfa/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_mfa okta_policy_mfa}.
type PolicyMfa interface {
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

// The jsii proxy struct for PolicyMfa
type jsiiProxy_PolicyMfa struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyMfa) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Duo() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"duo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) DuoInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"duoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) ExternalIdp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"externalIdp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) ExternalIdpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"externalIdpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) FidoU2F() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoU2F",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) FidoU2FInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoU2FInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) FidoWebauthn() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoWebauthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) FidoWebauthnInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoWebauthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) GoogleOtp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleOtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) GoogleOtpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleOtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) GroupsIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) GroupsIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Hotp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hotp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) HotpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hotpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) IsOie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) IsOieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaCall() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaCallInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaEmail() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaEmailInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaOtp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaOtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaOtpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaOtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaPassword() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaPasswordInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaPush() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaPushInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaQuestion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaQuestionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaSms() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaSms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaSmsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaSmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaVerify() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OktaVerifyInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OnpremMfa() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"onpremMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) OnpremMfaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"onpremMfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) PhoneNumber() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) PhoneNumberInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) RsaToken() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rsaToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) RsaTokenInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rsaTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) SecurityQuestion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"securityQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) SecurityQuestionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"securityQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) SymantecVip() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"symantecVip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) SymantecVipInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"symantecVipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) Webauthn() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webauthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) WebauthnInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webauthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) YubikeyToken() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"yubikeyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfa) YubikeyTokenInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"yubikeyTokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_mfa okta_policy_mfa} Resource.
func NewPolicyMfa(scope constructs.Construct, id *string, config *PolicyMfaConfig) PolicyMfa {
	_init_.Initialize()

	if err := validateNewPolicyMfaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyMfa{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyMfa.PolicyMfa",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_mfa okta_policy_mfa} Resource.
func NewPolicyMfa_Override(p PolicyMfa, scope constructs.Construct, id *string, config *PolicyMfaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyMfa.PolicyMfa",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyMfa)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetDuo(val *map[string]*string) {
	if err := j.validateSetDuoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duo",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetExternalIdp(val *map[string]*string) {
	if err := j.validateSetExternalIdpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIdp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetFidoU2F(val *map[string]*string) {
	if err := j.validateSetFidoU2FParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidoU2F",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetFidoWebauthn(val *map[string]*string) {
	if err := j.validateSetFidoWebauthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidoWebauthn",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetGoogleOtp(val *map[string]*string) {
	if err := j.validateSetGoogleOtpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleOtp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetGroupsIncluded(val *[]*string) {
	if err := j.validateSetGroupsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsIncluded",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetHotp(val *map[string]*string) {
	if err := j.validateSetHotpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetIsOie(val interface{}) {
	if err := j.validateSetIsOieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOie",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaCall(val *map[string]*string) {
	if err := j.validateSetOktaCallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaCall",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaEmail(val *map[string]*string) {
	if err := j.validateSetOktaEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaEmail",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaOtp(val *map[string]*string) {
	if err := j.validateSetOktaOtpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaOtp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaPassword(val *map[string]*string) {
	if err := j.validateSetOktaPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaPassword",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaPush(val *map[string]*string) {
	if err := j.validateSetOktaPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaPush",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaQuestion(val *map[string]*string) {
	if err := j.validateSetOktaQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaQuestion",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaSms(val *map[string]*string) {
	if err := j.validateSetOktaSmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaSms",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOktaVerify(val *map[string]*string) {
	if err := j.validateSetOktaVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaVerify",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetOnpremMfa(val *map[string]*string) {
	if err := j.validateSetOnpremMfaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onpremMfa",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetPhoneNumber(val *map[string]*string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetRsaToken(val *map[string]*string) {
	if err := j.validateSetRsaTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaToken",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetSecurityQuestion(val *map[string]*string) {
	if err := j.validateSetSecurityQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityQuestion",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetSymantecVip(val *map[string]*string) {
	if err := j.validateSetSymantecVipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symantecVip",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetWebauthn(val *map[string]*string) {
	if err := j.validateSetWebauthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webauthn",
		val,
	)
}

func (j *jsiiProxy_PolicyMfa)SetYubikeyToken(val *map[string]*string) {
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
func PolicyMfa_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyMfa_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfa.PolicyMfa",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyMfa_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyMfa_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfa.PolicyMfa",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyMfa_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyMfa_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfa.PolicyMfa",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyMfa_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyMfa.PolicyMfa",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyMfa) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyMfa) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PolicyMfa) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyMfa) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PolicyMfa) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PolicyMfa) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PolicyMfa) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PolicyMfa) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PolicyMfa) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PolicyMfa) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PolicyMfa) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyMfa) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyMfa) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetDuo() {
	_jsii_.InvokeVoid(
		p,
		"resetDuo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetExternalIdp() {
	_jsii_.InvokeVoid(
		p,
		"resetExternalIdp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetFidoU2F() {
	_jsii_.InvokeVoid(
		p,
		"resetFidoU2F",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetFidoWebauthn() {
	_jsii_.InvokeVoid(
		p,
		"resetFidoWebauthn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetGoogleOtp() {
	_jsii_.InvokeVoid(
		p,
		"resetGoogleOtp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetGroupsIncluded() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupsIncluded",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetHotp() {
	_jsii_.InvokeVoid(
		p,
		"resetHotp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetIsOie() {
	_jsii_.InvokeVoid(
		p,
		"resetIsOie",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaCall() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaCall",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaEmail() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaEmail",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaOtp() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaOtp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaPush() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaPush",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaQuestion() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaQuestion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaSms() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaSms",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOktaVerify() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaVerify",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOnpremMfa() {
	_jsii_.InvokeVoid(
		p,
		"resetOnpremMfa",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		p,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetPriority() {
	_jsii_.InvokeVoid(
		p,
		"resetPriority",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetRsaToken() {
	_jsii_.InvokeVoid(
		p,
		"resetRsaToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetSecurityQuestion() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuestion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetStatus() {
	_jsii_.InvokeVoid(
		p,
		"resetStatus",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetSymantecVip() {
	_jsii_.InvokeVoid(
		p,
		"resetSymantecVip",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetWebauthn() {
	_jsii_.InvokeVoid(
		p,
		"resetWebauthn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) ResetYubikeyToken() {
	_jsii_.InvokeVoid(
		p,
		"resetYubikeyToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfa) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfa) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfa) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfa) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

