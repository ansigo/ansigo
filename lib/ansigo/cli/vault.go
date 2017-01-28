package cli

import (
os
sys

from ansible.errors import AnsibleError, AnsibleOptionsError
from ansible.parsing.dataloader import DataLoader
from ansible.parsing.vault import VaultEditor
from ansible.cli import CLI
from ansible.module_utils._text import to_text

from ansible.utils.display import Display
)


""" Vault command line class """
type VaultCLI struct {
 CLI
}

VALID_ACTIONS = ("create", "decrypt", "edit", "encrypt", "rekey", "view")

func VaultCLI __init__(self, args):

func VaultCLI parse(self):

func VaultCLI run(self):


func VaultCLI execute_encrypt(self):


func VaultCLI execute_decrypt(self):


func VaultCLI execute_create(self):


func VaultCLI execute_edit(self):


func VaultCLI execute_view(self):


func VaultCLI execute_rekey(self):
