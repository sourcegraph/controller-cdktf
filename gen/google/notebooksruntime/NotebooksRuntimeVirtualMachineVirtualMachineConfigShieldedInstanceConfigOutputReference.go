package notebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/notebooksruntime/internal"
)

type NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference interface {
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
	EnableIntegrityMonitoring() any
	SetEnableIntegrityMonitoring(val any)
	EnableIntegrityMonitoringInput() any
	EnableSecureBoot() any
	SetEnableSecureBoot(val any)
	EnableSecureBootInput() any
	EnableVtpm() any
	SetEnableVtpm(val any)
	EnableVtpmInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig
	SetInternalValue(val *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig)
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
	ResetEnableIntegrityMonitoring()
	ResetEnableSecureBoot()
	ResetEnableVtpm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference
type jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) EnableIntegrityMonitoring() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIntegrityMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) EnableIntegrityMonitoringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIntegrityMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) EnableSecureBoot() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSecureBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) EnableSecureBootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSecureBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) EnableVtpm() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVtpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) EnableVtpmInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVtpmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) InternalValue() *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference_Override(n NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetEnableIntegrityMonitoring(val any) {
	if err := j.validateSetEnableIntegrityMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIntegrityMonitoring",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetEnableSecureBoot(val any) {
	if err := j.validateSetEnableSecureBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecureBoot",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetEnableVtpm(val any) {
	if err := j.validateSetEnableVtpmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVtpm",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetInternalValue(val *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ResetEnableIntegrityMonitoring() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableIntegrityMonitoring",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ResetEnableSecureBoot() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableSecureBoot",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ResetEnableVtpm() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableVtpm",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
