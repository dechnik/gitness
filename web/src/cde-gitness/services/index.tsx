/*
 * Copyright 2024 Harness Inc. All rights reserved.
 * Use of this source code is governed by the PolyForm Shield 1.0.0 license
 * that can be found in the licenses directory at the root of this repository, also available at
 * https://polyformproject.org/wp-content/uploads/2020/06/PolyForm-Shield-1.0.0.txt.
 */

/* Generated by restful-react */

import React from 'react'
import { Get, GetProps, useGet, UseGetProps, Mutate, MutateProps, useMutate, UseMutateProps } from 'restful-react'

import { getConfig } from '../../services/config'
export const SPEC_VERSION = '0.0.0'
export interface ApiFileDiffRequest {
  path?: string
  start_line?: number
}

export type EnumCIStatus = string

export type EnumCheckPayloadKind = '' | 'markdown' | 'pipeline' | 'raw'

export type EnumCheckStatus = 'error' | 'failure' | 'pending' | 'running' | 'success'

export type EnumCodeRepoType = 'github' | 'gitlab' | 'harness_code' | 'bitbucket' | 'unknown'

export type EnumContentEncodingType = 'base64' | 'utf8'

export type EnumFileDiffStatus = string

export type EnumGitspaceAccessType = 'jwt_token' | 'password' | 'ssh_key'

export type EnumGitspaceStateType = 'running' | 'stopped' | 'error' | 'uninitialized'

export type EnumIDEType = 'vs_code' | 'vs_code_web'

export type EnumMembershipRole = 'contributor' | 'executor' | 'reader' | 'space_owner'

export type EnumMergeCheckStatus = string

export type EnumMergeMethod = 'merge' | 'rebase' | 'squash'

export type EnumParentResourceType = 'space' | 'repo'

export type EnumPrincipalType = 'service' | 'serviceaccount' | 'user'

export type EnumProviderType = 'docker'

export type EnumPublicKeyUsage = 'auth'

export type EnumPullReqActivityKind = 'change-comment' | 'comment' | 'system'

export type EnumPullReqActivityType =
  | 'branch-delete'
  | 'branch-update'
  | 'code-comment'
  | 'comment'
  | 'merge'
  | 'review-submit'
  | 'state-change'
  | 'title-change'

export type EnumPullReqCommentStatus = 'active' | 'resolved'

export type EnumPullReqReviewDecision = 'approved' | 'changereq' | 'pending' | 'reviewed'

export type EnumPullReqReviewerType = 'assigned' | 'requested' | 'self_assigned'

export type EnumPullReqState = 'closed' | 'merged' | 'open'

export type EnumResolverType = string

export type EnumRuleState = 'active' | 'disabled' | 'monitor' | null

export type EnumTokenType = string

export type EnumTriggerAction =
  | 'branch_created'
  | 'branch_updated'
  | 'pullreq_branch_updated'
  | 'pullreq_closed'
  | 'pullreq_created'
  | 'pullreq_merged'
  | 'pullreq_reopened'
  | 'tag_created'
  | 'tag_updated'

export type EnumWebhookExecutionResult = 'fatal_error' | 'retriable_error' | 'success' | null

export type EnumWebhookParent = 'repo' | 'space'

export type EnumWebhookTrigger =
  | 'branch_created'
  | 'branch_deleted'
  | 'branch_updated'
  | 'pullreq_branch_updated'
  | 'pullreq_closed'
  | 'pullreq_comment_created'
  | 'pullreq_created'
  | 'pullreq_merged'
  | 'pullreq_reopened'
  | 'tag_created'
  | 'tag_deleted'
  | 'tag_updated'

export interface GitBlamePart {
  commit?: GitCommit
  lines?: string[] | null
}

export interface GitCommit {
  author?: GitSignature
  committer?: GitSignature
  file_stats?: GitCommitFileStats[]
  message?: string
  parent_shas?: ShaSHA[]
  sha?: ShaSHA
  title?: string
}

export interface GitCommitFileStats {
  [key: string]: any
}

export type GitFileAction = 'CREATE' | 'UPDATE' | 'DELETE' | 'MOVE' | 'PATCH_TEXT'

export interface GitFileDiff {
  additions?: number
  changes?: number
  deletions?: number
  is_binary?: boolean
  is_submodule?: boolean
  old_path?: string
  old_sha?: string
  patch?: number[]
  path?: string
  sha?: string
  status?: EnumFileDiffStatus
}

export interface GitIdentity {
  email?: string
  name?: string
}

export interface GitPathDetails {
  last_commit?: GitCommit
  path?: string
}

