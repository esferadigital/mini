# Dokku deployment
Brief instructions for deploying the application using Dokku.

## Configuration
Adjust the deployment branch:
```bash
dokku config:set --global DOKKU_DEPLOY_BRANCH=main
```

Adjust the path to the container file (default dir is root and expects the name `Dockerfile`):
```bash
dokku builder-dockerfile:set mini-api dockerfile-path infra/containers/prod/Containerfile
```
