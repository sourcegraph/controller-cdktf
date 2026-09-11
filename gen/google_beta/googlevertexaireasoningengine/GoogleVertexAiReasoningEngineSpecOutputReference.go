package googlevertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaireasoningengine/internal"
)

type GoogleVertexAiReasoningEngineSpecOutputReference interface {
	cdktn.ComplexObject
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
	ContainerSpec() GoogleVertexAiReasoningEngineSpecContainerSpecOutputReference
	ContainerSpecInput() *GoogleVertexAiReasoningEngineSpecContainerSpec
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentSpec() GoogleVertexAiReasoningEngineSpecDeploymentSpecOutputReference
	DeploymentSpecInput() *GoogleVertexAiReasoningEngineSpecDeploymentSpec
	EffectiveIdentity() *string
	// Experimental.
	Fqn() *string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	InternalValue() *GoogleVertexAiReasoningEngineSpec
	SetInternalValue(val *GoogleVertexAiReasoningEngineSpec)
	PackageSpec() GoogleVertexAiReasoningEngineSpecPackageSpecOutputReference
	PackageSpecInput() *GoogleVertexAiReasoningEngineSpecPackageSpec
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	SourceCodeSpec() GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference
	SourceCodeSpecInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpec
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutContainerSpec(value *GoogleVertexAiReasoningEngineSpecContainerSpec)
	PutDeploymentSpec(value *GoogleVertexAiReasoningEngineSpecDeploymentSpec)
	PutPackageSpec(value *GoogleVertexAiReasoningEngineSpecPackageSpec)
	PutSourceCodeSpec(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpec)
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiReasoningEngineSpecOutputReference
type jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) AgentFramework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentFramework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) AgentFrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentFrameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ClassMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ClassMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ContainerSpec() GoogleVertexAiReasoningEngineSpecContainerSpecOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ContainerSpecInput() *GoogleVertexAiReasoningEngineSpecContainerSpec {
	var returns *GoogleVertexAiReasoningEngineSpecContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) DeploymentSpec() GoogleVertexAiReasoningEngineSpecDeploymentSpecOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecDeploymentSpecOutputReference
	_jsii_.Get(
		j,
		"deploymentSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) DeploymentSpecInput() *GoogleVertexAiReasoningEngineSpecDeploymentSpec {
	var returns *GoogleVertexAiReasoningEngineSpecDeploymentSpec
	_jsii_.Get(
		j,
		"deploymentSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) EffectiveIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) InternalValue() *GoogleVertexAiReasoningEngineSpec {
	var returns *GoogleVertexAiReasoningEngineSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) PackageSpec() GoogleVertexAiReasoningEngineSpecPackageSpecOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecPackageSpecOutputReference
	_jsii_.Get(
		j,
		"packageSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) PackageSpecInput() *GoogleVertexAiReasoningEngineSpecPackageSpec {
	var returns *GoogleVertexAiReasoningEngineSpecPackageSpec
	_jsii_.Get(
		j,
		"packageSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) SourceCodeSpec() GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference
	_jsii_.Get(
		j,
		"sourceCodeSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) SourceCodeSpecInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpec {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpec
	_jsii_.Get(
		j,
		"sourceCodeSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiReasoningEngineSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiReasoningEngineSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiReasoningEngineSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiReasoningEngineSpecOutputReference_Override(g GoogleVertexAiReasoningEngineSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetAgentFramework(val *string) {
	if err := j.validateSetAgentFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentFramework",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetClassMethods(val *string) {
	if err := j.validateSetClassMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classMethods",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetInternalValue(val *GoogleVertexAiReasoningEngineSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) PutContainerSpec(value *GoogleVertexAiReasoningEngineSpecContainerSpec) {
	if err := g.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) PutDeploymentSpec(value *GoogleVertexAiReasoningEngineSpecDeploymentSpec) {
	if err := g.validatePutDeploymentSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) PutPackageSpec(value *GoogleVertexAiReasoningEngineSpecPackageSpec) {
	if err := g.validatePutPackageSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPackageSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) PutSourceCodeSpec(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpec) {
	if err := g.validatePutSourceCodeSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceCodeSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetAgentFramework() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentFramework",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetClassMethods() {
	_jsii_.InvokeVoid(
		g,
		"resetClassMethods",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetDeploymentSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetPackageSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPackageSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ResetSourceCodeSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceCodeSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

