// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentInitParameters struct {

	// Name of the agent.
	AgentName *string `json:"agentName,omitempty" tf:"agent_name,omitempty"`

	// ARN of the IAM role with permissions to invoke API operations on the agent.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	AgentResourceRoleArn *string `json:"agentResourceRoleArn,omitempty" tf:"agent_resource_role_arn,omitempty"`

	// Reference to a Role in iam to populate agentResourceRoleArn.
	// +kubebuilder:validation:Optional
	AgentResourceRoleArnRef *v1.Reference `json:"agentResourceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate agentResourceRoleArn.
	// +kubebuilder:validation:Optional
	AgentResourceRoleArnSelector *v1.Selector `json:"agentResourceRoleArnSelector,omitempty" tf:"-"`

	// ARN of the AWS KMS key that encrypts the agent.
	CustomerEncryptionKeyArn *string `json:"customerEncryptionKeyArn,omitempty" tf:"customer_encryption_key_arn,omitempty"`

	// Description of the agent.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Foundation model used for orchestration by the agent.
	FoundationModel *string `json:"foundationModel,omitempty" tf:"foundation_model,omitempty"`

	// Number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent. A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Bedrock deletes any data provided before the timeout.
	IdleSessionTTLInSeconds *float64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds,omitempty"`

	// Instructions that tell the agent what it should do and how it should interact with users.
	Instruction *string `json:"instruction,omitempty" tf:"instruction,omitempty"`

	// Whether to prepare the agent after creation or modification. Defaults to true.
	PrepareAgent *bool `json:"prepareAgent,omitempty" tf:"prepare_agent,omitempty"`

	// Configurations to override prompt templates in different parts of an agent sequence. For more information, see Advanced prompts. See prompt_override_configuration Block for details.
	PromptOverrideConfiguration []PromptOverrideConfigurationInitParameters `json:"promptOverrideConfiguration,omitempty" tf:"prompt_override_configuration,omitempty"`

	// Whether the in-use check is skipped when deleting the agent.
	SkipResourceInUseCheck *bool `json:"skipResourceInUseCheck,omitempty" tf:"skip_resource_in_use_check,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AgentObservation struct {

	// ARN of the agent.
	AgentArn *string `json:"agentArn,omitempty" tf:"agent_arn,omitempty"`

	// Unique identifier of the agent.
	AgentID *string `json:"agentId,omitempty" tf:"agent_id,omitempty"`

	// Name of the agent.
	AgentName *string `json:"agentName,omitempty" tf:"agent_name,omitempty"`

	// ARN of the IAM role with permissions to invoke API operations on the agent.
	AgentResourceRoleArn *string `json:"agentResourceRoleArn,omitempty" tf:"agent_resource_role_arn,omitempty"`

	// Version of the agent.
	AgentVersion *string `json:"agentVersion,omitempty" tf:"agent_version,omitempty"`

	// ARN of the AWS KMS key that encrypts the agent.
	CustomerEncryptionKeyArn *string `json:"customerEncryptionKeyArn,omitempty" tf:"customer_encryption_key_arn,omitempty"`

	// Description of the agent.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Foundation model used for orchestration by the agent.
	FoundationModel *string `json:"foundationModel,omitempty" tf:"foundation_model,omitempty"`

	// Unique identifier of the agent.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent. A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Bedrock deletes any data provided before the timeout.
	IdleSessionTTLInSeconds *float64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds,omitempty"`

	// Instructions that tell the agent what it should do and how it should interact with users.
	Instruction *string `json:"instruction,omitempty" tf:"instruction,omitempty"`

	// Whether to prepare the agent after creation or modification. Defaults to true.
	PrepareAgent *bool `json:"prepareAgent,omitempty" tf:"prepare_agent,omitempty"`

	// Configurations to override prompt templates in different parts of an agent sequence. For more information, see Advanced prompts. See prompt_override_configuration Block for details.
	PromptOverrideConfiguration []PromptOverrideConfigurationObservation `json:"promptOverrideConfiguration,omitempty" tf:"prompt_override_configuration,omitempty"`

	// Whether the in-use check is skipped when deleting the agent.
	SkipResourceInUseCheck *bool `json:"skipResourceInUseCheck,omitempty" tf:"skip_resource_in_use_check,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AgentParameters struct {

	// Name of the agent.
	// +kubebuilder:validation:Optional
	AgentName *string `json:"agentName,omitempty" tf:"agent_name,omitempty"`

	// ARN of the IAM role with permissions to invoke API operations on the agent.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	AgentResourceRoleArn *string `json:"agentResourceRoleArn,omitempty" tf:"agent_resource_role_arn,omitempty"`

	// Reference to a Role in iam to populate agentResourceRoleArn.
	// +kubebuilder:validation:Optional
	AgentResourceRoleArnRef *v1.Reference `json:"agentResourceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate agentResourceRoleArn.
	// +kubebuilder:validation:Optional
	AgentResourceRoleArnSelector *v1.Selector `json:"agentResourceRoleArnSelector,omitempty" tf:"-"`

	// ARN of the AWS KMS key that encrypts the agent.
	// +kubebuilder:validation:Optional
	CustomerEncryptionKeyArn *string `json:"customerEncryptionKeyArn,omitempty" tf:"customer_encryption_key_arn,omitempty"`

	// Description of the agent.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Foundation model used for orchestration by the agent.
	// +kubebuilder:validation:Optional
	FoundationModel *string `json:"foundationModel,omitempty" tf:"foundation_model,omitempty"`

	// Number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent. A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Bedrock deletes any data provided before the timeout.
	// +kubebuilder:validation:Optional
	IdleSessionTTLInSeconds *float64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds,omitempty"`

	// Instructions that tell the agent what it should do and how it should interact with users.
	// +kubebuilder:validation:Optional
	Instruction *string `json:"instruction,omitempty" tf:"instruction,omitempty"`

	// Whether to prepare the agent after creation or modification. Defaults to true.
	// +kubebuilder:validation:Optional
	PrepareAgent *bool `json:"prepareAgent,omitempty" tf:"prepare_agent,omitempty"`

	// Configurations to override prompt templates in different parts of an agent sequence. For more information, see Advanced prompts. See prompt_override_configuration Block for details.
	// +kubebuilder:validation:Optional
	PromptOverrideConfiguration []PromptOverrideConfigurationParameters `json:"promptOverrideConfiguration,omitempty" tf:"prompt_override_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether the in-use check is skipped when deleting the agent.
	// +kubebuilder:validation:Optional
	SkipResourceInUseCheck *bool `json:"skipResourceInUseCheck,omitempty" tf:"skip_resource_in_use_check,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InferenceConfigurationInitParameters struct {

	// Maximum number of tokens to allow in the generated response.
	MaxLength *float64 `json:"maxLength,omitempty" tf:"max_length"`

	// List of stop sequences. A stop sequence is a sequence of characters that causes the model to stop generating the response.
	StopSequences []*string `json:"stopSequences,omitempty" tf:"stop_sequences"`

	// Likelihood of the model selecting higher-probability options while generating a response. A lower value makes the model more likely to choose higher-probability options, while a higher value makes the model more likely to choose lower-probability options.
	Temperature *float64 `json:"temperature,omitempty" tf:"temperature"`

	// Number of top most-likely candidates, between 0 and 500, from which the model chooses the next token in the sequence.
	TopK *float64 `json:"topK,omitempty" tf:"top_k"`

	// Top percentage of the probability distribution of next tokens, between 0 and 1 (denoting 0% and 100%), from which the model chooses the next token in the sequence.
	TopP *float64 `json:"topP,omitempty" tf:"top_p"`
}

