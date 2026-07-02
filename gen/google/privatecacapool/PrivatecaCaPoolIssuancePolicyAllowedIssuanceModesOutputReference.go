package privatecacapool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/privatecacapool/internal"
)

type PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference interface {
	cdktf.ComplexObject
	AllowConfigBasedIssuance() any
	SetAllowConfigBasedIssuance(val any)
	AllowConfigBasedIssuanceInput() any
	AllowCsrBasedIssuance() any
	SetAllowCsrBasedIssuance(val any)
	AllowCsrBasedIssuanceInput() any
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	InternalValue() *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes
	SetInternalValue(val *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes)
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference
type jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) AllowConfigBasedIssuance() any {
	var returns any
	_jsii_.Get(
		j,
		"allowConfigBasedIssuance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) AllowConfigBasedIssuanceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowConfigBasedIssuanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) AllowCsrBasedIssuance() any {
	var returns any
	_jsii_.Get(
		j,
		"allowCsrBasedIssuance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) AllowCsrBasedIssuanceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowCsrBasedIssuanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) InternalValue() *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes {
	var returns *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewPrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCaPool.PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference_Override(p PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCaPool.PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetAllowConfigBasedIssuance(val any) {
	if err := j.validateSetAllowConfigBasedIssuanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowConfigBasedIssuance",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetAllowCsrBasedIssuance(val any) {
	if err := j.validateSetAllowCsrBasedIssuanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCsrBasedIssuance",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetInternalValue(val *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
