package googlecomputeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionurlmap/internal"
)

type GoogleComputeRegionUrlMapPathMatcherOutputReference interface {
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
	DefaultRouteAction() GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference
	DefaultRouteActionInput() *GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction
	DefaultService() *string
	SetDefaultService(val *string)
	DefaultServiceInput() *string
	DefaultUrlRedirect() GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirectOutputReference
	DefaultUrlRedirectInput() *GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirect
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	HeaderAction() GoogleComputeRegionUrlMapPathMatcherHeaderActionOutputReference
	HeaderActionInput() *GoogleComputeRegionUrlMapPathMatcherHeaderAction
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PathRule() GoogleComputeRegionUrlMapPathMatcherPathRuleList
	PathRuleInput() interface{}
	RouteRules() GoogleComputeRegionUrlMapPathMatcherRouteRulesList
	RouteRulesInput() interface{}
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
	PutDefaultRouteAction(value *GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction)
	PutDefaultUrlRedirect(value *GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirect)
	PutHeaderAction(value *GoogleComputeRegionUrlMapPathMatcherHeaderAction)
	PutPathRule(value interface{})
	PutRouteRules(value interface{})
	ResetDefaultRouteAction()
	ResetDefaultService()
	ResetDefaultUrlRedirect()
	ResetDescription()
	ResetHeaderAction()
	ResetPathRule()
	ResetRouteRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionUrlMapPathMatcherOutputReference
type jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DefaultRouteAction() GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference
	_jsii_.Get(
		j,
		"defaultRouteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DefaultRouteActionInput() *GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction {
	var returns *GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction
	_jsii_.Get(
		j,
		"defaultRouteActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DefaultService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DefaultServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DefaultUrlRedirect() GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirectOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"defaultUrlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DefaultUrlRedirectInput() *GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirect {
	var returns *GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirect
	_jsii_.Get(
		j,
		"defaultUrlRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) HeaderAction() GoogleComputeRegionUrlMapPathMatcherHeaderActionOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) HeaderActionInput() *GoogleComputeRegionUrlMapPathMatcherHeaderAction {
	var returns *GoogleComputeRegionUrlMapPathMatcherHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PathRule() GoogleComputeRegionUrlMapPathMatcherPathRuleList {
	var returns GoogleComputeRegionUrlMapPathMatcherPathRuleList
	_jsii_.Get(
		j,
		"pathRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PathRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) RouteRules() GoogleComputeRegionUrlMapPathMatcherRouteRulesList {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesList
	_jsii_.Get(
		j,
		"routeRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) RouteRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionUrlMapPathMatcherOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionUrlMapPathMatcherOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionUrlMapPathMatcherOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionUrlMapPathMatcherOutputReference_Override(g GoogleComputeRegionUrlMapPathMatcherOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetDefaultService(val *string) {
	if err := j.validateSetDefaultServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultService",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PutDefaultRouteAction(value *GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction) {
	if err := g.validatePutDefaultRouteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultRouteAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PutDefaultUrlRedirect(value *GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirect) {
	if err := g.validatePutDefaultUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultUrlRedirect",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PutHeaderAction(value *GoogleComputeRegionUrlMapPathMatcherHeaderAction) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PutPathRule(value interface{}) {
	if err := g.validatePutPathRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPathRule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) PutRouteRules(value interface{}) {
	if err := g.validatePutRouteRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRouteRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetDefaultRouteAction() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultRouteAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetDefaultService() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetDefaultUrlRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultUrlRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetPathRule() {
	_jsii_.InvokeVoid(
		g,
		"resetPathRule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ResetRouteRules() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

