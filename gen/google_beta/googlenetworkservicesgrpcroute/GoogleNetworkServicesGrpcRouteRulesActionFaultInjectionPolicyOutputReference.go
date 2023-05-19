package googlenetworkservicesgrpcroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesgrpcroute/internal"
)

type GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference interface {
	cdktf.ComplexObject
	Abort() GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortOutputReference
	AbortInput() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort
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
	Delay() GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayOutputReference
	DelayInput() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy
	SetInternalValue(val *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy)
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
	PutAbort(value *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort)
	PutDelay(value *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay)
	ResetAbort()
	ResetDelay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference
type jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Abort() GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortOutputReference {
	var returns GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortOutputReference
	_jsii_.Get(
		j,
		"abort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) AbortInput() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort {
	var returns *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort
	_jsii_.Get(
		j,
		"abortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Delay() GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayOutputReference {
	var returns GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayOutputReference
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) DelayInput() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay {
	var returns *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay
	_jsii_.Get(
		j,
		"delayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) InternalValue() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy {
	var returns *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkServicesGrpcRoute.GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference_Override(g GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkServicesGrpcRoute.GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetInternalValue(val *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) PutAbort(value *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort) {
	if err := g.validatePutAbortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAbort",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) PutDelay(value *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay) {
	if err := g.validatePutDelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDelay",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ResetAbort() {
	_jsii_.InvokeVoid(
		g,
		"resetAbort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ResetDelay() {
	_jsii_.InvokeVoid(
		g,
		"resetDelay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

