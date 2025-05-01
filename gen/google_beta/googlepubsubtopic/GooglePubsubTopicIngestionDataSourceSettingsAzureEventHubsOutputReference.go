package googlepubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubtopic/internal"
)

type GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
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
	EventHub() *string
	SetEventHub(val *string)
	EventHubInput() *string
	// Experimental.
	Fqn() *string
	GcpServiceAccount() *string
	SetGcpServiceAccount(val *string)
	GcpServiceAccountInput() *string
	InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs
	SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	ResourceGroup() *string
	SetResourceGroup(val *string)
	ResourceGroupInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetClientId()
	ResetEventHub()
	ResetGcpServiceAccount()
	ResetNamespace()
	ResetResourceGroup()
	ResetSubscriptionId()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference
type jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) EventHub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) EventHubInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GcpServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GcpServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference_Override(g GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetEventHub(val *string) {
	if err := j.validateSetEventHubParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHub",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetGcpServiceAccount(val *string) {
	if err := j.validateSetGcpServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetResourceGroup(val *string) {
	if err := j.validateSetResourceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroup",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		g,
		"resetClientId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetEventHub() {
	_jsii_.InvokeVoid(
		g,
		"resetEventHub",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetGcpServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetGcpServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		g,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		g,
		"resetTenantId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

