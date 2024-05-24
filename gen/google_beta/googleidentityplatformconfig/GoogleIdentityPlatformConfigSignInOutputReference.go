package googleidentityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleidentityplatformconfig/internal"
)

type GoogleIdentityPlatformConfigSignInOutputReference interface {
	cdktf.ComplexObject
	AllowDuplicateEmails() interface{}
	SetAllowDuplicateEmails(val interface{})
	AllowDuplicateEmailsInput() interface{}
	Anonymous() GoogleIdentityPlatformConfigSignInAnonymousOutputReference
	AnonymousInput() *GoogleIdentityPlatformConfigSignInAnonymous
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
	Email() GoogleIdentityPlatformConfigSignInEmailOutputReference
	EmailInput() *GoogleIdentityPlatformConfigSignInEmail
	// Experimental.
	Fqn() *string
	HashConfig() GoogleIdentityPlatformConfigSignInHashConfigList
	InternalValue() *GoogleIdentityPlatformConfigSignIn
	SetInternalValue(val *GoogleIdentityPlatformConfigSignIn)
	PhoneNumber() GoogleIdentityPlatformConfigSignInPhoneNumberOutputReference
	PhoneNumberInput() *GoogleIdentityPlatformConfigSignInPhoneNumber
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
	PutAnonymous(value *GoogleIdentityPlatformConfigSignInAnonymous)
	PutEmail(value *GoogleIdentityPlatformConfigSignInEmail)
	PutPhoneNumber(value *GoogleIdentityPlatformConfigSignInPhoneNumber)
	ResetAllowDuplicateEmails()
	ResetAnonymous()
	ResetEmail()
	ResetPhoneNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIdentityPlatformConfigSignInOutputReference
type jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) AllowDuplicateEmails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) AllowDuplicateEmailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) Anonymous() GoogleIdentityPlatformConfigSignInAnonymousOutputReference {
	var returns GoogleIdentityPlatformConfigSignInAnonymousOutputReference
	_jsii_.Get(
		j,
		"anonymous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) AnonymousInput() *GoogleIdentityPlatformConfigSignInAnonymous {
	var returns *GoogleIdentityPlatformConfigSignInAnonymous
	_jsii_.Get(
		j,
		"anonymousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) Email() GoogleIdentityPlatformConfigSignInEmailOutputReference {
	var returns GoogleIdentityPlatformConfigSignInEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) EmailInput() *GoogleIdentityPlatformConfigSignInEmail {
	var returns *GoogleIdentityPlatformConfigSignInEmail
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) HashConfig() GoogleIdentityPlatformConfigSignInHashConfigList {
	var returns GoogleIdentityPlatformConfigSignInHashConfigList
	_jsii_.Get(
		j,
		"hashConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) InternalValue() *GoogleIdentityPlatformConfigSignIn {
	var returns *GoogleIdentityPlatformConfigSignIn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) PhoneNumber() GoogleIdentityPlatformConfigSignInPhoneNumberOutputReference {
	var returns GoogleIdentityPlatformConfigSignInPhoneNumberOutputReference
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) PhoneNumberInput() *GoogleIdentityPlatformConfigSignInPhoneNumber {
	var returns *GoogleIdentityPlatformConfigSignInPhoneNumber
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIdentityPlatformConfigSignInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIdentityPlatformConfigSignInOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIdentityPlatformConfigSignInOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfigSignInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIdentityPlatformConfigSignInOutputReference_Override(g GoogleIdentityPlatformConfigSignInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfigSignInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference)SetAllowDuplicateEmails(val interface{}) {
	if err := j.validateSetAllowDuplicateEmailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicateEmails",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference)SetInternalValue(val *GoogleIdentityPlatformConfigSignIn) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) PutAnonymous(value *GoogleIdentityPlatformConfigSignInAnonymous) {
	if err := g.validatePutAnonymousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnonymous",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) PutEmail(value *GoogleIdentityPlatformConfigSignInEmail) {
	if err := g.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEmail",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) PutPhoneNumber(value *GoogleIdentityPlatformConfigSignInPhoneNumber) {
	if err := g.validatePutPhoneNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPhoneNumber",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ResetAllowDuplicateEmails() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowDuplicateEmails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ResetAnonymous() {
	_jsii_.InvokeVoid(
		g,
		"resetAnonymous",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		g,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigSignInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

