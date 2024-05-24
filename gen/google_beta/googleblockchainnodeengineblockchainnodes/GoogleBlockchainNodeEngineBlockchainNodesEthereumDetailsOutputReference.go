package googleblockchainnodeengineblockchainnodes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleblockchainnodeengineblockchainnodes/internal"
)

type GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference interface {
	cdktf.ComplexObject
	AdditionalEndpoints() GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsList
	ApiEnableAdmin() interface{}
	SetApiEnableAdmin(val interface{})
	ApiEnableAdminInput() interface{}
	ApiEnableDebug() interface{}
	SetApiEnableDebug(val interface{})
	ApiEnableDebugInput() interface{}
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
	ConsensusClient() *string
	SetConsensusClient(val *string)
	ConsensusClientInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExecutionClient() *string
	SetExecutionClient(val *string)
	ExecutionClientInput() *string
	FetchhDetails() GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetailsOutputReference
	FetchhDetailsInput() *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetails
	SetInternalValue(val *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetails)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidatorConfig() GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfigOutputReference
	ValidatorConfigInput() *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig
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
	PutFetchhDetails(value *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails)
	PutValidatorConfig(value *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig)
	ResetApiEnableAdmin()
	ResetApiEnableDebug()
	ResetConsensusClient()
	ResetExecutionClient()
	ResetFetchhDetails()
	ResetNetwork()
	ResetNodeType()
	ResetValidatorConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference
type jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) AdditionalEndpoints() GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsList {
	var returns GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsList
	_jsii_.Get(
		j,
		"additionalEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableDebug() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableDebug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableDebugInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableDebugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ConsensusClient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consensusClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ConsensusClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consensusClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ExecutionClient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ExecutionClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) FetchhDetails() GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetailsOutputReference {
	var returns GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetailsOutputReference
	_jsii_.Get(
		j,
		"fetchhDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) FetchhDetailsInput() *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails {
	var returns *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails
	_jsii_.Get(
		j,
		"fetchhDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) InternalValue() *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetails {
	var returns *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ValidatorConfig() GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfigOutputReference {
	var returns GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfigOutputReference
	_jsii_.Get(
		j,
		"validatorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ValidatorConfigInput() *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig {
	var returns *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig
	_jsii_.Get(
		j,
		"validatorConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBlockchainNodeEngineBlockchainNodes.GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference_Override(g GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBlockchainNodeEngineBlockchainNodes.GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetApiEnableAdmin(val interface{}) {
	if err := j.validateSetApiEnableAdminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiEnableAdmin",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetApiEnableDebug(val interface{}) {
	if err := j.validateSetApiEnableDebugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiEnableDebug",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetConsensusClient(val *string) {
	if err := j.validateSetConsensusClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consensusClient",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetExecutionClient(val *string) {
	if err := j.validateSetExecutionClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionClient",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetInternalValue(val *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) PutFetchhDetails(value *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails) {
	if err := g.validatePutFetchhDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFetchhDetails",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) PutValidatorConfig(value *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig) {
	if err := g.validatePutValidatorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putValidatorConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetApiEnableAdmin() {
	_jsii_.InvokeVoid(
		g,
		"resetApiEnableAdmin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetApiEnableDebug() {
	_jsii_.InvokeVoid(
		g,
		"resetApiEnableDebug",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetConsensusClient() {
	_jsii_.InvokeVoid(
		g,
		"resetConsensusClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetExecutionClient() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetFetchhDetails() {
	_jsii_.InvokeVoid(
		g,
		"resetFetchhDetails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetNodeType() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetValidatorConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetValidatorConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

