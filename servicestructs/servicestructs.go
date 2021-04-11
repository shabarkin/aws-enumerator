package servicestructs

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/chime"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codestar"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	"github.com/aws/aws-sdk-go-v2/service/macie"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
	"github.com/aws/aws-sdk-go-v2/service/mobile"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/sms"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/workdocs"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workmail"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/cloud-enumerator/servicemaster"
	"github.com/cloud-enumerator/utils"
)

func GetServices() []servicemaster.ServiceMaster {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalln(utils.Red("Error:"), utils.Yellow("Unable to load SDK config,"))
	}

	acm_svc := &servicemaster.ServiceMaster{
		Svc:     acm.NewFromConfig(cfg),
		SvcName: "acm",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListCertificates", "input_obj": &acm.ListCertificatesInput{}},
		}}

	amplify_svc := &servicemaster.ServiceMaster{
		Svc:     amplify.NewFromConfig(cfg),
		SvcName: "amplify",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListApps", "input_obj": &amplify.ListAppsInput{}},
		}}

	apigateway_svc := &servicemaster.ServiceMaster{
		Svc:     apigateway.NewFromConfig(cfg),
		SvcName: "apigateway",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetAccount", "input_obj": &apigateway.GetAccountInput{}},
			{"apicall": "GetDomainNames", "input_obj": &apigateway.GetDomainNamesInput{}},
			{"apicall": "GetUsagePlans", "input_obj": &apigateway.GetUsagePlansInput{}},
			{"apicall": "GetClientCertificates", "input_obj": &apigateway.GetClientCertificatesInput{}},
			{"apicall": "GetApiKeys", "input_obj": &apigateway.GetApiKeysInput{}},
			{"apicall": "GetSdkTypes", "input_obj": &apigateway.GetSdkTypesInput{}},
			{"apicall": "GetVpcLinks", "input_obj": &apigateway.GetVpcLinksInput{}},
			{"apicall": "GetRestApis", "input_obj": &apigateway.GetRestApisInput{}},
		}}

	appmesh_svc := &servicemaster.ServiceMaster{
		Svc:     appmesh.NewFromConfig(cfg),
		SvcName: "appmesh",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListMeshes", "input_obj": &appmesh.ListMeshesInput{}},
		}}

	appsync_svc := &servicemaster.ServiceMaster{
		Svc:     appsync.NewFromConfig(cfg),
		SvcName: "appsync",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListGraphqlApis", "input_obj": &appsync.ListGraphqlApisInput{}},
		}}

	athena_svc := &servicemaster.ServiceMaster{
		Svc:     athena.NewFromConfig(cfg),
		SvcName: "athena",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListQueryExecutions", "input_obj": &athena.ListQueryExecutionsInput{}},
			{"apicall": "ListNamedQueries", "input_obj": &athena.ListNamedQueriesInput{}},
			{"apicall": "ListWorkGroups", "input_obj": &athena.ListWorkGroupsInput{}},
		}}

	autoscaling_svc := &servicemaster.ServiceMaster{
		Svc:     autoscaling.NewFromConfig(cfg),
		SvcName: "autoscaling",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeAdjustmentTypes", "input_obj": &autoscaling.DescribeAdjustmentTypesInput{}},
			{"apicall": "DescribeScheduledActions", "input_obj": &autoscaling.DescribeScheduledActionsInput{}},
			{"apicall": "DescribeAutoScalingGroups", "input_obj": &autoscaling.DescribeAutoScalingGroupsInput{}},
			{"apicall": "DescribeNotificationConfigurations", "input_obj": &autoscaling.DescribeNotificationConfigurationsInput{}},
			{"apicall": "DescribeAccountLimits", "input_obj": &autoscaling.DescribeAccountLimitsInput{}},
			{"apicall": "DescribePolicies", "input_obj": &autoscaling.DescribePoliciesInput{}},
			{"apicall": "DescribeScalingProcessTypes", "input_obj": &autoscaling.DescribeScalingProcessTypesInput{}},
			{"apicall": "DescribeTerminationPolicyTypes", "input_obj": &autoscaling.DescribeTerminationPolicyTypesInput{}},
			{"apicall": "DescribeScalingActivities", "input_obj": &autoscaling.DescribeScalingActivitiesInput{}},
			{"apicall": "DescribeAutoScalingNotificationTypes", "input_obj": &autoscaling.DescribeAutoScalingNotificationTypesInput{}},
			{"apicall": "DescribeLaunchConfigurations", "input_obj": &autoscaling.DescribeLaunchConfigurationsInput{}},
			{"apicall": "DescribeLifecycleHookTypes", "input_obj": &autoscaling.DescribeLifecycleHookTypesInput{}},
			{"apicall": "DescribeMetricCollectionTypes", "input_obj": &autoscaling.DescribeMetricCollectionTypesInput{}},
			{"apicall": "DescribeTags", "input_obj": &autoscaling.DescribeTagsInput{}},
			{"apicall": "DescribeAutoScalingInstances", "input_obj": &autoscaling.DescribeAutoScalingInstancesInput{}},
		}}

	backup_svc := &servicemaster.ServiceMaster{
		Svc:     backup.NewFromConfig(cfg),
		SvcName: "backup",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListBackupVaults", "input_obj": &backup.ListBackupVaultsInput{}},
			{"apicall": "ListProtectedResources", "input_obj": &backup.ListProtectedResourcesInput{}},
			{"apicall": "ListBackupPlanTemplates", "input_obj": &backup.ListBackupPlanTemplatesInput{}},
			{"apicall": "ListBackupJobs", "input_obj": &backup.ListBackupJobsInput{}},
			{"apicall": "GetSupportedResourceTypes", "input_obj": &backup.GetSupportedResourceTypesInput{}},
			{"apicall": "ListRestoreJobs", "input_obj": &backup.ListRestoreJobsInput{}},
			{"apicall": "ListBackupPlans", "input_obj": &backup.ListBackupPlansInput{}},
		}}

	batch_svc := &servicemaster.ServiceMaster{
		Svc:     batch.NewFromConfig(cfg),
		SvcName: "batch",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListJobs", "input_obj": &batch.ListJobsInput{}},
			{"apicall": "DescribeComputeEnvironments", "input_obj": &batch.DescribeComputeEnvironmentsInput{}},
			{"apicall": "DescribeJobDefinitions", "input_obj": &batch.DescribeJobDefinitionsInput{}},
			{"apicall": "DescribeJobQueues", "input_obj": &batch.DescribeJobQueuesInput{}},
		}}

	chime_svc := &servicemaster.ServiceMaster{
		Svc:     chime.NewFromConfig(cfg),
		SvcName: "chime",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListAccounts", "input_obj": &chime.ListAccountsInput{}},
		}}

	cloud9_svc := &servicemaster.ServiceMaster{
		Svc:     cloud9.NewFromConfig(cfg),
		SvcName: "cloud9",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListEnvironments", "input_obj": &cloud9.ListEnvironmentsInput{}},
			{"apicall": "DescribeEnvironmentMemberships", "input_obj": &cloud9.DescribeEnvironmentMembershipsInput{}},
		}}

	clouddirectory_svc := &servicemaster.ServiceMaster{
		Svc:     clouddirectory.NewFromConfig(cfg),
		SvcName: "clouddirectory",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListDevelopmentSchemaArns", "input_obj": &clouddirectory.ListDevelopmentSchemaArnsInput{}},
			{"apicall": "ListManagedSchemaArns", "input_obj": &clouddirectory.ListManagedSchemaArnsInput{}},
			{"apicall": "ListDirectories", "input_obj": &clouddirectory.ListDirectoriesInput{}},
			{"apicall": "ListPublishedSchemaArns", "input_obj": &clouddirectory.ListPublishedSchemaArnsInput{}},
		}}

	cloudformation_svc := &servicemaster.ServiceMaster{
		Svc:     cloudformation.NewFromConfig(cfg),
		SvcName: "cloudformation",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeStackEvents", "input_obj": &cloudformation.DescribeStackEventsInput{}},
			{"apicall": "ListExports", "input_obj": &cloudformation.ListExportsInput{}},
			{"apicall": "DescribeStackResources", "input_obj": &cloudformation.DescribeStackResourcesInput{}},
			{"apicall": "DescribeAccountLimits", "input_obj": &cloudformation.DescribeAccountLimitsInput{}},
			{"apicall": "ListStackSets", "input_obj": &cloudformation.ListStackSetsInput{}},
			{"apicall": "ListStacks", "input_obj": &cloudformation.ListStacksInput{}},
			{"apicall": "GetTemplateSummary", "input_obj": &cloudformation.GetTemplateSummaryInput{}},
			{"apicall": "GetTemplate", "input_obj": &cloudformation.GetTemplateInput{}},
		}}

	cloudfront_svc := &servicemaster.ServiceMaster{
		Svc:     cloudfront.NewFromConfig(cfg),
		SvcName: "cloudfront",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListFieldLevelEncryptionProfiles", "input_obj": &cloudfront.ListFieldLevelEncryptionProfilesInput{}},
			{"apicall": "ListCloudFrontOriginAccessIdentities", "input_obj": &cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}},
			{"apicall": "ListFieldLevelEncryptionConfigs", "input_obj": &cloudfront.ListFieldLevelEncryptionConfigsInput{}},
			{"apicall": "ListDistributions", "input_obj": &cloudfront.ListDistributionsInput{}},
			{"apicall": "ListStreamingDistributions", "input_obj": &cloudfront.ListStreamingDistributionsInput{}},
		}}

	cloudhsm_svc := &servicemaster.ServiceMaster{
		Svc:     cloudhsm.NewFromConfig(cfg),
		SvcName: "cloudhsm",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListHapgs", "input_obj": &cloudhsm.ListHapgsInput{}},
			{"apicall": "ListHsms", "input_obj": &cloudhsm.ListHsmsInput{}},
			{"apicall": "DescribeLunaClient", "input_obj": &cloudhsm.DescribeLunaClientInput{}},
			{"apicall": "ListAvailableZones", "input_obj": &cloudhsm.ListAvailableZonesInput{}},
			{"apicall": "DescribeHsm", "input_obj": &cloudhsm.DescribeHsmInput{}},
			{"apicall": "ListLunaClients", "input_obj": &cloudhsm.ListLunaClientsInput{}},
		}}

	cloudhsmv2_svc := &servicemaster.ServiceMaster{
		Svc:     cloudhsmv2.NewFromConfig(cfg),
		SvcName: "cloudhsmv2",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeBackups", "input_obj": &cloudhsmv2.DescribeBackupsInput{}},
			{"apicall": "DescribeClusters", "input_obj": &cloudhsmv2.DescribeClustersInput{}},
		}}

	cloudsearch_svc := &servicemaster.ServiceMaster{
		Svc:     cloudsearch.NewFromConfig(cfg),
		SvcName: "cloudsearch",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeDomains", "input_obj": &cloudsearch.DescribeDomainsInput{}},
			{"apicall": "ListDomainNames", "input_obj": &cloudsearch.ListDomainNamesInput{}},
		}}

	cloudtrail_svc := &servicemaster.ServiceMaster{
		Svc:     cloudtrail.NewFromConfig(cfg),
		SvcName: "cloudtrail",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeTrails", "input_obj": &cloudtrail.DescribeTrailsInput{}},
		}}

	codebuild_svc := &servicemaster.ServiceMaster{
		Svc:     codebuild.NewFromConfig(cfg),
		SvcName: "codebuild",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListBuilds", "input_obj": &codebuild.ListBuildsInput{}},
			{"apicall": "ListProjects", "input_obj": &codebuild.ListProjectsInput{}},
			{"apicall": "ListCuratedEnvironmentImages", "input_obj": &codebuild.ListCuratedEnvironmentImagesInput{}},
			{"apicall": "ListSourceCredentials", "input_obj": &codebuild.ListSourceCredentialsInput{}},
		}}

	codecommit_svc := &servicemaster.ServiceMaster{
		Svc:     codecommit.NewFromConfig(cfg),
		SvcName: "codecommit",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListRepositories", "input_obj": &codecommit.ListRepositoriesInput{}},
			{"apicall": "GetBranch", "input_obj": &codecommit.GetBranchInput{}},
		}}

	codedeploy_svc := &servicemaster.ServiceMaster{
		Svc:     codedeploy.NewFromConfig(cfg),
		SvcName: "codedeploy",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetDeploymentTarget", "input_obj": &codedeploy.GetDeploymentTargetInput{}},
			{"apicall": "ListGitHubAccountTokenNames", "input_obj": &codedeploy.ListGitHubAccountTokenNamesInput{}},
			{"apicall": "ListDeploymentConfigs", "input_obj": &codedeploy.ListDeploymentConfigsInput{}},
			{"apicall": "BatchGetDeploymentTargets", "input_obj": &codedeploy.BatchGetDeploymentTargetsInput{}},
			{"apicall": "ListDeploymentTargets", "input_obj": &codedeploy.ListDeploymentTargetsInput{}},
			{"apicall": "ListApplications", "input_obj": &codedeploy.ListApplicationsInput{}},
			{"apicall": "ListOnPremisesInstances", "input_obj": &codedeploy.ListOnPremisesInstancesInput{}},
			{"apicall": "ListDeployments", "input_obj": &codedeploy.ListDeploymentsInput{}},
		}}

	codepipeline_svc := &servicemaster.ServiceMaster{
		Svc:     codepipeline.NewFromConfig(cfg),
		SvcName: "codepipeline",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListActionTypes", "input_obj": &codepipeline.ListActionTypesInput{}},
			{"apicall": "ListPipelines", "input_obj": &codepipeline.ListPipelinesInput{}},
			{"apicall": "ListWebhooks", "input_obj": &codepipeline.ListWebhooksInput{}},
		}}

	codestar_svc := &servicemaster.ServiceMaster{
		Svc:     codestar.NewFromConfig(cfg),
		SvcName: "codestar",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListProjects", "input_obj": &codestar.ListProjectsInput{}},
			{"apicall": "ListUserProfiles", "input_obj": &codestar.ListUserProfilesInput{}},
		}}

	comprehend_svc := &servicemaster.ServiceMaster{
		Svc:     comprehend.NewFromConfig(cfg),
		SvcName: "comprehend",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListDominantLanguageDetectionJobs", "input_obj": &comprehend.ListDominantLanguageDetectionJobsInput{}},
			{"apicall": "ListTopicsDetectionJobs", "input_obj": &comprehend.ListTopicsDetectionJobsInput{}},
			{"apicall": "ListEntitiesDetectionJobs", "input_obj": &comprehend.ListEntitiesDetectionJobsInput{}},
			{"apicall": "ListEntityRecognizers", "input_obj": &comprehend.ListEntityRecognizersInput{}},
			{"apicall": "ListDocumentClassifiers", "input_obj": &comprehend.ListDocumentClassifiersInput{}},
			{"apicall": "ListSentimentDetectionJobs", "input_obj": &comprehend.ListSentimentDetectionJobsInput{}},
			{"apicall": "ListKeyPhrasesDetectionJobs", "input_obj": &comprehend.ListKeyPhrasesDetectionJobsInput{}},
			{"apicall": "ListDocumentClassificationJobs", "input_obj": &comprehend.ListDocumentClassificationJobsInput{}},
		}}

	datapipeline_svc := &servicemaster.ServiceMaster{
		Svc:     datapipeline.NewFromConfig(cfg),
		SvcName: "datapipeline",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListPipelines", "input_obj": &datapipeline.ListPipelinesInput{}},
		}}

	datasync_svc := &servicemaster.ServiceMaster{
		Svc:     datasync.NewFromConfig(cfg),
		SvcName: "datasync",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListLocations", "input_obj": &datasync.ListLocationsInput{}},
			{"apicall": "ListTaskExecutions", "input_obj": &datasync.ListTaskExecutionsInput{}},
			{"apicall": "ListTasks", "input_obj": &datasync.ListTasksInput{}},
			{"apicall": "ListAgents", "input_obj": &datasync.ListAgentsInput{}},
		}}

	dax_svc := &servicemaster.ServiceMaster{
		Svc:     dax.NewFromConfig(cfg),
		SvcName: "dax",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeDefaultParameters", "input_obj": &dax.DescribeDefaultParametersInput{}},
			{"apicall": "DescribeSubnetGroups", "input_obj": &dax.DescribeSubnetGroupsInput{}},
			{"apicall": "DescribeClusters", "input_obj": &dax.DescribeClustersInput{}},
			{"apicall": "DescribeParameterGroups", "input_obj": &dax.DescribeParameterGroupsInput{}},
		}}

	devicefarm_svc := &servicemaster.ServiceMaster{
		Svc:     devicefarm.NewFromConfig(cfg),
		SvcName: "devicefarm",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListProjects", "input_obj": &devicefarm.ListProjectsInput{}},
			{"apicall": "ListOfferingPromotions", "input_obj": &devicefarm.ListOfferingPromotionsInput{}},
			{"apicall": "ListOfferings", "input_obj": &devicefarm.ListOfferingsInput{}},
			{"apicall": "GetOfferingStatus", "input_obj": &devicefarm.GetOfferingStatusInput{}},
			{"apicall": "ListOfferingTransactions", "input_obj": &devicefarm.ListOfferingTransactionsInput{}},
			{"apicall": "ListVPCEConfigurations", "input_obj": &devicefarm.ListVPCEConfigurationsInput{}},
			{"apicall": "ListDeviceInstances", "input_obj": &devicefarm.ListDeviceInstancesInput{}},
			{"apicall": "ListInstanceProfiles", "input_obj": &devicefarm.ListInstanceProfilesInput{}},
			{"apicall": "ListDevices", "input_obj": &devicefarm.ListDevicesInput{}},
			{"apicall": "GetAccountSettings", "input_obj": &devicefarm.GetAccountSettingsInput{}},
		}}

	directconnect_svc := &servicemaster.ServiceMaster{
		Svc:     directconnect.NewFromConfig(cfg),
		SvcName: "directconnect",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeDirectConnectGatewayAssociations", "input_obj": &directconnect.DescribeDirectConnectGatewayAssociationsInput{}},
			{"apicall": "DescribeConnections", "input_obj": &directconnect.DescribeConnectionsInput{}},
			{"apicall": "DescribeDirectConnectGatewayAttachments", "input_obj": &directconnect.DescribeDirectConnectGatewayAttachmentsInput{}},
			{"apicall": "DescribeLags", "input_obj": &directconnect.DescribeLagsInput{}},
			{"apicall": "DescribeDirectConnectGateways", "input_obj": &directconnect.DescribeDirectConnectGatewaysInput{}},
			{"apicall": "DescribeLocations", "input_obj": &directconnect.DescribeLocationsInput{}},
			{"apicall": "DescribeVirtualGateways", "input_obj": &directconnect.DescribeVirtualGatewaysInput{}},
			{"apicall": "DescribeInterconnects", "input_obj": &directconnect.DescribeInterconnectsInput{}},
			{"apicall": "DescribeVirtualInterfaces", "input_obj": &directconnect.DescribeVirtualInterfacesInput{}},
		}}

	dlm_svc := &servicemaster.ServiceMaster{
		Svc:     dlm.NewFromConfig(cfg),
		SvcName: "dlm",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetLifecyclePolicies", "input_obj": &dlm.GetLifecyclePoliciesInput{}},
		}}

	dynamodb_svc := &servicemaster.ServiceMaster{
		Svc:     dynamodb.NewFromConfig(cfg),
		SvcName: "dynamodb",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListBackups", "input_obj": &dynamodb.ListBackupsInput{}},
			{"apicall": "ListTables", "input_obj": &dynamodb.ListTablesInput{}},
			{"apicall": "DescribeLimits", "input_obj": &dynamodb.DescribeLimitsInput{}},
			{"apicall": "ListGlobalTables", "input_obj": &dynamodb.ListGlobalTablesInput{}},
			{"apicall": "DescribeEndpoints", "input_obj": &dynamodb.DescribeEndpointsInput{}},
		}}

	ec2_svc := &servicemaster.ServiceMaster{
		Svc:     ec2.NewFromConfig(cfg),
		SvcName: "ec2",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeImages", "input_obj": &ec2.DescribeImagesInput{}},
			{"apicall": "DescribeSubnets", "input_obj": &ec2.DescribeSubnetsInput{}},
			{"apicall": "DescribeIamInstanceProfileAssociations", "input_obj": &ec2.DescribeIamInstanceProfileAssociationsInput{}},
			{"apicall": "DescribeNatGateways", "input_obj": &ec2.DescribeNatGatewaysInput{}},
			{"apicall": "DescribeSpotFleetRequests", "input_obj": &ec2.DescribeSpotFleetRequestsInput{}},
			{"apicall": "DescribeVpcEndpointConnections", "input_obj": &ec2.DescribeVpcEndpointConnectionsInput{}},
			{"apicall": "DescribeVpcEndpointServiceConfigurations", "input_obj": &ec2.DescribeVpcEndpointServiceConfigurationsInput{}},
			{"apicall": "DescribeVpcClassicLink", "input_obj": &ec2.DescribeVpcClassicLinkInput{}},
			{"apicall": "DescribeReservedInstances", "input_obj": &ec2.DescribeReservedInstancesInput{}},
			{"apicall": "DescribeReservedInstancesModifications", "input_obj": &ec2.DescribeReservedInstancesModificationsInput{}},
			{"apicall": "DescribeSpotPriceHistory", "input_obj": &ec2.DescribeSpotPriceHistoryInput{}},
			{"apicall": "DescribeTransitGatewayAttachments", "input_obj": &ec2.DescribeTransitGatewayAttachmentsInput{}},
			{"apicall": "DescribeReservedInstancesOfferings", "input_obj": &ec2.DescribeReservedInstancesOfferingsInput{}},
			{"apicall": "DescribeVpcEndpoints", "input_obj": &ec2.DescribeVpcEndpointsInput{}},
			{"apicall": "DescribeSnapshots", "input_obj": &ec2.DescribeSnapshotsInput{}},
			{"apicall": "DescribeNetworkAcls", "input_obj": &ec2.DescribeNetworkAclsInput{}},
			{"apicall": "DescribeRouteTables", "input_obj": &ec2.DescribeRouteTablesInput{}},
			{"apicall": "DescribeRegions", "input_obj": &ec2.DescribeRegionsInput{}},
			{"apicall": "DescribeInstances", "input_obj": &ec2.DescribeInstancesInput{}},
			{"apicall": "DescribeFleets", "input_obj": &ec2.DescribeFleetsInput{}},
			{"apicall": "DescribeInstanceCreditSpecifications", "input_obj": &ec2.DescribeInstanceCreditSpecificationsInput{}},
			{"apicall": "DescribeImportSnapshotTasks", "input_obj": &ec2.DescribeImportSnapshotTasksInput{}},
			{"apicall": "DescribeClientVpnEndpoints", "input_obj": &ec2.DescribeClientVpnEndpointsInput{}},
			{"apicall": "DescribeKeyPairs", "input_obj": &ec2.DescribeKeyPairsInput{}},
			{"apicall": "DescribeIdFormat", "input_obj": &ec2.DescribeIdFormatInput{}},
			{"apicall": "DescribePublicIpv4Pools", "input_obj": &ec2.DescribePublicIpv4PoolsInput{}},
			{"apicall": "DescribeScheduledInstances", "input_obj": &ec2.DescribeScheduledInstancesInput{}},
			{"apicall": "DescribeConversionTasks", "input_obj": &ec2.DescribeConversionTasksInput{}},
			{"apicall": "DescribeEgressOnlyInternetGateways", "input_obj": &ec2.DescribeEgressOnlyInternetGatewaysInput{}},
			{"apicall": "DescribePrincipalIdFormat", "input_obj": &ec2.DescribePrincipalIdFormatInput{}},
			{"apicall": "DescribeNetworkInterfacePermissions", "input_obj": &ec2.DescribeNetworkInterfacePermissionsInput{}},
			{"apicall": "DescribeVpnConnections", "input_obj": &ec2.DescribeVpnConnectionsInput{}},
			{"apicall": "DescribeVpcs", "input_obj": &ec2.DescribeVpcsInput{}},
			{"apicall": "DescribeNetworkInterfaces", "input_obj": &ec2.DescribeNetworkInterfacesInput{}},
			{"apicall": "DescribeSecurityGroups", "input_obj": &ec2.DescribeSecurityGroupsInput{}},
			{"apicall": "DescribeHostReservations", "input_obj": &ec2.DescribeHostReservationsInput{}},
			{"apicall": "DescribePrefixLists", "input_obj": &ec2.DescribePrefixListsInput{}},
			{"apicall": "DescribeVpnGateways", "input_obj": &ec2.DescribeVpnGatewaysInput{}},
			{"apicall": "DescribeImportImageTasks", "input_obj": &ec2.DescribeImportImageTasksInput{}},
			{"apicall": "DescribeClassicLinkInstances", "input_obj": &ec2.DescribeClassicLinkInstancesInput{}},
			{"apicall": "DescribeVpcPeeringConnections", "input_obj": &ec2.DescribeVpcPeeringConnectionsInput{}},
			{"apicall": "DescribeSpotDatafeedSubscription", "input_obj": &ec2.DescribeSpotDatafeedSubscriptionInput{}},
			{"apicall": "DescribeMovingAddresses", "input_obj": &ec2.DescribeMovingAddressesInput{}},
			{"apicall": "DescribeAvailabilityZones", "input_obj": &ec2.DescribeAvailabilityZonesInput{}},
			{"apicall": "DescribeAddresses", "input_obj": &ec2.DescribeAddressesInput{}},
			{"apicall": "DescribeElasticGpus", "input_obj": &ec2.DescribeElasticGpusInput{}},
			{"apicall": "DescribeHosts", "input_obj": &ec2.DescribeHostsInput{}},
			{"apicall": "DescribeTransitGateways", "input_obj": &ec2.DescribeTransitGatewaysInput{}},
			{"apicall": "DescribeLaunchTemplateVersions", "input_obj": &ec2.DescribeLaunchTemplateVersionsInput{}},
			{"apicall": "DescribeTags", "input_obj": &ec2.DescribeTagsInput{}},
			{"apicall": "DescribeDhcpOptions", "input_obj": &ec2.DescribeDhcpOptionsInput{}},
			{"apicall": "DescribeLaunchTemplates", "input_obj": &ec2.DescribeLaunchTemplatesInput{}},
			{"apicall": "DescribeVolumeStatus", "input_obj": &ec2.DescribeVolumeStatusInput{}},
			{"apicall": "DescribePlacementGroups", "input_obj": &ec2.DescribePlacementGroupsInput{}},
			{"apicall": "DescribeAggregateIdFormat", "input_obj": &ec2.DescribeAggregateIdFormatInput{}},
			{"apicall": "DescribeTransitGatewayRouteTables", "input_obj": &ec2.DescribeTransitGatewayRouteTablesInput{}},
			{"apicall": "DescribeInstanceStatus", "input_obj": &ec2.DescribeInstanceStatusInput{}},
			{"apicall": "DescribeCustomerGateways", "input_obj": &ec2.DescribeCustomerGatewaysInput{}},
			{"apicall": "DescribeInternetGateways", "input_obj": &ec2.DescribeInternetGatewaysInput{}},
			{"apicall": "DescribeHostReservationOfferings", "input_obj": &ec2.DescribeHostReservationOfferingsInput{}},
			{"apicall": "DescribeBundleTasks", "input_obj": &ec2.DescribeBundleTasksInput{}},
			{"apicall": "DescribeVolumesModifications", "input_obj": &ec2.DescribeVolumesModificationsInput{}},
			{"apicall": "DescribeExportTasks", "input_obj": &ec2.DescribeExportTasksInput{}},
			{"apicall": "DescribeVolumes", "input_obj": &ec2.DescribeVolumesInput{}},
			{"apicall": "DescribeFlowLogs", "input_obj": &ec2.DescribeFlowLogsInput{}},
			{"apicall": "DescribeSpotInstanceRequests", "input_obj": &ec2.DescribeSpotInstanceRequestsInput{}},
			{"apicall": "DescribeVpcEndpointServices", "input_obj": &ec2.DescribeVpcEndpointServicesInput{}},
			{"apicall": "DescribeTransitGatewayVpcAttachments", "input_obj": &ec2.DescribeTransitGatewayVpcAttachmentsInput{}},
			{"apicall": "DescribeFpgaImages", "input_obj": &ec2.DescribeFpgaImagesInput{}},
			{"apicall": "DescribeCapacityReservations", "input_obj": &ec2.DescribeCapacityReservationsInput{}},
			{"apicall": "DescribeVpcClassicLinkDnsSupport", "input_obj": &ec2.DescribeVpcClassicLinkDnsSupportInput{}},
			{"apicall": "DescribeReservedInstancesListings", "input_obj": &ec2.DescribeReservedInstancesListingsInput{}},
			{"apicall": "DescribeVpcEndpointConnectionNotifications", "input_obj": &ec2.DescribeVpcEndpointConnectionNotificationsInput{}},
			{"apicall": "DescribeAccountAttributes", "input_obj": &ec2.DescribeAccountAttributesInput{}},
		}}

	ecr_svc := &servicemaster.ServiceMaster{
		Svc:     ecr.NewFromConfig(cfg),
		SvcName: "ecr",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeRepositories", "input_obj": &ecr.DescribeRepositoriesInput{}},
			{"apicall": "GetAuthorizationToken", "input_obj": &ecr.GetAuthorizationTokenInput{}},
		}}

	ecs_svc := &servicemaster.ServiceMaster{
		Svc:     ecs.NewFromConfig(cfg),
		SvcName: "ecs",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListServices", "input_obj": &ecs.ListServicesInput{}},
			{"apicall": "DescribeClusters", "input_obj": &ecs.DescribeClustersInput{}},
			{"apicall": "ListClusters", "input_obj": &ecs.ListClustersInput{}},
			{"apicall": "ListTasks", "input_obj": &ecs.ListTasksInput{}},
			{"apicall": "ListTaskDefinitions", "input_obj": &ecs.ListTaskDefinitionsInput{}},
			{"apicall": "ListContainerInstances", "input_obj": &ecs.ListContainerInstancesInput{}},
			{"apicall": "ListAccountSettings", "input_obj": &ecs.ListAccountSettingsInput{}},
			{"apicall": "ListTaskDefinitionFamilies", "input_obj": &ecs.ListTaskDefinitionFamiliesInput{}},
		}}

	eks_svc := &servicemaster.ServiceMaster{
		Svc:     eks.NewFromConfig(cfg),
		SvcName: "eks",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListClusters", "input_obj": &eks.ListClustersInput{}},
		}}

	elasticache_svc := &servicemaster.ServiceMaster{
		Svc:     elasticache.NewFromConfig(cfg),
		SvcName: "elasticache",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeReservedCacheNodes", "input_obj": &elasticache.DescribeReservedCacheNodesInput{}},
			{"apicall": "DescribeReservedCacheNodesOfferings", "input_obj": &elasticache.DescribeReservedCacheNodesOfferingsInput{}},
			{"apicall": "DescribeCacheSubnetGroups", "input_obj": &elasticache.DescribeCacheSubnetGroupsInput{}},
			{"apicall": "DescribeCacheEngineVersions", "input_obj": &elasticache.DescribeCacheEngineVersionsInput{}},
			{"apicall": "ListAllowedNodeTypeModifications", "input_obj": &elasticache.ListAllowedNodeTypeModificationsInput{}},
			{"apicall": "DescribeCacheSecurityGroups", "input_obj": &elasticache.DescribeCacheSecurityGroupsInput{}},
			{"apicall": "DescribeCacheParameterGroups", "input_obj": &elasticache.DescribeCacheParameterGroupsInput{}},
			{"apicall": "DescribeReplicationGroups", "input_obj": &elasticache.DescribeReplicationGroupsInput{}},
			{"apicall": "DescribeSnapshots", "input_obj": &elasticache.DescribeSnapshotsInput{}},
			{"apicall": "DescribeCacheClusters", "input_obj": &elasticache.DescribeCacheClustersInput{}},
		}}

	elasticbeanstalk_svc := &servicemaster.ServiceMaster{
		Svc:     elasticbeanstalk.NewFromConfig(cfg),
		SvcName: "elasticbeanstalk",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeEnvironmentManagedActionHistory", "input_obj": &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput{}},
			{"apicall": "DescribePlatformVersion", "input_obj": &elasticbeanstalk.DescribePlatformVersionInput{}},
			{"apicall": "DescribeInstancesHealth", "input_obj": &elasticbeanstalk.DescribeInstancesHealthInput{}},
			{"apicall": "DescribeEnvironmentHealth", "input_obj": &elasticbeanstalk.DescribeEnvironmentHealthInput{}},
			{"apicall": "DescribeConfigurationOptions", "input_obj": &elasticbeanstalk.DescribeConfigurationOptionsInput{}},
			{"apicall": "DescribeEnvironmentManagedActions", "input_obj": &elasticbeanstalk.DescribeEnvironmentManagedActionsInput{}},
			{"apicall": "DescribeEnvironmentResources", "input_obj": &elasticbeanstalk.DescribeEnvironmentResourcesInput{}},
			{"apicall": "DescribeAccountAttributes", "input_obj": &elasticbeanstalk.DescribeAccountAttributesInput{}},
		}}

	elastictranscoder_svc := &servicemaster.ServiceMaster{
		Svc:     elastictranscoder.NewFromConfig(cfg),
		SvcName: "elastictranscoder",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListPipelines", "input_obj": &elastictranscoder.ListPipelinesInput{}},
			{"apicall": "ListPresets", "input_obj": &elastictranscoder.ListPresetsInput{}},
		}}

	firehose_svc := &servicemaster.ServiceMaster{
		Svc:     firehose.NewFromConfig(cfg),
		SvcName: "firehose",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListDeliveryStreams", "input_obj": &firehose.ListDeliveryStreamsInput{}},
		}}

	fms_svc := &servicemaster.ServiceMaster{
		Svc:     fms.NewFromConfig(cfg),
		SvcName: "fms",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetNotificationChannel", "input_obj": &fms.GetNotificationChannelInput{}},
			{"apicall": "ListMemberAccounts", "input_obj": &fms.ListMemberAccountsInput{}},
			{"apicall": "ListPolicies", "input_obj": &fms.ListPoliciesInput{}},
			{"apicall": "GetAdminAccount", "input_obj": &fms.GetAdminAccountInput{}},
		}}

	fsx_svc := &servicemaster.ServiceMaster{
		Svc:     fsx.NewFromConfig(cfg),
		SvcName: "fsx",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeFileSystems", "input_obj": &fsx.DescribeFileSystemsInput{}},
			{"apicall": "DescribeBackups", "input_obj": &fsx.DescribeBackupsInput{}},
		}}

	gamelift_svc := &servicemaster.ServiceMaster{
		Svc:     gamelift.NewFromConfig(cfg),
		SvcName: "gamelift",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribePlayerSessions", "input_obj": &gamelift.DescribePlayerSessionsInput{}},
			{"apicall": "DescribeGameSessionDetails", "input_obj": &gamelift.DescribeGameSessionDetailsInput{}},
			{"apicall": "DescribeMatchmakingRuleSets", "input_obj": &gamelift.DescribeMatchmakingRuleSetsInput{}},
			{"apicall": "DescribeFleetUtilization", "input_obj": &gamelift.DescribeFleetUtilizationInput{}},
			{"apicall": "DescribeGameSessions", "input_obj": &gamelift.DescribeGameSessionsInput{}},
			{"apicall": "DescribeVpcPeeringAuthorizations", "input_obj": &gamelift.DescribeVpcPeeringAuthorizationsInput{}},
			{"apicall": "ListAliases", "input_obj": &gamelift.ListAliasesInput{}},
			{"apicall": "DescribeVpcPeeringConnections", "input_obj": &gamelift.DescribeVpcPeeringConnectionsInput{}},
			{"apicall": "DescribeMatchmakingConfigurations", "input_obj": &gamelift.DescribeMatchmakingConfigurationsInput{}},
			{"apicall": "DescribeFleetAttributes", "input_obj": &gamelift.DescribeFleetAttributesInput{}},
			{"apicall": "ListFleets", "input_obj": &gamelift.ListFleetsInput{}},
			{"apicall": "ListBuilds", "input_obj": &gamelift.ListBuildsInput{}},
			{"apicall": "DescribeGameSessionQueues", "input_obj": &gamelift.DescribeGameSessionQueuesInput{}},
			{"apicall": "DescribeEC2InstanceLimits", "input_obj": &gamelift.DescribeEC2InstanceLimitsInput{}},
			{"apicall": "DescribeFleetCapacity", "input_obj": &gamelift.DescribeFleetCapacityInput{}},
		}}

	globalaccelerator_svc := &servicemaster.ServiceMaster{
		Svc:     globalaccelerator.NewFromConfig(cfg),
		SvcName: "globalaccelerator",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListAccelerators", "input_obj": &globalaccelerator.ListAcceleratorsInput{}},
			{"apicall": "DescribeAcceleratorAttributes", "input_obj": &globalaccelerator.DescribeAcceleratorAttributesInput{}},
		}}

	glue_svc := &servicemaster.ServiceMaster{
		Svc:     glue.NewFromConfig(cfg),
		SvcName: "glue",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetClassifiers", "input_obj": &glue.GetClassifiersInput{}},
			{"apicall": "GetDatabases", "input_obj": &glue.GetDatabasesInput{}},
			{"apicall": "GetCrawlers", "input_obj": &glue.GetCrawlersInput{}},
			{"apicall": "ListDevEndpoints", "input_obj": &glue.ListDevEndpointsInput{}},
			{"apicall": "ListTriggers", "input_obj": &glue.ListTriggersInput{}},
			{"apicall": "GetTriggers", "input_obj": &glue.GetTriggersInput{}},
			{"apicall": "ListCrawlers", "input_obj": &glue.ListCrawlersInput{}},
			{"apicall": "ListJobs", "input_obj": &glue.ListJobsInput{}},
			{"apicall": "GetDataCatalogEncryptionSettings", "input_obj": &glue.GetDataCatalogEncryptionSettingsInput{}},
			{"apicall": "GetConnections", "input_obj": &glue.GetConnectionsInput{}},
			{"apicall": "GetJobs", "input_obj": &glue.GetJobsInput{}},
			{"apicall": "GetCrawlerMetrics", "input_obj": &glue.GetCrawlerMetricsInput{}},
			{"apicall": "GetDataflowGraph", "input_obj": &glue.GetDataflowGraphInput{}},
			{"apicall": "GetResourcePolicy", "input_obj": &glue.GetResourcePolicyInput{}},
			{"apicall": "GetSecurityConfigurations", "input_obj": &glue.GetSecurityConfigurationsInput{}},
			{"apicall": "GetCatalogImportStatus", "input_obj": &glue.GetCatalogImportStatusInput{}},
			{"apicall": "GetDevEndpoints", "input_obj": &glue.GetDevEndpointsInput{}},
		}}

	greengrass_svc := &servicemaster.ServiceMaster{
		Svc:     greengrass.NewFromConfig(cfg),
		SvcName: "greengrass",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListDeviceDefinitions", "input_obj": &greengrass.ListDeviceDefinitionsInput{}},
			{"apicall": "GetServiceRoleForAccount", "input_obj": &greengrass.GetServiceRoleForAccountInput{}},
			{"apicall": "ListCoreDefinitions", "input_obj": &greengrass.ListCoreDefinitionsInput{}},
			{"apicall": "ListLoggerDefinitions", "input_obj": &greengrass.ListLoggerDefinitionsInput{}},
			{"apicall": "ListBulkDeployments", "input_obj": &greengrass.ListBulkDeploymentsInput{}},
			{"apicall": "ListConnectorDefinitions", "input_obj": &greengrass.ListConnectorDefinitionsInput{}},
			{"apicall": "ListGroups", "input_obj": &greengrass.ListGroupsInput{}},
			{"apicall": "ListSubscriptionDefinitions", "input_obj": &greengrass.ListSubscriptionDefinitionsInput{}},
			{"apicall": "ListFunctionDefinitions", "input_obj": &greengrass.ListFunctionDefinitionsInput{}},
			{"apicall": "ListResourceDefinitions", "input_obj": &greengrass.ListResourceDefinitionsInput{}},
		}}

	guardduty_svc := &servicemaster.ServiceMaster{
		Svc:     guardduty.NewFromConfig(cfg),
		SvcName: "guardduty",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListInvitations", "input_obj": &guardduty.ListInvitationsInput{}},
			{"apicall": "GetInvitationsCount", "input_obj": &guardduty.GetInvitationsCountInput{}},
			{"apicall": "ListDetectors", "input_obj": &guardduty.ListDetectorsInput{}},
		}}

	health_svc := &servicemaster.ServiceMaster{
		Svc:     health.NewFromConfig(cfg),
		SvcName: "health",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeEntityAggregates", "input_obj": &health.DescribeEntityAggregatesInput{}},
			{"apicall": "DescribeEventTypes", "input_obj": &health.DescribeEventTypesInput{}},
		}}

	iam_svc := &servicemaster.ServiceMaster{
		Svc:     iam.NewFromConfig(cfg),
		SvcName: "iam",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListRoles", "input_obj": &iam.ListRolesInput{}},
			{"apicall": "ListAccessKeys", "input_obj": &iam.ListAccessKeysInput{}},
			{"apicall": "ListGroups", "input_obj": &iam.ListGroupsInput{}},
			{"apicall": "ListOpenIDConnectProviders", "input_obj": &iam.ListOpenIDConnectProvidersInput{}},
			{"apicall": "GetUser", "input_obj": &iam.GetUserInput{}},
			{"apicall": "ListSAMLProviders", "input_obj": &iam.ListSAMLProvidersInput{}},
			{"apicall": "ListAccountAliases", "input_obj": &iam.ListAccountAliasesInput{}},
			{"apicall": "GetAccountSummary", "input_obj": &iam.GetAccountSummaryInput{}},
			{"apicall": "ListMFADevices", "input_obj": &iam.ListMFADevicesInput{}},
			{"apicall": "ListServerCertificates", "input_obj": &iam.ListServerCertificatesInput{}},
			{"apicall": "ListServiceSpecificCredentials", "input_obj": &iam.ListServiceSpecificCredentialsInput{}},
			{"apicall": "GetAccountAuthorizationDetails", "input_obj": &iam.GetAccountAuthorizationDetailsInput{}},
			{"apicall": "ListSSHPublicKeys", "input_obj": &iam.ListSSHPublicKeysInput{}},
			{"apicall": "ListSigningCertificates", "input_obj": &iam.ListSigningCertificatesInput{}},
			{"apicall": "ListPolicies", "input_obj": &iam.ListPoliciesInput{}},
			{"apicall": "ListVirtualMFADevices", "input_obj": &iam.ListVirtualMFADevicesInput{}},
			{"apicall": "ListInstanceProfiles", "input_obj": &iam.ListInstanceProfilesInput{}},
			{"apicall": "ListUsers", "input_obj": &iam.ListUsersInput{}},
			{"apicall": "GetCredentialReport", "input_obj": &iam.GetCredentialReportInput{}},
			{"apicall": "GetAccountPasswordPolicy", "input_obj": &iam.GetAccountPasswordPolicyInput{}},
		}}

	inspector_svc := &servicemaster.ServiceMaster{
		Svc:     inspector.NewFromConfig(cfg),
		SvcName: "inspector",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListEventSubscriptions", "input_obj": &inspector.ListEventSubscriptionsInput{}},
			{"apicall": "DescribeCrossAccountAccessRole", "input_obj": &inspector.DescribeCrossAccountAccessRoleInput{}},
			{"apicall": "ListAssessmentTemplates", "input_obj": &inspector.ListAssessmentTemplatesInput{}},
			{"apicall": "ListRulesPackages", "input_obj": &inspector.ListRulesPackagesInput{}},
			{"apicall": "ListAssessmentRuns", "input_obj": &inspector.ListAssessmentRunsInput{}},
			{"apicall": "ListFindings", "input_obj": &inspector.ListFindingsInput{}},
			{"apicall": "ListAssessmentTargets", "input_obj": &inspector.ListAssessmentTargetsInput{}},
		}}

	iot_svc := &servicemaster.ServiceMaster{
		Svc:     iot.NewFromConfig(cfg),
		SvcName: "iot",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListCertificates", "input_obj": &iot.ListCertificatesInput{}},
			{"apicall": "ListStreams", "input_obj": &iot.ListStreamsInput{}},
			{"apicall": "DescribeEventConfigurations", "input_obj": &iot.DescribeEventConfigurationsInput{}},
			{"apicall": "ListRoleAliases", "input_obj": &iot.ListRoleAliasesInput{}},
			{"apicall": "ListOutgoingCertificates", "input_obj": &iot.ListOutgoingCertificatesInput{}},
			{"apicall": "ListV2LoggingLevels", "input_obj": &iot.ListV2LoggingLevelsInput{}},
			{"apicall": "DescribeDefaultAuthorizer", "input_obj": &iot.DescribeDefaultAuthorizerInput{}},
			{"apicall": "ListThings", "input_obj": &iot.ListThingsInput{}},
			{"apicall": "GetLoggingOptions", "input_obj": &iot.GetLoggingOptionsInput{}},
			{"apicall": "DescribeEndpoint", "input_obj": &iot.DescribeEndpointInput{}},
			{"apicall": "ListCACertificates", "input_obj": &iot.ListCACertificatesInput{}},
			{"apicall": "GetIndexingConfiguration", "input_obj": &iot.GetIndexingConfigurationInput{}},
			{"apicall": "ListJobs", "input_obj": &iot.ListJobsInput{}},
			{"apicall": "ListActiveViolations", "input_obj": &iot.ListActiveViolationsInput{}},
			{"apicall": "ListAuditFindings", "input_obj": &iot.ListAuditFindingsInput{}},
			{"apicall": "ListIndices", "input_obj": &iot.ListIndicesInput{}},
			{"apicall": "ListThingTypes", "input_obj": &iot.ListThingTypesInput{}},
			{"apicall": "GetEffectivePolicies", "input_obj": &iot.GetEffectivePoliciesInput{}},
			{"apicall": "DescribeAccountAuditConfiguration", "input_obj": &iot.DescribeAccountAuditConfigurationInput{}},
			{"apicall": "ListOTAUpdates", "input_obj": &iot.ListOTAUpdatesInput{}},
			{"apicall": "ListBillingGroups", "input_obj": &iot.ListBillingGroupsInput{}},
			{"apicall": "ListAuthorizers", "input_obj": &iot.ListAuthorizersInput{}},
			{"apicall": "ListThingGroups", "input_obj": &iot.ListThingGroupsInput{}},
			{"apicall": "GetRegistrationCode", "input_obj": &iot.GetRegistrationCodeInput{}},
			{"apicall": "ListThingRegistrationTasks", "input_obj": &iot.ListThingRegistrationTasksInput{}},
			{"apicall": "ListScheduledAudits", "input_obj": &iot.ListScheduledAuditsInput{}},
			{"apicall": "GetV2LoggingOptions", "input_obj": &iot.GetV2LoggingOptionsInput{}},
			{"apicall": "ListPolicies", "input_obj": &iot.ListPoliciesInput{}},
			{"apicall": "ListTopicRules", "input_obj": &iot.ListTopicRulesInput{}},
			{"apicall": "ListSecurityProfiles", "input_obj": &iot.ListSecurityProfilesInput{}},
		}}

	iotanalytics_svc := &servicemaster.ServiceMaster{
		Svc:     iotanalytics.NewFromConfig(cfg),
		SvcName: "iotanalytics",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeLoggingOptions", "input_obj": &iotanalytics.DescribeLoggingOptionsInput{}},
			{"apicall": "ListDatasets", "input_obj": &iotanalytics.ListDatasetsInput{}},
			{"apicall": "ListChannels", "input_obj": &iotanalytics.ListChannelsInput{}},
			{"apicall": "ListPipelines", "input_obj": &iotanalytics.ListPipelinesInput{}},
			{"apicall": "ListDatastores", "input_obj": &iotanalytics.ListDatastoresInput{}},
		}}

	kafka_svc := &servicemaster.ServiceMaster{
		Svc:     kafka.NewFromConfig(cfg),
		SvcName: "kafka",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListClusters", "input_obj": &kafka.ListClustersInput{}},
		}}

	kinesis_svc := &servicemaster.ServiceMaster{
		Svc:     kinesis.NewFromConfig(cfg),
		SvcName: "kinesis",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeStreamConsumer", "input_obj": &kinesis.DescribeStreamConsumerInput{}},
			{"apicall": "ListStreams", "input_obj": &kinesis.ListStreamsInput{}},
			{"apicall": "ListShards", "input_obj": &kinesis.ListShardsInput{}},
			{"apicall": "DescribeLimits", "input_obj": &kinesis.DescribeLimitsInput{}},
		}}

	kinesisanalytics_svc := &servicemaster.ServiceMaster{
		Svc:     kinesisanalytics.NewFromConfig(cfg),
		SvcName: "kinesisanalytics",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListApplications", "input_obj": &kinesisanalytics.ListApplicationsInput{}},
		}}

	kinesisvideo_svc := &servicemaster.ServiceMaster{
		Svc:     kinesisvideo.NewFromConfig(cfg),
		SvcName: "kinesisvideo",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListStreams", "input_obj": &kinesisvideo.ListStreamsInput{}},
			{"apicall": "ListTagsForStream", "input_obj": &kinesisvideo.ListTagsForStreamInput{}},
			{"apicall": "DescribeStream", "input_obj": &kinesisvideo.DescribeStreamInput{}},
		}}

	kms_svc := &servicemaster.ServiceMaster{
		Svc:     kms.NewFromConfig(cfg),
		SvcName: "kms",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListKeys", "input_obj": &kms.ListKeysInput{}},
			{"apicall": "ListAliases", "input_obj": &kms.ListAliasesInput{}},
			{"apicall": "DescribeCustomKeyStores", "input_obj": &kms.DescribeCustomKeyStoresInput{}},
		}}

	lambda_svc := &servicemaster.ServiceMaster{
		Svc:     lambda.NewFromConfig(cfg),
		SvcName: "lambda",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListLayers", "input_obj": &lambda.ListLayersInput{}},
			{"apicall": "ListEventSourceMappings", "input_obj": &lambda.ListEventSourceMappingsInput{}},
			{"apicall": "GetAccountSettings", "input_obj": &lambda.GetAccountSettingsInput{}},
			{"apicall": "ListFunctions", "input_obj": &lambda.ListFunctionsInput{}},
		}}

	lightsail_svc := &servicemaster.ServiceMaster{
		Svc:     lightsail.NewFromConfig(cfg),
		SvcName: "lightsail",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetInstanceSnapshots", "input_obj": &lightsail.GetInstanceSnapshotsInput{}},
			{"apicall": "GetRelationalDatabaseSnapshots", "input_obj": &lightsail.GetRelationalDatabaseSnapshotsInput{}},
			{"apicall": "GetActiveNames", "input_obj": &lightsail.GetActiveNamesInput{}},
			{"apicall": "GetCloudFormationStackRecords", "input_obj": &lightsail.GetCloudFormationStackRecordsInput{}},
			{"apicall": "GetRelationalDatabases", "input_obj": &lightsail.GetRelationalDatabasesInput{}},
			{"apicall": "GetKeyPairs", "input_obj": &lightsail.GetKeyPairsInput{}},
			{"apicall": "GetLoadBalancers", "input_obj": &lightsail.GetLoadBalancersInput{}},
			{"apicall": "GetInstances", "input_obj": &lightsail.GetInstancesInput{}},
			{"apicall": "GetRegions", "input_obj": &lightsail.GetRegionsInput{}},
			{"apicall": "GetExportSnapshotRecords", "input_obj": &lightsail.GetExportSnapshotRecordsInput{}},
			{"apicall": "GetRelationalDatabaseBlueprints", "input_obj": &lightsail.GetRelationalDatabaseBlueprintsInput{}},
			{"apicall": "GetRelationalDatabaseBundles", "input_obj": &lightsail.GetRelationalDatabaseBundlesInput{}},
			{"apicall": "GetOperations", "input_obj": &lightsail.GetOperationsInput{}},
			{"apicall": "GetBundles", "input_obj": &lightsail.GetBundlesInput{}},
			{"apicall": "GetBlueprints", "input_obj": &lightsail.GetBlueprintsInput{}},
			{"apicall": "GetDisks", "input_obj": &lightsail.GetDisksInput{}},
			{"apicall": "GetDomains", "input_obj": &lightsail.GetDomainsInput{}},
			{"apicall": "GetStaticIps", "input_obj": &lightsail.GetStaticIpsInput{}},
			{"apicall": "GetDiskSnapshots", "input_obj": &lightsail.GetDiskSnapshotsInput{}},
		}}

	machinelearning_svc := &servicemaster.ServiceMaster{
		Svc:     machinelearning.NewFromConfig(cfg),
		SvcName: "machinelearning",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeMLModels", "input_obj": &machinelearning.DescribeMLModelsInput{}},
			{"apicall": "DescribeDataSources", "input_obj": &machinelearning.DescribeDataSourcesInput{}},
			{"apicall": "DescribeEvaluations", "input_obj": &machinelearning.DescribeEvaluationsInput{}},
			{"apicall": "DescribeBatchPredictions", "input_obj": &machinelearning.DescribeBatchPredictionsInput{}},
		}}

	macie_svc := &servicemaster.ServiceMaster{
		Svc:     macie.NewFromConfig(cfg),
		SvcName: "macie",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListS3Resources", "input_obj": &macie.ListS3ResourcesInput{}},
			{"apicall": "ListMemberAccounts", "input_obj": &macie.ListMemberAccountsInput{}},
		}}

	mediaconnect_svc := &servicemaster.ServiceMaster{
		Svc:     mediaconnect.NewFromConfig(cfg),
		SvcName: "mediaconnect",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListEntitlements", "input_obj": &mediaconnect.ListEntitlementsInput{}},
			{"apicall": "ListFlows", "input_obj": &mediaconnect.ListFlowsInput{}},
		}}

	mediaconvert_svc := &servicemaster.ServiceMaster{
		Svc:     mediaconvert.NewFromConfig(cfg),
		SvcName: "mediaconvert",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListPresets", "input_obj": &mediaconvert.ListPresetsInput{}},
			{"apicall": "ListJobs", "input_obj": &mediaconvert.ListJobsInput{}},
			{"apicall": "ListJobTemplates", "input_obj": &mediaconvert.ListJobTemplatesInput{}},
			{"apicall": "ListQueues", "input_obj": &mediaconvert.ListQueuesInput{}},
			{"apicall": "DescribeEndpoints", "input_obj": &mediaconvert.DescribeEndpointsInput{}},
		}}

	medialive_svc := &servicemaster.ServiceMaster{
		Svc:     medialive.NewFromConfig(cfg),
		SvcName: "medialive",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListInputs", "input_obj": &medialive.ListInputsInput{}},
			{"apicall": "ListChannels", "input_obj": &medialive.ListChannelsInput{}},
			{"apicall": "ListOfferings", "input_obj": &medialive.ListOfferingsInput{}},
			{"apicall": "ListReservations", "input_obj": &medialive.ListReservationsInput{}},
			{"apicall": "ListInputSecurityGroups", "input_obj": &medialive.ListInputSecurityGroupsInput{}},
		}}

	mediapackage_svc := &servicemaster.ServiceMaster{
		Svc:     mediapackage.NewFromConfig(cfg),
		SvcName: "mediapackage",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListChannels", "input_obj": &mediapackage.ListChannelsInput{}},
			{"apicall": "ListOriginEndpoints", "input_obj": &mediapackage.ListOriginEndpointsInput{}},
		}}

	mediastore_svc := &servicemaster.ServiceMaster{
		Svc:     mediastore.NewFromConfig(cfg),
		SvcName: "mediastore",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeContainer", "input_obj": &mediastore.DescribeContainerInput{}},
			{"apicall": "ListContainers", "input_obj": &mediastore.ListContainersInput{}},
		}}

	mediatailor_svc := &servicemaster.ServiceMaster{
		Svc:     mediatailor.NewFromConfig(cfg),
		SvcName: "mediatailor",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListPlaybackConfigurations", "input_obj": &mediatailor.ListPlaybackConfigurationsInput{}},
		}}

	mobile_svc := &servicemaster.ServiceMaster{
		Svc:     mobile.NewFromConfig(cfg),
		SvcName: "mobile",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListBundles", "input_obj": &mobile.ListBundlesInput{}},
			{"apicall": "ListProjects", "input_obj": &mobile.ListProjectsInput{}},
		}}

	mq_svc := &servicemaster.ServiceMaster{
		Svc:     mq.NewFromConfig(cfg),
		SvcName: "mq",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListConfigurations", "input_obj": &mq.ListConfigurationsInput{}},
			{"apicall": "ListBrokers", "input_obj": &mq.ListBrokersInput{}},
		}}

	opsworks_svc := &servicemaster.ServiceMaster{
		Svc:     opsworks.NewFromConfig(cfg),
		SvcName: "opsworks",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeMyUserProfile", "input_obj": &opsworks.DescribeMyUserProfileInput{}},
			{"apicall": "DescribeRaidArrays", "input_obj": &opsworks.DescribeRaidArraysInput{}},
			{"apicall": "DescribeUserProfiles", "input_obj": &opsworks.DescribeUserProfilesInput{}},
			{"apicall": "DescribeOperatingSystems", "input_obj": &opsworks.DescribeOperatingSystemsInput{}},
			{"apicall": "DescribeElasticLoadBalancers", "input_obj": &opsworks.DescribeElasticLoadBalancersInput{}},
			{"apicall": "DescribePermissions", "input_obj": &opsworks.DescribePermissionsInput{}},
			{"apicall": "DescribeVolumes", "input_obj": &opsworks.DescribeVolumesInput{}},
			{"apicall": "DescribeDeployments", "input_obj": &opsworks.DescribeDeploymentsInput{}},
			{"apicall": "DescribeEcsClusters", "input_obj": &opsworks.DescribeEcsClustersInput{}},
			{"apicall": "DescribeElasticIps", "input_obj": &opsworks.DescribeElasticIpsInput{}},
			{"apicall": "DescribeAgentVersions", "input_obj": &opsworks.DescribeAgentVersionsInput{}},
			{"apicall": "DescribeLayers", "input_obj": &opsworks.DescribeLayersInput{}},
			{"apicall": "DescribeApps", "input_obj": &opsworks.DescribeAppsInput{}},
			{"apicall": "DescribeCommands", "input_obj": &opsworks.DescribeCommandsInput{}},
			{"apicall": "DescribeInstances", "input_obj": &opsworks.DescribeInstancesInput{}},
		}}

	organizations_svc := &servicemaster.ServiceMaster{
		Svc:     organizations.NewFromConfig(cfg),
		SvcName: "organizations",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListAWSServiceAccessForOrganization", "input_obj": &organizations.ListAWSServiceAccessForOrganizationInput{}},
			{"apicall": "ListRoots", "input_obj": &organizations.ListRootsInput{}},
			{"apicall": "ListAccounts", "input_obj": &organizations.ListAccountsInput{}},
			{"apicall": "ListCreateAccountStatus", "input_obj": &organizations.ListCreateAccountStatusInput{}},
			{"apicall": "DescribeOrganization", "input_obj": &organizations.DescribeOrganizationInput{}},
			{"apicall": "ListHandshakesForAccount", "input_obj": &organizations.ListHandshakesForAccountInput{}},
			{"apicall": "ListHandshakesForOrganization", "input_obj": &organizations.ListHandshakesForOrganizationInput{}},
		}}

	pinpoint_svc := &servicemaster.ServiceMaster{
		Svc:     pinpoint.NewFromConfig(cfg),
		SvcName: "pinpoint",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetApps", "input_obj": &pinpoint.GetAppsInput{}},
		}}

	polly_svc := &servicemaster.ServiceMaster{
		Svc:     polly.NewFromConfig(cfg),
		SvcName: "polly",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListSpeechSynthesisTasks", "input_obj": &polly.ListSpeechSynthesisTasksInput{}},
			{"apicall": "DescribeVoices", "input_obj": &polly.DescribeVoicesInput{}},
			{"apicall": "ListLexicons", "input_obj": &polly.ListLexiconsInput{}},
		}}

	pricing_svc := &servicemaster.ServiceMaster{
		Svc:     pricing.NewFromConfig(cfg),
		SvcName: "pricing",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeServices", "input_obj": &pricing.DescribeServicesInput{}},
		}}

	ram_svc := &servicemaster.ServiceMaster{
		Svc:     ram.NewFromConfig(cfg),
		SvcName: "ram",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetResourceShareInvitations", "input_obj": &ram.GetResourceShareInvitationsInput{}},
		}}

	rds_svc := &servicemaster.ServiceMaster{
		Svc:     rds.NewFromConfig(cfg),
		SvcName: "rds",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeGlobalClusters", "input_obj": &rds.DescribeGlobalClustersInput{}},
			{"apicall": "DescribeDBClusterEndpoints", "input_obj": &rds.DescribeDBClusterEndpointsInput{}},
			{"apicall": "DescribeDBInstances", "input_obj": &rds.DescribeDBInstancesInput{}},
			{"apicall": "DescribeDBSecurityGroups", "input_obj": &rds.DescribeDBSecurityGroupsInput{}},
			{"apicall": "DescribePendingMaintenanceActions", "input_obj": &rds.DescribePendingMaintenanceActionsInput{}},
			{"apicall": "DescribeSourceRegions", "input_obj": &rds.DescribeSourceRegionsInput{}},
			{"apicall": "DescribeDBEngineVersions", "input_obj": &rds.DescribeDBEngineVersionsInput{}},
			{"apicall": "DescribeDBClusterSnapshots", "input_obj": &rds.DescribeDBClusterSnapshotsInput{}},
			{"apicall": "DescribeReservedDBInstances", "input_obj": &rds.DescribeReservedDBInstancesInput{}},
			{"apicall": "DescribeDBClusterParameterGroups", "input_obj": &rds.DescribeDBClusterParameterGroupsInput{}},
			{"apicall": "DescribeDBSubnetGroups", "input_obj": &rds.DescribeDBSubnetGroupsInput{}},
			{"apicall": "DescribeCertificates", "input_obj": &rds.DescribeCertificatesInput{}},
			{"apicall": "DescribeDBInstanceAutomatedBackups", "input_obj": &rds.DescribeDBInstanceAutomatedBackupsInput{}},
			{"apicall": "DescribeDBParameterGroups", "input_obj": &rds.DescribeDBParameterGroupsInput{}},
			{"apicall": "DescribeOptionGroups", "input_obj": &rds.DescribeOptionGroupsInput{}},
			{"apicall": "DescribeDBClusters", "input_obj": &rds.DescribeDBClustersInput{}},
			{"apicall": "DescribeReservedDBInstancesOfferings", "input_obj": &rds.DescribeReservedDBInstancesOfferingsInput{}},
			{"apicall": "DescribeEventCategories", "input_obj": &rds.DescribeEventCategoriesInput{}},
			{"apicall": "DescribeEventSubscriptions", "input_obj": &rds.DescribeEventSubscriptionsInput{}},
			{"apicall": "DescribeDBSnapshots", "input_obj": &rds.DescribeDBSnapshotsInput{}},
			{"apicall": "DescribeAccountAttributes", "input_obj": &rds.DescribeAccountAttributesInput{}},
		}}

	redshift_svc := &servicemaster.ServiceMaster{
		Svc:     redshift.NewFromConfig(cfg),
		SvcName: "redshift",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeHsmClientCertificates", "input_obj": &redshift.DescribeHsmClientCertificatesInput{}},
			{"apicall": "DescribeClusterSubnetGroups", "input_obj": &redshift.DescribeClusterSubnetGroupsInput{}},
			{"apicall": "DescribeTableRestoreStatus", "input_obj": &redshift.DescribeTableRestoreStatusInput{}},
			{"apicall": "DescribeSnapshotSchedules", "input_obj": &redshift.DescribeSnapshotSchedulesInput{}},
			{"apicall": "DescribeClusterTracks", "input_obj": &redshift.DescribeClusterTracksInput{}},
			{"apicall": "DescribeClusterSecurityGroups", "input_obj": &redshift.DescribeClusterSecurityGroupsInput{}},
			{"apicall": "DescribeClusters", "input_obj": &redshift.DescribeClustersInput{}},
			{"apicall": "DescribeReservedNodeOfferings", "input_obj": &redshift.DescribeReservedNodeOfferingsInput{}},
			{"apicall": "DescribeOrderableClusterOptions", "input_obj": &redshift.DescribeOrderableClusterOptionsInput{}},
			{"apicall": "DescribeClusterDbRevisions", "input_obj": &redshift.DescribeClusterDbRevisionsInput{}},
			{"apicall": "DescribeClusterParameterGroups", "input_obj": &redshift.DescribeClusterParameterGroupsInput{}},
			{"apicall": "DescribeTags", "input_obj": &redshift.DescribeTagsInput{}},
			{"apicall": "DescribeSnapshotCopyGrants", "input_obj": &redshift.DescribeSnapshotCopyGrantsInput{}},
			{"apicall": "DescribeClusterVersions", "input_obj": &redshift.DescribeClusterVersionsInput{}},
			{"apicall": "DescribeEventCategories", "input_obj": &redshift.DescribeEventCategoriesInput{}},
			{"apicall": "DescribeHsmConfigurations", "input_obj": &redshift.DescribeHsmConfigurationsInput{}},
			{"apicall": "DescribeReservedNodes", "input_obj": &redshift.DescribeReservedNodesInput{}},
			{"apicall": "DescribeStorage", "input_obj": &redshift.DescribeStorageInput{}},
			{"apicall": "DescribeEventSubscriptions", "input_obj": &redshift.DescribeEventSubscriptionsInput{}},
			{"apicall": "DescribeAccountAttributes", "input_obj": &redshift.DescribeAccountAttributesInput{}},
		}}

	rekognition_svc := &servicemaster.ServiceMaster{
		Svc:     rekognition.NewFromConfig(cfg),
		SvcName: "rekognition",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListStreamProcessors", "input_obj": &rekognition.ListStreamProcessorsInput{}},
			{"apicall": "ListCollections", "input_obj": &rekognition.ListCollectionsInput{}},
		}}

	robomaker_svc := &servicemaster.ServiceMaster{
		Svc:     robomaker.NewFromConfig(cfg),
		SvcName: "robomaker",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListSimulationApplications", "input_obj": &robomaker.ListSimulationApplicationsInput{}},
			{"apicall": "ListSimulationJobs", "input_obj": &robomaker.ListSimulationJobsInput{}},
			{"apicall": "ListRobotApplications", "input_obj": &robomaker.ListRobotApplicationsInput{}},
			{"apicall": "ListRobots", "input_obj": &robomaker.ListRobotsInput{}},
			{"apicall": "ListFleets", "input_obj": &robomaker.ListFleetsInput{}},
			{"apicall": "ListDeploymentJobs", "input_obj": &robomaker.ListDeploymentJobsInput{}},
		}}

	route53_svc := &servicemaster.ServiceMaster{
		Svc:     route53.NewFromConfig(cfg),
		SvcName: "route53",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetHealthCheckCount", "input_obj": &route53.GetHealthCheckCountInput{}},
			{"apicall": "GetTrafficPolicyInstanceCount", "input_obj": &route53.GetTrafficPolicyInstanceCountInput{}},
			{"apicall": "ListHealthChecks", "input_obj": &route53.ListHealthChecksInput{}},
			{"apicall": "ListHostedZonesByName", "input_obj": &route53.ListHostedZonesByNameInput{}},
			{"apicall": "ListTrafficPolicyInstances", "input_obj": &route53.ListTrafficPolicyInstancesInput{}},
			{"apicall": "ListHostedZones", "input_obj": &route53.ListHostedZonesInput{}},
			{"apicall": "ListQueryLoggingConfigs", "input_obj": &route53.ListQueryLoggingConfigsInput{}},
			{"apicall": "GetHostedZoneCount", "input_obj": &route53.GetHostedZoneCountInput{}},
			{"apicall": "ListTrafficPolicies", "input_obj": &route53.ListTrafficPoliciesInput{}},
			{"apicall": "ListReusableDelegationSets", "input_obj": &route53.ListReusableDelegationSetsInput{}},
		}}

	route53domains_svc := &servicemaster.ServiceMaster{
		Svc:     route53domains.NewFromConfig(cfg),
		SvcName: "route53domains",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListOperations", "input_obj": &route53domains.ListOperationsInput{}},
			{"apicall": "GetContactReachabilityStatus", "input_obj": &route53domains.GetContactReachabilityStatusInput{}},
			{"apicall": "ListDomains", "input_obj": &route53domains.ListDomainsInput{}},
		}}

	route53resolver_svc := &servicemaster.ServiceMaster{
		Svc:     route53resolver.NewFromConfig(cfg),
		SvcName: "route53resolver",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListResolverEndpoints", "input_obj": &route53resolver.ListResolverEndpointsInput{}},
			{"apicall": "ListResolverRules", "input_obj": &route53resolver.ListResolverRulesInput{}},
			{"apicall": "ListResolverRuleAssociations", "input_obj": &route53resolver.ListResolverRuleAssociationsInput{}},
		}}

	s3_svc := &servicemaster.ServiceMaster{
		Svc:     s3.NewFromConfig(cfg),
		SvcName: "s3",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListBuckets", "input_obj": &s3.ListBucketsInput{}},
		}}

	sagemaker_svc := &servicemaster.ServiceMaster{
		Svc:     sagemaker.NewFromConfig(cfg),
		SvcName: "sagemaker",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListWorkteams", "input_obj": &sagemaker.ListWorkteamsInput{}},
			{"apicall": "ListAlgorithms", "input_obj": &sagemaker.ListAlgorithmsInput{}},
			{"apicall": "ListNotebookInstanceLifecycleConfigs", "input_obj": &sagemaker.ListNotebookInstanceLifecycleConfigsInput{}},
			{"apicall": "ListCodeRepositories", "input_obj": &sagemaker.ListCodeRepositoriesInput{}},
			{"apicall": "ListLabelingJobs", "input_obj": &sagemaker.ListLabelingJobsInput{}},
			{"apicall": "ListEndpoints", "input_obj": &sagemaker.ListEndpointsInput{}},
			{"apicall": "ListEndpointConfigs", "input_obj": &sagemaker.ListEndpointConfigsInput{}},
			{"apicall": "ListSubscribedWorkteams", "input_obj": &sagemaker.ListSubscribedWorkteamsInput{}},
			{"apicall": "ListTransformJobs", "input_obj": &sagemaker.ListTransformJobsInput{}},
			{"apicall": "ListTrainingJobs", "input_obj": &sagemaker.ListTrainingJobsInput{}},
			{"apicall": "ListNotebookInstances", "input_obj": &sagemaker.ListNotebookInstancesInput{}},
			{"apicall": "ListModelPackages", "input_obj": &sagemaker.ListModelPackagesInput{}},
			{"apicall": "ListCompilationJobs", "input_obj": &sagemaker.ListCompilationJobsInput{}},
			{"apicall": "ListModels", "input_obj": &sagemaker.ListModelsInput{}},
			{"apicall": "ListHyperParameterTuningJobs", "input_obj": &sagemaker.ListHyperParameterTuningJobsInput{}},
		}}

	secretsmanager_svc := &servicemaster.ServiceMaster{
		Svc:     secretsmanager.NewFromConfig(cfg),
		SvcName: "secretsmanager",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetRandomPassword", "input_obj": &secretsmanager.GetRandomPasswordInput{}},
			{"apicall": "ListSecrets", "input_obj": &secretsmanager.ListSecretsInput{}},
		}}

	securityhub_svc := &servicemaster.ServiceMaster{
		Svc:     securityhub.NewFromConfig(cfg),
		SvcName: "securityhub",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListInvitations", "input_obj": &securityhub.ListInvitationsInput{}},
			{"apicall": "ListEnabledProductsForImport", "input_obj": &securityhub.ListEnabledProductsForImportInput{}},
			{"apicall": "GetFindings", "input_obj": &securityhub.GetFindingsInput{}},
			{"apicall": "GetInvitationsCount", "input_obj": &securityhub.GetInvitationsCountInput{}},
			{"apicall": "GetEnabledStandards", "input_obj": &securityhub.GetEnabledStandardsInput{}},
			{"apicall": "GetInsights", "input_obj": &securityhub.GetInsightsInput{}},
			{"apicall": "GetMasterAccount", "input_obj": &securityhub.GetMasterAccountInput{}},
			{"apicall": "ListMembers", "input_obj": &securityhub.ListMembersInput{}},
		}}

	servicecatalog_svc := &servicemaster.ServiceMaster{
		Svc:     servicecatalog.NewFromConfig(cfg),
		SvcName: "servicecatalog",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListTagOptions", "input_obj": &servicecatalog.ListTagOptionsInput{}},
			{"apicall": "ListRecordHistory", "input_obj": &servicecatalog.ListRecordHistoryInput{}},
			{"apicall": "ListProvisionedProductPlans", "input_obj": &servicecatalog.ListProvisionedProductPlansInput{}},
			{"apicall": "ListServiceActions", "input_obj": &servicecatalog.ListServiceActionsInput{}},
			{"apicall": "ListPortfolios", "input_obj": &servicecatalog.ListPortfoliosInput{}},
			{"apicall": "ListAcceptedPortfolioShares", "input_obj": &servicecatalog.ListAcceptedPortfolioSharesInput{}},
			{"apicall": "GetAWSOrganizationsAccessStatus", "input_obj": &servicecatalog.GetAWSOrganizationsAccessStatusInput{}},
		}}

	shield_svc := &servicemaster.ServiceMaster{
		Svc:     shield.NewFromConfig(cfg),
		SvcName: "shield",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeEmergencyContactSettings", "input_obj": &shield.DescribeEmergencyContactSettingsInput{}},
			{"apicall": "DescribeDRTAccess", "input_obj": &shield.DescribeDRTAccessInput{}},
			{"apicall": "ListAttacks", "input_obj": &shield.ListAttacksInput{}},
			{"apicall": "ListProtections", "input_obj": &shield.ListProtectionsInput{}},
			{"apicall": "GetSubscriptionState", "input_obj": &shield.GetSubscriptionStateInput{}},
			{"apicall": "DescribeProtection", "input_obj": &shield.DescribeProtectionInput{}},
			{"apicall": "DescribeSubscription", "input_obj": &shield.DescribeSubscriptionInput{}},
		}}

	signer_svc := &servicemaster.ServiceMaster{
		Svc:     signer.NewFromConfig(cfg),
		SvcName: "signer",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListSigningPlatforms", "input_obj": &signer.ListSigningPlatformsInput{}},
			{"apicall": "ListSigningProfiles", "input_obj": &signer.ListSigningProfilesInput{}},
			{"apicall": "ListSigningJobs", "input_obj": &signer.ListSigningJobsInput{}},
		}}

	sms_svc := &servicemaster.ServiceMaster{
		Svc:     sms.NewFromConfig(cfg),
		SvcName: "sms",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetAppReplicationConfiguration", "input_obj": &sms.GetAppReplicationConfigurationInput{}},
			{"apicall": "GetAppLaunchConfiguration", "input_obj": &sms.GetAppLaunchConfigurationInput{}},
			{"apicall": "ListApps", "input_obj": &sms.ListAppsInput{}},
			{"apicall": "GetReplicationJobs", "input_obj": &sms.GetReplicationJobsInput{}},
			{"apicall": "GetApp", "input_obj": &sms.GetAppInput{}},
			{"apicall": "GetServers", "input_obj": &sms.GetServersInput{}},
			{"apicall": "GetConnectors", "input_obj": &sms.GetConnectorsInput{}},
		}}

	snowball_svc := &servicemaster.ServiceMaster{
		Svc:     snowball.NewFromConfig(cfg),
		SvcName: "snowball",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListCompatibleImages", "input_obj": &snowball.ListCompatibleImagesInput{}},
			{"apicall": "ListClusters", "input_obj": &snowball.ListClustersInput{}},
			{"apicall": "ListJobs", "input_obj": &snowball.ListJobsInput{}},
			{"apicall": "GetSnowballUsage", "input_obj": &snowball.GetSnowballUsageInput{}},
			{"apicall": "DescribeAddresses", "input_obj": &snowball.DescribeAddressesInput{}},
		}}

	sns_svc := &servicemaster.ServiceMaster{
		Svc:     sns.NewFromConfig(cfg),
		SvcName: "sns",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetSMSAttributes", "input_obj": &sns.GetSMSAttributesInput{}},
			{"apicall": "ListPhoneNumbersOptedOut", "input_obj": &sns.ListPhoneNumbersOptedOutInput{}},
			{"apicall": "ListTopics", "input_obj": &sns.ListTopicsInput{}},
			{"apicall": "ListPlatformApplications", "input_obj": &sns.ListPlatformApplicationsInput{}},
			{"apicall": "ListSubscriptions", "input_obj": &sns.ListSubscriptionsInput{}},
		}}

	sqs_svc := &servicemaster.ServiceMaster{
		Svc:     sqs.NewFromConfig(cfg),
		SvcName: "sqs",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListQueues", "input_obj": &sqs.ListQueuesInput{}},
		}}

	ssm_svc := &servicemaster.ServiceMaster{
		Svc:     ssm.NewFromConfig(cfg),
		SvcName: "ssm",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListCommands", "input_obj": &ssm.ListCommandsInput{}},
			{"apicall": "ListResourceComplianceSummaries", "input_obj": &ssm.ListResourceComplianceSummariesInput{}},
			{"apicall": "DescribeAvailablePatches", "input_obj": &ssm.DescribeAvailablePatchesInput{}},
			{"apicall": "GetInventorySchema", "input_obj": &ssm.GetInventorySchemaInput{}},
			{"apicall": "ListComplianceSummaries", "input_obj": &ssm.ListComplianceSummariesInput{}},
			{"apicall": "DescribeAssociation", "input_obj": &ssm.DescribeAssociationInput{}},
			{"apicall": "DescribeActivations", "input_obj": &ssm.DescribeActivationsInput{}},
			{"apicall": "ListResourceDataSync", "input_obj": &ssm.ListResourceDataSyncInput{}},
			{"apicall": "DescribeMaintenanceWindows", "input_obj": &ssm.DescribeMaintenanceWindowsInput{}},
			{"apicall": "DescribeMaintenanceWindowSchedule", "input_obj": &ssm.DescribeMaintenanceWindowScheduleInput{}},
			{"apicall": "DescribePatchBaselines", "input_obj": &ssm.DescribePatchBaselinesInput{}},
			{"apicall": "DescribePatchGroups", "input_obj": &ssm.DescribePatchGroupsInput{}},
			{"apicall": "ListComplianceItems", "input_obj": &ssm.ListComplianceItemsInput{}},
			{"apicall": "GetDefaultPatchBaseline", "input_obj": &ssm.GetDefaultPatchBaselineInput{}},
			{"apicall": "DescribeInventoryDeletions", "input_obj": &ssm.DescribeInventoryDeletionsInput{}},
			{"apicall": "ListCommandInvocations", "input_obj": &ssm.ListCommandInvocationsInput{}},
		}}

	storagegateway_svc := &servicemaster.ServiceMaster{
		Svc:     storagegateway.NewFromConfig(cfg),
		SvcName: "storagegateway",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListVolumes", "input_obj": &storagegateway.ListVolumesInput{}},
			{"apicall": "ListGateways", "input_obj": &storagegateway.ListGatewaysInput{}},
			{"apicall": "DescribeTapeArchives", "input_obj": &storagegateway.DescribeTapeArchivesInput{}},
			{"apicall": "ListFileShares", "input_obj": &storagegateway.ListFileSharesInput{}},
			{"apicall": "ListTapes", "input_obj": &storagegateway.ListTapesInput{}},
		}}

	sts_svc := &servicemaster.ServiceMaster{
		Svc:     sts.NewFromConfig(cfg),
		SvcName: "sts",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetSessionToken", "input_obj": &sts.GetSessionTokenInput{}},
			{"apicall": "GetCallerIdentity", "input_obj": &sts.GetCallerIdentityInput{}},
		}}

	support_svc := &servicemaster.ServiceMaster{
		Svc:     support.NewFromConfig(cfg),
		SvcName: "support",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeServices", "input_obj": &support.DescribeServicesInput{}},
			{"apicall": "DescribeSeverityLevels", "input_obj": &support.DescribeSeverityLevelsInput{}},
			{"apicall": "DescribeCases", "input_obj": &support.DescribeCasesInput{}},
		}}

	transcribe_svc := &servicemaster.ServiceMaster{
		Svc:     transcribe.NewFromConfig(cfg),
		SvcName: "transcribe",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListTranscriptionJobs", "input_obj": &transcribe.ListTranscriptionJobsInput{}},
			{"apicall": "ListVocabularies", "input_obj": &transcribe.ListVocabulariesInput{}},
		}}

	transfer_svc := &servicemaster.ServiceMaster{
		Svc:     transfer.NewFromConfig(cfg),
		SvcName: "transfer",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListServers", "input_obj": &transfer.ListServersInput{}},
		}}

	translate_svc := &servicemaster.ServiceMaster{
		Svc:     translate.NewFromConfig(cfg),
		SvcName: "translate",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListTerminologies", "input_obj": &translate.ListTerminologiesInput{}},
		}}

	waf_svc := &servicemaster.ServiceMaster{
		Svc:     waf.NewFromConfig(cfg),
		SvcName: "waf",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListIPSets", "input_obj": &waf.ListIPSetsInput{}},
			{"apicall": "ListByteMatchSets", "input_obj": &waf.ListByteMatchSetsInput{}},
			{"apicall": "ListRegexMatchSets", "input_obj": &waf.ListRegexMatchSetsInput{}},
			{"apicall": "GetChangeToken", "input_obj": &waf.GetChangeTokenInput{}},
			{"apicall": "ListXssMatchSets", "input_obj": &waf.ListXssMatchSetsInput{}},
			{"apicall": "ListGeoMatchSets", "input_obj": &waf.ListGeoMatchSetsInput{}},
			{"apicall": "ListRateBasedRules", "input_obj": &waf.ListRateBasedRulesInput{}},
			{"apicall": "ListRuleGroups", "input_obj": &waf.ListRuleGroupsInput{}},
			{"apicall": "ListRegexPatternSets", "input_obj": &waf.ListRegexPatternSetsInput{}},
			{"apicall": "ListRules", "input_obj": &waf.ListRulesInput{}},
			{"apicall": "ListSizeConstraintSets", "input_obj": &waf.ListSizeConstraintSetsInput{}},
			{"apicall": "ListLoggingConfigurations", "input_obj": &waf.ListLoggingConfigurationsInput{}},
			{"apicall": "ListActivatedRulesInRuleGroup", "input_obj": &waf.ListActivatedRulesInRuleGroupInput{}},
			{"apicall": "ListSubscribedRuleGroups", "input_obj": &waf.ListSubscribedRuleGroupsInput{}},
			{"apicall": "ListSqlInjectionMatchSets", "input_obj": &waf.ListSqlInjectionMatchSetsInput{}},
		}}

	workdocs_svc := &servicemaster.ServiceMaster{
		Svc:     workdocs.NewFromConfig(cfg),
		SvcName: "workdocs",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeActivities", "input_obj": &workdocs.DescribeActivitiesInput{}},
			{"apicall": "DescribeUsers", "input_obj": &workdocs.DescribeUsersInput{}},
			{"apicall": "GetResources", "input_obj": &workdocs.GetResourcesInput{}},
		}}

	worklink_svc := &servicemaster.ServiceMaster{
		Svc:     worklink.NewFromConfig(cfg),
		SvcName: "worklink",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListFleets", "input_obj": &worklink.ListFleetsInput{}},
		}}

	workmail_svc := &servicemaster.ServiceMaster{
		Svc:     workmail.NewFromConfig(cfg),
		SvcName: "workmail",
		ApiCalls: []map[string]interface{}{
			{"apicall": "ListOrganizations", "input_obj": &workmail.ListOrganizationsInput{}},
		}}

	workspaces_svc := &servicemaster.ServiceMaster{
		Svc:     workspaces.NewFromConfig(cfg),
		SvcName: "workspaces",
		ApiCalls: []map[string]interface{}{
			{"apicall": "DescribeWorkspacesConnectionStatus", "input_obj": &workspaces.DescribeWorkspacesConnectionStatusInput{}},
			{"apicall": "DescribeWorkspaceDirectories", "input_obj": &workspaces.DescribeWorkspaceDirectoriesInput{}},
			{"apicall": "DescribeWorkspaceBundles", "input_obj": &workspaces.DescribeWorkspaceBundlesInput{}},
			{"apicall": "DescribeIpGroups", "input_obj": &workspaces.DescribeIpGroupsInput{}},
			{"apicall": "DescribeAccountModifications", "input_obj": &workspaces.DescribeAccountModificationsInput{}},
			{"apicall": "DescribeWorkspaces", "input_obj": &workspaces.DescribeWorkspacesInput{}},
			{"apicall": "DescribeAccount", "input_obj": &workspaces.DescribeAccountInput{}},
			{"apicall": "DescribeWorkspaceImages", "input_obj": &workspaces.DescribeWorkspaceImagesInput{}},
		}}

	xray_svc := &servicemaster.ServiceMaster{
		Svc:     xray.NewFromConfig(cfg),
		SvcName: "xray",
		ApiCalls: []map[string]interface{}{
			{"apicall": "GetSamplingRules", "input_obj": &xray.GetSamplingRulesInput{}},
			{"apicall": "GetGroups", "input_obj": &xray.GetGroupsInput{}},
			{"apicall": "GetSamplingStatisticSummaries", "input_obj": &xray.GetSamplingStatisticSummariesInput{}},
			{"apicall": "GetEncryptionConfig", "input_obj": &xray.GetEncryptionConfigInput{}},
			{"apicall": "GetGroup", "input_obj": &xray.GetGroupInput{}},
		}}

	services := []servicemaster.ServiceMaster{*acm_svc, *amplify_svc, *apigateway_svc, *appmesh_svc, *appsync_svc, *athena_svc, *autoscaling_svc, *backup_svc, *batch_svc, *chime_svc, *cloud9_svc, *clouddirectory_svc, *cloudformation_svc, *cloudfront_svc, *cloudhsm_svc, *cloudhsmv2_svc, *cloudsearch_svc, *cloudtrail_svc, *codebuild_svc, *codecommit_svc, *codedeploy_svc, *codepipeline_svc, *codestar_svc, *comprehend_svc, *datapipeline_svc, *datasync_svc, *dax_svc, *devicefarm_svc, *directconnect_svc, *dlm_svc, *dynamodb_svc, *ec2_svc, *ecr_svc, *ecs_svc, *eks_svc, *elasticache_svc, *elasticbeanstalk_svc, *elastictranscoder_svc, *firehose_svc, *fms_svc, *fsx_svc, *gamelift_svc, *globalaccelerator_svc, *glue_svc, *greengrass_svc, *guardduty_svc, *health_svc, *iam_svc, *inspector_svc, *iot_svc, *iotanalytics_svc, *kafka_svc, *kinesis_svc, *kinesisanalytics_svc, *kinesisvideo_svc, *kms_svc, *lambda_svc, *lightsail_svc, *machinelearning_svc, *macie_svc, *mediaconnect_svc, *mediaconvert_svc, *medialive_svc, *mediapackage_svc, *mediastore_svc, *mediatailor_svc, *mobile_svc, *mq_svc, *opsworks_svc, *organizations_svc, *pinpoint_svc, *polly_svc, *pricing_svc, *ram_svc, *rds_svc, *redshift_svc, *rekognition_svc, *robomaker_svc, *route53_svc, *route53domains_svc, *route53resolver_svc, *s3_svc, *sagemaker_svc, *secretsmanager_svc, *securityhub_svc, *servicecatalog_svc, *shield_svc, *signer_svc, *sms_svc, *snowball_svc, *sns_svc, *sqs_svc, *ssm_svc, *storagegateway_svc, *sts_svc, *support_svc, *transcribe_svc, *transfer_svc, *translate_svc, *waf_svc, *workdocs_svc, *worklink_svc, *workmail_svc, *workspaces_svc, *xray_svc}

	return services
}
