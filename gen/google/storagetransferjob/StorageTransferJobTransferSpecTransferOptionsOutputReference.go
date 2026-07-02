package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/storagetransferjob/internal"
)

type StorageTransferJobTransferSpecTransferOptionsOutputReference interface {
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
	DeleteObjectsFromSourceAfterTransfer() any
	SetDeleteObjectsFromSourceAfterTransfer(val any)
	DeleteObjectsFromSourceAfterTransferInput() any
	DeleteObjectsUniqueInSink() any
	SetDeleteObjectsUniqueInSink(val any)
	DeleteObjectsUniqueInSinkInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *StorageTransferJobTransferSpecTransferOptions
	SetInternalValue(val *StorageTransferJobTransferSpecTransferOptions)
	MetadataOptions() StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference
	MetadataOptionsInput() *StorageTransferJobTransferSpecTransferOptionsMetadataOptions
	OverwriteObjectsAlreadyExistingInSink() any
	SetOverwriteObjectsAlreadyExistingInSink(val any)
	OverwriteObjectsAlreadyExistingInSinkInput() any
	OverwriteWhen() *string
	SetOverwriteWhen(val *string)
	OverwriteWhenInput() *string
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
	PutMetadataOptions(value *StorageTransferJobTransferSpecTransferOptionsMetadataOptions)
	ResetDeleteObjectsFromSourceAfterTransfer()
	ResetDeleteObjectsUniqueInSink()
	ResetMetadataOptions()
	ResetOverwriteObjectsAlreadyExistingInSink()
	ResetOverwriteWhen()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobTransferSpecTransferOptionsOutputReference
type jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) DeleteObjectsFromSourceAfterTransfer() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteObjectsFromSourceAfterTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) DeleteObjectsFromSourceAfterTransferInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteObjectsFromSourceAfterTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) DeleteObjectsUniqueInSink() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteObjectsUniqueInSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) DeleteObjectsUniqueInSinkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteObjectsUniqueInSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) InternalValue() *StorageTransferJobTransferSpecTransferOptions {
	var returns *StorageTransferJobTransferSpecTransferOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) MetadataOptions() StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference {
	var returns StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) MetadataOptionsInput() *StorageTransferJobTransferSpecTransferOptionsMetadataOptions {
	var returns *StorageTransferJobTransferSpecTransferOptionsMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) OverwriteObjectsAlreadyExistingInSink() any {
	var returns any
	_jsii_.Get(
		j,
		"overwriteObjectsAlreadyExistingInSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) OverwriteObjectsAlreadyExistingInSinkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"overwriteObjectsAlreadyExistingInSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) OverwriteWhen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteWhen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) OverwriteWhenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteWhenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStorageTransferJobTransferSpecTransferOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobTransferSpecTransferOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobTransferSpecTransferOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecTransferOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobTransferSpecTransferOptionsOutputReference_Override(s StorageTransferJobTransferSpecTransferOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecTransferOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetDeleteObjectsFromSourceAfterTransfer(val any) {
	if err := j.validateSetDeleteObjectsFromSourceAfterTransferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteObjectsFromSourceAfterTransfer",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetDeleteObjectsUniqueInSink(val any) {
	if err := j.validateSetDeleteObjectsUniqueInSinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteObjectsUniqueInSink",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetInternalValue(val *StorageTransferJobTransferSpecTransferOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetOverwriteObjectsAlreadyExistingInSink(val any) {
	if err := j.validateSetOverwriteObjectsAlreadyExistingInSinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteObjectsAlreadyExistingInSink",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetOverwriteWhen(val *string) {
	if err := j.validateSetOverwriteWhenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteWhen",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) PutMetadataOptions(value *StorageTransferJobTransferSpecTransferOptionsMetadataOptions) {
	if err := s.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetadataOptions",
		[]any{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ResetDeleteObjectsFromSourceAfterTransfer() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteObjectsFromSourceAfterTransfer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ResetDeleteObjectsUniqueInSink() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteObjectsUniqueInSink",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ResetOverwriteObjectsAlreadyExistingInSink() {
	_jsii_.InvokeVoid(
		s,
		"resetOverwriteObjectsAlreadyExistingInSink",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ResetOverwriteWhen() {
	_jsii_.InvokeVoid(
		s,
		"resetOverwriteWhen",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		s,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
