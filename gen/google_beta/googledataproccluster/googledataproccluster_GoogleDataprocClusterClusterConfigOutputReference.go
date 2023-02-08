package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataproccluster/internal"
)

type GoogleDataprocClusterClusterConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingConfig() GoogleDataprocClusterClusterConfigAutoscalingConfigOutputReference
	AutoscalingConfigInput() *GoogleDataprocClusterClusterConfigAutoscalingConfig
	Bucket() *string
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
	EncryptionConfig() GoogleDataprocClusterClusterConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleDataprocClusterClusterConfigEncryptionConfig
	EndpointConfig() GoogleDataprocClusterClusterConfigEndpointConfigOutputReference
	EndpointConfigInput() *GoogleDataprocClusterClusterConfigEndpointConfig
	// Experimental.
	Fqn() *string
	GceClusterConfig() GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference
	GceClusterConfigInput() *GoogleDataprocClusterClusterConfigGceClusterConfig
	InitializationAction() GoogleDataprocClusterClusterConfigInitializationActionList
	InitializationActionInput() interface{}
	InternalValue() *GoogleDataprocClusterClusterConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfig)
	LifecycleConfig() GoogleDataprocClusterClusterConfigLifecycleConfigOutputReference
	LifecycleConfigInput() *GoogleDataprocClusterClusterConfigLifecycleConfig
	MasterConfig() GoogleDataprocClusterClusterConfigMasterConfigOutputReference
	MasterConfigInput() *GoogleDataprocClusterClusterConfigMasterConfig
	MetastoreConfig() GoogleDataprocClusterClusterConfigMetastoreConfigOutputReference
	MetastoreConfigInput() *GoogleDataprocClusterClusterConfigMetastoreConfig
	PreemptibleWorkerConfig() GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigOutputReference
	PreemptibleWorkerConfigInput() *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig
	SecurityConfig() GoogleDataprocClusterClusterConfigSecurityConfigOutputReference
	SecurityConfigInput() *GoogleDataprocClusterClusterConfigSecurityConfig
	SoftwareConfig() GoogleDataprocClusterClusterConfigSoftwareConfigOutputReference
	SoftwareConfigInput() *GoogleDataprocClusterClusterConfigSoftwareConfig
	StagingBucket() *string
	SetStagingBucket(val *string)
	StagingBucketInput() *string
	TempBucket() *string
	SetTempBucket(val *string)
	TempBucketInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerConfig() GoogleDataprocClusterClusterConfigWorkerConfigOutputReference
	WorkerConfigInput() *GoogleDataprocClusterClusterConfigWorkerConfig
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
	PutAutoscalingConfig(value *GoogleDataprocClusterClusterConfigAutoscalingConfig)
	PutEncryptionConfig(value *GoogleDataprocClusterClusterConfigEncryptionConfig)
	PutEndpointConfig(value *GoogleDataprocClusterClusterConfigEndpointConfig)
	PutGceClusterConfig(value *GoogleDataprocClusterClusterConfigGceClusterConfig)
	PutInitializationAction(value interface{})
	PutLifecycleConfig(value *GoogleDataprocClusterClusterConfigLifecycleConfig)
	PutMasterConfig(value *GoogleDataprocClusterClusterConfigMasterConfig)
	PutMetastoreConfig(value *GoogleDataprocClusterClusterConfigMetastoreConfig)
	PutPreemptibleWorkerConfig(value *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig)
	PutSecurityConfig(value *GoogleDataprocClusterClusterConfigSecurityConfig)
	PutSoftwareConfig(value *GoogleDataprocClusterClusterConfigSoftwareConfig)
	PutWorkerConfig(value *GoogleDataprocClusterClusterConfigWorkerConfig)
	ResetAutoscalingConfig()
	ResetEncryptionConfig()
	ResetEndpointConfig()
	ResetGceClusterConfig()
	ResetInitializationAction()
	ResetLifecycleConfig()
	ResetMasterConfig()
	ResetMetastoreConfig()
	ResetPreemptibleWorkerConfig()
	ResetSecurityConfig()
	ResetSoftwareConfig()
	ResetStagingBucket()
	ResetTempBucket()
	ResetWorkerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) AutoscalingConfig() GoogleDataprocClusterClusterConfigAutoscalingConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigAutoscalingConfigOutputReference
	_jsii_.Get(
		j,
		"autoscalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) AutoscalingConfigInput() *GoogleDataprocClusterClusterConfigAutoscalingConfig {
	var returns *GoogleDataprocClusterClusterConfigAutoscalingConfig
	_jsii_.Get(
		j,
		"autoscalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) EncryptionConfig() GoogleDataprocClusterClusterConfigEncryptionConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) EncryptionConfigInput() *GoogleDataprocClusterClusterConfigEncryptionConfig {
	var returns *GoogleDataprocClusterClusterConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) EndpointConfig() GoogleDataprocClusterClusterConfigEndpointConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"endpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) EndpointConfigInput() *GoogleDataprocClusterClusterConfigEndpointConfig {
	var returns *GoogleDataprocClusterClusterConfigEndpointConfig
	_jsii_.Get(
		j,
		"endpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GceClusterConfig() GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gceClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GceClusterConfigInput() *GoogleDataprocClusterClusterConfigGceClusterConfig {
	var returns *GoogleDataprocClusterClusterConfigGceClusterConfig
	_jsii_.Get(
		j,
		"gceClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) InitializationAction() GoogleDataprocClusterClusterConfigInitializationActionList {
	var returns GoogleDataprocClusterClusterConfigInitializationActionList
	_jsii_.Get(
		j,
		"initializationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) InitializationActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfig {
	var returns *GoogleDataprocClusterClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) LifecycleConfig() GoogleDataprocClusterClusterConfigLifecycleConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigLifecycleConfigOutputReference
	_jsii_.Get(
		j,
		"lifecycleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) LifecycleConfigInput() *GoogleDataprocClusterClusterConfigLifecycleConfig {
	var returns *GoogleDataprocClusterClusterConfigLifecycleConfig
	_jsii_.Get(
		j,
		"lifecycleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) MasterConfig() GoogleDataprocClusterClusterConfigMasterConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigMasterConfigOutputReference
	_jsii_.Get(
		j,
		"masterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) MasterConfigInput() *GoogleDataprocClusterClusterConfigMasterConfig {
	var returns *GoogleDataprocClusterClusterConfigMasterConfig
	_jsii_.Get(
		j,
		"masterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) MetastoreConfig() GoogleDataprocClusterClusterConfigMetastoreConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigMetastoreConfigOutputReference
	_jsii_.Get(
		j,
		"metastoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) MetastoreConfigInput() *GoogleDataprocClusterClusterConfigMetastoreConfig {
	var returns *GoogleDataprocClusterClusterConfigMetastoreConfig
	_jsii_.Get(
		j,
		"metastoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PreemptibleWorkerConfig() GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"preemptibleWorkerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PreemptibleWorkerConfigInput() *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig {
	var returns *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig
	_jsii_.Get(
		j,
		"preemptibleWorkerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) SecurityConfig() GoogleDataprocClusterClusterConfigSecurityConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) SecurityConfigInput() *GoogleDataprocClusterClusterConfigSecurityConfig {
	var returns *GoogleDataprocClusterClusterConfigSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) SoftwareConfig() GoogleDataprocClusterClusterConfigSoftwareConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) SoftwareConfigInput() *GoogleDataprocClusterClusterConfigSoftwareConfig {
	var returns *GoogleDataprocClusterClusterConfigSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) TempBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) TempBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) WorkerConfig() GoogleDataprocClusterClusterConfigWorkerConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"workerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) WorkerConfigInput() *GoogleDataprocClusterClusterConfigWorkerConfig {
	var returns *GoogleDataprocClusterClusterConfigWorkerConfig
	_jsii_.Get(
		j,
		"workerConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetTempBucket(val *string) {
	if err := j.validateSetTempBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutAutoscalingConfig(value *GoogleDataprocClusterClusterConfigAutoscalingConfig) {
	if err := g.validatePutAutoscalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutEncryptionConfig(value *GoogleDataprocClusterClusterConfigEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutEndpointConfig(value *GoogleDataprocClusterClusterConfigEndpointConfig) {
	if err := g.validatePutEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpointConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutGceClusterConfig(value *GoogleDataprocClusterClusterConfigGceClusterConfig) {
	if err := g.validatePutGceClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGceClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutInitializationAction(value interface{}) {
	if err := g.validatePutInitializationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInitializationAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutLifecycleConfig(value *GoogleDataprocClusterClusterConfigLifecycleConfig) {
	if err := g.validatePutLifecycleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLifecycleConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutMasterConfig(value *GoogleDataprocClusterClusterConfigMasterConfig) {
	if err := g.validatePutMasterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutMetastoreConfig(value *GoogleDataprocClusterClusterConfigMetastoreConfig) {
	if err := g.validatePutMetastoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetastoreConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutPreemptibleWorkerConfig(value *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfig) {
	if err := g.validatePutPreemptibleWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPreemptibleWorkerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutSecurityConfig(value *GoogleDataprocClusterClusterConfigSecurityConfig) {
	if err := g.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutSoftwareConfig(value *GoogleDataprocClusterClusterConfigSoftwareConfig) {
	if err := g.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) PutWorkerConfig(value *GoogleDataprocClusterClusterConfigWorkerConfig) {
	if err := g.validatePutWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetAutoscalingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetEndpointConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetGceClusterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGceClusterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetInitializationAction() {
	_jsii_.InvokeVoid(
		g,
		"resetInitializationAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetLifecycleConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLifecycleConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetMasterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetMetastoreConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMetastoreConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetPreemptibleWorkerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptibleWorkerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetTempBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetTempBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ResetWorkerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

