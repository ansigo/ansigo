package cli

import (

github.com/ansigo/constants as C
github.com/ansigo/cli import CLI
github.com/ansigo/errors import AnsibleError, AnsibleOptionsError
github.com/ansigo/executor/task_queue_manager import TaskQueueManager
github.com/ansigo/inventory import Inventory
github.com/ansigo/module_utils/_text import to_text
github.com/ansigo/parsing/dataloader import DataLoader
github.com/ansigo/parsing/splitter import parse_kv
github.com/ansigo/playbook/play import Play
github.com/ansigo/plugins import get_all_plugin_loaders
github.com/ansigo/utils/vars import load_extra_vars
github.com/ansigo/utils/vars import load_options_vars
github.com/ansigo/vars import VariableManager
github.com/ansigo/utils/display import Display
)

type AdHocCLI(CLI):
    ''' code behind ansible ad-hoc cli'''

func AdHocCLI(CLI) parse(self):
        ''' create an options parser for bin/ansible '''

func AdHocCLI(CLI) run(self):
        ''' use Runner lib to do SSH things '''
