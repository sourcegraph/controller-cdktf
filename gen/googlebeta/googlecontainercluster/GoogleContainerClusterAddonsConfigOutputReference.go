package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecontainercluster/internal"
)

type GoogleContainerClusterAddonsConfigOutputReference interface {
	cdktf.ComplexObject
	CloudrunConfig() GoogleContainerClusterAddonsConfigCloudrunConfigOutputReference
	CloudrunConfigInput() *GoogleContainerClusterAddonsConfigCloudrunConfig
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
	ConfigConnectorConfig() GoogleContainerClusterAddonsConfigConfigConnectorConfigOutputReference
	ConfigConnectorConfigInput() *GoogleContainerClusterAddonsConfigConfigConnectorConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsCacheConfig() GoogleContainerClusterAddonsConfigDnsCacheConfigOutputReference
	DnsCacheConfigInput() *GoogleContainerClusterAddonsConfigDnsCacheConfig
	// Experimental.
	Fqn() *string
	GcePersistentDiskCsiDriverConfig() GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigOutputReference
	GcePersistentDiskCsiDriverConfigInput() *GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig
	GcpFilestoreCsiDriverConfig() GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfigOutputReference
	GcpFilestoreCsiDriverConfigInput() *GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig
	GkeBackupAgentConfig() GoogleContainerClusterAddonsConfigGkeBackupAgentConfigOutputReference
	GkeBackupAgentConfigInput() *GoogleContainerClusterAddonsConfigGkeBackupAgentConfig
	HorizontalPodAutoscaling() GoogleContainerClusterAddonsConfigHorizontalPodAutoscalingOutputReference
	HorizontalPodAutoscalingInput() *GoogleContainerClusterAddonsConfigHorizontalPodAutoscaling
	HttpLoadBalancing() GoogleContainerClusterAddonsConfigHttpLoadBalancingOutputReference
	HttpLoadBalancingInput() *GoogleContainerClusterAddonsConfigHttpLoadBalancing
	InternalValue() *GoogleContainerClusterAddonsConfig
	SetInternalValue(val *GoogleContainerClusterAddonsConfig)
	IstioConfig() GoogleContainerClusterAddonsConfigIstioConfigOutputReference
	IstioConfigInput() *GoogleContainerClusterAddonsConfigIstioConfig
	KalmConfig() GoogleContainerClusterAddonsConfigKalmConfigOutputReference
	KalmConfigInput() *GoogleContainerClusterAddonsConfigKalmConfig
	NetworkPolicyConfig() GoogleContainerClusterAddonsConfigNetworkPolicyConfigOutputReference
	NetworkPolicyConfigInput() *GoogleContainerClusterAddonsConfigNetworkPolicyConfig
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
	PutCloudrunConfig(value *GoogleContainerClusterAddonsConfigCloudrunConfig)
	PutConfigConnectorConfig(value *GoogleContainerClusterAddonsConfigConfigConnectorConfig)
	PutDnsCacheConfig(value *GoogleContainerClusterAddonsConfigDnsCacheConfig)
	PutGcePersistentDiskCsiDriverConfig(value *GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
	PutGcpFilestoreCsiDriverConfig(value *GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig)
	PutGkeBackupAgentConfig(value *GoogleContainerClusterAddonsConfigGkeBackupAgentConfig)
	PutHorizontalPodAutoscaling(value *GoogleContainerClusterAddonsConfigHorizontalPodAutoscaling)
	PutHttpLoadBalancing(value *GoogleContainerClusterAddonsConfigHttpLoadBalancing)
	PutIstioConfig(value *GoogleContainerClusterAddonsConfigIstioConfig)
	PutKalmConfig(value *GoogleContainerClusterAddonsConfigKalmConfig)
	PutNetworkPolicyConfig(value *GoogleContainerClusterAddonsConfigNetworkPolicyConfig)
	ResetCloudrunConfig()
	ResetConfigConnectorConfig()
	ResetDnsCacheConfig()
	ResetGcePersistentDiskCsiDriverConfig()
	ResetGcpFilestoreCsiDriverConfig()
	ResetGkeBackupAgentConfig()
	ResetHorizontalPodAutoscaling()
	ResetHttpLoadBalancing()
	ResetIstioConfig()
	ResetKalmConfig()
	ResetNetworkPolicyConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterAddonsConfigOutputReference
type jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) CloudrunConfig() GoogleContainerClusterAddonsConfigCloudrunConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigCloudrunConfigOutputReference
	_jsii_.Get(
		j,
		"cloudrunConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) CloudrunConfigInput() *GoogleContainerClusterAddonsConfigCloudrunConfig {
	var returns *GoogleContainerClusterAddonsConfigCloudrunConfig
	_jsii_.Get(
		j,
		"cloudrunConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ConfigConnectorConfig() GoogleContainerClusterAddonsConfigConfigConnectorConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigConfigConnectorConfigOutputReference
	_jsii_.Get(
		j,
		"configConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ConfigConnectorConfigInput() *GoogleContainerClusterAddonsConfigConfigConnectorConfig {
	var returns *GoogleContainerClusterAddonsConfigConfigConnectorConfig
	_jsii_.Get(
		j,
		"configConnectorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) DnsCacheConfig() GoogleContainerClusterAddonsConfigDnsCacheConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigDnsCacheConfigOutputReference
	_jsii_.Get(
		j,
		"dnsCacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) DnsCacheConfigInput() *GoogleContainerClusterAddonsConfigDnsCacheConfig {
	var returns *GoogleContainerClusterAddonsConfigDnsCacheConfig
	_jsii_.Get(
		j,
		"dnsCacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GcePersistentDiskCsiDriverConfig() GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDiskCsiDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GcePersistentDiskCsiDriverConfigInput() *GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	var returns *GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig
	_jsii_.Get(
		j,
		"gcePersistentDiskCsiDriverConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GcpFilestoreCsiDriverConfig() GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfigOutputReference
	_jsii_.Get(
		j,
		"gcpFilestoreCsiDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GcpFilestoreCsiDriverConfigInput() *GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig {
	var returns *GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig
	_jsii_.Get(
		j,
		"gcpFilestoreCsiDriverConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GkeBackupAgentConfig() GoogleContainerClusterAddonsConfigGkeBackupAgentConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigGkeBackupAgentConfigOutputReference
	_jsii_.Get(
		j,
		"gkeBackupAgentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GkeBackupAgentConfigInput() *GoogleContainerClusterAddonsConfigGkeBackupAgentConfig {
	var returns *GoogleContainerClusterAddonsConfigGkeBackupAgentConfig
	_jsii_.Get(
		j,
		"gkeBackupAgentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) HorizontalPodAutoscaling() GoogleContainerClusterAddonsConfigHorizontalPodAutoscalingOutputReference {
	var returns GoogleContainerClusterAddonsConfigHorizontalPodAutoscalingOutputReference
	_jsii_.Get(
		j,
		"horizontalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) HorizontalPodAutoscalingInput() *GoogleContainerClusterAddonsConfigHorizontalPodAutoscaling {
	var returns *GoogleContainerClusterAddonsConfigHorizontalPodAutoscaling
	_jsii_.Get(
		j,
		"horizontalPodAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) HttpLoadBalancing() GoogleContainerClusterAddonsConfigHttpLoadBalancingOutputReference {
	var returns GoogleContainerClusterAddonsConfigHttpLoadBalancingOutputReference
	_jsii_.Get(
		j,
		"httpLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) HttpLoadBalancingInput() *GoogleContainerClusterAddonsConfigHttpLoadBalancing {
	var returns *GoogleContainerClusterAddonsConfigHttpLoadBalancing
	_jsii_.Get(
		j,
		"httpLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) InternalValue() *GoogleContainerClusterAddonsConfig {
	var returns *GoogleContainerClusterAddonsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) IstioConfig() GoogleContainerClusterAddonsConfigIstioConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigIstioConfigOutputReference
	_jsii_.Get(
		j,
		"istioConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) IstioConfigInput() *GoogleContainerClusterAddonsConfigIstioConfig {
	var returns *GoogleContainerClusterAddonsConfigIstioConfig
	_jsii_.Get(
		j,
		"istioConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) KalmConfig() GoogleContainerClusterAddonsConfigKalmConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigKalmConfigOutputReference
	_jsii_.Get(
		j,
		"kalmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) KalmConfigInput() *GoogleContainerClusterAddonsConfigKalmConfig {
	var returns *GoogleContainerClusterAddonsConfigKalmConfig
	_jsii_.Get(
		j,
		"kalmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) NetworkPolicyConfig() GoogleContainerClusterAddonsConfigNetworkPolicyConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigNetworkPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"networkPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) NetworkPolicyConfigInput() *GoogleContainerClusterAddonsConfigNetworkPolicyConfig {
	var returns *GoogleContainerClusterAddonsConfigNetworkPolicyConfig
	_jsii_.Get(
		j,
		"networkPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterAddonsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterAddonsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterAddonsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerCluster.GoogleContainerClusterAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterAddonsConfigOutputReference_Override(g GoogleContainerClusterAddonsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerCluster.GoogleContainerClusterAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference)SetInternalValue(val *GoogleContainerClusterAddonsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutCloudrunConfig(value *GoogleContainerClusterAddonsConfigCloudrunConfig) {
	if err := g.validatePutCloudrunConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudrunConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutConfigConnectorConfig(value *GoogleContainerClusterAddonsConfigConfigConnectorConfig) {
	if err := g.validatePutConfigConnectorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfigConnectorConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutDnsCacheConfig(value *GoogleContainerClusterAddonsConfigDnsCacheConfig) {
	if err := g.validatePutDnsCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsCacheConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutGcePersistentDiskCsiDriverConfig(value *GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig) {
	if err := g.validatePutGcePersistentDiskCsiDriverConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcePersistentDiskCsiDriverConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutGcpFilestoreCsiDriverConfig(value *GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig) {
	if err := g.validatePutGcpFilestoreCsiDriverConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcpFilestoreCsiDriverConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutGkeBackupAgentConfig(value *GoogleContainerClusterAddonsConfigGkeBackupAgentConfig) {
	if err := g.validatePutGkeBackupAgentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGkeBackupAgentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutHorizontalPodAutoscaling(value *GoogleContainerClusterAddonsConfigHorizontalPodAutoscaling) {
	if err := g.validatePutHorizontalPodAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHorizontalPodAutoscaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutHttpLoadBalancing(value *GoogleContainerClusterAddonsConfigHttpLoadBalancing) {
	if err := g.validatePutHttpLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpLoadBalancing",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutIstioConfig(value *GoogleContainerClusterAddonsConfigIstioConfig) {
	if err := g.validatePutIstioConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIstioConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutKalmConfig(value *GoogleContainerClusterAddonsConfigKalmConfig) {
	if err := g.validatePutKalmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKalmConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) PutNetworkPolicyConfig(value *GoogleContainerClusterAddonsConfigNetworkPolicyConfig) {
	if err := g.validatePutNetworkPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPolicyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetCloudrunConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudrunConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetConfigConnectorConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigConnectorConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetDnsCacheConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsCacheConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetGcePersistentDiskCsiDriverConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcePersistentDiskCsiDriverConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetGcpFilestoreCsiDriverConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcpFilestoreCsiDriverConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetGkeBackupAgentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeBackupAgentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetHorizontalPodAutoscaling() {
	_jsii_.InvokeVoid(
		g,
		"resetHorizontalPodAutoscaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetHttpLoadBalancing() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpLoadBalancing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetIstioConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIstioConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetKalmConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKalmConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ResetNetworkPolicyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPolicyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterAddonsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

