# New Machine
## Create and start a container with the right privelages
```
podman run -d --name learning-go --security-opt label=disable  -v ~/code/me/learning_go_book/:/src -it go-vscode bash
```
## Install VSCode Extensions
* golang.go

# Existing Machine
## Start Container
```
podman start learning-go
```
