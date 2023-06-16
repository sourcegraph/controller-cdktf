package googlenetworkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkserviceshttproute/internal"
)

type GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference interface {
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
	ExactMatch() *string
	SetExactMatch(val *string)
	ExactMatchInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PresentMatch() interface{}
	SetPresentMatch(val interface{})
	PresentMatchInput() interface{}
	QueryParameter() *string
	SetQueryParameter(val *string)
	QueryParameterInput() *string
	RegexMatch() *string
	SetRegexMatch(val *string)
	RegexMatchInput() *string
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
	ResetExactMatch()
	ResetPresentMatch()
	ResetQueryParameter()
	ResetRegexMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference
type jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ExactMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ExactMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) PresentMatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presentMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) PresentMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presentMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) QueryParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) QueryParameterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) RegexMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) RegexMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference_Override(g GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetExactMatch(val *string) {
	if err := j.validateSetExactMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactMatch",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetPresentMatch(val interface{}) {
	if err := j.validateSetPresentMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presentMatch",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetQueryParameter(val *string) {
	if err := j.validateSetQueryParameterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryParameter",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetRegexMatch(val *string) {
	if err := j.validateSetRegexMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexMatch",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ResetExactMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetExactMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ResetPresentMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetPresentMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ResetQueryParameter() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryParameter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ResetRegexMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetRegexMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

