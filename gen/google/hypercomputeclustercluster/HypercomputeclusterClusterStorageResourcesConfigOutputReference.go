package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterStorageResourcesConfigOutputReference interface {
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
	ExistingBucket() HypercomputeclusterClusterStorageResourcesConfigExistingBucketOutputReference
	ExistingBucketInput() *HypercomputeclusterClusterStorageResourcesConfigExistingBucket
	ExistingFilestore() HypercomputeclusterClusterStorageResourcesConfigExistingFilestoreOutputReference
	ExistingFilestoreInput() *HypercomputeclusterClusterStorageResourcesConfigExistingFilestore
	ExistingLustre() HypercomputeclusterClusterStorageResourcesConfigExistingLustreOutputReference
	ExistingLustreInput() *HypercomputeclusterClusterStorageResourcesConfigExistingLustre
	// Experimental.
	Fqn() *string
	InternalValue() *HypercomputeclusterClusterStorageResourcesConfig
	SetInternalValue(val *HypercomputeclusterClusterStorageResourcesConfig)
	NewBucket() HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference
	NewBucketInput() *HypercomputeclusterClusterStorageResourcesConfigNewBucket
	NewFilestore() HypercomputeclusterClusterStorageResourcesConfigNewFilestoreOutputReference
	NewFilestoreInput() *HypercomputeclusterClusterStorageResourcesConfigNewFilestore
	NewLustre() HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference
	NewLustreInput() *HypercomputeclusterClusterStorageResourcesConfigNewLustre
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
	PutExistingBucket(value *HypercomputeclusterClusterStorageResourcesConfigExistingBucket)
	PutExistingFilestore(value *HypercomputeclusterClusterStorageResourcesConfigExistingFilestore)
	PutExistingLustre(value *HypercomputeclusterClusterStorageResourcesConfigExistingLustre)
	PutNewBucket(value *HypercomputeclusterClusterStorageResourcesConfigNewBucket)
	PutNewFilestore(value *HypercomputeclusterClusterStorageResourcesConfigNewFilestore)
	PutNewLustre(value *HypercomputeclusterClusterStorageResourcesConfigNewLustre)
	ResetExistingBucket()
	ResetExistingFilestore()
	ResetExistingLustre()
	ResetNewBucket()
	ResetNewFilestore()
	ResetNewLustre()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterStorageResourcesConfigOutputReference
type jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingBucket() HypercomputeclusterClusterStorageResourcesConfigExistingBucketOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigExistingBucketOutputReference
	_jsii_.Get(
		j,
		"existingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingBucketInput() *HypercomputeclusterClusterStorageResourcesConfigExistingBucket {
	var returns *HypercomputeclusterClusterStorageResourcesConfigExistingBucket
	_jsii_.Get(
		j,
		"existingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingFilestore() HypercomputeclusterClusterStorageResourcesConfigExistingFilestoreOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigExistingFilestoreOutputReference
	_jsii_.Get(
		j,
		"existingFilestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingFilestoreInput() *HypercomputeclusterClusterStorageResourcesConfigExistingFilestore {
	var returns *HypercomputeclusterClusterStorageResourcesConfigExistingFilestore
	_jsii_.Get(
		j,
		"existingFilestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingLustre() HypercomputeclusterClusterStorageResourcesConfigExistingLustreOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigExistingLustreOutputReference
	_jsii_.Get(
		j,
		"existingLustre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingLustreInput() *HypercomputeclusterClusterStorageResourcesConfigExistingLustre {
	var returns *HypercomputeclusterClusterStorageResourcesConfigExistingLustre
	_jsii_.Get(
		j,
		"existingLustreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) InternalValue() *HypercomputeclusterClusterStorageResourcesConfig {
	var returns *HypercomputeclusterClusterStorageResourcesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) NewBucket() HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference
	_jsii_.Get(
		j,
		"newBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) NewBucketInput() *HypercomputeclusterClusterStorageResourcesConfigNewBucket {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewBucket
	_jsii_.Get(
		j,
		"newBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) NewFilestore() HypercomputeclusterClusterStorageResourcesConfigNewFilestoreOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigNewFilestoreOutputReference
	_jsii_.Get(
		j,
		"newFilestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) NewFilestoreInput() *HypercomputeclusterClusterStorageResourcesConfigNewFilestore {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewFilestore
	_jsii_.Get(
		j,
		"newFilestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) NewLustre() HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference
	_jsii_.Get(
		j,
		"newLustre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) NewLustreInput() *HypercomputeclusterClusterStorageResourcesConfigNewLustre {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewLustre
	_jsii_.Get(
		j,
		"newLustreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewHypercomputeclusterClusterStorageResourcesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterStorageResourcesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterStorageResourcesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterStorageResourcesConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterStorageResourcesConfigOutputReference_Override(h HypercomputeclusterClusterStorageResourcesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterStorageResourcesConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) SetInternalValue(val *HypercomputeclusterClusterStorageResourcesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) PutExistingBucket(value *HypercomputeclusterClusterStorageResourcesConfigExistingBucket) {
	if err := h.validatePutExistingBucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExistingBucket",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) PutExistingFilestore(value *HypercomputeclusterClusterStorageResourcesConfigExistingFilestore) {
	if err := h.validatePutExistingFilestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExistingFilestore",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) PutExistingLustre(value *HypercomputeclusterClusterStorageResourcesConfigExistingLustre) {
	if err := h.validatePutExistingLustreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExistingLustre",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) PutNewBucket(value *HypercomputeclusterClusterStorageResourcesConfigNewBucket) {
	if err := h.validatePutNewBucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewBucket",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) PutNewFilestore(value *HypercomputeclusterClusterStorageResourcesConfigNewFilestore) {
	if err := h.validatePutNewFilestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewFilestore",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) PutNewLustre(value *HypercomputeclusterClusterStorageResourcesConfigNewLustre) {
	if err := h.validatePutNewLustreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewLustre",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetExistingBucket() {
	_jsii_.InvokeVoid(
		h,
		"resetExistingBucket",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetExistingFilestore() {
	_jsii_.InvokeVoid(
		h,
		"resetExistingFilestore",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetExistingLustre() {
	_jsii_.InvokeVoid(
		h,
		"resetExistingLustre",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetNewBucket() {
	_jsii_.InvokeVoid(
		h,
		"resetNewBucket",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetNewFilestore() {
	_jsii_.InvokeVoid(
		h,
		"resetNewFilestore",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetNewLustre() {
	_jsii_.InvokeVoid(
		h,
		"resetNewLustre",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		h,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
