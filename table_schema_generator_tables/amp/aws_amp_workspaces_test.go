

package amp







import (
	"testing"









	"github.com/aws/aws-sdk-go-v2/service/amp"


	"github.com/aws/aws-sdk-go-v2/service/amp/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"


)







func buildRuleGroupsNamespaces(t *testing.T, m *mocks.MockAmpClient) {




	var summary types.RuleGroupsNamespaceSummary
	if err := faker.FakeObject(&summary); err != nil {
		t.Fatal(err)


	}



	m.EXPECT().ListRuleGroupsNamespaces(gomock.Any(), gomock.Any()).Return(


		&amp.ListRuleGroupsNamespacesOutput{


			RuleGroupsNamespaces: []types.RuleGroupsNamespaceSummary{summary},


		},


		nil,


	)





	var description types.RuleGroupsNamespaceDescription


	if err := faker.FakeObject(&description); err != nil {
		t.Fatal(err)
	}



	m.EXPECT().DescribeRuleGroupsNamespace(gomock.Any(), gomock.Any()).Return(




		&amp.DescribeRuleGroupsNamespaceOutput{


			RuleGroupsNamespace: &description,




		},




		nil,




	)


}





func buildWorkspaces(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockAmpClient(ctrl)





	var summary types.WorkspaceSummary
	if err := faker.FakeObject(&summary); err != nil {


		t.Fatal(err)
	}



	m.EXPECT().ListWorkspaces(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&amp.ListWorkspacesOutput{


			Workspaces: []types.WorkspaceSummary{summary},
		},
		nil,




	)





	var description types.WorkspaceDescription
	if err := faker.FakeObject(&description); err != nil {
		t.Fatal(err)


	}





	m.EXPECT().DescribeWorkspace(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&amp.DescribeWorkspaceOutput{
			Workspace: &description,
		},
		nil,




	)





	var alertManagerDefinition types.AlertManagerDefinitionDescription




	if err := faker.FakeObject(&alertManagerDefinition); err != nil {
		t.Fatal(err)
	}





	m.EXPECT().DescribeAlertManagerDefinition(gomock.Any(), gomock.Any()).AnyTimes().Return(




		&amp.DescribeAlertManagerDefinitionOutput{
			AlertManagerDefinition: &alertManagerDefinition,


		},


		nil,




	)







	var loggingConfiguration types.LoggingConfigurationMetadata


	if err := faker.FakeObject(&loggingConfiguration); err != nil {
		t.Fatal(err)
	}





	m.EXPECT().DescribeLoggingConfiguration(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&amp.DescribeLoggingConfigurationOutput{
			LoggingConfiguration: &loggingConfiguration,
		},




		nil,




	)





	buildRuleGroupsNamespaces(t, m)



	return aws_client.AwsServices{Amp: m}
}







func TestWorkspaces(t *testing.T) {




	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAmpWorkspacesGenerator{}), buildWorkspaces, aws_client.TestOptions{})
}