export interface GitSignature {
  identity?: GitIdentity
  when?: string
}

export type ImporterPipelineOption = 'convert' | 'ignore'

export interface ImporterProvider {
  host?: string
  password?: string
  type?: ImporterProviderType
  username?: string
}

export type ImporterProviderType = 'github' | 'gitlab' | 'bitbucket' | 'stash' | 'gitea' | 'gogs' | 'azure'

export interface JobProgress {
  failure?: string
  progress?: number
  result?: string
  state?: JobState
}

export type JobState = 'canceled' | 'failed' | 'finished' | 'running' | 'scheduled'

export interface LivelogLine {
  out?: string
  pos?: number
  time?: number
}

export interface OpenapiAdminUsersCreateRequest {
  display_name?: string
  email?: string
  password?: string
  uid?: string
}

export interface OpenapiAdminUsersUpdateRequest {
  display_name?: string | null
  email?: string | null
  password?: string | null
}

export interface OpenapiCalculateCommitDivergenceRequest {
  max_count?: number
  requests?: RepoCommitDivergenceRequest[] | null
}

export interface OpenapiCommentApplySuggestionstRequest {
  bypass_rules?: boolean
  dry_run_rules?: boolean
  message?: string
  suggestions?: PullreqSuggestionReference[] | null
  title?: string
}

export interface OpenapiCommentCreatePullReqRequest {
  line_end?: number
  line_end_new?: boolean
  line_start?: number
  line_start_new?: boolean
  parent_id?: number
  path?: string
  source_commit_sha?: string
  target_commit_sha?: string
  text?: string
}

export interface OpenapiCommentStatusPullReqRequest {
  status?: EnumPullReqCommentStatus
}

export interface OpenapiCommentUpdatePullReqRequest {
  text?: string
}

export interface OpenapiCommitFilesRequest {
  actions?: RepoCommitFileAction[] | null
  branch?: string
  bypass_rules?: boolean
  dry_run_rules?: boolean
  message?: string
  new_branch?: string
  title?: string
}

export type OpenapiContent = RepoFileContent | OpenapiDirContent | RepoSymlinkContent | RepoSubmoduleContent

export interface OpenapiContentInfo {
  latest_commit?: TypesCommit
  name?: string
  path?: string
  sha?: string
  type?: OpenapiContentType
}

export type OpenapiContentType = 'file' | 'dir' | 'symlink' | 'submodule'

export interface OpenapiCreateBranchRequest {
  bypass_rules?: boolean
  name?: string
  target?: string
}

export interface OpenapiCreateConnectorRequest {
  data?: string
  description?: string
  identifier?: string
  space_ref?: string
  type?: string
  uid?: string
}

export interface OpenapiCreateGitspaceRequest {
  branch?: string
  code_repo_url?: string
  devcontainer_path?: string
  id?: string
  ide?: EnumIDEType
  infra_provider_resource_id?: string
  metadata?: {
    [key: string]: string
  } | null
  name?: string
  provider?: ImporterProvider
  provider_repo?: string
}

export interface OpenapiCreatePipelineRequest {
  config_path?: string
  default_branch?: string
  description?: string
  disabled?: boolean
  identifier?: string
  uid?: string
}

export interface OpenapiCreatePullReqRequest {
  description?: string
  is_draft?: boolean
  source_branch?: string
  source_repo_ref?: string
  target_branch?: string
  title?: string
}

export interface OpenapiCreateRepositoryRequest {
  default_branch?: string
  description?: string
  fork_id?: number
  git_ignore?: string
  identifier?: string
  is_public?: boolean
  license?: string
  parent_ref?: string
  readme?: boolean
  uid?: string
}

export interface OpenapiCreateSecretRequest {
  data?: string
  description?: string
  identifier?: string
  space_ref?: string
  uid?: string
}

export interface OpenapiCreateSpaceRequest {
  description?: string
  identifier?: string
  is_public?: boolean
  parent_ref?: string
  uid?: string
}

export interface OpenapiCreateTagRequest {
  bypass_rules?: boolean
  message?: string
  name?: string
  target?: string
}

export interface OpenapiCreateTemplateRequest {
  data?: string
  description?: string
  identifier?: string
  space_ref?: string
  uid?: string
}

export interface OpenapiCreateTokenRequest {
  identifier?: string
  lifetime?: TimeDuration
  uid?: string
}

export interface OpenapiCreateTriggerRequest {
  actions?: EnumTriggerAction[] | null
  description?: string
  disabled?: boolean
  identifier?: string
  secret?: string
  uid?: string
}

