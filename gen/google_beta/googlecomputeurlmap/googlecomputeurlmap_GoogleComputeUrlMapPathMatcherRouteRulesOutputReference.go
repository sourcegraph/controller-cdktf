package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapPathMatcherRouteRulesOutputReference interface {
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
	HeaderAction() GoogleComputeUrlMapPathMatcherRouteRulesHeaderActionOutputReference
	HeaderActionInput() *GoogleComputeUrlMapPathMatcherRouteRulesHeaderAction
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchRules() GoogleComputeUrlMapPathMatcherRouteRulesMatchRulesList
	MatchRulesInput() interface{}
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	RouteAction() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference
	RouteActionInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRedirect() GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirectOutputReference
	UrlRedirectInput() *GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirect
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
	PutHeaderAction(value *GoogleComputeUrlMapPathMatcherRouteRulesHeaderAction)
	PutMatchRules(value interface{})
	PutRouteAction(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction)
	PutUrlRedirect(value *GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirect)
	ResetHeaderAction()
	ResetMatchRules()
	ResetRouteAction()
	ResetService()
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

// The jsii proxy struct for GoogleComputeUrlMapPathMatcherRouteRulesOutputReference
type jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) HeaderAction() GoogleComputeUrlMapPathMatcherRouteRulesHeaderActionOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) HeaderActionInput() *GoogleComputeUrlMapPathMatcherRouteRulesHeaderAction {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) MatchRules() GoogleComputeUrlMapPathMatcherRouteRulesMatchRulesList {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesMatchRulesList
	_jsii_.Get(
		j,
		"matchRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) MatchRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) RouteAction() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference
	_jsii_.Get(
		j,
		"routeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) RouteActionInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction
	_jsii_.Get(
		j,
		"routeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) UrlRedirect() GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirectOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"urlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) UrlRedirectInput() *GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirect {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirect
	_jsii_.Get(
		j,
		"urlRedirectInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapPathMatcherRouteRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeUrlMapPathMatcherRouteRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapPathMatcherRouteRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherRouteRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapPathMatcherRouteRulesOutputReference_Override(g GoogleComputeUrlMapPathMatcherRouteRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherRouteRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) PutHeaderAction(value *GoogleComputeUrlMapPathMatcherRouteRulesHeaderAction) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) PutMatchRules(value interface{}) {
	if err := g.validatePutMatchRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMatchRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) PutRouteAction(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction) {
	if err := g.validatePutRouteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRouteAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) PutUrlRedirect(value *GoogleComputeUrlMapPathMatcherRouteRulesUrlRedirect) {
	if err := g.validatePutUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRedirect",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ResetMatchRules() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ResetRouteAction() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		g,
		"resetService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ResetUrlRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

