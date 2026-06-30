package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataproccluster/internal"
)

type DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference interface {
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
	InstanceSelectionList() DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListStructList
	InstanceSelectionListInput() any
	InstanceSelectionResults() DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionResultsList
	InternalValue() *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy
	SetInternalValue(val *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy)
	ProvisioningModelMix() DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMixOutputReference
	ProvisioningModelMixInput() *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix
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
	PutInstanceSelectionList(value any)
	PutProvisioningModelMix(value *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix)
	ResetInstanceSelectionList()
	ResetProvisioningModelMix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference
type jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) InstanceSelectionList() DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListStructList {
	var returns DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListStructList
	_jsii_.Get(
		j,
		"instanceSelectionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) InstanceSelectionListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"instanceSelectionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) InstanceSelectionResults() DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionResultsList {
	var returns DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionResultsList
	_jsii_.Get(
		j,
		"instanceSelectionResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) InternalValue() *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy {
	var returns *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ProvisioningModelMix() DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMixOutputReference {
	var returns DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMixOutputReference
	_jsii_.Get(
		j,
		"provisioningModelMix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ProvisioningModelMixInput() *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix {
	var returns *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix
	_jsii_.Get(
		j,
		"provisioningModelMixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference_Override(d DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) SetInternalValue(val *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) PutInstanceSelectionList(value any) {
	if err := d.validatePutInstanceSelectionListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstanceSelectionList",
		[]any{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) PutProvisioningModelMix(value *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix) {
	if err := d.validatePutProvisioningModelMixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisioningModelMix",
		[]any{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ResetInstanceSelectionList() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceSelectionList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ResetProvisioningModelMix() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisioningModelMix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
