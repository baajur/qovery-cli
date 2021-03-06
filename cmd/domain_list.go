package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"qovery.go/io"
)

var domainListCmd = &cobra.Command{
	Use:   "list",
	Short: "List domains",
	Long: `LIST show all linked domains. For example:

	qovery domain list`,
	Run: func(cmd *cobra.Command, args []string) {
		LoadCommandOptions(cmd, true, true, true, false)
		ShowDomainList(OrganizationName, ProjectName, BranchName)
	},
}

func init() {
	domainListCmd.PersistentFlags().StringVarP(&OrganizationName, "organization", "o", "", "Your organization name")
	domainListCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")
	domainListCmd.PersistentFlags().StringVarP(&BranchName, "branch", "b", "", "Your branch name")
	domainCmd.AddCommand(domainListCmd)
}

func ShowDomainList(organizationName string, projectName string, branchName string) {
	table := io.GetTable()
	table.SetHeader([]string{"branch", "domain", "status", "validation domain", "router name"})

	projectId := io.GetProjectByName(projectName, organizationName).Id
	environment := io.GetEnvironmentByName(projectId, branchName)

	routers := io.ListRouters(projectId, environment.Id)
	if routers.Results == nil || len(routers.Results) == 0 {
		table.Append([]string{"", "", ""})
	} else {
		for _, r := range routers.Results {

			for _, cd := range r.CustomDomains {
				table.Append([]string{
					branchName,
					cd.Domain,
					cd.Status.GetColoredStatus(),
					cd.GetTargetDomain(),
					r.Name,
				})
			}

			table.Append([]string{
				branchName,
				r.ConnectionURI,
				color.GreenString("live"),
				"none",
				r.Name,
			})
		}
	}

	table.Render()
	fmt.Printf("\n")
}
