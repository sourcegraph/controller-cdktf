package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestool/internal"
)

type CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference interface {
	cdktf.ComplexObject
	Boost() *float64
	SetBoost(val *float64)
	BoostControlSpec() CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference
	BoostControlSpecInput() *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec
	BoostInput() *float64
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
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
	PutBoostControlSpec(value *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec)
	ResetBoost()
	ResetBoostControlSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference
type jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) Boost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"boost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) BoostControlSpec() CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference {
	var returns CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference
	_jsii_.Get(
		j,
		"boostControlSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) BoostControlSpecInput() *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec {
	var returns *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec
	_jsii_.Get(
		j,
		"boostControlSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) BoostInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"boostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference_Override(c CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetBoost(val *float64) {
	if err := j.validateSetBoostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boost",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) PutBoostControlSpec(value *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec) {
	if err := c.validatePutBoostControlSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBoostControlSpec",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ResetBoost() {
	_jsii_.InvokeVoid(
		c,
		"resetBoost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ResetBoostControlSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetBoostControlSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
