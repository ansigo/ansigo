package cli



import (
atexit
cmd
getpass
readline
os
sys

github.com/ansigo/constants as C
github.com/ansigo/cli import CLI
github.com/ansigo/errors import AnsibleError
github.com/ansigo/executor/task_queue_manager import TaskQueueManager
github.com/ansigo/inventory import Inventory
github.com/ansigo/module_utils/_text import to_native, to_text
github.com/ansigo/parsing/dataloader import DataLoader
github.com/ansigo/parsing/splitter import parse_kv
github.com/ansigo/playbook/play import Play
github.com/ansigo/plugins import module_loader
github.com/ansigo/utils import module_docs
github.com/ansigo/utils/color import stringc
github.com/ansigo/vars import VariableManager
github.com/ansigo/utils/display import Display
)

type ConsoleCLI struct {
  CLI
  cmd.Cmd
  }

    func ConsoleCLI init(self, args)

    func ConsoleCLI parse(self):

    func ConsoleCLI get_names(self):


    func ConsoleCLI cmdloop(self):


    func ConsoleCLI set_prompt(self):


    func ConsoleCLI list_modules(self) modules :


    func ConsoleCLI _find_modules_in_path(self, path):


    func ConsoleCLI default(self, arg, forceshell=False):
        """ actually runs modules """


    func ConsoleCLI emptyline(self):

    func ConsoleCLI do_shell(self, arg):
        """
        You can run shell commands through the shell module.

        eg.:
        shell ps uax | grep java | wc -l
        shell killall python
        shell halt -n

        You can use the ! to force the shell module. eg.:
        !ps aux | grep java | wc -l
        """

    func ConsoleCLI do_forks(self, arg):
        """Set the number of forks"""


    func ConsoleCLI do_verbosity(self, arg):
        """Set verbosity level"""

    func ConsoleCLI do_cd(self, arg):
        """
            Change active host/group. You can use hosts patterns as well eg.:
            cd webservers
            cd webservers:dbservers
            cd webservers:!phoenix
            cd webservers:&staging
            cd webservers:dbservers:&staging:!phoenix
        """


    func ConsoleCLI do_list(self, arg):
        """List the hosts in the current group"""


    func ConsoleCLI do_become(self, arg):
        """Toggle whether plays run with become"""

    func ConsoleCLI do_remote_user(self, arg):
        """Given a username, set the remote user plays are run by"""

    func ConsoleCLI do_become_user(self, arg):
        """Given a username, set the user that plays are run by when using become"""

    func ConsoleCLI do_become_method(self, arg):
        """Given a become_method, set the privilege escalation method when using become"""

    func ConsoleCLI do_check(self, arg):
        """Toggle whether plays run with check mode"""

    func ConsoleCLI do_diff(self, arg):
        """Toggle whether plays run with diff"""

    func ConsoleCLI do_exit(self, args):
        """Exits from the console"""

    func ConsoleCLI helpdefault(self, module_name):

    func ConsoleCLI complete_cd(self, text, line, begidx, endidx):

    func ConsoleCLI completedefault(self, text, line, begidx, endidx):

    func ConsoleCLI module_args(self, module_name):

    func ConsoleCLI run(self):