export interface OpenapiCreateWebhookRequest {
  description?: string
  display_name?: string
  enabled?: boolean
  identifier?: string
  insecure?: boolean
  secret?: string
  triggers?: EnumWebhookTrigger[] | null
  uid?: string
  url?: string
}

export interface OpenapiDirContent {
  entries?: OpenapiContentInfo[] | null
}

export interface OpenapiExportSpaceRequest {
  account_id?: string
  org_identifier?: string
  project_identifier?: string
  token?: string
}

export interface OpenapiFileViewAddPullReqRequest {
  commit_sha?: string
  path?: string
}

export interface OpenapiGeneralSettingsRequest {
  file_size_limit?: number | null
}

export interface OpenapiGetContentOutput {
  content?: OpenapiContent
  latest_commit?: TypesCommit
  name?: string
  path?: string
  sha?: string
  type?: OpenapiContentType
}

export interface OpenapiLoginRequest {
  login_identifier?: string
  password?: string
}

export interface OpenapiMergePullReq {
  bypass_rules?: boolean
  dry_run?: boolean
  message?: string
  method?: EnumMergeMethod
  source_sha?: string
  title?: string
}

export interface OpenapiMoveRepoRequest {
  identifier?: string | null
  uid?: string | null
}

export interface OpenapiMoveSpaceRequest {
  identifier?: string | null
  uid?: string | null
}

export interface OpenapiPathsDetailsRequest {
  paths?: string[] | null
}

export type OpenapiPostRawDiffRequest = ApiFileDiffRequest[] | null

export type OpenapiPostRawPRDiffRequest = ApiFileDiffRequest[] | null

export interface OpenapiRegisterRequest {
  display_name?: string
  email?: string
  password?: string
  uid?: string
}

export interface OpenapiRestoreRequest {
  new_identifier?: string | null
  new_parent_ref?: string | null
}

export interface OpenapiRestoreSpaceRequest {
  new_identifier?: string | null
  new_parent_ref?: string | null
}

export interface OpenapiReviewSubmitPullReqRequest {
  commit_sha?: string
  decision?: EnumPullReqReviewDecision
}

export interface OpenapiReviewerAddPullReqRequest {
  reviewer_id?: number
}

export interface OpenapiRule {
  created?: number
  created_by?: TypesPrincipalInfo
  definition?: OpenapiRuleDefinition
  description?: string
  identifier?: string
  pattern?: ProtectionPattern
  state?: EnumRuleState
  type?: OpenapiRuleType
  updated?: number
  users?: {
    [key: string]: TypesPrincipalInfo
  } | null
}

export type OpenapiRuleDefinition = ProtectionBranch

export type OpenapiRuleType = 'branch'

export interface OpenapiSecuritySettingsRequest {
  secret_scanning_enabled?: boolean | null
}

export interface OpenapiStatePullReqRequest {
  is_draft?: boolean
  state?: EnumPullReqState
}

export interface OpenapiUpdateAdminRequest {
  admin?: boolean
}

export interface OpenapiUpdateConnectorRequest {
  data?: string | null
  description?: string | null
  identifier?: string | null
  uid?: string | null
}

export interface OpenapiUpdatePipelineRequest {
  config_path?: string | null
  description?: string | null
  disabled?: boolean | null
  identifier?: string | null
  uid?: string | null
}

export interface OpenapiUpdatePullReqRequest {
  description?: string
  title?: string
}

export interface OpenapiUpdateRepoPublicAccessRequest {
  is_public?: boolean
}

export interface OpenapiUpdateRepoRequest {
  description?: string | null
}

export interface OpenapiUpdateSecretRequest {
  data?: string | null
  description?: string | null
  identifier?: string | null
  uid?: string | null
}

export interface OpenapiUpdateSpacePublicAccessRequest {
  is_public?: boolean
}

export interface OpenapiUpdateSpaceRequest {
  description?: string | null
}

export interface OpenapiUpdateTemplateRequest {
  data?: string | null
  description?: string | null
  identifier?: string | null
  uid?: string | null
}

export interface OpenapiUpdateTriggerRequest {
  actions?: EnumTriggerAction[] | null
  description?: string | null
  disabled?: boolean | null
  identifier?: string | null
  secret?: string | null
  uid?: string | null
}

export interface OpenapiUpdateWebhookRequest {
  description?: string | null
  display_name?: string | null
  enabled?: boolean | null
  identifier?: string | null
  insecure?: boolean | null
  secret?: string | null
  triggers?: EnumWebhookTrigger[] | null
  uid?: string | null
  url?: string | null
}

