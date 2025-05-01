package googlegkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkehubfeaturemembership/internal"
)

type GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference interface {
	cdktf.ComplexObject
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
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	CpuLimit() *string
	SetCpuLimit(val *string)
	CpuLimitInput() *string
	CpuRequest() *string
	SetCpuRequest(val *string)
	CpuRequestInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MemoryLimit() *string
	SetMemoryLimit(val *string)
	MemoryLimitInput() *string
	MemoryRequest() *string
	SetMemoryRequest(val *string)
	MemoryRequestInput() *string
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
	ResetContainerName()
	ResetCpuLimit()
	ResetCpuRequest()
	ResetMemoryLimit()
	ResetMemoryRequest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference
type jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) CpuLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) CpuLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) CpuRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) CpuRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) MemoryLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) MemoryLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) MemoryRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) MemoryRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference_Override(g GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetCpuLimit(val *string) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetCpuRequest(val *string) {
	if err := j.validateSetCpuRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuRequest",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetMemoryLimit(val *string) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetMemoryRequest(val *string) {
	if err := j.validateSetMemoryRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryRequest",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ResetContainerName() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ResetCpuRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ResetMemoryRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetMemoryRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

