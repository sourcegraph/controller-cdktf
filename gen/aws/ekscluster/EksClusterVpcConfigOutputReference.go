package ekscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/ekscluster/internal"
)

type EksClusterVpcConfigOutputReference interface {
	cdktf.ComplexObject
	ClusterSecurityGroupId() *string
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
	EndpointPrivateAccess() any
	SetEndpointPrivateAccess(val any)
	EndpointPrivateAccessInput() any
	EndpointPublicAccess() any
	SetEndpointPublicAccess(val any)
	EndpointPublicAccessInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *EksClusterVpcConfig
	SetInternalValue(val *EksClusterVpcConfig)
	PublicAccessCidrs() *[]*string
	SetPublicAccessCidrs(val *[]*string)
	PublicAccessCidrsInput() *[]*string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
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
	ResetEndpointPrivateAccess()
	ResetEndpointPublicAccess()
	ResetPublicAccessCidrs()
	ResetSecurityGroupIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EksClusterVpcConfigOutputReference
type jsiiProxy_EksClusterVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) EndpointPrivateAccess() any {
	var returns any
	_jsii_.Get(
		j,
		"endpointPrivateAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) EndpointPrivateAccessInput() any {
	var returns any
	_jsii_.Get(
		j,
		"endpointPrivateAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) EndpointPublicAccess() any {
	var returns any
	_jsii_.Get(
		j,
		"endpointPublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) EndpointPublicAccessInput() any {
	var returns any
	_jsii_.Get(
		j,
		"endpointPublicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) InternalValue() *EksClusterVpcConfig {
	var returns *EksClusterVpcConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) PublicAccessCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicAccessCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) PublicAccessCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicAccessCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func NewEksClusterVpcConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EksClusterVpcConfigOutputReference {
	_init_.Initialize()

	if err := validateNewEksClusterVpcConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksClusterVpcConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.eksCluster.EksClusterVpcConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEksClusterVpcConfigOutputReference_Override(e EksClusterVpcConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.eksCluster.EksClusterVpcConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetEndpointPrivateAccess(val any) {
	if err := j.validateSetEndpointPrivateAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointPrivateAccess",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetEndpointPublicAccess(val any) {
	if err := j.validateSetEndpointPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointPublicAccess",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetInternalValue(val *EksClusterVpcConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetPublicAccessCidrs(val *[]*string) {
	if err := j.validateSetPublicAccessCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccessCidrs",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EksClusterVpcConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) ResetEndpointPrivateAccess() {
	_jsii_.InvokeVoid(
		e,
		"resetEndpointPrivateAccess",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) ResetEndpointPublicAccess() {
	_jsii_.InvokeVoid(
		e,
		"resetEndpointPublicAccess",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) ResetPublicAccessCidrs() {
	_jsii_.InvokeVoid(
		e,
		"resetPublicAccessCidrs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (e *jsiiProxy_EksClusterVpcConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
