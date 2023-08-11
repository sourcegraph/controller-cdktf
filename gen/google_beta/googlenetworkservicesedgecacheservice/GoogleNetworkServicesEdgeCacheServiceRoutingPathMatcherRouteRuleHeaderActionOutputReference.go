package googlenetworkservicesedgecacheservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesedgecacheservice/internal"
)

type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction
	SetInternalValue(val *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction)
	RequestHeaderToAdd() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToAddList
	RequestHeaderToAddInput() interface{}
	RequestHeaderToRemove() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToRemoveList
	RequestHeaderToRemoveInput() interface{}
	ResponseHeaderToAdd() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToAddList
	ResponseHeaderToAddInput() interface{}
	ResponseHeaderToRemove() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToRemoveList
	ResponseHeaderToRemoveInput() interface{}
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
	PutRequestHeaderToAdd(value interface{})
	PutRequestHeaderToRemove(value interface{})
	PutResponseHeaderToAdd(value interface{})
	PutResponseHeaderToRemove(value interface{})
	ResetRequestHeaderToAdd()
	ResetRequestHeaderToRemove()
	ResetResponseHeaderToAdd()
	ResetResponseHeaderToRemove()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference
type jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) InternalValue() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction {
	var returns *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) RequestHeaderToAdd() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToAddList {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToAddList
	_jsii_.Get(
		j,
		"requestHeaderToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) RequestHeaderToAddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) RequestHeaderToRemove() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToRemoveList {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToRemoveList
	_jsii_.Get(
		j,
		"requestHeaderToRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) RequestHeaderToRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderToRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResponseHeaderToAdd() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToAddList {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToAddList
	_jsii_.Get(
		j,
		"responseHeaderToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResponseHeaderToAddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeaderToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResponseHeaderToRemove() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToRemoveList {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToRemoveList
	_jsii_.Get(
		j,
		"responseHeaderToRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResponseHeaderToRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeaderToRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesEdgeCacheService.GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference_Override(g GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesEdgeCacheService.GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference)SetInternalValue(val *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) PutRequestHeaderToAdd(value interface{}) {
	if err := g.validatePutRequestHeaderToAddParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestHeaderToAdd",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) PutRequestHeaderToRemove(value interface{}) {
	if err := g.validatePutRequestHeaderToRemoveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestHeaderToRemove",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) PutResponseHeaderToAdd(value interface{}) {
	if err := g.validatePutResponseHeaderToAddParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResponseHeaderToAdd",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) PutResponseHeaderToRemove(value interface{}) {
	if err := g.validatePutResponseHeaderToRemoveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResponseHeaderToRemove",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResetRequestHeaderToAdd() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeaderToAdd",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResetRequestHeaderToRemove() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeaderToRemove",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResetResponseHeaderToAdd() {
	_jsii_.InvokeVoid(
		g,
		"resetResponseHeaderToAdd",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ResetResponseHeaderToRemove() {
	_jsii_.InvokeVoid(
		g,
		"resetResponseHeaderToRemove",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

