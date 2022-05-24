## Example

```yaml
name: "Run Make Check"
on: push

jobs:
  dockerfile-test:
    runs-on: nexient-llc/platform-images
    steps:
    - name: checkout source
      uses: actions/checkout@master

    - name: Run Container Tests
      uses: ./.github/actions/check
```
