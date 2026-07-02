package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/jobv1/internal"
)

type JobV1SpecTemplateSpecContainerSecurityContextOutputReference interface {
	cdktf.ComplexObject
	AllowPrivilegeEscalation() any
	SetAllowPrivilegeEscalation(val any)
	AllowPrivilegeEscalationInput() any
	Capabilities() JobV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference
	CapabilitiesInput() *JobV1SpecTemplateSpecContainerSecurityContextCapabilities
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
	InternalValue() *JobV1SpecTemplateSpecContainerSecurityContext
	SetInternalValue(val *JobV1SpecTemplateSpecContainerSecurityContext)
	Privileged() any
	SetPrivileged(val any)
	PrivilegedInput() any
	ReadOnlyRootFilesystem() any
	SetReadOnlyRootFilesystem(val any)
	ReadOnlyRootFilesystemInput() any
	RunAsGroup() *string
	SetRunAsGroup(val *string)
	RunAsGroupInput() *string
	RunAsNonRoot() any
	SetRunAsNonRoot(val any)
	RunAsNonRootInput() any
	RunAsUser() *string
	SetRunAsUser(val *string)
	RunAsUserInput() *string
	SeccompProfile() JobV1SpecTemplateSpecContainerSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *JobV1SpecTemplateSpecContainerSecurityContextSeccompProfile
	SeLinuxOptions() JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptions
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
	PutCapabilities(value *JobV1SpecTemplateSpecContainerSecurityContextCapabilities)
	PutSeccompProfile(value *JobV1SpecTemplateSpecContainerSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptions)
	ResetAllowPrivilegeEscalation()
	ResetCapabilities()
	ResetPrivileged()
	ResetReadOnlyRootFilesystem()
	ResetRunAsGroup()
	ResetRunAsNonRoot()
	ResetRunAsUser()
	ResetSeccompProfile()
	ResetSeLinuxOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecContainerSecurityContextOutputReference
type jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) AllowPrivilegeEscalation() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) AllowPrivilegeEscalationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) Capabilities() JobV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference {
	var returns JobV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) CapabilitiesInput() *JobV1SpecTemplateSpecContainerSecurityContextCapabilities {
	var returns *JobV1SpecTemplateSpecContainerSecurityContextCapabilities
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) InternalValue() *JobV1SpecTemplateSpecContainerSecurityContext {
	var returns *JobV1SpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) Privileged() any {
	var returns any
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) PrivilegedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ReadOnlyRootFilesystem() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ReadOnlyRootFilesystemInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) RunAsNonRoot() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) RunAsNonRootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SeccompProfile() JobV1SpecTemplateSpecContainerSecurityContextSeccompProfileOutputReference {
	var returns JobV1SpecTemplateSpecContainerSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SeccompProfileInput() *JobV1SpecTemplateSpecContainerSecurityContextSeccompProfile {
	var returns *JobV1SpecTemplateSpecContainerSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SeLinuxOptions() JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptionsOutputReference {
	var returns JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SeLinuxOptionsInput() *JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptions {
	var returns *JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewJobV1SpecTemplateSpecContainerSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecTemplateSpecContainerSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecContainerSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecContainerSecurityContextOutputReference_Override(j JobV1SpecTemplateSpecContainerSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetAllowPrivilegeEscalation(val any) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetInternalValue(val *JobV1SpecTemplateSpecContainerSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetPrivileged(val any) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetReadOnlyRootFilesystem(val any) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetRunAsNonRoot(val any) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) PutCapabilities(value *JobV1SpecTemplateSpecContainerSecurityContextCapabilities) {
	if err := j.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCapabilities",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) PutSeccompProfile(value *JobV1SpecTemplateSpecContainerSecurityContextSeccompProfile) {
	if err := j.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSeccompProfile",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) PutSeLinuxOptions(value *JobV1SpecTemplateSpecContainerSecurityContextSeLinuxOptions) {
	if err := j.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSeLinuxOptions",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		j,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		j,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		j,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		j,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		j,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		j,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