export interface OpenapiWebhookType {
  created?: number
  created_by?: number
  description?: string
  display_name?: string
  enabled?: boolean
  has_secret?: boolean
  id?: number
  identifier?: string
  insecure?: boolean
  latest_execution_result?: EnumWebhookExecutionResult
  parent_id?: number
  parent_type?: EnumWebhookParent
  triggers?: EnumWebhookTrigger[] | null
  updated?: number
  url?: string
  version?: number
}

export interface ProtectionBranch {
  bypass?: ProtectionDefBypass
  lifecycle?: ProtectionDefLifecycle
  pullreq?: ProtectionDefPullReq
}

export interface ProtectionDefApprovals {
  require_code_owners?: boolean
  require_latest_commit?: boolean
  require_minimum_count?: number
  require_no_change_request?: boolean
}

export interface ProtectionDefBypass {
  repo_owners?: boolean
  user_ids?: number[]
}

export interface ProtectionDefComments {
  require_resolve_all?: boolean
}

export interface ProtectionDefLifecycle {
  create_forbidden?: boolean
  delete_forbidden?: boolean
  update_forbidden?: boolean
}

export interface ProtectionDefMerge {
  delete_branch?: boolean
  strategies_allowed?: EnumMergeMethod[]
}

export interface ProtectionDefPullReq {
  approvals?: ProtectionDefApprovals
  comments?: ProtectionDefComments
  merge?: ProtectionDefMerge
  status_checks?: ProtectionDefStatusChecks
}

export interface ProtectionDefStatusChecks {
  require_identifiers?: string[]
}

export type ProtectionPattern = {
  default?: boolean
  exclude?: string[]
  include?: string[]
} | null

export interface PullreqCommentApplySuggestionsOutput {
  commit_id?: string
  dry_run_rules?: boolean
  rule_violations?: TypesRuleViolations[]
}

export interface PullreqSuggestionReference {
  check_sum?: string
  comment_id?: number
}

export interface RepoBranch {
  commit?: TypesCommit
  name?: string
  sha?: string
}

export interface RepoCommitDivergence {
  ahead?: number
  behind?: number
}

export interface RepoCommitDivergenceRequest {
  from?: string
  to?: string
}

export interface RepoCommitFileAction {
  action?: GitFileAction
  encoding?: EnumContentEncodingType
  path?: string
  payload?: string
  sha?: ShaSHA
}

export interface RepoCommitTag {
  commit?: TypesCommit
  is_annotated?: boolean
  message?: string
  name?: string
  sha?: string
  tagger?: TypesSignature
  title?: string
}

export interface RepoContentInfo {
  latest_commit?: TypesCommit
  name?: string
  path?: string
  sha?: string
  type?: RepoContentType
}

export type RepoContentType = string

export interface RepoFileContent {
  data?: string
  data_size?: number
  encoding?: EnumContentEncodingType
  size?: number
}

export interface RepoListPathsOutput {
  directories?: string[]
  files?: string[]
}

export interface RepoMergeCheck {
  conflict_files?: string[]
  mergeable?: boolean
}

export interface RepoPathsDetailsOutput {
  details?: GitPathDetails[] | null
}

export interface RepoRepositoryOutput {
  created?: number
  created_by?: number
  default_branch?: string
  deleted?: number | null
  description?: string
  fork_id?: number
  git_ssh_url?: string
  git_url?: string
  id?: number
  identifier?: string
  importing?: boolean
  is_empty?: boolean
  is_public?: boolean
  num_closed_pulls?: number
  num_forks?: number
  num_merged_pulls?: number
  num_open_pulls?: number
  num_pulls?: number
  parent_id?: number
  path?: string
  size?: number
  size_updated?: number
  updated?: number
}

export interface RepoSoftDeleteResponse {
  deleted_at?: number
}

export interface RepoSubmoduleContent {
  commit_sha?: string
  url?: string
}

export interface RepoSymlinkContent {
  size?: number
  target?: string
}

export interface ReposettingsGeneralSettings {
  file_size_limit?: number | null
}

export interface ReposettingsSecuritySettings {
  secret_scanning_enabled?: boolean | null
}

/**
 * Git object hash
 */
export type ShaSHA = string

export interface SpaceExportProgressOutput {
  repos?: JobProgress[] | null
}

export interface SpaceImportRepositoriesOutput {
  duplicate_repos?: RepoRepositoryOutput[] | null
  importing_repos?: RepoRepositoryOutput[] | null
}

export interface SpaceSoftDeleteResponse {
  deleted_at?: number
}

export interface SpaceSpaceOutput {
  created?: number
  created_by?: number
  deleted?: number | null
  description?: string
  id?: number
  identifier?: string
  is_public?: boolean
  parent_id?: number
  path?: string
  updated?: number
}

