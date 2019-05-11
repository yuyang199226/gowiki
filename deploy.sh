#!/usr/bin/env bash
echo change work dir
cd provisioning

ansible-playbook -i hosts --ask-vault-pass --become --become-method sudo --user ubuntu  --key-file=//Users/yang.yu/.ssh/yuyangaws.pem wiki.yml
