package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesapp/internal"
)

type CesAppDefaultChannelProfileWebWidgetConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppDefaultChannelProfileWebWidgetConfig
	SetInternalValue(val *CesAppDefaultChannelProfileWebWidgetConfig)
	Modality() *string
	SetModality(val *string)
	ModalityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Theme() *string
	SetTheme(val *string)
	ThemeInput() *string
	WebWidgetTitle() *string
	SetWebWidgetTitle(val *string)
	WebWidgetTitleInput() *string
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
	ResetModality()
	ResetTheme()
	ResetWebWidgetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppDefaultChannelProfileWebWidgetConfigOutputReference
type jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) InternalValue() *CesAppDefaultChannelProfileWebWidgetConfig {
	var returns *CesAppDefaultChannelProfileWebWidgetConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) Modality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ModalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) Theme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"theme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) WebWidgetTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webWidgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) WebWidgetTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webWidgetTitleInput",
		&returns,
	)
	return returns
}

func NewCesAppDefaultChannelProfileWebWidgetConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesAppDefaultChannelProfileWebWidgetConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppDefaultChannelProfileWebWidgetConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppDefaultChannelProfileWebWidgetConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppDefaultChannelProfileWebWidgetConfigOutputReference_Override(c CesAppDefaultChannelProfileWebWidgetConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppDefaultChannelProfileWebWidgetConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetInternalValue(val *CesAppDefaultChannelProfileWebWidgetConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetModality(val *string) {
	if err := j.validateSetModalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modality",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetTheme(val *string) {
	if err := j.validateSetThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"theme",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) SetWebWidgetTitle(val *string) {
	if err := j.validateSetWebWidgetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webWidgetTitle",
		val,
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ResetModality() {
	_jsii_.InvokeVoid(
		c,
		"resetModality",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ResetTheme() {
	_jsii_.InvokeVoid(
		c,
		"resetTheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ResetWebWidgetTitle() {
	_jsii_.InvokeVoid(
		c,
		"resetWebWidgetTitle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileWebWidgetConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
