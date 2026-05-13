package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference interface {
	cdktf.ComplexObject
	AccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuthOutputReference
	AccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth
	AwsIamRoleAuth() GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference
	AwsIamRoleAuthInput() *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth
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
	InternalValue() *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication
	SetInternalValue(val *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication)
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
	PutAccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth)
	PutAwsIamRoleAuth(value *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth)
	ResetAccessKeySecretAuth()
	ResetAwsIamRoleAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) AccessKeySecretAuth() GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuthOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuthOutputReference
	_jsii_.Get(
		j,
		"accessKeySecretAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) AccessKeySecretAuthInput() *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth {
	var returns *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth
	_jsii_.Get(
		j,
		"accessKeySecretAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) AwsIamRoleAuth() GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference
	_jsii_.Get(
		j,
		"awsIamRoleAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) AwsIamRoleAuthInput() *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth {
	var returns *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth
	_jsii_.Get(
		j,
		"awsIamRoleAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) InternalValue() *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication {
	var returns *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference_Override(g GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference)SetInternalValue(val *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) PutAccessKeySecretAuth(value *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth) {
	if err := g.validatePutAccessKeySecretAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessKeySecretAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) PutAwsIamRoleAuth(value *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth) {
	if err := g.validatePutAwsIamRoleAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsIamRoleAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) ResetAccessKeySecretAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessKeySecretAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) ResetAwsIamRoleAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsIamRoleAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

