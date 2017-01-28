package cli

import (

datetime
os
traceback
textwrap

github.com/ansigo/compat.six import iteritems, string_types
github.com/ansigo/constants as C
github.com/ansigo/errors import AnsibleError, AnsibleOptionsError
github.com/ansigo/plugins import module_loader, action_loader
github.com/ansigo/cli import CLI
github.com/ansigo/utils import module_docs
github.com/ansigo/utils.display import Display

)

//""" Vault command line class """
type DocCLI struct {
CLI
}

func DocCLI __init__(self, args):

func DocCLI parse(self):


func DocCLI run(self):
func DocCLI find_modules(self, path):

func DocCLI get_module_list_text(self):
@staticmethod
func DocCLI print_paths(finder):
    ''' Returns a string suitable for printing of the search path '''
func DocCLI get_snippet_text(self, doc):
func DocCLI get_man_text(self, doc):
