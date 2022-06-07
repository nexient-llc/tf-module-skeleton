## Inputs

* `branch`: target branch to conduct pre-merge check with (default: github.event.pull_request.base.ref)

## Notes

- When using this action you need to make sure to checkout the entire history of the branch otherwise this step will fail. To do this make sure to change the actions/checkout fetch-depth input to 0.

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
      with:
        fetch-depth: 0
    - name: Run pre-merge action
      uses: ./.github/actions/pre-merge
      with:
        branch: "origin/main"
```
