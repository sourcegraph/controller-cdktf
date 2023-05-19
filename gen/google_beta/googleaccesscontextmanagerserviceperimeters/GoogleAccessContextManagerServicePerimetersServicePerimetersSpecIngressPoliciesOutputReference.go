package googleaccesscontextmanagerserviceperimeters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleaccesscontextmanagerserviceperimeters/internal"
)

type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference interface {
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
	IngressFrom() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference
	IngressFromInput() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom
	IngressTo() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOutputReference
	IngressToInput() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutIngressFrom(value *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom)
	PutIngressTo(value *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo)
	ResetIngressFrom()
	ResetIngressTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference
type jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) IngressFrom() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference {
	var returns GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference
	_jsii_.Get(
		j,
		"ingressFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) IngressFromInput() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom {
	var returns *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom
	_jsii_.Get(
		j,
		"ingressFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) IngressTo() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOutputReference {
	var returns GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOutputReference
	_jsii_.Get(
		j,
		"ingressTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) IngressToInput() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo {
	var returns *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo
	_jsii_.Get(
		j,
		"ingressToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference{}

	_jsii_.Create(
		"google-beta.googleAccessContextManagerServicePerimeters.GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference_Override(g GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleAccessContextManagerServicePerimeters.GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) PutIngressFrom(value *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom) {
	if err := g.validatePutIngressFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIngressFrom",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) PutIngressTo(value *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo) {
	if err := g.validatePutIngressToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIngressTo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) ResetIngressFrom() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressFrom",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) ResetIngressTo() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressTo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

