package googlenetworkservicestcproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicestcproute/internal"
)

type GoogleNetworkServicesTcpRouteRulesActionOutputReference interface {
	cdktn.ComplexObject
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
	Destinations() GoogleNetworkServicesTcpRouteRulesActionDestinationsList
	DestinationsInput() interface{}
	// Experimental.
	Fqn() *string
	IdleTimeout() *string
	SetIdleTimeout(val *string)
	IdleTimeoutInput() *string
	InternalValue() *GoogleNetworkServicesTcpRouteRulesAction
	SetInternalValue(val *GoogleNetworkServicesTcpRouteRulesAction)
	OriginalDestination() interface{}
	SetOriginalDestination(val interface{})
	OriginalDestinationInput() interface{}
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
	PutDestinations(value interface{})
	ResetDestinations()
	ResetIdleTimeout()
	ResetOriginalDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesTcpRouteRulesActionOutputReference
type jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) Destinations() GoogleNetworkServicesTcpRouteRulesActionDestinationsList {
	var returns GoogleNetworkServicesTcpRouteRulesActionDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) IdleTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) IdleTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) InternalValue() *GoogleNetworkServicesTcpRouteRulesAction {
	var returns *GoogleNetworkServicesTcpRouteRulesAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) OriginalDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originalDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) OriginalDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originalDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesTcpRouteRulesActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesTcpRouteRulesActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesTcpRouteRulesActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkServicesTcpRoute.GoogleNetworkServicesTcpRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesTcpRouteRulesActionOutputReference_Override(g GoogleNetworkServicesTcpRouteRulesActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkServicesTcpRoute.GoogleNetworkServicesTcpRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetIdleTimeout(val *string) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetInternalValue(val *GoogleNetworkServicesTcpRouteRulesAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetOriginalDestination(val interface{}) {
	if err := j.validateSetOriginalDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originalDestination",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) PutDestinations(value interface{}) {
	if err := g.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ResetOriginalDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetOriginalDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesTcpRouteRulesActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

