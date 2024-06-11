package policymfadefault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policymfadefault/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_mfa_default okta_policy_mfa_default}.
type PolicyMfaDefault interface {
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
	DefaultIncludedGroupId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
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
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetDuo()
	ResetExternalIdp()
	ResetFidoU2F()
	ResetFidoWebauthn()
	ResetGoogleOtp()
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
	ResetRsaToken()
	ResetSecurityQuestion()
	ResetSymantecVip()
	ResetWebauthn()
	ResetYubikeyToken()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PolicyMfaDefault
type jsiiProxy_PolicyMfaDefault struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyMfaDefault) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) DefaultIncludedGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIncludedGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Duo() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"duo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) DuoInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"duoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) ExternalIdp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"externalIdp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) ExternalIdpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"externalIdpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) FidoU2F() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoU2F",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) FidoU2FInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoU2FInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) FidoWebauthn() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoWebauthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) FidoWebauthnInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fidoWebauthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) GoogleOtp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleOtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) GoogleOtpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleOtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Hotp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hotp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) HotpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hotpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) IsOie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) IsOieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaCall() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaCallInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaEmail() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaEmailInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaOtp() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaOtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaOtpInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaOtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaPassword() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaPasswordInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaPush() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaPushInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaQuestion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaQuestionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaSms() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaSms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaSmsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaSmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaVerify() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OktaVerifyInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"oktaVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OnpremMfa() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"onpremMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) OnpremMfaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"onpremMfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) PhoneNumber() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) PhoneNumberInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) RsaToken() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rsaToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) RsaTokenInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rsaTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) SecurityQuestion() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"securityQuestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) SecurityQuestionInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"securityQuestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) SymantecVip() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"symantecVip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) SymantecVipInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"symantecVipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) Webauthn() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webauthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) WebauthnInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webauthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) YubikeyToken() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"yubikeyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyMfaDefault) YubikeyTokenInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"yubikeyTokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_mfa_default okta_policy_mfa_default} Resource.
func NewPolicyMfaDefault(scope constructs.Construct, id *string, config *PolicyMfaDefaultConfig) PolicyMfaDefault {
	_init_.Initialize()

	if err := validateNewPolicyMfaDefaultParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyMfaDefault{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_mfa_default okta_policy_mfa_default} Resource.
func NewPolicyMfaDefault_Override(p PolicyMfaDefault, scope constructs.Construct, id *string, config *PolicyMfaDefaultConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetDuo(val *map[string]*string) {
	if err := j.validateSetDuoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duo",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetExternalIdp(val *map[string]*string) {
	if err := j.validateSetExternalIdpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIdp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetFidoU2F(val *map[string]*string) {
	if err := j.validateSetFidoU2FParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidoU2F",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetFidoWebauthn(val *map[string]*string) {
	if err := j.validateSetFidoWebauthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidoWebauthn",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetGoogleOtp(val *map[string]*string) {
	if err := j.validateSetGoogleOtpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleOtp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetHotp(val *map[string]*string) {
	if err := j.validateSetHotpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetIsOie(val interface{}) {
	if err := j.validateSetIsOieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOie",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaCall(val *map[string]*string) {
	if err := j.validateSetOktaCallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaCall",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaEmail(val *map[string]*string) {
	if err := j.validateSetOktaEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaEmail",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaOtp(val *map[string]*string) {
	if err := j.validateSetOktaOtpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaOtp",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaPassword(val *map[string]*string) {
	if err := j.validateSetOktaPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaPassword",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaPush(val *map[string]*string) {
	if err := j.validateSetOktaPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaPush",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaQuestion(val *map[string]*string) {
	if err := j.validateSetOktaQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaQuestion",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaSms(val *map[string]*string) {
	if err := j.validateSetOktaSmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaSms",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOktaVerify(val *map[string]*string) {
	if err := j.validateSetOktaVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oktaVerify",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetOnpremMfa(val *map[string]*string) {
	if err := j.validateSetOnpremMfaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onpremMfa",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetPhoneNumber(val *map[string]*string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetRsaToken(val *map[string]*string) {
	if err := j.validateSetRsaTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaToken",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetSecurityQuestion(val *map[string]*string) {
	if err := j.validateSetSecurityQuestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityQuestion",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetSymantecVip(val *map[string]*string) {
	if err := j.validateSetSymantecVipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symantecVip",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetWebauthn(val *map[string]*string) {
	if err := j.validateSetWebauthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webauthn",
		val,
	)
}

func (j *jsiiProxy_PolicyMfaDefault)SetYubikeyToken(val *map[string]*string) {
	if err := j.validateSetYubikeyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yubikeyToken",
		val,
	)
}

// Generates CDKTF code for importing a PolicyMfaDefault resource upon running "cdktf plan <stack-name>".
func PolicyMfaDefault_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePolicyMfaDefault_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func PolicyMfaDefault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyMfaDefault_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyMfaDefault_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyMfaDefault_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyMfaDefault_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyMfaDefault_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyMfaDefault_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyMfaDefault.PolicyMfaDefault",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PolicyMfaDefault) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyMfaDefault) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PolicyMfaDefault) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PolicyMfaDefault) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PolicyMfaDefault) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PolicyMfaDefault) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PolicyMfaDefault) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PolicyMfaDefault) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PolicyMfaDefault) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyMfaDefault) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetDuo() {
	_jsii_.InvokeVoid(
		p,
		"resetDuo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetExternalIdp() {
	_jsii_.InvokeVoid(
		p,
		"resetExternalIdp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetFidoU2F() {
	_jsii_.InvokeVoid(
		p,
		"resetFidoU2F",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetFidoWebauthn() {
	_jsii_.InvokeVoid(
		p,
		"resetFidoWebauthn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetGoogleOtp() {
	_jsii_.InvokeVoid(
		p,
		"resetGoogleOtp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetHotp() {
	_jsii_.InvokeVoid(
		p,
		"resetHotp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetIsOie() {
	_jsii_.InvokeVoid(
		p,
		"resetIsOie",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaCall() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaCall",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaEmail() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaEmail",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaOtp() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaOtp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaPush() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaPush",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaQuestion() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaQuestion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaSms() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaSms",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOktaVerify() {
	_jsii_.InvokeVoid(
		p,
		"resetOktaVerify",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOnpremMfa() {
	_jsii_.InvokeVoid(
		p,
		"resetOnpremMfa",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		p,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetRsaToken() {
	_jsii_.InvokeVoid(
		p,
		"resetRsaToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetSecurityQuestion() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuestion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetSymantecVip() {
	_jsii_.InvokeVoid(
		p,
		"resetSymantecVip",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetWebauthn() {
	_jsii_.InvokeVoid(
		p,
		"resetWebauthn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) ResetYubikeyToken() {
	_jsii_.InvokeVoid(
		p,
		"resetYubikeyToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyMfaDefault) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyMfaDefault) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

