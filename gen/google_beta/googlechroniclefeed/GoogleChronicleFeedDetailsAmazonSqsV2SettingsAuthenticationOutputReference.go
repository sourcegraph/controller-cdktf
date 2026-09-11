package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference interface {
	cdktn.ComplexObject
	AwsIamRoleAuth() GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuthOutputReference
	AwsIamRoleAuthInput() *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth
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
	InternalValue() *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthentication
	SetInternalValue(val *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthentication)
	SqsV2AccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuthOutputReference
	SqsV2AccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth
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
	PutAwsIamRoleAuth(value *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth)
	PutSqsV2AccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) AwsIamRoleAuth() GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuthOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuthOutputReference
	_jsii_.Get(
		j,
		"awsIamRoleAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) AwsIamRoleAuthInput() *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth {
	var returns *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth
	_jsii_.Get(
		j,
		"awsIamRoleAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) InternalValue() *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthentication {
	var returns *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) SqsV2AccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuthOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuthOutputReference
	_jsii_.Get(
		j,
		"sqsV2AccessKeySecretAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) SqsV2AccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth {
	var returns *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth
	_jsii_.Get(
		j,
		"sqsV2AccessKeySecretAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference_Override(g GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetInternalValue(val *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) PutAwsIamRoleAuth(value *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth) {
	if err := g.validatePutAwsIamRoleAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsIamRoleAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) PutSqsV2AccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth) {
	if err := g.validatePutSqsV2AccessKeySecretAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqsV2AccessKeySecretAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

