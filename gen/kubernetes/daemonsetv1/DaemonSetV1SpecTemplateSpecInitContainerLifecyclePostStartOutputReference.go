package daemonsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/daemonsetv1/internal"
)

type DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference interface {
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
	Exec() DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExecOutputReference
	ExecInput() *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExec
	// Experimental.
	Fqn() *string
	HttpGet() DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGetOutputReference
	HttpGetInput() *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGet
	InternalValue() any
	SetInternalValue(val any)
	TcpSocket() DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartTcpSocketList
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
	PutExec(value *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExec)
	PutHttpGet(value *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGet)
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

// The jsii proxy struct for DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference
type jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) Exec() DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExecOutputReference {
	var returns DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ExecInput() *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExec {
	var returns *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) HttpGet() DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGetOutputReference {
	var returns DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) HttpGetInput() *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGet {
	var returns *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) TcpSocket() DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartTcpSocketList {
	var returns DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartTcpSocketList
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) TcpSocketInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonSetV1.DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference_Override(d DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonSetV1.DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) PutExec(value *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartExec) {
	if err := d.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExec",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) PutHttpGet(value *DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartHttpGet) {
	if err := d.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHttpGet",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) PutTcpSocket(value any) {
	if err := d.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTcpSocket",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		d,
		"resetExec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		d,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecInitContainerLifecyclePostStartOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
