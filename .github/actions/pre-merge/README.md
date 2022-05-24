## Inputs

* `branch`: target branch to conduct pre-merge check with (default: github.event.pull_request.base.ref)

## Example

```yaml
name: "Example Workflow"
on: [pull_request]

jobs:
  configure:
    runs-on: nexient-llc/platform-images
    steps:
    - name: checkout source
      uses: actions/checkout@master
    - name: Run pre-merge action
      uses: ./.github/actions/pre-merge
      with:
        branch: "main"
```
