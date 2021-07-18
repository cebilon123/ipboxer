package cmd

import (
	"fmt"
	"github.com/cebilon123/ipboxer/provider"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var providerName *string

// providerCmd represents the provider command
var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Set provider of time data. Default: Clockify",
	Long:  `Set provider of time data. Default: Clockify`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provider called")
	},
}

// providerListCmd represents the provider list command
var providerListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show list of available providers",
	Long:  "Show list of available providers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provider list called")
	},
}

// providerCreateCmd represents the provider create command, which creates provider config
var providerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates provider config in .yaml format. Default (\"clockify\")",
	Long:  "Creates provider config in .yaml format. Default (\"clockify\")",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := provider.NewConfig("clockify", "asdxcvAs241@asdx4%", "header:{X-Api-Key}", "/workspaces/workspaceId_54123/user/userId_51234/time-entries?start=10.12.2024&end=12.12.2024")
		bts, _ := yaml.Marshal(cfg)
		fmt.Println("provider create called")
		fmt.Println(string(bts))
	},
}

func init() {
	rootCmd.AddCommand(providerCmd)

	providerCmd.AddCommand(providerListCmd, providerCreateCmd)
	providerName = providerCmd.Flags().String("name", "clockify", "Set time provider")
}
