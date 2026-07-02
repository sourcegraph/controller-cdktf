package replicationcontroller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/replicationcontroller/internal"
)

type ReplicationControllerSpecTemplateSpecSecurityContextOutputReference interface {
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
	FsGroup() *string
	SetFsGroup(val *string)
	FsGroupInput() *string
	InternalValue() *ReplicationControllerSpecTemplateSpecSecurityContext
	SetInternalValue(val *ReplicationControllerSpecTemplateSpecSecurityContext)
	RunAsGroup() *string
	SetRunAsGroup(val *string)
	RunAsGroupInput() *string
	RunAsNonRoot() any
	SetRunAsNonRoot(val any)
	RunAsNonRootInput() any
	RunAsUser() *string
	SetRunAsUser(val *string)
	RunAsUserInput() *string
	SeccompProfile() ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfile
	SeLinuxOptions() ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptions
	SupplementalGroups() *[]*float64
	SetSupplementalGroups(val *[]*float64)
	SupplementalGroupsInput() *[]*float64
	Sysctl() ReplicationControllerSpecTemplateSpecSecurityContextSysctlList
	SysctlInput() any
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
	PutSeccompProfile(value *ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptions)
	PutSysctl(value any)
	ResetFsGroup()
	ResetRunAsGroup()
	ResetRunAsNonRoot()
	ResetRunAsUser()
	ResetSeccompProfile()
	ResetSeLinuxOptions()
	ResetSupplementalGroups()
	ResetSysctl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerSpecTemplateSpecSecurityContextOutputReference
type jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) FsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) FsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) InternalValue() *ReplicationControllerSpecTemplateSpecSecurityContext {
	var returns *ReplicationControllerSpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) RunAsNonRoot() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) RunAsNonRootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SeccompProfile() ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfileOutputReference {
	var returns ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SeccompProfileInput() *ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfile {
	var returns *ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SeLinuxOptions() ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptionsOutputReference {
	var returns ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SeLinuxOptionsInput() *ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptions {
	var returns *ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SupplementalGroups() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"supplementalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SupplementalGroupsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"supplementalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) Sysctl() ReplicationControllerSpecTemplateSpecSecurityContextSysctlList {
	var returns ReplicationControllerSpecTemplateSpecSecurityContextSysctlList
	_jsii_.Get(
		j,
		"sysctl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SysctlInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sysctlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewReplicationControllerSpecTemplateSpecSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerSpecTemplateSpecSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerSpecTemplateSpecSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerSpecTemplateSpecSecurityContextOutputReference_Override(r ReplicationControllerSpecTemplateSpecSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetFsGroup(val *string) {
	if err := j.validateSetFsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsGroup",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetInternalValue(val *ReplicationControllerSpecTemplateSpecSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetRunAsNonRoot(val any) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetSupplementalGroups(val *[]*float64) {
	if err := j.validateSetSupplementalGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplementalGroups",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) PutSeccompProfile(value *ReplicationControllerSpecTemplateSpecSecurityContextSeccompProfile) {
	if err := r.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSeccompProfile",
		[]any{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) PutSeLinuxOptions(value *ReplicationControllerSpecTemplateSpecSecurityContextSeLinuxOptions) {
	if err := r.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSeLinuxOptions",
		[]any{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) PutSysctl(value any) {
	if err := r.validatePutSysctlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSysctl",
		[]any{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetFsGroup() {
	_jsii_.InvokeVoid(
		r,
		"resetFsGroup",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		r,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		r,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		r,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		r,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		r,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetSupplementalGroups() {
	_jsii_.InvokeVoid(
		r,
		"resetSupplementalGroups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ResetSysctl() {
	_jsii_.InvokeVoid(
		r,
		"resetSysctl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		r,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
