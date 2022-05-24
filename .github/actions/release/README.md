## Inputs

* ``: (default: )

## Example

```yaml
name: "Run Make Release"
on: pull_request

jobs:
  dockerfile-test:
    runs-on: nexient-llc/platform-images
    steps:
    - name: checkout source
      uses: actions/checkout@master
    - name: Run Make Release
      uses: ./.github/actions/release
```
