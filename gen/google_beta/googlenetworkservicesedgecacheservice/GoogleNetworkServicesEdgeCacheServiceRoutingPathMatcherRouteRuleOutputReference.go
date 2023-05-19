package googlenetworkservicesedgecacheservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesedgecacheservice/internal"
)

type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	HeaderAction() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference
	HeaderActionInput() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchRule() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleList
	MatchRuleInput() interface{}
	Origin() *string
	SetOrigin(val *string)
	OriginInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	RouteAction() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionOutputReference
	RouteActionInput() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRedirect() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirectOutputReference
	UrlRedirectInput() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect
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
	PutHeaderAction(value *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction)
	PutMatchRule(value interface{})
	PutRouteAction(value *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction)
	PutUrlRedirect(value *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect)
	ResetDescription()
	ResetHeaderAction()
	ResetOrigin()
	ResetRouteAction()
	ResetUrlRedirect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference
type jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) HeaderAction() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) HeaderActionInput() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction {
	var returns *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) MatchRule() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleList {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleList
	_jsii_.Get(
		j,
		"matchRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) MatchRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) Origin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) OriginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) RouteAction() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionOutputReference {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionOutputReference
	_jsii_.Get(
		j,
		"routeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) RouteActionInput() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction {
	var returns *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction
	_jsii_.Get(
		j,
		"routeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) UrlRedirect() GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirectOutputReference {
	var returns GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"urlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) UrlRedirectInput() *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect {
	var returns *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect
	_jsii_.Get(
		j,
		"urlRedirectInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkServicesEdgeCacheService.GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference_Override(g GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkServicesEdgeCacheService.GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetOrigin(val *string) {
	if err := j.validateSetOriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"origin",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) PutHeaderAction(value *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) PutMatchRule(value interface{}) {
	if err := g.validatePutMatchRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMatchRule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) PutRouteAction(value *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction) {
	if err := g.validatePutRouteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRouteAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) PutUrlRedirect(value *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect) {
	if err := g.validatePutUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRedirect",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ResetOrigin() {
	_jsii_.InvokeVoid(
		g,
		"resetOrigin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ResetRouteAction() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ResetUrlRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

