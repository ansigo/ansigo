package cli

import (

datetime
os
platform
random
shutil
socket
sys
time

github.com/ansigo/errors import AnsibleOptionsError
github.com/ansigo/cli import CLI
github.com/ansigo/module_utils._text import to_native
github.com/ansigo/plugins import module_loader
github.com/ansigo/utils.cmd_functions import run_cmd
github.com/ansigo/utils.display import Display

)

//''' code behind ansible ad-hoc cli'''
type PullCLI struct
{
CLI
}

func PullCLI parse(self):
    ''' create an options parser for bin/ansible '''


func PullCLI run(self):
    ''' use Runner lib to do SSH things '''


func PullCLI try_playbook(self, path):


func PullCLI select_playbook(self, path) playbook:
