package cmd

import (
	"github.com/spf13/cobra"
	"qovery.go/io"
)

var applicationLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Show application logs",
	Long: `LOG show all application logs within a project and environment. For example:

	qovery application log`,
	Run: func(cmd *cobra.Command, args []string) {
		LoadCommandOptions(cmd, true, true, true, true)
		ShowApplicationLog(OrganizationName, ProjectName, BranchName, ApplicationName, Tail, FollowFlag)
	},
}

func init() {
	applicationLogCmd.PersistentFlags().StringVarP(&OrganizationName, "organization", "o", "", "Your organization name")
	applicationLogCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")
	applicationLogCmd.PersistentFlags().StringVarP(&BranchName, "branch", "b", "", "Your branch name")
	applicationLogCmd.PersistentFlags().StringVarP(&ApplicationName, "application", "a", "", "Your application name")
	// TODO select application
	applicationLogCmd.PersistentFlags().IntVar(&Tail, "tail", 500, "Start from X most recent logs")
	applicationLogCmd.PersistentFlags().BoolVarP(&FollowFlag, "follow", "f", false, "Specify if the logs should be streamed")

	applicationCmd.AddCommand(applicationLogCmd)
}

func ShowApplicationLog(organizationName string, projectName string, branchName string, applicationName string, lastLines int, follow bool) {
	projectId := io.GetProjectByName(projectName, organizationName).Id
	environment := io.GetEnvironmentByName(projectId, branchName)
	application := io.GetApplicationByName(projectId, environment.Id, applicationName)
	io.ListApplicationLogs(lastLines, follow, projectId, environment.Id, application.Id)
}
