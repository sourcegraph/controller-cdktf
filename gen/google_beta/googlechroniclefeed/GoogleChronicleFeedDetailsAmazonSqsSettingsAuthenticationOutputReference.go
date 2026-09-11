package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference interface {
	cdktn.ComplexObject
	AdditionalS3AccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuthOutputReference
	AdditionalS3AccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication
	SetInternalValue(val *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication)
	SqsAccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference
	SqsAccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAdditionalS3AccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth)
	PutSqsAccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth)
	ResetAdditionalS3AccessKeySecretAuth()
	ResetSqsAccessKeySecretAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) AdditionalS3AccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuthOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuthOutputReference
	_jsii_.Get(
		j,
		"additionalS3AccessKeySecretAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) AdditionalS3AccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth {
	var returns *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth
	_jsii_.Get(
		j,
		"additionalS3AccessKeySecretAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) InternalValue() *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication {
	var returns *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) SqsAccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference
	_jsii_.Get(
		j,
		"sqsAccessKeySecretAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) SqsAccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth {
	var returns *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth
	_jsii_.Get(
		j,
		"sqsAccessKeySecretAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference_Override(g GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference)SetInternalValue(val *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) PutAdditionalS3AccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth) {
	if err := g.validatePutAdditionalS3AccessKeySecretAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalS3AccessKeySecretAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) PutSqsAccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth) {
	if err := g.validatePutSqsAccessKeySecretAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqsAccessKeySecretAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) ResetAdditionalS3AccessKeySecretAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalS3AccessKeySecretAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) ResetSqsAccessKeySecretAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetSqsAccessKeySecretAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

