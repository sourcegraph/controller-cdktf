package emrcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/emrcluster/internal"
)

type EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference interface {
	cdktf.ComplexObject
	BidPrice() *string
	SetBidPrice(val *string)
	BidPriceAsPercentageOfOnDemandPrice() *float64
	SetBidPriceAsPercentageOfOnDemandPrice(val *float64)
	BidPriceAsPercentageOfOnDemandPriceInput() *float64
	BidPriceInput() *string
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
	Configurations() EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurationsList
	ConfigurationsInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EbsConfig() EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfigList
	EbsConfigInput() any
	// Experimental.
	Fqn() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() any
	SetInternalValue(val any)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeightedCapacity() *float64
	SetWeightedCapacity(val *float64)
	WeightedCapacityInput() *float64
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
	PutConfigurations(value any)
	PutEbsConfig(value any)
	ResetBidPrice()
	ResetBidPriceAsPercentageOfOnDemandPrice()
	ResetConfigurations()
	ResetEbsConfig()
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference
type jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) BidPriceAsPercentageOfOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPriceAsPercentageOfOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) BidPriceAsPercentageOfOnDemandPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPriceAsPercentageOfOnDemandPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) BidPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) Configurations() EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurationsList {
	var returns EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurationsList
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ConfigurationsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) EbsConfig() EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfigList {
	var returns EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfigList
	_jsii_.Get(
		j,
		"ebsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) EbsConfigInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ebsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) WeightedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) WeightedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}

func NewEmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewEmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference_Override(e EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetBidPrice(val *string) {
	if err := j.validateSetBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetBidPriceAsPercentageOfOnDemandPrice(val *float64) {
	if err := j.validateSetBidPriceAsPercentageOfOnDemandPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPriceAsPercentageOfOnDemandPrice",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) SetWeightedCapacity(val *float64) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) PutConfigurations(value any) {
	if err := e.validatePutConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConfigurations",
		[]any{value},
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) PutEbsConfig(value any) {
	if err := e.validatePutEbsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEbsConfig",
		[]any{value},
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ResetBidPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ResetBidPriceAsPercentageOfOnDemandPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPriceAsPercentageOfOnDemandPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ResetConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ResetEbsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		e,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetInstanceTypeConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
