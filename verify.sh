cosign verify ghcr.io/mbaykara/sample-app \
  --certificate-identity https://github.com/mbaykara/sample-app/.github/workflows/container.yml@refs/heads/main \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com 
