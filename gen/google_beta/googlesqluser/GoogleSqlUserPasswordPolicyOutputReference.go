package googlesqluser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqluser/internal"
)

type GoogleSqlUserPasswordPolicyOutputReference interface {
	cdktf.ComplexObject
	AllowedFailedAttempts() *float64
	SetAllowedFailedAttempts(val *float64)
	AllowedFailedAttemptsInput() *float64
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
	EnableFailedAttemptsCheck() interface{}
	SetEnableFailedAttemptsCheck(val interface{})
	EnableFailedAttemptsCheckInput() interface{}
	EnablePasswordVerification() interface{}
	SetEnablePasswordVerification(val interface{})
	EnablePasswordVerificationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlUserPasswordPolicy
	SetInternalValue(val *GoogleSqlUserPasswordPolicy)
	PasswordExpirationDuration() *string
	SetPasswordExpirationDuration(val *string)
	PasswordExpirationDurationInput() *string
	Status() GoogleSqlUserPasswordPolicyStatusList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAllowedFailedAttempts()
	ResetEnableFailedAttemptsCheck()
	ResetEnablePasswordVerification()
	ResetPasswordExpirationDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlUserPasswordPolicyOutputReference
type jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) AllowedFailedAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedFailedAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) AllowedFailedAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedFailedAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) EnableFailedAttemptsCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFailedAttemptsCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) EnableFailedAttemptsCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFailedAttemptsCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) EnablePasswordVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePasswordVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) EnablePasswordVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePasswordVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) InternalValue() *GoogleSqlUserPasswordPolicy {
	var returns *GoogleSqlUserPasswordPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) PasswordExpirationDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordExpirationDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) PasswordExpirationDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordExpirationDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) Status() GoogleSqlUserPasswordPolicyStatusList {
	var returns GoogleSqlUserPasswordPolicyStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSqlUserPasswordPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlUserPasswordPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlUserPasswordPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlUser.GoogleSqlUserPasswordPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlUserPasswordPolicyOutputReference_Override(g GoogleSqlUserPasswordPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlUser.GoogleSqlUserPasswordPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetAllowedFailedAttempts(val *float64) {
	if err := j.validateSetAllowedFailedAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedFailedAttempts",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetEnableFailedAttemptsCheck(val interface{}) {
	if err := j.validateSetEnableFailedAttemptsCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFailedAttemptsCheck",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetEnablePasswordVerification(val interface{}) {
	if err := j.validateSetEnablePasswordVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePasswordVerification",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetInternalValue(val *GoogleSqlUserPasswordPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetPasswordExpirationDuration(val *string) {
	if err := j.validateSetPasswordExpirationDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordExpirationDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ResetAllowedFailedAttempts() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedFailedAttempts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ResetEnableFailedAttemptsCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableFailedAttemptsCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ResetEnablePasswordVerification() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePasswordVerification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ResetPasswordExpirationDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordExpirationDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlUserPasswordPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

