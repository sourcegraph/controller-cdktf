package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestoolset/internal"
)

type CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference interface {
	cdktf.ComplexObject
	ApiKeySecretVersion() *string
	SetApiKeySecretVersion(val *string)
	ApiKeySecretVersionInput() *string
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
	InternalValue() *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig
	SetInternalValue(val *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig)
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	RequestLocation() *string
	SetRequestLocation(val *string)
	RequestLocationInput() *string
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

// The jsii proxy struct for CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference
type jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) ApiKeySecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) ApiKeySecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) InternalValue() *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig {
	var returns *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) RequestLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) RequestLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference_Override(c CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetApiKeySecretVersion(val *string) {
	if err := j.validateSetApiKeySecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeySecretVersion",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetInternalValue(val *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetRequestLocation(val *string) {
	if err := j.validateSetRequestLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestLocation",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
