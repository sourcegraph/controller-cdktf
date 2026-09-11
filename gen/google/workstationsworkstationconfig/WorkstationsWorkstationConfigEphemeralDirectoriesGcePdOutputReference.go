package workstationsworkstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/workstationsworkstationconfig/internal"
)

type WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
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
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WorkstationsWorkstationConfigEphemeralDirectoriesGcePd
	SetInternalValue(val *WorkstationsWorkstationConfigEphemeralDirectoriesGcePd)
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SourceImage() *string
	SetSourceImage(val *string)
	SourceImageInput() *string
	SourceSnapshot() *string
	SetSourceSnapshot(val *string)
	SourceSnapshotInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetDiskType()
	ResetReadOnly()
	ResetSourceImage()
	ResetSourceSnapshot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference
type jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) InternalValue() *WorkstationsWorkstationConfigEphemeralDirectoriesGcePd {
	var returns *WorkstationsWorkstationConfigEphemeralDirectoriesGcePd
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) SourceImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) SourceImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) SourceSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) SourceSnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference {
	_init_.Initialize()

	if err := validateNewWorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference_Override(w WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetInternalValue(val *WorkstationsWorkstationConfigEphemeralDirectoriesGcePd) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetSourceImage(val *string) {
	if err := j.validateSetSourceImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImage",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetSourceSnapshot(val *string) {
	if err := j.validateSetSourceSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSnapshot",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		w,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ResetSourceImage() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ResetSourceSnapshot() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceSnapshot",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigEphemeralDirectoriesGcePdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

