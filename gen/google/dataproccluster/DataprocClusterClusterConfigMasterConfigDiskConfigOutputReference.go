package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataproccluster/internal"
)

type DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference interface {
	cdktf.ComplexObject
	BootDiskProvisionedIops() *float64
	SetBootDiskProvisionedIops(val *float64)
	BootDiskProvisionedIopsInput() *float64
	BootDiskProvisionedThroughput() *float64
	SetBootDiskProvisionedThroughput(val *float64)
	BootDiskProvisionedThroughputInput() *float64
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
	BootDiskType() *string
	SetBootDiskType(val *string)
	BootDiskTypeInput() *string
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
	InternalValue() *DataprocClusterClusterConfigMasterConfigDiskConfig
	SetInternalValue(val *DataprocClusterClusterConfigMasterConfigDiskConfig)
	LocalSsdInterface() *string
	SetLocalSsdInterface(val *string)
	LocalSsdInterfaceInput() *string
	NumLocalSsds() *float64
	SetNumLocalSsds(val *float64)
	NumLocalSsdsInput() *float64
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
	ResetBootDiskProvisionedIops()
	ResetBootDiskProvisionedThroughput()
	ResetBootDiskSizeGb()
	ResetBootDiskType()
	ResetLocalSsdInterface()
	ResetNumLocalSsds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) InternalValue() *DataprocClusterClusterConfigMasterConfigDiskConfig {
	var returns *DataprocClusterClusterConfigMasterConfigDiskConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) LocalSsdInterface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) LocalSsdInterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) NumLocalSsds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) NumLocalSsdsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataprocClusterClusterConfigMasterConfigDiskConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigMasterConfigDiskConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigMasterConfigDiskConfigOutputReference_Override(d DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetBootDiskProvisionedIops(val *float64) {
	if err := j.validateSetBootDiskProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedIops",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetBootDiskProvisionedThroughput(val *float64) {
	if err := j.validateSetBootDiskProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetInternalValue(val *DataprocClusterClusterConfigMasterConfigDiskConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetLocalSsdInterface(val *string) {
	if err := j.validateSetLocalSsdInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdInterface",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetNumLocalSsds(val *float64) {
	if err := j.validateSetNumLocalSsdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numLocalSsds",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ResetBootDiskProvisionedIops() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskProvisionedIops",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ResetBootDiskProvisionedThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskProvisionedThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ResetLocalSsdInterface() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSsdInterface",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ResetNumLocalSsds() {
	_jsii_.InvokeVoid(
		d,
		"resetNumLocalSsds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
