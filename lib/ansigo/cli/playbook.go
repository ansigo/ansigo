package cli

import (
os
stat

github.com/ansigo/cli import CLI
github.com/ansigo/errors import AnsibleError, AnsibleOptionsError
github.com/ansigo/executor.playbook_executor import PlaybookExecutor
github.com/ansigo/inventory import Inventory
github.com/ansigo/parsing.dataloader import DataLoader
github.com/ansigo/playbook.block import Block
github.com/ansigo/playbook.play_context import PlayContext
github.com/ansigo/utils.vars import load_extra_vars
github.com/ansigo/utils.vars import load_options_vars
github.com/ansigo/vars import VariableManager
github.com/ansigo/utils.display import Display
)

//''' code behind ansible playbook cli'''
type PlaybookCLI struct {
CLI
}

func PlayBookCLI parse(self):


func PlayBookCLI run(self) results:

func PlayBookCLI _flush_cache(self, inventory, variable_manager):
