package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/pod/internal"
)

type PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference interface {
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
	InternalValue() *PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution
	SetInternalValue(val *PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)
	NodeSelectorTerm() PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermList
	NodeSelectorTermInput() any
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
	PutNodeSelectorTerm(value any)
	ResetNodeSelectorTerm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference
type jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) InternalValue() *PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution {
	var returns *PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) NodeSelectorTerm() PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermList {
	var returns PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermList
	_jsii_.Get(
		j,
		"nodeSelectorTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) NodeSelectorTermInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nodeSelectorTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewPodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference_Override(p PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) SetInternalValue(val *PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) PutNodeSelectorTerm(value any) {
	if err := p.validatePutNodeSelectorTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNodeSelectorTerm",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ResetNodeSelectorTerm() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeSelectorTerm",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (p *jsiiProxy_PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
