variable "auth-project-tag" {
  value = "$NVY_VAR{repo.namespace}/mrs-service-cab-connector-auth:$NVY_VAR{repo.tag}"
}

variable "catalog-project-tag" {
  value = "$NVY_VAR{repo.namespace}/mrs-service-cab-connector-catalog:$NVY_VAR{repo.tag}"
}

/*variable "tasks-project-tag" {
  value = "$NVY_VAR{repo.namespace}/mrs-service-cab-connector-tasks:$NVY_VAR{repo.tag}"
}*/

# Full build image name ex: buuild-simpleserver-v1.0.0
variable "project-build-name-connector" {
  value = "build-$NVY_VAR{repo.name}-connector-$NVY_VAR{repo.tag}"
}

variable "project-tag-build" {
  value = "build-$NVY_VAR{repo.namespace}/$NVY_VAR{repo.name}:$NVY_VAR{repo.tag}"
}

# Build docker build image
docker "build" {
  dockerfile = "scripts/Dockerfile.build"
  path = "$NVY_VAR{sys_workspace}"
  tag = "$NVY_VAR{project-tag-build}"
  no_cache = "false"
}

# Run build image and mount output
docker "run" {
  name = "$NVY_VAR{project-build-name-connector}"
  image = "$NVY_VAR{project-tag-build}"
  cmd = [
    "env",
    "GOOS=linux",
    "GOARCH=amd64",
    "go",
    "build",
    "-o",
    "bin/auth",
    "cmd/connector/auth/main.go"]
  volume_binds = [
    "$NVY_VAR{out_workspace}/bin:/go/src/github.com/Tlantic/mrs-service-cab-connector/bin"]
}

# Remove build container
docker "rm" {
  container = "$NVY_VAR{project-build-name-connector}"
}

docker "run" {
  name = "$NVY_VAR{project-build-name-connector}"
  image = "$NVY_VAR{project-tag-build}"
  cmd = [
    "env",
    "GOOS=linux",
    "GOARCH=amd64",
    "go",
    "build",
    "-o",
    "bin/catalog",
    "cmd/connector/catalog/main.go"]
  volume_binds = [
    "$NVY_VAR{out_workspace}/bin:/go/src/github.com/Tlantic/mrs-service-cab-connector/bin"]
}


docker "rm" {
  container = "$NVY_VAR{project-build-name-connector}"
}

/*docker "run" {
  name = "$NVY_VAR{project-build-name-connector}"
  image = "$NVY_VAR{project-tag-build}"
  cmd = [
    "env",
    "GOOS=linux",
    "GOARCH=amd64",
    "go",
    "build",
    "-o",
    "bin/tasks",
    "cmd/connector/tasks/main.go"]
  volume_binds = [
    "$NVY_VAR{out_workspace}/bin:/go/src/github.com/Tlantic/mrs-service-cab-connector/bin"]
}

docker "rm" {
  container = "$NVY_VAR{project-build-name-connector}"
}*/

# Remove build image
docker "rmi" {
  name = "$NVY_VAR{project-tag-build}"
}


# Run dist docker image
docker "build" {
  dockerfile = "scripts/Dockerfile.auth.dist"
  path = "$NVY_VAR{sys_workspace}"
  tag = "$NVY_VAR{auth-project-tag}"
  no_cache = "false"
}

# Run dist docker image
docker "build" {
  dockerfile = "scripts/Dockerfile.catalog.dist"
  path = "$NVY_VAR{sys_workspace}"
  tag = "$NVY_VAR{catalog-project-tag}"
  no_cache = "false"
}

# Run dist docker image
/*docker "build" {
  dockerfile = "scripts/Dockerfile.tasks.dist"
  path = "$NVY_VAR{sys_workspace}"
  tag = "$NVY_VAR{tasks-project-tag}"
  no_cache = "false"
}*/


# Push dist image
docker "push" {
  image = "$NVY_VAR{auth-project-tag}"
  username = "$NVY_VAR{ecr.username}"
  password = "$NVY_VAR{ecr.password}"
}

# Push dist image
docker "push" {
  image = "$NVY_VAR{catalog-project-tag}"
  username = "$NVY_VAR{ecr.username}"
  password = "$NVY_VAR{ecr.password}"
}

# Push dist image
/*docker "push" {
  image = "$NVY_VAR{tasks-project-tag}"
  username = "$NVY_VAR{ecr.username}"
  password = "$NVY_VAR{ecr.password}"
}*/


# Remove dist image
docker "rmi" {
  name = "$NVY_VAR{auth-project-tag}"
}

# Remove dist image
docker "rmi" {
  name = "$NVY_VAR{catalog-project-tag}"
}

# Remove dist image
/*docker "rmi" {
  name = "$NVY_VAR{tasks-project-tag}"
}*/

plugin "navy-helm" "build"{
  path = "$NVY_VAR{sys_workspace}/navyhook/$NVY_VAR{repo.name}-auth"
  destination ="$NVY_VAR{sys_workspace}"
  repository = "http://54.171.51.193:8000/envs/$NVY_VAR{repo.env}/charts/upload/"
  image = "$NVY_VAR{repo.namespace}/$NVY_VAR{repo.name}-auth"
  tag = "$NVY_VAR{repo.tag}"
  name = "$NVY_VAR{repo.name}-auth"
}

plugin "navy-helm" "build"{
  path = "$NVY_VAR{sys_workspace}/navyhook/$NVY_VAR{repo.name}-catalog"
  destination ="$NVY_VAR{sys_workspace}"
  repository = "http://54.171.51.193:8000/envs/$NVY_VAR{repo.env}/charts/upload/"
  image = "$NVY_VAR{repo.namespace}/$NVY_VAR{repo.name}-catalog"
  tag = "$NVY_VAR{repo.tag}"
  name = "$NVY_VAR{repo.name}-catalog"
}

/*plugin "navy-helm" "build"{
  path = "$NVY_VAR{sys_workspace}/navyhook/$NVY_VAR{repo.name}-tasks"
  destination ="$NVY_VAR{sys_workspace}"
  repository = "http://54.171.51.193:8000/envs/$NVY_VAR{repo.env}/charts/upload/"
  image = "$NVY_VAR{repo.namespace}/$NVY_VAR{repo.name}-tasks"
  tag = "$NVY_VAR{repo.tag}"
  name = "$NVY_VAR{repo.name}-tasks"
}*/