type InferenceConfigurationObservation struct {

	// Maximum number of tokens to allow in the generated response.
	MaxLength *float64 `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// List of stop sequences. A stop sequence is a sequence of characters that causes the model to stop generating the response.
	StopSequences []*string `json:"stopSequences,omitempty" tf:"stop_sequences,omitempty"`

	// Likelihood of the model selecting higher-probability options while generating a response. A lower value makes the model more likely to choose higher-probability options, while a higher value makes the model more likely to choose lower-probability options.
	Temperature *float64 `json:"temperature,omitempty" tf:"temperature,omitempty"`

	// Number of top most-likely candidates, between 0 and 500, from which the model chooses the next token in the sequence.
	TopK *float64 `json:"topK,omitempty" tf:"top_k,omitempty"`

	// Top percentage of the probability distribution of next tokens, between 0 and 1 (denoting 0% and 100%), from which the model chooses the next token in the sequence.
	TopP *float64 `json:"topP,omitempty" tf:"top_p,omitempty"`
}

type InferenceConfigurationParameters struct {

	// Maximum number of tokens to allow in the generated response.
	// +kubebuilder:validation:Optional
	MaxLength *float64 `json:"maxLength,omitempty" tf:"max_length"`

	// List of stop sequences. A stop sequence is a sequence of characters that causes the model to stop generating the response.
	// +kubebuilder:validation:Optional
	StopSequences []*string `json:"stopSequences,omitempty" tf:"stop_sequences"`

	// Likelihood of the model selecting higher-probability options while generating a response. A lower value makes the model more likely to choose higher-probability options, while a higher value makes the model more likely to choose lower-probability options.
	// +kubebuilder:validation:Optional
	Temperature *float64 `json:"temperature,omitempty" tf:"temperature"`

	// Number of top most-likely candidates, between 0 and 500, from which the model chooses the next token in the sequence.
	// +kubebuilder:validation:Optional
	TopK *float64 `json:"topK,omitempty" tf:"top_k"`

	// Top percentage of the probability distribution of next tokens, between 0 and 1 (denoting 0% and 100%), from which the model chooses the next token in the sequence.
	// +kubebuilder:validation:Optional
	TopP *float64 `json:"topP,omitempty" tf:"top_p"`
}

type PromptConfigurationsInitParameters struct {

	// prompt template with which to replace the default prompt template. You can use placeholder variables in the base prompt template to customize the prompt. For more information, see Prompt template placeholder variables.
	BasePromptTemplate *string `json:"basePromptTemplate,omitempty" tf:"base_prompt_template"`

	// Inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the prompt_type. For more information, see Inference parameters for foundation models. See inference_configuration Block for details.
	InferenceConfiguration []InferenceConfigurationInitParameters `json:"inferenceConfiguration,omitempty" tf:"inference_configuration"`

	// Whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the prompt_type. If you set the argument as OVERRIDDEN, the override_lambda argument in the prompt_override_configuration block must be specified with the ARN of a Lambda function. Valid values: DEFAULT, OVERRIDDEN.
	ParserMode *string `json:"parserMode,omitempty" tf:"parser_mode"`

	// Whether to override the default prompt template for this prompt_type. Set this argument to OVERRIDDEN to use the prompt that you provide in the base_prompt_template. If you leave it as DEFAULT, the agent uses a default prompt template. Valid values: DEFAULT, OVERRIDDEN.
	PromptCreationMode *string `json:"promptCreationMode,omitempty" tf:"prompt_creation_mode"`

	// Whether to allow the agent to carry out the step specified in the prompt_type. If you set this argument to DISABLED, the agent skips that step. Valid Values: ENABLED, DISABLED.
	PromptState *string `json:"promptState,omitempty" tf:"prompt_state"`

	// Step in the agent sequence that this prompt configuration applies to. Valid values: PRE_PROCESSING, ORCHESTRATION, POST_PROCESSING, KNOWLEDGE_BASE_RESPONSE_GENERATION.
	PromptType *string `json:"promptType,omitempty" tf:"prompt_type"`
}

type PromptConfigurationsObservation struct {

	// prompt template with which to replace the default prompt template. You can use placeholder variables in the base prompt template to customize the prompt. For more information, see Prompt template placeholder variables.
	BasePromptTemplate *string `json:"basePromptTemplate,omitempty" tf:"base_prompt_template,omitempty"`

	// Inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the prompt_type. For more information, see Inference parameters for foundation models. See inference_configuration Block for details.
	InferenceConfiguration []InferenceConfigurationObservation `json:"inferenceConfiguration,omitempty" tf:"inference_configuration,omitempty"`

	// Whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the prompt_type. If you set the argument as OVERRIDDEN, the override_lambda argument in the prompt_override_configuration block must be specified with the ARN of a Lambda function. Valid values: DEFAULT, OVERRIDDEN.
	ParserMode *string `json:"parserMode,omitempty" tf:"parser_mode,omitempty"`

	// Whether to override the default prompt template for this prompt_type. Set this argument to OVERRIDDEN to use the prompt that you provide in the base_prompt_template. If you leave it as DEFAULT, the agent uses a default prompt template. Valid values: DEFAULT, OVERRIDDEN.
	PromptCreationMode *string `json:"promptCreationMode,omitempty" tf:"prompt_creation_mode,omitempty"`

	// Whether to allow the agent to carry out the step specified in the prompt_type. If you set this argument to DISABLED, the agent skips that step. Valid Values: ENABLED, DISABLED.
	PromptState *string `json:"promptState,omitempty" tf:"prompt_state,omitempty"`

	// Step in the agent sequence that this prompt configuration applies to. Valid values: PRE_PROCESSING, ORCHESTRATION, POST_PROCESSING, KNOWLEDGE_BASE_RESPONSE_GENERATION.
	PromptType *string `json:"promptType,omitempty" tf:"prompt_type,omitempty"`
}

type PromptConfigurationsParameters struct {

	// prompt template with which to replace the default prompt template. You can use placeholder variables in the base prompt template to customize the prompt. For more information, see Prompt template placeholder variables.
	// +kubebuilder:validation:Optional
	BasePromptTemplate *string `json:"basePromptTemplate,omitempty" tf:"base_prompt_template"`

	// Inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the prompt_type. For more information, see Inference parameters for foundation models. See inference_configuration Block for details.
	// +kubebuilder:validation:Optional
	InferenceConfiguration []InferenceConfigurationParameters `json:"inferenceConfiguration,omitempty" tf:"inference_configuration"`

	// Whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the prompt_type. If you set the argument as OVERRIDDEN, the override_lambda argument in the prompt_override_configuration block must be specified with the ARN of a Lambda function. Valid values: DEFAULT, OVERRIDDEN.
	// +kubebuilder:validation:Optional
	ParserMode *string `json:"parserMode,omitempty" tf:"parser_mode"`

	// Whether to override the default prompt template for this prompt_type. Set this argument to OVERRIDDEN to use the prompt that you provide in the base_prompt_template. If you leave it as DEFAULT, the agent uses a default prompt template. Valid values: DEFAULT, OVERRIDDEN.
	// +kubebuilder:validation:Optional
	PromptCreationMode *string `json:"promptCreationMode,omitempty" tf:"prompt_creation_mode"`

	// Whether to allow the agent to carry out the step specified in the prompt_type. If you set this argument to DISABLED, the agent skips that step. Valid Values: ENABLED, DISABLED.
	// +kubebuilder:validation:Optional
	PromptState *string `json:"promptState,omitempty" tf:"prompt_state"`

	// Step in the agent sequence that this prompt configuration applies to. Valid values: PRE_PROCESSING, ORCHESTRATION, POST_PROCESSING, KNOWLEDGE_BASE_RESPONSE_GENERATION.
	// +kubebuilder:validation:Optional
	PromptType *string `json:"promptType,omitempty" tf:"prompt_type"`
}

type PromptOverrideConfigurationInitParameters struct {

	// ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence. If you specify this field, at least one of the prompt_configurations block must contain a parser_mode value that is set to OVERRIDDEN.
	OverrideLambda *string `json:"overrideLambda,omitempty" tf:"override_lambda"`

	// Configurations to override a prompt template in one part of an agent sequence. See prompt_configurations Block for details.
	PromptConfigurations []PromptConfigurationsInitParameters `json:"promptConfigurations,omitempty" tf:"prompt_configurations"`
}

type PromptOverrideConfigurationObservation struct {

	// ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence. If you specify this field, at least one of the prompt_configurations block must contain a parser_mode value that is set to OVERRIDDEN.
	OverrideLambda *string `json:"overrideLambda,omitempty" tf:"override_lambda,omitempty"`

	// Configurations to override a prompt template in one part of an agent sequence. See prompt_configurations Block for details.
	PromptConfigurations []PromptConfigurationsObservation `json:"promptConfigurations,omitempty" tf:"prompt_configurations,omitempty"`
}

type PromptOverrideConfigurationParameters struct {

	// ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence. If you specify this field, at least one of the prompt_configurations block must contain a parser_mode value that is set to OVERRIDDEN.
	// +kubebuilder:validation:Optional
	OverrideLambda *string `json:"overrideLambda,omitempty" tf:"override_lambda"`

	// Configurations to override a prompt template in one part of an agent sequence. See prompt_configurations Block for details.
	// +kubebuilder:validation:Optional
	PromptConfigurations []PromptConfigurationsParameters `json:"promptConfigurations,omitempty" tf:"prompt_configurations"`
}

// AgentSpec defines the desired state of Agent
type AgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AgentInitParameters `json:"initProvider,omitempty"`
}

// AgentStatus defines the observed state of Agent.
type AgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Agent is the Schema for the Agents API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.agentName) || (has(self.initProvider) && has(self.initProvider.agentName))",message="spec.forProvider.agentName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.foundationModel) || (has(self.initProvider) && has(self.initProvider.foundationModel))",message="spec.forProvider.foundationModel is a required parameter"
	Spec   AgentSpec   `json:"spec"`
	Status AgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentList contains a list of Agents
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agent `json:"items"`
}

// Repository type metadata.
var (
	Agent_Kind             = "Agent"
	Agent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Agent_Kind}.String()
	Agent_KindAPIVersion   = Agent_Kind + "." + CRDGroupVersion.String()
	Agent_GroupVersionKind = CRDGroupVersion.WithKind(Agent_Kind)
)

func init() {
	SchemeBuilder.Register(&Agent{}, &AgentList{})
}