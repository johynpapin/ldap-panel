# LDAP Panel

## Introduction

A very simple LDAP panel that allows users to change their password. Administrators can create new accounts or delete existing ones.

Currently, the interface is in French. If you are interested in a translation in another language, do not hesitate to open an issue.

## Configuration

The configuration is done through environment variables. All parameters are required.

| Name | Description |
| ---- | ----------- |
| LDAP_PANEL_SESSION_KEY | A random key to secure sessions. |
| LDAP_PANEL_LDAP_URL | LDAP server URL |
| LDAP_PANEL_LDAP_USERNAME | Username of an administrator |
| LDAP_PANEL_LDAP_PASSWORD | Password of an administrator |
