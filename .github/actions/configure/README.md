## Inputs

- `JOB_NAME`: Sets the git user's name for the job (REQUIRED)
- `JOB_EMAIL`: Sets the git user's email for the job (REQUIRED)
- `IS_PIPELINE`: Denote if running in a pipeline context (default: true)

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