export interface SystemConfigOutput {
  public_resource_creation_enabled?: boolean
  ssh_enabled?: boolean
  user_signup_allowed?: boolean
}

export type TimeDuration = number | null

export interface TypesChangeStats {
  changes?: number
  deletions?: number
  insertions?: number
}

export interface TypesCheck {
  created?: number
  ended?: number
  id?: number
  identifier?: string
  link?: string
  metadata?: {}
  payload?: TypesCheckPayload
  reported_by?: TypesPrincipalInfo
  started?: number
  status?: EnumCheckStatus
  summary?: string
  updated?: number
}

export interface TypesCheckPayload {
  data?: {}
  kind?: EnumCheckPayloadKind
  version?: string
}

export interface TypesCodeCommentFields {
  line_new?: number
  line_old?: number
  merge_base_sha?: string
  outdated?: boolean
  path?: string
  source_sha?: string
  span_new?: number
  span_old?: number
}

export interface TypesCodeOwnerEvaluation {
  evaluation_entries?: TypesCodeOwnerEvaluationEntry[] | null
  file_sha?: string
}

export interface TypesCodeOwnerEvaluationEntry {
  line_number?: number
  owner_evaluations?: TypesOwnerEvaluation[] | null
  pattern?: string
  user_group_owner_evaluations?: TypesUserGroupOwnerEvaluation[] | null
}

export interface TypesCommit {
  author?: TypesSignature
  committer?: TypesSignature
  message?: string
  parent_shas?: string[]
  sha?: string
  stats?: TypesCommitStats
  title?: string
}

export interface TypesCommitFileStats {
  changes?: number
  deletions?: number
  insertions?: number
  old_path?: string
  path?: string
  status?: EnumFileDiffStatus
}

export interface TypesCommitFilesResponse {
  commit_id?: string
  dry_run_rules?: boolean
  rule_violations?: TypesRuleViolations[]
}

export interface TypesCommitStats {
  files?: TypesCommitFileStats[]
  total?: TypesChangeStats
}

export interface TypesConnector {
  created?: number
  data?: string
  description?: string
  identifier?: string
  space_id?: number
  type?: string
  updated?: number
}

export interface TypesDiffStats {
  commits?: number | null
  files_changed?: number | null
}

export interface TypesExecution {
  action?: string
  after?: string
  author_avatar?: string
  author_email?: string
  author_login?: string
  author_name?: string
  before?: string
  created?: number
  created_by?: number
  cron?: string
  debug?: boolean
  deploy_id?: number
  deploy_to?: string
  error?: string
  event?: string
  finished?: number
  link?: string
  message?: string
  number?: number
  params?: {
    [key: string]: string
  }
  parent?: number
  pipeline_id?: number
  ref?: string
  repo_id?: number
  sender?: string
  source?: string
  source_repo?: string
  stages?: TypesStage[]
  started?: number
  status?: EnumCIStatus
  target?: string
  timestamp?: number
  title?: string
  trigger?: string
  updated?: number
}

export interface TypesGitspaceConfig {
  branch?: string
  code_repo_id?: string
  code_repo_type?: EnumCodeRepoType
  code_repo_url?: string
  created?: number
  devcontainer_path?: string
  id?: string
  ide?: EnumIDEType
  infra_provider_resource_id?: string
  instance?: TypesGitspaceInstance
  name?: string
  space_path?: string
  updated?: number
  user_id?: string
}

export interface TypesGitspaceInstance {
  access_key?: string
  access_type?: EnumGitspaceAccessType
  created?: number
  id?: string
  last_used?: number
  machine_user?: string
  resource_usage?: string
  space_path?: string
  state?: EnumGitspaceStateType
  total_time_used?: number
  tracked_changes?: string
  updated?: number
  url?: string
}

export interface TypesIdentity {
  email?: string
  name?: string
}

export interface TypesInfraProviderResource {
  cpu?: string
  created?: number
  disk?: string
  gateway_host?: string
  gateway_port?: string
  id?: string
  infra_provider_config_id?: string
  infra_provider_type?: EnumProviderType
  memory?: string
  metadata?: {
    [key: string]: string
  } | null
  name?: string
  network?: string
  region?: string
  space_path?: string
  template_id?: string
  updated?: number
}

export interface TypesListCommitResponse {
  commits?: TypesCommit[] | null
  rename_details?: TypesRenameDetails[] | null
  total_commits?: number
}

