package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterIpAllocationPolicyOutputReference interface {
	cdktf.ComplexObject
	ClusterIpv4CidrBlock() *string
	SetClusterIpv4CidrBlock(val *string)
	ClusterIpv4CidrBlockInput() *string
	ClusterSecondaryRangeName() *string
	SetClusterSecondaryRangeName(val *string)
	ClusterSecondaryRangeNameInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerClusterIpAllocationPolicy
	SetInternalValue(val *GoogleContainerClusterIpAllocationPolicy)
	ServicesIpv4CidrBlock() *string
	SetServicesIpv4CidrBlock(val *string)
	ServicesIpv4CidrBlockInput() *string
	ServicesSecondaryRangeName() *string
	SetServicesSecondaryRangeName(val *string)
	ServicesSecondaryRangeNameInput() *string
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	ResetClusterIpv4CidrBlock()
	ResetClusterSecondaryRangeName()
	ResetServicesIpv4CidrBlock()
	ResetServicesSecondaryRangeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterIpAllocationPolicyOutputReference
type jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ClusterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ClusterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ClusterSecondaryRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecondaryRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ClusterSecondaryRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecondaryRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) InternalValue() *GoogleContainerClusterIpAllocationPolicy {
	var returns *GoogleContainerClusterIpAllocationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ServicesIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ServicesIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ServicesSecondaryRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesSecondaryRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ServicesSecondaryRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesSecondaryRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterIpAllocationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterIpAllocationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterIpAllocationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterIpAllocationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterIpAllocationPolicyOutputReference_Override(g GoogleContainerClusterIpAllocationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterIpAllocationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetClusterIpv4CidrBlock(val *string) {
	if err := j.validateSetClusterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetClusterSecondaryRangeName(val *string) {
	if err := j.validateSetClusterSecondaryRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSecondaryRangeName",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetInternalValue(val *GoogleContainerClusterIpAllocationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetServicesIpv4CidrBlock(val *string) {
	if err := j.validateSetServicesIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetServicesSecondaryRangeName(val *string) {
	if err := j.validateSetServicesSecondaryRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesSecondaryRangeName",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ResetClusterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ResetClusterSecondaryRangeName() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterSecondaryRangeName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ResetServicesIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetServicesIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ResetServicesSecondaryRangeName() {
	_jsii_.InvokeVoid(
		g,
		"resetServicesSecondaryRangeName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterIpAllocationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

