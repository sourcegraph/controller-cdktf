package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference interface {
	cdktf.ComplexObject
	Autoclass() HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclassOutputReference
	AutoclassInput() *HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	HierarchicalNamespace() HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespaceOutputReference
	HierarchicalNamespaceInput() *HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace
	InternalValue() *HypercomputeclusterClusterStorageResourcesConfigNewBucket
	SetInternalValue(val *HypercomputeclusterClusterStorageResourcesConfigNewBucket)
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	PutAutoclass(value *HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass)
	PutHierarchicalNamespace(value *HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace)
	ResetAutoclass()
	ResetHierarchicalNamespace()
	ResetStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference
type jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Autoclass() HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclassOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclassOutputReference
	_jsii_.Get(
		j,
		"autoclass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) AutoclassInput() *HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass
	_jsii_.Get(
		j,
		"autoclassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) HierarchicalNamespace() HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespaceOutputReference {
	var returns HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespaceOutputReference
	_jsii_.Get(
		j,
		"hierarchicalNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) HierarchicalNamespaceInput() *HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace
	_jsii_.Get(
		j,
		"hierarchicalNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) InternalValue() *HypercomputeclusterClusterStorageResourcesConfigNewBucket {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewBucket
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference_Override(h HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference",
		[]any{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetInternalValue(val *HypercomputeclusterClusterStorageResourcesConfigNewBucket) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) PutAutoclass(value *HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass) {
	if err := h.validatePutAutoclassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putAutoclass",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) PutHierarchicalNamespace(value *HypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace) {
	if err := h.validatePutHierarchicalNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putHierarchicalNamespace",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ResetAutoclass() {
	_jsii_.InvokeVoid(
		h,
		"resetAutoclass",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ResetHierarchicalNamespace() {
	_jsii_.InvokeVoid(
		h,
		"resetHierarchicalNamespace",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
