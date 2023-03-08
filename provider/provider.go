package provider

import (
	"github.com/selefra/selefra-provider-aws/constants"
	"context"
	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

var Version = "v0.0.5"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:		constants.Aws,
		Version:	Version,
		TableList:	GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var awsConfig aws_client.AwsProviderConfig
				err := config.Unmarshal(&awsConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				clients, err := aws_client.NewClients(awsConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				hash := make(map[string]bool)
				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					if hash[clients[i].GetAccount()] {
						continue
					}
					res = append(res, clients[i])
					hash[clients[i].GetAccount()] = true
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# Optional. by default assumes all regions
# regions:
#   - us-east-1
#   - us-west-2`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var awsConfig aws_client.AwsProviderConfig
				err := config.Unmarshal(&awsConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_34,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand:	true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorAll},
		},
	}
}
