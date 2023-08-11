package teamsrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsrule/internal"
)

type TeamsRuleRuleSettingsOutputReference interface {
	cdktf.ComplexObject
	AddHeaders() *map[string]*string
	SetAddHeaders(val *map[string]*string)
	AddHeadersInput() *map[string]*string
	AllowChildBypass() interface{}
	SetAllowChildBypass(val interface{})
	AllowChildBypassInput() interface{}
	AuditSsh() TeamsRuleRuleSettingsAuditSshOutputReference
	AuditSshInput() *TeamsRuleRuleSettingsAuditSsh
	BisoAdminControls() TeamsRuleRuleSettingsBisoAdminControlsOutputReference
	BisoAdminControlsInput() *TeamsRuleRuleSettingsBisoAdminControls
	BlockPageEnabled() interface{}
	SetBlockPageEnabled(val interface{})
	BlockPageEnabledInput() interface{}
	BlockPageReason() *string
	SetBlockPageReason(val *string)
	BlockPageReasonInput() *string
	BypassParentRule() interface{}
	SetBypassParentRule(val interface{})
	BypassParentRuleInput() interface{}
	CheckSession() TeamsRuleRuleSettingsCheckSessionOutputReference
	CheckSessionInput() *TeamsRuleRuleSettingsCheckSession
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Egress() TeamsRuleRuleSettingsEgressOutputReference
	EgressInput() *TeamsRuleRuleSettingsEgress
	// Experimental.
	Fqn() *string
	InsecureDisableDnssecValidation() interface{}
	SetInsecureDisableDnssecValidation(val interface{})
	InsecureDisableDnssecValidationInput() interface{}
	InternalValue() *TeamsRuleRuleSettings
	SetInternalValue(val *TeamsRuleRuleSettings)
	IpCategories() interface{}
	SetIpCategories(val interface{})
	IpCategoriesInput() interface{}
	L4Override() TeamsRuleRuleSettingsL4OverrideOutputReference
	L4OverrideInput() *TeamsRuleRuleSettingsL4Override
	OverrideHost() *string
	SetOverrideHost(val *string)
	OverrideHostInput() *string
	OverrideIps() *[]*string
	SetOverrideIps(val *[]*string)
	OverrideIpsInput() *[]*string
	PayloadLog() TeamsRuleRuleSettingsPayloadLogOutputReference
	PayloadLogInput() *TeamsRuleRuleSettingsPayloadLog
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UntrustedCert() TeamsRuleRuleSettingsUntrustedCertOutputReference
	UntrustedCertInput() *TeamsRuleRuleSettingsUntrustedCert
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAuditSsh(value *TeamsRuleRuleSettingsAuditSsh)
	PutBisoAdminControls(value *TeamsRuleRuleSettingsBisoAdminControls)
	PutCheckSession(value *TeamsRuleRuleSettingsCheckSession)
	PutEgress(value *TeamsRuleRuleSettingsEgress)
	PutL4Override(value *TeamsRuleRuleSettingsL4Override)
	PutPayloadLog(value *TeamsRuleRuleSettingsPayloadLog)
	PutUntrustedCert(value *TeamsRuleRuleSettingsUntrustedCert)
	ResetAddHeaders()
	ResetAllowChildBypass()
	ResetAuditSsh()
	ResetBisoAdminControls()
	ResetBlockPageEnabled()
	ResetBlockPageReason()
	ResetBypassParentRule()
	ResetCheckSession()
	ResetEgress()
	ResetInsecureDisableDnssecValidation()
	ResetIpCategories()
	ResetL4Override()
	ResetOverrideHost()
	ResetOverrideIps()
	ResetPayloadLog()
	ResetUntrustedCert()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamsRuleRuleSettingsOutputReference
type jsiiProxy_TeamsRuleRuleSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) AddHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"addHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) AddHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"addHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) AllowChildBypass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowChildBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) AllowChildBypassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowChildBypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) AuditSsh() TeamsRuleRuleSettingsAuditSshOutputReference {
	var returns TeamsRuleRuleSettingsAuditSshOutputReference
	_jsii_.Get(
		j,
		"auditSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) AuditSshInput() *TeamsRuleRuleSettingsAuditSsh {
	var returns *TeamsRuleRuleSettingsAuditSsh
	_jsii_.Get(
		j,
		"auditSshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BisoAdminControls() TeamsRuleRuleSettingsBisoAdminControlsOutputReference {
	var returns TeamsRuleRuleSettingsBisoAdminControlsOutputReference
	_jsii_.Get(
		j,
		"bisoAdminControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BisoAdminControlsInput() *TeamsRuleRuleSettingsBisoAdminControls {
	var returns *TeamsRuleRuleSettingsBisoAdminControls
	_jsii_.Get(
		j,
		"bisoAdminControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BlockPageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BlockPageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BlockPageReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockPageReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BlockPageReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockPageReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BypassParentRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassParentRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) BypassParentRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassParentRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) CheckSession() TeamsRuleRuleSettingsCheckSessionOutputReference {
	var returns TeamsRuleRuleSettingsCheckSessionOutputReference
	_jsii_.Get(
		j,
		"checkSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) CheckSessionInput() *TeamsRuleRuleSettingsCheckSession {
	var returns *TeamsRuleRuleSettingsCheckSession
	_jsii_.Get(
		j,
		"checkSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) Egress() TeamsRuleRuleSettingsEgressOutputReference {
	var returns TeamsRuleRuleSettingsEgressOutputReference
	_jsii_.Get(
		j,
		"egress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) EgressInput() *TeamsRuleRuleSettingsEgress {
	var returns *TeamsRuleRuleSettingsEgress
	_jsii_.Get(
		j,
		"egressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) InsecureDisableDnssecValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureDisableDnssecValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) InsecureDisableDnssecValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureDisableDnssecValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) InternalValue() *TeamsRuleRuleSettings {
	var returns *TeamsRuleRuleSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) IpCategories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) IpCategoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) L4Override() TeamsRuleRuleSettingsL4OverrideOutputReference {
	var returns TeamsRuleRuleSettingsL4OverrideOutputReference
	_jsii_.Get(
		j,
		"l4Override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) L4OverrideInput() *TeamsRuleRuleSettingsL4Override {
	var returns *TeamsRuleRuleSettingsL4Override
	_jsii_.Get(
		j,
		"l4OverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) OverrideHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) OverrideHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) OverrideIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overrideIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) OverrideIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overrideIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PayloadLog() TeamsRuleRuleSettingsPayloadLogOutputReference {
	var returns TeamsRuleRuleSettingsPayloadLogOutputReference
	_jsii_.Get(
		j,
		"payloadLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PayloadLogInput() *TeamsRuleRuleSettingsPayloadLog {
	var returns *TeamsRuleRuleSettingsPayloadLog
	_jsii_.Get(
		j,
		"payloadLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) UntrustedCert() TeamsRuleRuleSettingsUntrustedCertOutputReference {
	var returns TeamsRuleRuleSettingsUntrustedCertOutputReference
	_jsii_.Get(
		j,
		"untrustedCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference) UntrustedCertInput() *TeamsRuleRuleSettingsUntrustedCert {
	var returns *TeamsRuleRuleSettingsUntrustedCert
	_jsii_.Get(
		j,
		"untrustedCertInput",
		&returns,
	)
	return returns
}


func NewTeamsRuleRuleSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsRuleRuleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsRuleRuleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsRuleRuleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsRule.TeamsRuleRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsRuleRuleSettingsOutputReference_Override(t TeamsRuleRuleSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsRule.TeamsRuleRuleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetAddHeaders(val *map[string]*string) {
	if err := j.validateSetAddHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addHeaders",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetAllowChildBypass(val interface{}) {
	if err := j.validateSetAllowChildBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowChildBypass",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetBlockPageEnabled(val interface{}) {
	if err := j.validateSetBlockPageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPageEnabled",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetBlockPageReason(val *string) {
	if err := j.validateSetBlockPageReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPageReason",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetBypassParentRule(val interface{}) {
	if err := j.validateSetBypassParentRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassParentRule",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetInsecureDisableDnssecValidation(val interface{}) {
	if err := j.validateSetInsecureDisableDnssecValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureDisableDnssecValidation",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetInternalValue(val *TeamsRuleRuleSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetIpCategories(val interface{}) {
	if err := j.validateSetIpCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipCategories",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetOverrideHost(val *string) {
	if err := j.validateSetOverrideHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideHost",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetOverrideIps(val *[]*string) {
	if err := j.validateSetOverrideIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideIps",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsRuleRuleSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutAuditSsh(value *TeamsRuleRuleSettingsAuditSsh) {
	if err := t.validatePutAuditSshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuditSsh",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutBisoAdminControls(value *TeamsRuleRuleSettingsBisoAdminControls) {
	if err := t.validatePutBisoAdminControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBisoAdminControls",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutCheckSession(value *TeamsRuleRuleSettingsCheckSession) {
	if err := t.validatePutCheckSessionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCheckSession",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutEgress(value *TeamsRuleRuleSettingsEgress) {
	if err := t.validatePutEgressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEgress",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutL4Override(value *TeamsRuleRuleSettingsL4Override) {
	if err := t.validatePutL4OverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putL4Override",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutPayloadLog(value *TeamsRuleRuleSettingsPayloadLog) {
	if err := t.validatePutPayloadLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPayloadLog",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) PutUntrustedCert(value *TeamsRuleRuleSettingsUntrustedCert) {
	if err := t.validatePutUntrustedCertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUntrustedCert",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetAddHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetAddHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetAllowChildBypass() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowChildBypass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetAuditSsh() {
	_jsii_.InvokeVoid(
		t,
		"resetAuditSsh",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetBisoAdminControls() {
	_jsii_.InvokeVoid(
		t,
		"resetBisoAdminControls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetBlockPageEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockPageEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetBlockPageReason() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockPageReason",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetBypassParentRule() {
	_jsii_.InvokeVoid(
		t,
		"resetBypassParentRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetCheckSession() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckSession",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetEgress() {
	_jsii_.InvokeVoid(
		t,
		"resetEgress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetInsecureDisableDnssecValidation() {
	_jsii_.InvokeVoid(
		t,
		"resetInsecureDisableDnssecValidation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetIpCategories() {
	_jsii_.InvokeVoid(
		t,
		"resetIpCategories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetL4Override() {
	_jsii_.InvokeVoid(
		t,
		"resetL4Override",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetOverrideHost() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideHost",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetOverrideIps() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideIps",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetPayloadLog() {
	_jsii_.InvokeVoid(
		t,
		"resetPayloadLog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ResetUntrustedCert() {
	_jsii_.InvokeVoid(
		t,
		"resetUntrustedCert",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsRuleRuleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

