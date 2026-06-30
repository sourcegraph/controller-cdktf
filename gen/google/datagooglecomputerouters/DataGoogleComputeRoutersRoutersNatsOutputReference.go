package datagooglecomputerouters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglecomputerouters/internal"
)

type DataGoogleComputeRoutersRoutersNatsOutputReference interface {
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
	EnableEndpointIndependentMapping() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	IcmpIdleTimeoutSec() *float64
	InternalValue() *DataGoogleComputeRoutersRoutersNats
	SetInternalValue(val *DataGoogleComputeRoutersRoutersNats)
	MinPortsPerVm() *float64
	Name() *string
	NatIpAllocateOption() *string
	NatIps() *[]*string
	SourceSubnetworkIpRangesToNat() *string
	TcpEstablishedIdleTimeoutSec() *float64
	TcpTransitoryIdleTimeoutSec() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UdpIdleTimeoutSec() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComputeRoutersRoutersNatsOutputReference
type jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) EnableEndpointIndependentMapping() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableEndpointIndependentMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) IcmpIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) InternalValue() *DataGoogleComputeRoutersRoutersNats {
	var returns *DataGoogleComputeRoutersRoutersNats
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) MinPortsPerVm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPortsPerVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) NatIpAllocateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAllocateOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) NatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) SourceSubnetworkIpRangesToNat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) TcpEstablishedIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) TcpTransitoryIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTransitoryIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) UdpIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpIdleTimeoutSec",
		&returns,
	)
	return returns
}

func NewDataGoogleComputeRoutersRoutersNatsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeRoutersRoutersNatsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeRoutersRoutersNatsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeRouters.DataGoogleComputeRoutersRoutersNatsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeRoutersRoutersNatsOutputReference_Override(d DataGoogleComputeRoutersRoutersNatsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeRouters.DataGoogleComputeRoutersRoutersNatsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) SetInternalValue(val *DataGoogleComputeRoutersRoutersNats) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataGoogleComputeRoutersRoutersNatsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
