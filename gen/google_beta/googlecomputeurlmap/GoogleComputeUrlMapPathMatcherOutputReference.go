package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapPathMatcherOutputReference interface {
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
	DefaultCustomErrorResponsePolicy() GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicyOutputReference
	DefaultCustomErrorResponsePolicyInput() *GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicy
	DefaultRouteAction() GoogleComputeUrlMapPathMatcherDefaultRouteActionOutputReference
	DefaultRouteActionInput() *GoogleComputeUrlMapPathMatcherDefaultRouteAction
	DefaultService() *string
	SetDefaultService(val *string)
	DefaultServiceInput() *string
	DefaultUrlRedirect() GoogleComputeUrlMapPathMatcherDefaultUrlRedirectOutputReference
	DefaultUrlRedirectInput() *GoogleComputeUrlMapPathMatcherDefaultUrlRedirect
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	HeaderAction() GoogleComputeUrlMapPathMatcherHeaderActionOutputReference
	HeaderActionInput() *GoogleComputeUrlMapPathMatcherHeaderAction
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PathRule() GoogleComputeUrlMapPathMatcherPathRuleList
	PathRuleInput() interface{}
	RouteRules() GoogleComputeUrlMapPathMatcherRouteRulesList
	RouteRulesInput() interface{}
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
	PutDefaultCustomErrorResponsePolicy(value *GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicy)
	PutDefaultRouteAction(value *GoogleComputeUrlMapPathMatcherDefaultRouteAction)
	PutDefaultUrlRedirect(value *GoogleComputeUrlMapPathMatcherDefaultUrlRedirect)
	PutHeaderAction(value *GoogleComputeUrlMapPathMatcherHeaderAction)
	PutPathRule(value interface{})
	PutRouteRules(value interface{})
	ResetDefaultCustomErrorResponsePolicy()
	ResetDefaultRouteAction()
	ResetDefaultService()
	ResetDefaultUrlRedirect()
	ResetDescription()
	ResetHeaderAction()
	ResetPathRule()
	ResetRouteRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeUrlMapPathMatcherOutputReference
type jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultCustomErrorResponsePolicy() GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicyOutputReference
	_jsii_.Get(
		j,
		"defaultCustomErrorResponsePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultCustomErrorResponsePolicyInput() *GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicy {
	var returns *GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicy
	_jsii_.Get(
		j,
		"defaultCustomErrorResponsePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultRouteAction() GoogleComputeUrlMapPathMatcherDefaultRouteActionOutputReference {
	var returns GoogleComputeUrlMapPathMatcherDefaultRouteActionOutputReference
	_jsii_.Get(
		j,
		"defaultRouteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultRouteActionInput() *GoogleComputeUrlMapPathMatcherDefaultRouteAction {
	var returns *GoogleComputeUrlMapPathMatcherDefaultRouteAction
	_jsii_.Get(
		j,
		"defaultRouteActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultUrlRedirect() GoogleComputeUrlMapPathMatcherDefaultUrlRedirectOutputReference {
	var returns GoogleComputeUrlMapPathMatcherDefaultUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"defaultUrlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DefaultUrlRedirectInput() *GoogleComputeUrlMapPathMatcherDefaultUrlRedirect {
	var returns *GoogleComputeUrlMapPathMatcherDefaultUrlRedirect
	_jsii_.Get(
		j,
		"defaultUrlRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) HeaderAction() GoogleComputeUrlMapPathMatcherHeaderActionOutputReference {
	var returns GoogleComputeUrlMapPathMatcherHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) HeaderActionInput() *GoogleComputeUrlMapPathMatcherHeaderAction {
	var returns *GoogleComputeUrlMapPathMatcherHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PathRule() GoogleComputeUrlMapPathMatcherPathRuleList {
	var returns GoogleComputeUrlMapPathMatcherPathRuleList
	_jsii_.Get(
		j,
		"pathRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PathRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) RouteRules() GoogleComputeUrlMapPathMatcherRouteRulesList {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesList
	_jsii_.Get(
		j,
		"routeRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) RouteRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapPathMatcherOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeUrlMapPathMatcherOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapPathMatcherOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapPathMatcherOutputReference_Override(g GoogleComputeUrlMapPathMatcherOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetDefaultService(val *string) {
	if err := j.validateSetDefaultServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultService",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PutDefaultCustomErrorResponsePolicy(value *GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicy) {
	if err := g.validatePutDefaultCustomErrorResponsePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultCustomErrorResponsePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PutDefaultRouteAction(value *GoogleComputeUrlMapPathMatcherDefaultRouteAction) {
	if err := g.validatePutDefaultRouteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultRouteAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PutDefaultUrlRedirect(value *GoogleComputeUrlMapPathMatcherDefaultUrlRedirect) {
	if err := g.validatePutDefaultUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultUrlRedirect",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PutHeaderAction(value *GoogleComputeUrlMapPathMatcherHeaderAction) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PutPathRule(value interface{}) {
	if err := g.validatePutPathRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPathRule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) PutRouteRules(value interface{}) {
	if err := g.validatePutRouteRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRouteRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetDefaultCustomErrorResponsePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultCustomErrorResponsePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetDefaultRouteAction() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultRouteAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetDefaultService() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetDefaultUrlRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultUrlRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetPathRule() {
	_jsii_.InvokeVoid(
		g,
		"resetPathRule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ResetRouteRules() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

