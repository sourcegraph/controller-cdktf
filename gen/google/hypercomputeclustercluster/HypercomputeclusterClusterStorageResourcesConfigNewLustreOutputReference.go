package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference interface {
	cdktf.ComplexObject
	CapacityGb() *string
	SetCapacityGb(val *string)
	CapacityGbInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Filesystem() *string
	SetFilesystem(val *string)
	FilesystemInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HypercomputeclusterClusterStorageResourcesConfigNewLustre
	SetInternalValue(val *HypercomputeclusterClusterStorageResourcesConfigNewLustre)
	Lustre() *string
	SetLustre(val *string)
	LustreInput() *string
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
	ResetDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference
type jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) CapacityGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) CapacityGbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) Filesystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) FilesystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) InternalValue() *HypercomputeclusterClusterStorageResourcesConfigNewLustre {
	var returns *HypercomputeclusterClusterStorageResourcesConfigNewLustre
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) Lustre() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lustre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) LustreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lustreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewHypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference_Override(h HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference",
		[]any{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetCapacityGb(val *string) {
	if err := j.validateSetCapacityGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityGb",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetFilesystem(val *string) {
	if err := j.validateSetFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filesystem",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetInternalValue(val *HypercomputeclusterClusterStorageResourcesConfigNewLustre) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetLustre(val *string) {
	if err := j.validateSetLustreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lustre",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		h,
		"resetDescription",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (h *jsiiProxy_HypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
