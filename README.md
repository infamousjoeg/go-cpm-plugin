# go-cpm-plugin
A Golang-based template for creating @CyberArk CPM Plugins

## Development

Reference of arguments needed for plugin.exe:

[https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/TPC-test-plugin.htm?tocpath=Developer%7CCreate%20extensions%7CCPM%20Plugins%7CTerminal%20Plugin%20Controller%7C_____8](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/TPC-test-plugin.htm?tocpath=Developer%7CCreate%20extensions%7CCPM%20Plugins%7CTerminal%20Plugin%20Controller%7C_____8)

### TL;DR

`plugin.exe <XML File Path> <action>`

Action can be:
* logon
* verifypass
* changepass
* prereconcilepass
* reconcilepass
* deletepass
* deleteafterreconcile