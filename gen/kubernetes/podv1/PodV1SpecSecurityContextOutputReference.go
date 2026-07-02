package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/podv1/internal"
)

type PodV1SpecSecurityContextOutputReference interface {
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
	InternalValue() *PodV1SpecSecurityContext
	SetInternalValue(val *PodV1SpecSecurityContext)
	RunAsGroup() *string
	SetRunAsGroup(val *string)
	RunAsGroupInput() *string
	RunAsNonRoot() any
	SetRunAsNonRoot(val any)
	RunAsNonRootInput() any
	RunAsUser() *string
	SetRunAsUser(val *string)
	RunAsUserInput() *string
	SeccompProfile() PodV1SpecSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *PodV1SpecSecurityContextSeccompProfile
	SeLinuxOptions() PodV1SpecSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *PodV1SpecSecurityContextSeLinuxOptions
	SupplementalGroups() *[]*float64
	SetSupplementalGroups(val *[]*float64)
	SupplementalGroupsInput() *[]*float64
	Sysctl() PodV1SpecSecurityContextSysctlList
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
	PutSeccompProfile(value *PodV1SpecSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *PodV1SpecSecurityContextSeLinuxOptions)
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

// The jsii proxy struct for PodV1SpecSecurityContextOutputReference
type jsiiProxy_PodV1SpecSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) FsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) FsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) InternalValue() *PodV1SpecSecurityContext {
	var returns *PodV1SpecSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) RunAsNonRoot() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) RunAsNonRootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SeccompProfile() PodV1SpecSecurityContextSeccompProfileOutputReference {
	var returns PodV1SpecSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SeccompProfileInput() *PodV1SpecSecurityContextSeccompProfile {
	var returns *PodV1SpecSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SeLinuxOptions() PodV1SpecSecurityContextSeLinuxOptionsOutputReference {
	var returns PodV1SpecSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SeLinuxOptionsInput() *PodV1SpecSecurityContextSeLinuxOptions {
	var returns *PodV1SpecSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SupplementalGroups() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"supplementalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SupplementalGroupsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"supplementalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) Sysctl() PodV1SpecSecurityContextSysctlList {
	var returns PodV1SpecSecurityContextSysctlList
	_jsii_.Get(
		j,
		"sysctl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SysctlInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sysctlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewPodV1SpecSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodV1SpecSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodV1SpecSecurityContextOutputReference_Override(p PodV1SpecSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecSecurityContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetFsGroup(val *string) {
	if err := j.validateSetFsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsGroup",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetInternalValue(val *PodV1SpecSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetRunAsNonRoot(val any) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetSupplementalGroups(val *[]*float64) {
	if err := j.validateSetSupplementalGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplementalGroups",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) PutSeccompProfile(value *PodV1SpecSecurityContextSeccompProfile) {
	if err := p.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSeccompProfile",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) PutSeLinuxOptions(value *PodV1SpecSecurityContextSeLinuxOptions) {
	if err := p.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSeLinuxOptions",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) PutSysctl(value any) {
	if err := p.validatePutSysctlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSysctl",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetFsGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetFsGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		p,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetSupplementalGroups() {
	_jsii_.InvokeVoid(
		p,
		"resetSupplementalGroups",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ResetSysctl() {
	_jsii_.InvokeVoid(
		p,
		"resetSysctl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
