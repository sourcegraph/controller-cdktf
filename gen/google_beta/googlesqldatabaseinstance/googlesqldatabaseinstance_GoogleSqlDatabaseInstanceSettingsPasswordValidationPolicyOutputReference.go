package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference interface {
	cdktf.ComplexObject
	Complexity() *string
	SetComplexity(val *string)
	ComplexityInput() *string
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
	DisallowUsernameSubstring() interface{}
	SetDisallowUsernameSubstring(val interface{})
	DisallowUsernameSubstringInput() interface{}
	EnablePasswordPolicy() interface{}
	SetEnablePasswordPolicy(val interface{})
	EnablePasswordPolicyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy
	SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy)
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	PasswordChangeInterval() *string
	SetPasswordChangeInterval(val *string)
	PasswordChangeIntervalInput() *string
	ReuseInterval() *float64
	SetReuseInterval(val *float64)
	ReuseIntervalInput() *float64
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
	ResetComplexity()
	ResetDisallowUsernameSubstring()
	ResetMinLength()
	ResetPasswordChangeInterval()
	ResetReuseInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) Complexity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ComplexityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) DisallowUsernameSubstring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowUsernameSubstring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) DisallowUsernameSubstringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowUsernameSubstringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) EnablePasswordPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePasswordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) EnablePasswordPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePasswordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) InternalValue() *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy {
	var returns *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) PasswordChangeInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordChangeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) PasswordChangeIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordChangeIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ReuseInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reuseInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ReuseIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reuseIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetComplexity(val *string) {
	if err := j.validateSetComplexityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexity",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetDisallowUsernameSubstring(val interface{}) {
	if err := j.validateSetDisallowUsernameSubstringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowUsernameSubstring",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetEnablePasswordPolicy(val interface{}) {
	if err := j.validateSetEnablePasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePasswordPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetMinLength(val *float64) {
	if err := j.validateSetMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetPasswordChangeInterval(val *string) {
	if err := j.validateSetPasswordChangeIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordChangeInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetReuseInterval(val *float64) {
	if err := j.validateSetReuseIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reuseInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ResetComplexity() {
	_jsii_.InvokeVoid(
		g,
		"resetComplexity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ResetDisallowUsernameSubstring() {
	_jsii_.InvokeVoid(
		g,
		"resetDisallowUsernameSubstring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ResetMinLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMinLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ResetPasswordChangeInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordChangeInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ResetReuseInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetReuseInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

