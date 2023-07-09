package Cmd

import "github.com/spf13/cobra"

const (
	rootGroupID    string = "root-group-id"
	rootGroupTitle string = "Commands:"
)

var _ Cmd = (*rootCmd)(nil)

// Cmd 命令行解析工具对象
type Cmd interface {
	i()

	// RegisterCmds
	//@Title RegisterCmds
	// @Description 命令行工具注册函数
	// @Tags Cmd
	// @param *cobra.Command "cobra对象数组"
	RegisterCmds(subCmds ...*cobra.Command)

	//Execute
	// @Title Execute
	// @Tags Cmd
	// @Description "执行函数,调用cobra的Execute命令"
	// @Return error "执行时可能出现的错误"
	Execute() error
}

type rootCmd struct {
	root *cobra.Command
}

func (*rootCmd) i() {}

func (r *rootCmd) RegisterCmds(subCmds ...*cobra.Command) {
	r.root.AddCommand(subCmds...)
}

func (r *rootCmd) Execute() error {
	return r.root.Execute()
}

// NewCmd Cmd工厂函数
func NewCmd() Cmd {
	r := new(rootCmd)
	r.root = &cobra.Command{
		Use:   "Jnote-Server",
		Short: "run Jonas Note Server",
		Long:  "run Jonas Note Server",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				return
			}
			cmd.Help()
		},
		Args:    cobra.NoArgs,
		GroupID: "root",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		DisableSuggestions: true,
	}

	r.root.AddGroup(&cobra.Group{
		ID:    rootGroupID,
		Title: rootGroupTitle,
	})

	r.root.SetHelpCommandGroupID(rootGroupID)
	return r
}
