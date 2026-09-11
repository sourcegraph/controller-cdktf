package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference interface {
	cdktn.ComplexObject
	Authentication() GoogleChronicleFeedDetailsFoxItStixSettingsAuthenticationOutputReference
	AuthenticationInput() *GoogleChronicleFeedDetailsFoxItStixSettingsAuthentication
	Collection() *string
	SetCollection(val *string)
	CollectionInput() *string
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
	InternalValue() *GoogleChronicleFeedDetailsFoxItStixSettings
	SetInternalValue(val *GoogleChronicleFeedDetailsFoxItStixSettings)
	PollServiceUri() *string
	SetPollServiceUri(val *string)
	PollServiceUriInput() *string
	Ssl() GoogleChronicleFeedDetailsFoxItStixSettingsSslOutputReference
	SslInput() *GoogleChronicleFeedDetailsFoxItStixSettingsSsl
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
	PutAuthentication(value *GoogleChronicleFeedDetailsFoxItStixSettingsAuthentication)
	PutSsl(value *GoogleChronicleFeedDetailsFoxItStixSettingsSsl)
	ResetAuthentication()
	ResetCollection()
	ResetPollServiceUri()
	ResetSsl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) Authentication() GoogleChronicleFeedDetailsFoxItStixSettingsAuthenticationOutputReference {
	var returns GoogleChronicleFeedDetailsFoxItStixSettingsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) AuthenticationInput() *GoogleChronicleFeedDetailsFoxItStixSettingsAuthentication {
	var returns *GoogleChronicleFeedDetailsFoxItStixSettingsAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) Collection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) CollectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) InternalValue() *GoogleChronicleFeedDetailsFoxItStixSettings {
	var returns *GoogleChronicleFeedDetailsFoxItStixSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) PollServiceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollServiceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) PollServiceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollServiceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) Ssl() GoogleChronicleFeedDetailsFoxItStixSettingsSslOutputReference {
	var returns GoogleChronicleFeedDetailsFoxItStixSettingsSslOutputReference
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) SslInput() *GoogleChronicleFeedDetailsFoxItStixSettingsSsl {
	var returns *GoogleChronicleFeedDetailsFoxItStixSettingsSsl
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsFoxItStixSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsFoxItStixSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsFoxItStixSettingsOutputReference_Override(g GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetCollection(val *string) {
	if err := j.validateSetCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collection",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetInternalValue(val *GoogleChronicleFeedDetailsFoxItStixSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetPollServiceUri(val *string) {
	if err := j.validateSetPollServiceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pollServiceUri",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) PutAuthentication(value *GoogleChronicleFeedDetailsFoxItStixSettingsAuthentication) {
	if err := g.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) PutSsl(value *GoogleChronicleFeedDetailsFoxItStixSettingsSsl) {
	if err := g.validatePutSslParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSsl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ResetCollection() {
	_jsii_.InvokeVoid(
		g,
		"resetCollection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ResetPollServiceUri() {
	_jsii_.InvokeVoid(
		g,
		"resetPollServiceUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		g,
		"resetSsl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

