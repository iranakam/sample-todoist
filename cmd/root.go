package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/iranakam/tdicli/todoist"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Args is an object contained within Todoist command.
type Args struct {
	ID              int         `json:"id,oimitempty"`
	Name            string      `json:"name,omitempty"`
	Color           int         `json:"color,omitempty"`
	ParentID        int         `json:"parent_id,omitempty"`
	ChildOrder      int         `json:"child_order,omitempty"`
	Collapsed       int         `json:"collapsed,omitempty"`
	IsFavorite      int         `json:"is_favorite,omitempty"`
	UserID          int         `json:"user_id,omitempty"`
	ProjectID       int         `json:"project_id,omitempty"`
	Content         string      `json:"content,omitempty"`
	Indent          int         `json:"indent,omitempty"`
	Priority        int         `json:"priority,omitempty"`
	ItemOrder       int         `json:"item_order,omitempty"`
	DayOrder        int         `json:"day_order,omitempty"`
	Children        interface{} `json:"children,omitempty"`
	AssignedByUID   int         `json:"assigned_by_uid,omitempty"`
	Checked         int         `json:"checked,omitempty"`
	InHistory       int         `json:"in_history,omitempty"`
	IsDeleted       int         `json:"is_deleted,omitempty"`
	IsArchived      int         `json:"is_archived,omitempty"`
	DateAdded       time.Time   `json:"date_added,omitempty"`
	AutoReminder    bool        `json:"auto_reminder,omitempty"`
	AutoParseLabels bool        `json:"auto_parse_labels,omitempty"`
	ForceHistory    bool        `json:"force_history,omitempty"`
}

// client is Todoist client.
var client todoist.Client

// output is Tabwriter writer.
var output tabwriter.Writer

// structArgs is an initialized args.
var structArgs Args

// configFile is a config file path.
var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tdicli",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	output.Init(os.Stdout, 0, 4, 12, ' ', 0)
	rootCmd.SetOutput(&output)
	err := rootCmd.Execute()
	if err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
	output.Flush()
	os.Exit(0)
}

// init works to initial cobra and to set flags to root command.
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.tdicli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".tdicli")
	}

	if existConfig := viper.ReadInConfig(); existConfig != nil {
		rootCmd.Println(existConfig)
		os.Exit(1)
	}

	existToken := viper.Get("token")

	if existToken == nil {
		rootCmd.Println("Token is especially necessary for use.")
		os.Exit(1)
	}

	token := existToken.(string)
	client = todoist.NewClient(token).SetNewID()
}
