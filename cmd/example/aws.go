package example

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/cliutils"

	"github.com/Excoriate/golang-cli-boilerplate/internal/output"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/cloudaws"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"

	"github.com/Excoriate/golang-cli-boilerplate/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var commandName = "aws-ecs"

// GetClient returns the client instance from the context.
// Update GetClient Function as below

// GetClient returns the client instance from the context.
func GetClient(cmd *cobra.Command) *cli.Client {
	ctx := cmd.Context().Value(cliutils.GetCtxKey())
	if ctx == nil {
		log.Fatal("Unable to get the client context.")
	}
	client, ok := ctx.(*cli.Client)
	if !ok {
		log.Fatal("Unable to assert client.")
	}
	return client
}

var Cmd = &cobra.Command{
	Use: commandName,
	Long: fmt.Sprintf(`The '%s' command is used to deploy a service to ECS, or perform
different type of actions on top of it.`, commandName),
	Example: cliutils.GenerateExampleInCMD([]cliutils.ExampleTemplateOptions{
		{
			CLIName: "golang-cli-boilerplate",
			Command: commandName,
			Options: "--aws-region=us-east-1 --output=yaml --save",
			Explanation: `This command will list all the ECS clusters in the us-east-1 region,
and it'll save the output in a file called ecs-clusters.yaml.`,
		},
	}),

	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient(cmd)

		client.UX.Titles.ShowTitle(viper.GetString("CLI_NAME_TITLE"))
		client.UX.Messages.ShowInfo(tui.MessageOptions{
			Message: "CLI successfully started. Configuration loaded successfully.",
		})

		// Get arguments.
		awsAccessKeyPassed := viper.GetString("awsAccessKeyId")
		awsSecretKeyPassed := viper.GetString("awsSecretAccessKey")
		awsRegionPassed := viper.GetString("awsRegion")

		// Instantiating client.
		aws, err := cloudaws.NewClient(client.Ctx, client.Logger, cloudaws.InitAWSAdapterOptions{
			Region: awsRegionPassed,
			Creds: cloudaws.AWSCreds{
				AccessKeyID:     awsAccessKeyPassed,
				SecretAccessKey: awsSecretKeyPassed,
			},
		})

		if err != nil {
			client.UX.Messages.ShowError(tui.MessageOptions{
				Message: err.Error(),
			})
			os.Exit(1)
		}

		adapter, err := aws.Build(aws.WithECS())
		if err != nil {
			client.UX.Messages.ShowError(tui.MessageOptions{
				Message: err.Error(),
			})
			os.Exit(1)
		}

		// Connector instantiation.j
		ecsConnector := cloudaws.NewECSConnector(client.Ctx, adapter.ECSClient, client.Logger)

		// List ECS clusters.
		result, _ := ecsConnector.ListECSClusters()

		// Checking if the output option (-o or --output) was passed.
		option := viper.GetString("output")
		save := viper.GetBool("save")

		clusters := struct {
			ClusterArns []string `json:"clusterArns" yaml:"clusterArns"`
		}{}

		clusters.ClusterArns = append(clusters.ClusterArns, result.ClusterArns...)

		out := output.NewTerminalOutput(client.UX.Messages, client.Logger)
		var data [][]string
		for i, cluster := range result.ClusterArns {
			data = append(data, []string{strconv.Itoa(i), cluster})
		}

		if err := out.Show(output.Options{
			Data:         data,
			OutputType:   option,
			SaveInDisk:   save,
			Filename:     "ecs-clusters",
			TableHeaders: []string{"#", "CLUSTER"},
		}); err != nil {
			client.UX.Messages.ShowError(tui.MessageOptions{
				Message: err.Error(),
			})
			os.Exit(1)
		}
	},
}

func addPersistentFlags() {
	err := cliutils.AddFlags(Cmd, []cliutils.CobraFlagOptions{
		{
			LongName:     "aws-region",
			Usage:        "The AWS region to carry out operations in.",
			IsPersistent: true,
			IsRequired:   true, // It'll fail if the flag is not passed.
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "awsRegion",
				EnvVariableNameInHost:  "AWS_REGION", // replace with actual env variable name
				IsEnabled:              true,
			},
		},
		{
			LongName:     "aws-access-key-id",
			Usage:        "The AWS Access Key ID. If it's not set, it'll be read from the AWS_ACCESS_KEY_ID environment variable.",
			IsPersistent: true,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "awsAccessKeyId",
				EnvVariableNameInHost:  "AWS_ACCESS_KEY_ID", // replace with actual env variable name
				IsEnabled:              true,
			},
		},
		{
			LongName:     "aws-secret-access-key",
			Usage:        "The AWS Secret Access Key. If it's not set, it'll be read from the AWS_SECRET_ACCESS_KEY environment variable.",
			IsPersistent: true,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "awsSecretAccessKey",
				EnvVariableNameInHost:  "AWS_SECRET_ACCESS_KEY", // replace with actual env variable name
				IsEnabled:              true,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	addPersistentFlags()
}
