package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/job/internal"
)

type JobSpecTemplateSpecContainerSecurityContextOutputReference interface {
	cdktf.ComplexObject
	AllowPrivilegeEscalation() any
	SetAllowPrivilegeEscalation(val any)
	AllowPrivilegeEscalationInput() any
	Capabilities() JobSpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference
	CapabilitiesInput() *JobSpecTemplateSpecContainerSecurityContextCapabilities
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
	InternalValue() *JobSpecTemplateSpecContainerSecurityContext
	SetInternalValue(val *JobSpecTemplateSpecContainerSecurityContext)
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
	SeccompProfile() JobSpecTemplateSpecContainerSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *JobSpecTemplateSpecContainerSecurityContextSeccompProfile
	SeLinuxOptions() JobSpecTemplateSpecContainerSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *JobSpecTemplateSpecContainerSecurityContextSeLinuxOptions
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
	PutCapabilities(value *JobSpecTemplateSpecContainerSecurityContextCapabilities)
	PutSeccompProfile(value *JobSpecTemplateSpecContainerSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *JobSpecTemplateSpecContainerSecurityContextSeLinuxOptions)
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

// The jsii proxy struct for JobSpecTemplateSpecContainerSecurityContextOutputReference
type jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) AllowPrivilegeEscalation() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) AllowPrivilegeEscalationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) Capabilities() JobSpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference {
	var returns JobSpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) CapabilitiesInput() *JobSpecTemplateSpecContainerSecurityContextCapabilities {
	var returns *JobSpecTemplateSpecContainerSecurityContextCapabilities
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) InternalValue() *JobSpecTemplateSpecContainerSecurityContext {
	var returns *JobSpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) Privileged() any {
	var returns any
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) PrivilegedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ReadOnlyRootFilesystem() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ReadOnlyRootFilesystemInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) RunAsNonRoot() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) RunAsNonRootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SeccompProfile() JobSpecTemplateSpecContainerSecurityContextSeccompProfileOutputReference {
	var returns JobSpecTemplateSpecContainerSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SeccompProfileInput() *JobSpecTemplateSpecContainerSecurityContextSeccompProfile {
	var returns *JobSpecTemplateSpecContainerSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SeLinuxOptions() JobSpecTemplateSpecContainerSecurityContextSeLinuxOptionsOutputReference {
	var returns JobSpecTemplateSpecContainerSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SeLinuxOptionsInput() *JobSpecTemplateSpecContainerSecurityContextSeLinuxOptions {
	var returns *JobSpecTemplateSpecContainerSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewJobSpecTemplateSpecContainerSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobSpecTemplateSpecContainerSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecContainerSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecContainerSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecContainerSecurityContextOutputReference_Override(j JobSpecTemplateSpecContainerSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecContainerSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetAllowPrivilegeEscalation(val any) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetInternalValue(val *JobSpecTemplateSpecContainerSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetPrivileged(val any) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetReadOnlyRootFilesystem(val any) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetRunAsNonRoot(val any) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) PutCapabilities(value *JobSpecTemplateSpecContainerSecurityContextCapabilities) {
	if err := j.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCapabilities",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) PutSeccompProfile(value *JobSpecTemplateSpecContainerSecurityContextSeccompProfile) {
	if err := j.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSeccompProfile",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) PutSeLinuxOptions(value *JobSpecTemplateSpecContainerSecurityContextSeLinuxOptions) {
	if err := j.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSeLinuxOptions",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		j,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		j,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		j,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		j,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		j,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
