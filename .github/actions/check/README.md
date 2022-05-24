## Example

```yaml
name: "Example Workflow"
on: [push, pull_request]

jobs:
  check:
    runs-on: nexient-llc/platform-images
    steps:
    - name: checkout source
      uses: actions/checkout@master
    - name: Run make check
      uses: ./.github/actions/check
```
