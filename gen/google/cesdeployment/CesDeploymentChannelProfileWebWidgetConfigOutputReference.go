package cesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesdeployment/internal"
)

type CesDeploymentChannelProfileWebWidgetConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *CesDeploymentChannelProfileWebWidgetConfig
	SetInternalValue(val *CesDeploymentChannelProfileWebWidgetConfig)
	Modality() *string
	SetModality(val *string)
	ModalityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Theme() *string
	SetTheme(val *string)
	ThemeInput() *string
	WebWidgetTitle() *string
	SetWebWidgetTitle(val *string)
	WebWidgetTitleInput() *string
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
	ResetModality()
	ResetTheme()
	ResetWebWidgetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesDeploymentChannelProfileWebWidgetConfigOutputReference
type jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) InternalValue() *CesDeploymentChannelProfileWebWidgetConfig {
	var returns *CesDeploymentChannelProfileWebWidgetConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) Modality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ModalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) Theme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"theme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) WebWidgetTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webWidgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) WebWidgetTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webWidgetTitleInput",
		&returns,
	)
	return returns
}


func NewCesDeploymentChannelProfileWebWidgetConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesDeploymentChannelProfileWebWidgetConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesDeploymentChannelProfileWebWidgetConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentChannelProfileWebWidgetConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesDeploymentChannelProfileWebWidgetConfigOutputReference_Override(c CesDeploymentChannelProfileWebWidgetConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentChannelProfileWebWidgetConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetInternalValue(val *CesDeploymentChannelProfileWebWidgetConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetModality(val *string) {
	if err := j.validateSetModalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modality",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetTheme(val *string) {
	if err := j.validateSetThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"theme",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference)SetWebWidgetTitle(val *string) {
	if err := j.validateSetWebWidgetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webWidgetTitle",
		val,
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ResetModality() {
	_jsii_.InvokeVoid(
		c,
		"resetModality",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ResetTheme() {
	_jsii_.InvokeVoid(
		c,
		"resetTheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ResetWebWidgetTitle() {
	_jsii_.InvokeVoid(
		c,
		"resetWebWidgetTitle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

