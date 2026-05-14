package vertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaireasoningengine/internal"
)

type VertexAiReasoningEngineSpecOutputReference interface {
	cdktf.ComplexObject
	AgentFramework() *string
	SetAgentFramework(val *string)
	AgentFrameworkInput() *string
	ClassMethods() *string
	SetClassMethods(val *string)
	ClassMethodsInput() *string
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
	ContainerSpec() VertexAiReasoningEngineSpecContainerSpecOutputReference
	ContainerSpecInput() *VertexAiReasoningEngineSpecContainerSpec
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentSpec() VertexAiReasoningEngineSpecDeploymentSpecOutputReference
	DeploymentSpecInput() *VertexAiReasoningEngineSpecDeploymentSpec
	EffectiveIdentity() *string
	// Experimental.
	Fqn() *string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	InternalValue() *VertexAiReasoningEngineSpec
	SetInternalValue(val *VertexAiReasoningEngineSpec)
	PackageSpec() VertexAiReasoningEngineSpecPackageSpecOutputReference
	PackageSpecInput() *VertexAiReasoningEngineSpecPackageSpec
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	SourceCodeSpec() VertexAiReasoningEngineSpecSourceCodeSpecOutputReference
	SourceCodeSpecInput() *VertexAiReasoningEngineSpecSourceCodeSpec
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
	PutContainerSpec(value *VertexAiReasoningEngineSpecContainerSpec)
	PutDeploymentSpec(value *VertexAiReasoningEngineSpecDeploymentSpec)
	PutPackageSpec(value *VertexAiReasoningEngineSpecPackageSpec)
	PutSourceCodeSpec(value *VertexAiReasoningEngineSpecSourceCodeSpec)
	ResetAgentFramework()
	ResetClassMethods()
	ResetContainerSpec()
	ResetDeploymentSpec()
	ResetIdentityType()
	ResetPackageSpec()
	ResetServiceAccount()
	ResetSourceCodeSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiReasoningEngineSpecOutputReference
type jsiiProxy_VertexAiReasoningEngineSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) AgentFramework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentFramework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) AgentFrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentFrameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ClassMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ClassMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ContainerSpec() VertexAiReasoningEngineSpecContainerSpecOutputReference {
	var returns VertexAiReasoningEngineSpecContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ContainerSpecInput() *VertexAiReasoningEngineSpecContainerSpec {
	var returns *VertexAiReasoningEngineSpecContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) DeploymentSpec() VertexAiReasoningEngineSpecDeploymentSpecOutputReference {
	var returns VertexAiReasoningEngineSpecDeploymentSpecOutputReference
	_jsii_.Get(
		j,
		"deploymentSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) DeploymentSpecInput() *VertexAiReasoningEngineSpecDeploymentSpec {
	var returns *VertexAiReasoningEngineSpecDeploymentSpec
	_jsii_.Get(
		j,
		"deploymentSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) EffectiveIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) InternalValue() *VertexAiReasoningEngineSpec {
	var returns *VertexAiReasoningEngineSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) PackageSpec() VertexAiReasoningEngineSpecPackageSpecOutputReference {
	var returns VertexAiReasoningEngineSpecPackageSpecOutputReference
	_jsii_.Get(
		j,
		"packageSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) PackageSpecInput() *VertexAiReasoningEngineSpecPackageSpec {
	var returns *VertexAiReasoningEngineSpecPackageSpec
	_jsii_.Get(
		j,
		"packageSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) SourceCodeSpec() VertexAiReasoningEngineSpecSourceCodeSpecOutputReference {
	var returns VertexAiReasoningEngineSpecSourceCodeSpecOutputReference
	_jsii_.Get(
		j,
		"sourceCodeSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) SourceCodeSpecInput() *VertexAiReasoningEngineSpecSourceCodeSpec {
	var returns *VertexAiReasoningEngineSpecSourceCodeSpec
	_jsii_.Get(
		j,
		"sourceCodeSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiReasoningEngineSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiReasoningEngineSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiReasoningEngineSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiReasoningEngineSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiReasoningEngineSpecOutputReference_Override(v VertexAiReasoningEngineSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetAgentFramework(val *string) {
	if err := j.validateSetAgentFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentFramework",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetClassMethods(val *string) {
	if err := j.validateSetClassMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classMethods",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetInternalValue(val *VertexAiReasoningEngineSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) PutContainerSpec(value *VertexAiReasoningEngineSpecContainerSpec) {
	if err := v.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putContainerSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) PutDeploymentSpec(value *VertexAiReasoningEngineSpecDeploymentSpec) {
	if err := v.validatePutDeploymentSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDeploymentSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) PutPackageSpec(value *VertexAiReasoningEngineSpecPackageSpec) {
	if err := v.validatePutPackageSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPackageSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) PutSourceCodeSpec(value *VertexAiReasoningEngineSpecSourceCodeSpec) {
	if err := v.validatePutSourceCodeSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSourceCodeSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetAgentFramework() {
	_jsii_.InvokeVoid(
		v,
		"resetAgentFramework",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetClassMethods() {
	_jsii_.InvokeVoid(
		v,
		"resetClassMethods",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetDeploymentSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetDeploymentSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		v,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetPackageSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetPackageSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ResetSourceCodeSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceCodeSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

