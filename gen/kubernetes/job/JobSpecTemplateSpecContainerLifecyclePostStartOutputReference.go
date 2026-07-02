package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/job/internal"
)

type JobSpecTemplateSpecContainerLifecyclePostStartOutputReference interface {
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
	Exec() JobSpecTemplateSpecContainerLifecyclePostStartExecOutputReference
	ExecInput() *JobSpecTemplateSpecContainerLifecyclePostStartExec
	// Experimental.
	Fqn() *string
	HttpGet() JobSpecTemplateSpecContainerLifecyclePostStartHttpGetOutputReference
	HttpGetInput() *JobSpecTemplateSpecContainerLifecyclePostStartHttpGet
	InternalValue() any
	SetInternalValue(val any)
	TcpSocket() JobSpecTemplateSpecContainerLifecyclePostStartTcpSocketList
	TcpSocketInput() any
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
	PutExec(value *JobSpecTemplateSpecContainerLifecyclePostStartExec)
	PutHttpGet(value *JobSpecTemplateSpecContainerLifecyclePostStartHttpGet)
	PutTcpSocket(value any)
	ResetExec()
	ResetHttpGet()
	ResetTcpSocket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobSpecTemplateSpecContainerLifecyclePostStartOutputReference
type jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) Exec() JobSpecTemplateSpecContainerLifecyclePostStartExecOutputReference {
	var returns JobSpecTemplateSpecContainerLifecyclePostStartExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ExecInput() *JobSpecTemplateSpecContainerLifecyclePostStartExec {
	var returns *JobSpecTemplateSpecContainerLifecyclePostStartExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) HttpGet() JobSpecTemplateSpecContainerLifecyclePostStartHttpGetOutputReference {
	var returns JobSpecTemplateSpecContainerLifecyclePostStartHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) HttpGetInput() *JobSpecTemplateSpecContainerLifecyclePostStartHttpGet {
	var returns *JobSpecTemplateSpecContainerLifecyclePostStartHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) TcpSocket() JobSpecTemplateSpecContainerLifecyclePostStartTcpSocketList {
	var returns JobSpecTemplateSpecContainerLifecyclePostStartTcpSocketList
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) TcpSocketInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewJobSpecTemplateSpecContainerLifecyclePostStartOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobSpecTemplateSpecContainerLifecyclePostStartOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecContainerLifecyclePostStartOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecContainerLifecyclePostStartOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecContainerLifecyclePostStartOutputReference_Override(j JobSpecTemplateSpecContainerLifecyclePostStartOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecContainerLifecyclePostStartOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) PutExec(value *JobSpecTemplateSpecContainerLifecyclePostStartExec) {
	if err := j.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putExec",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) PutHttpGet(value *JobSpecTemplateSpecContainerLifecyclePostStartHttpGet) {
	if err := j.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHttpGet",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) PutTcpSocket(value any) {
	if err := j.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTcpSocket",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		j,
		"resetExec",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		j,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		j,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (j *jsiiProxy_JobSpecTemplateSpecContainerLifecyclePostStartOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
