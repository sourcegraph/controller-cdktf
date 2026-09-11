package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference interface {
	cdktn.ComplexObject
	Authentication() GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthenticationOutputReference
	AuthenticationInput() *GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthentication
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
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InitialStartTime() *string
	SetInitialStartTime(val *string)
	InitialStartTimeInput() *string
	InternalValue() *GoogleChronicleFeedDetailsSentineloneAlertSettings
	SetInternalValue(val *GoogleChronicleFeedDetailsSentineloneAlertSettings)
	IsAlertApiSubscribed() interface{}
	SetIsAlertApiSubscribed(val interface{})
	IsAlertApiSubscribedInput() interface{}
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
	PutAuthentication(value *GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthentication)
	ResetAuthentication()
	ResetHostname()
	ResetInitialStartTime()
	ResetIsAlertApiSubscribed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) Authentication() GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthenticationOutputReference {
	var returns GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) AuthenticationInput() *GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthentication {
	var returns *GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) InitialStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) InitialStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) InternalValue() *GoogleChronicleFeedDetailsSentineloneAlertSettings {
	var returns *GoogleChronicleFeedDetailsSentineloneAlertSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) IsAlertApiSubscribed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAlertApiSubscribed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) IsAlertApiSubscribedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAlertApiSubscribedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference_Override(g GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetInitialStartTime(val *string) {
	if err := j.validateSetInitialStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialStartTime",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetInternalValue(val *GoogleChronicleFeedDetailsSentineloneAlertSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetIsAlertApiSubscribed(val interface{}) {
	if err := j.validateSetIsAlertApiSubscribedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAlertApiSubscribed",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) PutAuthentication(value *GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthentication) {
	if err := g.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		g,
		"resetHostname",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ResetInitialStartTime() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialStartTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ResetIsAlertApiSubscribed() {
	_jsii_.InvokeVoid(
		g,
		"resetIsAlertApiSubscribed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