export interface TypesMembershipSpace {
  added_by?: TypesPrincipalInfo
  created?: number
  role?: EnumMembershipRole
  space?: TypesSpace
  updated?: number
}

export interface TypesMembershipUser {
  added_by?: TypesPrincipalInfo
  created?: number
  principal?: TypesPrincipalInfo
  role?: EnumMembershipRole
  updated?: number
}

export interface TypesMergeResponse {
  allowed_methods?: EnumMergeMethod[]
  branch_deleted?: boolean
  conflict_files?: string[]
  dry_run?: boolean
  minimum_required_approvals_count?: number
  minimum_required_approvals_count_latest?: number
  requires_code_owners_approval?: boolean
  requires_code_owners_approval_latest?: boolean
  requires_comment_resolution?: boolean
  requires_no_change_requests?: boolean
  rule_violations?: TypesRuleViolations[]
  sha?: string
}

export interface TypesMergeViolations {
  conflict_files?: string[]
  rule_violations?: TypesRuleViolations[]
}

export interface TypesOwnerEvaluation {
  owner?: TypesPrincipalInfo
  review_decision?: EnumPullReqReviewDecision
  review_sha?: string
}

export interface TypesPipeline {
  config_path?: string
  created?: number
  created_by?: number
  default_branch?: string
  description?: string
  disabled?: boolean
  execution?: TypesExecution
  identifier?: string
  repo_id?: number
  seq?: number
  updated?: number
}

export interface TypesPlugin {
  description?: string
  identifier?: string
  logo?: string
  spec?: string
  type?: string
  version?: string
}

export interface TypesPrincipalInfo {
  created?: number
  display_name?: string
  email?: string
  id?: number
  type?: EnumPrincipalType
  uid?: string
  updated?: number
}

export interface TypesPublicKey {
  comment?: string
  created?: number
  fingerprint?: string
  identifier?: string
  type?: string
  usage?: EnumPublicKeyUsage
  verified?: number | null
}

export interface TypesPullReq {
  author?: TypesPrincipalInfo
  closed?: number | null
  created?: number
  description?: string
  edited?: number
  is_draft?: boolean
  merge_base_sha?: string
  merge_check_status?: EnumMergeCheckStatus
  merge_conflicts?: string[]
  merge_method?: EnumMergeMethod
  merge_target_sha?: string | null
  merged?: number | null
  merger?: TypesPrincipalInfo
  number?: number
  source_branch?: string
  source_repo_id?: number
  source_sha?: string
  state?: EnumPullReqState
  stats?: TypesPullReqStats
  target_branch?: string
  target_repo_id?: number
  title?: string
}

export interface TypesPullReqActivity {
  author?: TypesPrincipalInfo
  code_comment?: TypesCodeCommentFields
  created?: number
  deleted?: number | null
  edited?: number
  id?: number
  kind?: EnumPullReqActivityKind
  mentions?: {
    [key: string]: TypesPrincipalInfo
  }
  metadata?: TypesPullReqActivityMetadata
  order?: number
  parent_id?: number | null
  payload?: {}
  pullreq_id?: number
  repo_id?: number
  resolved?: number | null
  resolver?: TypesPrincipalInfo
  sub_order?: number
  text?: string
  type?: EnumPullReqActivityType
  updated?: number
}

export interface TypesPullReqActivityMentionsMetadata {
  ids?: number[]
}

export interface TypesPullReqActivityMetadata {
  mentions?: TypesPullReqActivityMentionsMetadata
  suggestions?: TypesPullReqActivitySuggestionsMetadata
}

export interface TypesPullReqActivitySuggestionsMetadata {
  applied_check_sum?: string
  applied_commit_sha?: string
  check_sums?: string[]
}

export interface TypesPullReqCheck {
  bypassable?: boolean
  check?: TypesCheck
  required?: boolean
}

export interface TypesPullReqChecks {
  checks?: TypesPullReqCheck[] | null
  commit_sha?: string
}

export interface TypesPullReqFileView {
  obsolete?: boolean
  path?: string
  sha?: string
}

export interface TypesPullReqReviewer {
  added_by?: TypesPrincipalInfo
  created?: number
  latest_review_id?: number | null
  review_decision?: EnumPullReqReviewDecision
  reviewer?: TypesPrincipalInfo
  sha?: string
  type?: EnumPullReqReviewerType
  updated?: number
}

export interface TypesPullReqStats {
  commits?: number | null
  conversations?: number
  files_changed?: number | null
  unresolved_count?: number
}

export interface TypesRenameDetails {
  commit_sha_after?: string
  commit_sha_before?: string
  new_path?: string
  old_path?: string
}

