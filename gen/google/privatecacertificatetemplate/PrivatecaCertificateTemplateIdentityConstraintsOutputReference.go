package privatecacertificatetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/privatecacertificatetemplate/internal"
)

type PrivatecaCertificateTemplateIdentityConstraintsOutputReference interface {
	cdktn.ComplexObject
	AllowSubjectAltNamesPassthrough() interface{}
	SetAllowSubjectAltNamesPassthrough(val interface{})
	AllowSubjectAltNamesPassthroughInput() interface{}
	AllowSubjectPassthrough() interface{}
	SetAllowSubjectPassthrough(val interface{})
	AllowSubjectPassthroughInput() interface{}
	CelExpression() PrivatecaCertificateTemplateIdentityConstraintsCelExpressionOutputReference
	CelExpressionInput() *PrivatecaCertificateTemplateIdentityConstraintsCelExpression
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
	InternalValue() *PrivatecaCertificateTemplateIdentityConstraints
	SetInternalValue(val *PrivatecaCertificateTemplateIdentityConstraints)
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
	PutCelExpression(value *PrivatecaCertificateTemplateIdentityConstraintsCelExpression)
	ResetCelExpression()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateTemplateIdentityConstraintsOutputReference
type jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) AllowSubjectAltNamesPassthrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubjectAltNamesPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) AllowSubjectAltNamesPassthroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubjectAltNamesPassthroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) AllowSubjectPassthrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubjectPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) AllowSubjectPassthroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubjectPassthroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) CelExpression() PrivatecaCertificateTemplateIdentityConstraintsCelExpressionOutputReference {
	var returns PrivatecaCertificateTemplateIdentityConstraintsCelExpressionOutputReference
	_jsii_.Get(
		j,
		"celExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) CelExpressionInput() *PrivatecaCertificateTemplateIdentityConstraintsCelExpression {
	var returns *PrivatecaCertificateTemplateIdentityConstraintsCelExpression
	_jsii_.Get(
		j,
		"celExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) InternalValue() *PrivatecaCertificateTemplateIdentityConstraints {
	var returns *PrivatecaCertificateTemplateIdentityConstraints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateTemplateIdentityConstraintsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateTemplateIdentityConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateTemplateIdentityConstraintsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.privatecaCertificateTemplate.PrivatecaCertificateTemplateIdentityConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateTemplateIdentityConstraintsOutputReference_Override(p PrivatecaCertificateTemplateIdentityConstraintsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.privatecaCertificateTemplate.PrivatecaCertificateTemplateIdentityConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetAllowSubjectAltNamesPassthrough(val interface{}) {
	if err := j.validateSetAllowSubjectAltNamesPassthroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSubjectAltNamesPassthrough",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetAllowSubjectPassthrough(val interface{}) {
	if err := j.validateSetAllowSubjectPassthroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSubjectPassthrough",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetInternalValue(val *PrivatecaCertificateTemplateIdentityConstraints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) PutCelExpression(value *PrivatecaCertificateTemplateIdentityConstraintsCelExpression) {
	if err := p.validatePutCelExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCelExpression",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) ResetCelExpression() {
	_jsii_.InvokeVoid(
		p,
		"resetCelExpression",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplateIdentityConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

