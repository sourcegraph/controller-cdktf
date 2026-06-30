package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestool/internal"
)

type CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference interface {
	cdktf.ComplexObject
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
	Disabled() any
	SetDisabled(val any)
	DisabledInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolDataStoreToolModalityConfigsRewriterConfig
	SetInternalValue(val *CesToolDataStoreToolModalityConfigsRewriterConfig)
	ModelSettings() CesToolDataStoreToolModalityConfigsRewriterConfigModelSettingsOutputReference
	ModelSettingsInput() *CesToolDataStoreToolModalityConfigsRewriterConfigModelSettings
	Prompt() *string
	SetPrompt(val *string)
	PromptInput() *string
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
	PutModelSettings(value *CesToolDataStoreToolModalityConfigsRewriterConfigModelSettings)
	ResetDisabled()
	ResetPrompt()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference
type jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) Disabled() any {
	var returns any
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) DisabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) InternalValue() *CesToolDataStoreToolModalityConfigsRewriterConfig {
	var returns *CesToolDataStoreToolModalityConfigsRewriterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ModelSettings() CesToolDataStoreToolModalityConfigsRewriterConfigModelSettingsOutputReference {
	var returns CesToolDataStoreToolModalityConfigsRewriterConfigModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ModelSettingsInput() *CesToolDataStoreToolModalityConfigsRewriterConfigModelSettings {
	var returns *CesToolDataStoreToolModalityConfigsRewriterConfigModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) Prompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) PromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesToolDataStoreToolModalityConfigsRewriterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolDataStoreToolModalityConfigsRewriterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolDataStoreToolModalityConfigsRewriterConfigOutputReference_Override(c CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetDisabled(val any) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetInternalValue(val *CesToolDataStoreToolModalityConfigsRewriterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetPrompt(val *string) {
	if err := j.validateSetPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prompt",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) PutModelSettings(value *CesToolDataStoreToolModalityConfigsRewriterConfigModelSettings) {
	if err := c.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		c,
		"resetPrompt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