export interface TypesRepository {
  created?: number
  created_by?: number
  default_branch?: string
  deleted?: number | null
  description?: string
  fork_id?: number
  git_ssh_url?: string
  git_url?: string
  id?: number
  identifier?: string
  importing?: boolean
  is_empty?: boolean
  num_closed_pulls?: number
  num_forks?: number
  num_merged_pulls?: number
  num_open_pulls?: number
  num_pulls?: number
  parent_id?: number
  path?: string
  size?: number
  size_updated?: number
  updated?: number
}

export interface TypesRepositoryPullReqSummary {
  closed_count?: number
  merged_count?: number
  open_count?: number
}

export interface TypesRepositorySummary {
  branch_count?: number
  default_branch_commit_count?: number
  pull_req_summary?: TypesRepositoryPullReqSummary
  tag_count?: number
}

export interface TypesRuleInfo {
  identifier?: string
  repo_path?: string
  space_path?: string
  state?: EnumRuleState
  type?: TypesRuleType
}

export type TypesRuleType = string

export interface TypesRuleViolations {
  bypassable?: boolean
  bypassed?: boolean
  rule?: TypesRuleInfo
  violations?: TypesViolation[] | null
}

export interface TypesRulesViolations {
  violations?: TypesRuleViolations[] | null
}

export interface TypesSecret {
  created?: number
  created_by?: number
  description?: string
  identifier?: string
  space_id?: number
  updated?: number
}

export interface TypesServiceAccount {
  admin?: boolean
  blocked?: boolean
  created?: number
  display_name?: string
  email?: string
  parent_id?: number
  parent_type?: EnumParentResourceType
  uid?: string
  updated?: number
}

export interface TypesSignature {
  identity?: TypesIdentity
  when?: string
}

export interface TypesSpace {
  created?: number
  created_by?: number
  deleted?: number | null
  description?: string
  id?: number
  identifier?: string
  parent_id?: number
  path?: string
  updated?: number
}

export interface TypesStage {
  arch?: string
  depends_on?: string[]
  errignore?: boolean
  error?: string
  execution_id?: number
  exit_code?: number
  kernel?: string
  kind?: string
  labels?: {
    [key: string]: string
  }
  limit?: number
  machine?: string
  name?: string
  number?: number
  on_failure?: boolean
  on_success?: boolean
  os?: string
  repo_id?: number
  started?: number
  status?: EnumCIStatus
  steps?: TypesStep[]
  stopped?: number
  throttle?: number
  type?: string
  variant?: string
}

export interface TypesStep {
  depends_on?: string[]
  detached?: boolean
  errignore?: boolean
  error?: string
  exit_code?: number
  image?: string
  name?: string
  number?: number
  schema?: string
  started?: number
  status?: EnumCIStatus
  stopped?: number
}

export interface TypesTemplate {
  created?: number
  data?: string
  description?: string
  identifier?: string
  space_id?: number
  type?: EnumResolverType
  updated?: number
}

export interface TypesToken {
  created_by?: number
  expires_at?: number | null
  identifier?: string
  issued_at?: number
  principal_id?: number
  type?: EnumTokenType
}

export interface TypesTokenResponse {
  access_token?: string
  token?: TypesToken
}

export interface TypesTrigger {
  actions?: EnumTriggerAction[] | null
  created?: number
  created_by?: number
  description?: string
  disabled?: boolean
  identifier?: string
  pipeline_id?: number
  repo_id?: number
  trigger_type?: string
  updated?: number
}

export interface TypesUser {
  admin?: boolean
  blocked?: boolean
  created?: number
  display_name?: string
  email?: string
  uid?: string
  updated?: number
}

export interface TypesUserGroupOwnerEvaluation {
  evaluations?: TypesOwnerEvaluation[] | null
  id?: string
  name?: string
}

export interface TypesViolation {
  code?: string
  message?: string
  params?: {}[] | null
}

export interface TypesWebhookExecution {
  created?: number
  duration?: number
  error?: string
  id?: number
  request?: TypesWebhookExecutionRequest
  response?: TypesWebhookExecutionResponse
  result?: EnumWebhookExecutionResult
  retrigger_of?: number | null
  retriggerable?: boolean
  trigger_type?: EnumWebhookTrigger
  webhook_id?: number
}

export interface TypesWebhookExecutionRequest {
  body?: string
  headers?: string
  url?: string
}

export interface TypesWebhookExecutionResponse {
  body?: string
  headers?: string
  status?: string
  status_code?: number
}

export interface UploadResult {
  file_path?: string
}

export interface UserCreatePublicKeyInput {
  content?: string
  identifier?: string
  usage?: EnumPublicKeyUsage
}

