package aws

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/cliutils"

	"github.com/Excoriate/golang-cli-boilerplate/internal/sdout"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/cloudaws"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"

	"github.com/Excoriate/golang-cli-boilerplate/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cliClient *cli.Client

var Cmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "aws",
	Long: `The 'aws' command is used to deploy a service to ECS, or perform
different type of actions on top of it.`,

	Example: `
	  ecs-deployer service deploy --service=myservice --cluster=mycluster
	  ecs-deployer service scale --service=myservice --cluster=mycluster --desired-count=2
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize the CLI client.
		cliInit := cli.New(context.TODO(), cli.InitOptions{
			DotEnvFile:          viper.GetString("dotEnvFile"),
			ScanEnvVarsFromHost: viper.GetBool("scanEnvVarsFromHost"),
			ShowEnvVars:         viper.GetBool("showEnvVars"),
		})

		// Build the CLI client.
		cliClientCpy, err := cliInit.Build(cliInit.WithScannedEnvVarsFromHost(),
			cliInit.WithPrintedHostEnvVars(), cliInit.WithDotEnvFile())

		if err != nil {
			os.Exit(1)
		}

		cliClient = cliClientCpy
	},
	Run: func(cmd *cobra.Command, args []string) {
		cliClient.UX.Titles.ShowTitle(viper.GetString("CLI_NAME_TITLE"))
		cliClient.UX.Messages.ShowInfo(tui.MessageOptions{
			Message: "CLI successfully started. Configuration loaded successfully.",
		})

		// Get arguments.
		awsAccessKeyPassed := viper.GetString("awsAccessKeyId")
		awsSecretKeyPassed := viper.GetString("awsSecretAccessKey")
		awsRegionPassed := viper.GetString("awsRegion")

		// Instantiating client.
		aws, err := cloudaws.NewClient(cliClient.Ctx, cliClient.Logger, cloudaws.InitAWSAdapterOptions{
			Region: awsRegionPassed,
			Creds: cloudaws.AWSCreds{
				AccessKeyID:     awsAccessKeyPassed,
				SecretAccessKey: awsSecretKeyPassed,
			},
		})

		if err != nil {
			cliClient.UX.Messages.ShowError(tui.MessageOptions{
				Message: err.Error(),
			})
			os.Exit(1)
		}

		adapter, err := aws.Build(aws.WithECS())
		if err != nil {
			cliClient.UX.Messages.ShowError(tui.MessageOptions{
				Message: err.Error(),
			})
			os.Exit(1)
		}

		// Connector instantiation.j
		ecsConnector := cloudaws.NewECSConnector(cliClient.Ctx, adapter.ECSClient, cliClient.Logger)

		// List ECS clusters.
		result, _ := ecsConnector.ListECSClusters()

		// Checking if the output option (-o or --output) was passed.
		option := viper.GetString("output")
		save := viper.GetBool("save")
		if option == "" {
			option = "table"
		}

		// if option is yaml or yml, it's equivalent. So, always set yaml
		if option == "yaml" {
			option = "yml"
		}

		clusters := struct {
			ClusterArns []string `json:"clusterArns" yaml:"clusterArns"`
		}{}

		clusters.ClusterArns = append(clusters.ClusterArns, result.ClusterArns...)

		switch option {
		case "table":
			// Show table.
			headers := []string{"#", "CLUSTER"}
			var data [][]string
			for i, cluster := range result.ClusterArns {
				data = append(data, []string{strconv.Itoa(i), cluster})
			}

			_ = tui.ShowTable(tui.TableOptions{
				Headers: headers,
				Data:    data,
			})
		case "yml":
			if err := sdout.ShowYAML(sdout.StOutOptions{
				Data:       clusters,
				SaveInDisk: save,
				Filename:   "ecs-clusters.yml",
			}); err != nil {
				cliClient.UX.Messages.ShowError(tui.MessageOptions{
					Message: err.Error(),
				})
				os.Exit(1)
			}

		case "json":
			if err := sdout.ShowJSON(sdout.StOutOptions{
				Data:       clusters,
				SaveInDisk: save,
				Filename:   "ecs-clusters.json",
			}); err != nil {
				cliClient.UX.Messages.ShowError(tui.MessageOptions{
					Message: err.Error(),
				})
				os.Exit(1)
			}
		}

		if save {
			// Convert the struct to an array of bytes []byte
			cliClient.UX.Messages.ShowInfo(tui.MessageOptions{
				Message: "File ecs-clusters successfully saved.",
			})
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
