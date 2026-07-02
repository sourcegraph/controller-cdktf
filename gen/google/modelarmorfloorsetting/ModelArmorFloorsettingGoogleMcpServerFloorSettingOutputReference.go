package modelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/modelarmorfloorsetting/internal"
)

type ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference interface {
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
	EnableCloudLogging() any
	SetEnableCloudLogging(val any)
	EnableCloudLoggingInput() any
	// Experimental.
	Fqn() *string
	InspectAndBlock() any
	SetInspectAndBlock(val any)
	InspectAndBlockInput() any
	InspectOnly() any
	SetInspectOnly(val any)
	InspectOnlyInput() any
	InternalValue() *ModelArmorFloorsettingGoogleMcpServerFloorSetting
	SetInternalValue(val *ModelArmorFloorsettingGoogleMcpServerFloorSetting)
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
	ResetEnableCloudLogging()
	ResetInspectAndBlock()
	ResetInspectOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference
type jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) EnableCloudLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableCloudLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) EnableCloudLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableCloudLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InspectAndBlock() any {
	var returns any
	_jsii_.Get(
		j,
		"inspectAndBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InspectAndBlockInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inspectAndBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InspectOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"inspectOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InspectOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inspectOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InternalValue() *ModelArmorFloorsettingGoogleMcpServerFloorSetting {
	var returns *ModelArmorFloorsettingGoogleMcpServerFloorSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference_Override(m ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference",
		[]any{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetEnableCloudLogging(val any) {
	if err := j.validateSetEnableCloudLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudLogging",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetInspectAndBlock(val any) {
	if err := j.validateSetInspectAndBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectAndBlock",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetInspectOnly(val any) {
	if err := j.validateSetInspectOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectOnly",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetInternalValue(val *ModelArmorFloorsettingGoogleMcpServerFloorSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ResetEnableCloudLogging() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableCloudLogging",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ResetInspectAndBlock() {
	_jsii_.InvokeVoid(
		m,
		"resetInspectAndBlock",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ResetInspectOnly() {
	_jsii_.InvokeVoid(
		m,
		"resetInspectOnly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		m,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
