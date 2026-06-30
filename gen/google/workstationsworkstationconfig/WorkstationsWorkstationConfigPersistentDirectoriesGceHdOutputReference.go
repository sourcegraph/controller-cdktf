package workstationsworkstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/workstationsworkstationconfig/internal"
)

type WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference interface {
	cdktf.ComplexObject
	ArchiveTimeout() *string
	SetArchiveTimeout(val *string)
	ArchiveTimeoutInput() *string
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
	InternalValue() *WorkstationsWorkstationConfigPersistentDirectoriesGceHd
	SetInternalValue(val *WorkstationsWorkstationConfigPersistentDirectoriesGceHd)
	ReclaimPolicy() *string
	SetReclaimPolicy(val *string)
	ReclaimPolicyInput() *string
	SizeGb() *float64
	SetSizeGb(val *float64)
	SizeGbInput() *float64
	SourceSnapshot() *string
	SetSourceSnapshot(val *string)
	SourceSnapshotInput() *string
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
	ResetArchiveTimeout()
	ResetReclaimPolicy()
	ResetSizeGb()
	ResetSourceSnapshot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference
type jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ArchiveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ArchiveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) InternalValue() *WorkstationsWorkstationConfigPersistentDirectoriesGceHd {
	var returns *WorkstationsWorkstationConfigPersistentDirectoriesGceHd
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ReclaimPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ReclaimPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reclaimPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SourceSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SourceSnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewWorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference {
	_init_.Initialize()

	if err := validateNewWorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference_Override(w WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference",
		[]any{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetArchiveTimeout(val *string) {
	if err := j.validateSetArchiveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveTimeout",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetInternalValue(val *WorkstationsWorkstationConfigPersistentDirectoriesGceHd) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetReclaimPolicy(val *string) {
	if err := j.validateSetReclaimPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reclaimPolicy",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetSizeGb(val *float64) {
	if err := j.validateSetSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeGb",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetSourceSnapshot(val *string) {
	if err := j.validateSetSourceSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSnapshot",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ResetArchiveTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetArchiveTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ResetReclaimPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetReclaimPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ResetSizeGb() {
	_jsii_.InvokeVoid(
		w,
		"resetSizeGb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ResetSourceSnapshot() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceSnapshot",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (w *jsiiProxy_WorkstationsWorkstationConfigPersistentDirectoriesGceHdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
