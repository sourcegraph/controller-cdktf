package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference
	AutoscalingConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig
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
	EncryptionConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig
	EndpointConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference
	EndpointConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig
	// Experimental.
	Fqn() *string
	GceClusterConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference
	GceClusterConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig
	GkeClusterConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigOutputReference
	GkeClusterConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig
	InitializationActions() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList
	InitializationActionsInput() interface{}
	InternalValue() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig
	SetInternalValue(val *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig)
	LifecycleConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference
	LifecycleConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig
	MasterConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference
	MasterConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig
	MetastoreConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigOutputReference
	MetastoreConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig
	SecondaryWorkerConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference
	SecondaryWorkerConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
	SecurityConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference
	SecurityConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig
	SoftwareConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference
	SoftwareConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig
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
	WorkerConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference
	WorkerConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig
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
	PutAutoscalingConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig)
	PutEncryptionConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig)
	PutEndpointConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig)
	PutGceClusterConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig)
	PutGkeClusterConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig)
	PutInitializationActions(value interface{})
	PutLifecycleConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig)
	PutMasterConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig)
	PutMetastoreConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig)
	PutSecondaryWorkerConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig)
	PutSecurityConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig)
	PutSoftwareConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig)
	PutWorkerConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig)
	ResetAutoscalingConfig()
	ResetEncryptionConfig()
	ResetEndpointConfig()
	ResetGceClusterConfig()
	ResetGkeClusterConfig()
	ResetInitializationActions()
	ResetLifecycleConfig()
	ResetMasterConfig()
	ResetMetastoreConfig()
	ResetSecondaryWorkerConfig()
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

// The jsii proxy struct for GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference
type jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) AutoscalingConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference
	_jsii_.Get(
		j,
		"autoscalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) AutoscalingConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig
	_jsii_.Get(
		j,
		"autoscalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EncryptionConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EncryptionConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EndpointConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"endpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EndpointConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig
	_jsii_.Get(
		j,
		"endpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GceClusterConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gceClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GceClusterConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig
	_jsii_.Get(
		j,
		"gceClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GkeClusterConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gkeClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GkeClusterConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig
	_jsii_.Get(
		j,
		"gkeClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InitializationActions() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList
	_jsii_.Get(
		j,
		"initializationActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InitializationActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializationActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InternalValue() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) LifecycleConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference
	_jsii_.Get(
		j,
		"lifecycleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) LifecycleConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig
	_jsii_.Get(
		j,
		"lifecycleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) MasterConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference
	_jsii_.Get(
		j,
		"masterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) MasterConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig
	_jsii_.Get(
		j,
		"masterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) MetastoreConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigOutputReference
	_jsii_.Get(
		j,
		"metastoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) MetastoreConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig
	_jsii_.Get(
		j,
		"metastoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecondaryWorkerConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"secondaryWorkerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecondaryWorkerConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
	_jsii_.Get(
		j,
		"secondaryWorkerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecurityConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecurityConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SoftwareConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SoftwareConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TempBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TempBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) WorkerConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"workerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) WorkerConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig
	_jsii_.Get(
		j,
		"workerConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference_Override(g GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetInternalValue(val *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetTempBucket(val *string) {
	if err := j.validateSetTempBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutAutoscalingConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) {
	if err := g.validatePutAutoscalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutEncryptionConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutEndpointConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig) {
	if err := g.validatePutEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpointConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutGceClusterConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) {
	if err := g.validatePutGceClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGceClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutGkeClusterConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig) {
	if err := g.validatePutGkeClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGkeClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutInitializationActions(value interface{}) {
	if err := g.validatePutInitializationActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInitializationActions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutLifecycleConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) {
	if err := g.validatePutLifecycleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLifecycleConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutMasterConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig) {
	if err := g.validatePutMasterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutMetastoreConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig) {
	if err := g.validatePutMetastoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetastoreConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutSecondaryWorkerConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) {
	if err := g.validatePutSecondaryWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryWorkerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutSecurityConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig) {
	if err := g.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutSoftwareConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) {
	if err := g.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutWorkerConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig) {
	if err := g.validatePutWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetAutoscalingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetEndpointConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetGceClusterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGceClusterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetGkeClusterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeClusterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetInitializationActions() {
	_jsii_.InvokeVoid(
		g,
		"resetInitializationActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetLifecycleConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLifecycleConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetMasterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetMetastoreConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMetastoreConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetSecondaryWorkerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryWorkerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetTempBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetTempBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetWorkerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

