## Inputs

- `IS_AUTHENTICATED`: Denotes if git-auth target should run (default: false)
- `IS_PIPELINE`: Denote if running in a pipeline context (default: true)
- `JOB_NAME`: Sets the git user's name for the job (REQUIRED)
- `JOB_EMAIL`: Sets the git user's email for the job (REQUIRED)
- `REPO_MANIFESTS_URL`: The repository containing the repo manifest (default: "https://github.com/nexient-llc/module-manifests.git")
- `REPO_BRANCH`: The branch to point to for the repo manifest (default: "main")
- `REPO_MANIFEST`: The file in the repo manifest repository that contains the manifest's index (default: "tf_modules.xml")
- `COMPONENTS_DIR`: Directory for components to be installed to (default: "components")

## Notes

- There was issues related to CVE-2022-24765 resulting in the action instructing git that the parent directory created by github is safe to traverse with the container user.
- Occasionally github.com is not found within the known_hosts file resulting in the pipeline to fail. This is fixed by prefetching the github.com ssh key and pre-emptively adding it to the known_hosts file.

## Example

```yaml
name: "Example Workflow"
on: [push, pull_request]

jobs:
  configure:
    runs-on: nexient-llc/platform-images
    steps:
    - name: checkout source
      uses: actions/checkout@master

    - name: Run make configure
      uses: ./.github/actions/configure
      with:
        IS_PIPELINE: true
        JOB_NAME: "job"
        JOB_EMAIL: "job@job.job"
```
