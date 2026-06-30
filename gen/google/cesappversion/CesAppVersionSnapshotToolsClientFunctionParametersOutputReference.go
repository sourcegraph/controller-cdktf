package cesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesappversion/internal"
)

type CesAppVersionSnapshotToolsClientFunctionParametersOutputReference interface {
	cdktf.ComplexObject
	AdditionalProperties() *string
	AnyOf() *string
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
	Default() *string
	Defs() *string
	Description() *string
	Enum() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppVersionSnapshotToolsClientFunctionParameters
	SetInternalValue(val *CesAppVersionSnapshotToolsClientFunctionParameters)
	Items() *string
	Nullable() cdktf.IResolvable
	PrefixItems() *string
	Properties() *string
	Ref() *string
	Required() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	UniqueItems() cdktf.IResolvable
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

// The jsii proxy struct for CesAppVersionSnapshotToolsClientFunctionParametersOutputReference
type jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) AdditionalProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) AnyOf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Defs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) InternalValue() *CesAppVersionSnapshotToolsClientFunctionParameters {
	var returns *CesAppVersionSnapshotToolsClientFunctionParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Items() *string {
	var returns *string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Nullable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) PrefixItems() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Required() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) UniqueItems() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"uniqueItems",
		&returns,
	)
	return returns
}

func NewCesAppVersionSnapshotToolsClientFunctionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesAppVersionSnapshotToolsClientFunctionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppVersionSnapshotToolsClientFunctionParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesAppVersion.CesAppVersionSnapshotToolsClientFunctionParametersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesAppVersionSnapshotToolsClientFunctionParametersOutputReference_Override(c CesAppVersionSnapshotToolsClientFunctionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesAppVersion.CesAppVersionSnapshotToolsClientFunctionParametersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) SetInternalValue(val *CesAppVersionSnapshotToolsClientFunctionParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsClientFunctionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