export interface UserUpdateInput {
  display_name?: string | null
  email?: string | null
  password?: string | null
}

export interface UsererrorError {
  message?: string
  values?: { [key: string]: any }
}

export interface ListGitspacesQueryParams {
  sort?: 'id' | 'created' | 'updated'
  order?: 'asc' | 'desc'
  page?: number
  limit?: number
}

export type ListGitspacesProps = Omit<
  GetProps<TypesGitspaceConfig[], UsererrorError, ListGitspacesQueryParams, void>,
  'path'
>

/**
 * List gitspaces
 */
export const ListGitspaces = (props: ListGitspacesProps) => (
  <Get<TypesGitspaceConfig[], UsererrorError, ListGitspacesQueryParams, void>
    path={`/gitspaces`}
    base={getConfig('code/api/v1')}
    {...props}
  />
)

export type UseListGitspacesProps = Omit<
  UseGetProps<TypesGitspaceConfig[], UsererrorError, ListGitspacesQueryParams, void>,
  'path'
>

/**
 * List gitspaces
 */
export const useListGitspaces = (props: UseListGitspacesProps) =>
  useGet<TypesGitspaceConfig[], UsererrorError, ListGitspacesQueryParams, void>(`/gitspaces`, {
    base: getConfig('code/api/v1'),
    ...props
  })

export type CreateGitspaceProps = Omit<
  MutateProps<TypesGitspaceConfig, UsererrorError, void, OpenapiCreateGitspaceRequest, void>,
  'path' | 'verb'
>

/**
 * Create gitspace config
 */
export const CreateGitspace = (props: CreateGitspaceProps) => (
  <Mutate<TypesGitspaceConfig, UsererrorError, void, OpenapiCreateGitspaceRequest, void>
    verb="POST"
    path={`/gitspaces`}
    base={getConfig('code/api/v1')}
    {...props}
  />
)

export type UseCreateGitspaceProps = Omit<
  UseMutateProps<TypesGitspaceConfig, UsererrorError, void, OpenapiCreateGitspaceRequest, void>,
  'path' | 'verb'
>

/**
 * Create gitspace config
 */
export const useCreateGitspace = (props: UseCreateGitspaceProps) =>
  useMutate<TypesGitspaceConfig, UsererrorError, void, OpenapiCreateGitspaceRequest, void>('POST', `/gitspaces`, {
    base: getConfig('code/api/v1'),
    ...props
  })

export type DeleteGitspaceProps = Omit<MutateProps<void, UsererrorError, void, string, void>, 'path' | 'verb'>

/**
 * Delete gitspace config
 */
export const DeleteGitspace = (props: DeleteGitspaceProps) => (
  <Mutate<void, UsererrorError, void, string, void>
    verb="DELETE"
    path={`/gitspaces`}
    base={getConfig('code/api/v1')}
    {...props}
  />
)

export type UseDeleteGitspaceProps = Omit<UseMutateProps<void, UsererrorError, void, string, void>, 'path' | 'verb'>

/**
 * Delete gitspace config
 */
export const useDeleteGitspace = (props: UseDeleteGitspaceProps) =>
  useMutate<void, UsererrorError, void, string, void>('DELETE', `/gitspaces`, {
    base: getConfig('code/api/v1'),
    ...props
  })

export interface FindGitspacePathParams {
  gitspace_ref: string
}

export type FindGitspaceProps = Omit<
  GetProps<TypesGitspaceConfig, UsererrorError, void, FindGitspacePathParams>,
  'path'
> &
  FindGitspacePathParams

/**
 * Get gitspace
 */
export const FindGitspace = ({ gitspace_ref, ...props }: FindGitspaceProps) => (
  <Get<TypesGitspaceConfig, UsererrorError, void, FindGitspacePathParams>
    path={`/gitspaces/${gitspace_ref}`}
    base={getConfig('code/api/v1')}
    {...props}
  />
)

export type UseFindGitspaceProps = Omit<
  UseGetProps<TypesGitspaceConfig, UsererrorError, void, FindGitspacePathParams>,
  'path'
> &
  FindGitspacePathParams

/**
 * Get gitspace
 */
export const useFindGitspace = ({ gitspace_ref, ...props }: UseFindGitspaceProps) =>
  useGet<TypesGitspaceConfig, UsererrorError, void, FindGitspacePathParams>(
    (paramsInPath: FindGitspacePathParams) => `/gitspaces/${paramsInPath.gitspace_ref}`,
    { base: getConfig('code/api/v1'), pathParams: { gitspace_ref }, ...props }
  )