package workstationsworkstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/workstationsworkstationconfig/internal"
)

type WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference interface {
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
	InternalValue() *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig
	SetInternalValue(val *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig)
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

// The jsii proxy struct for WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference
type jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) EnableIntegrityMonitoring() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIntegrityMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) EnableIntegrityMonitoringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIntegrityMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) EnableSecureBoot() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSecureBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) EnableSecureBootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSecureBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) EnableVtpm() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVtpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) EnableVtpmInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVtpmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) InternalValue() *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig {
	var returns *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference_Override(w WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetEnableIntegrityMonitoring(val any) {
	if err := j.validateSetEnableIntegrityMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIntegrityMonitoring",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetEnableSecureBoot(val any) {
	if err := j.validateSetEnableSecureBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecureBoot",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetEnableVtpm(val any) {
	if err := j.validateSetEnableVtpmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVtpm",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetInternalValue(val *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ResetEnableIntegrityMonitoring() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableIntegrityMonitoring",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ResetEnableSecureBoot() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableSecureBoot",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ResetEnableVtpm() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableVtpm",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		w,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
