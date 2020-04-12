# MVS Processing

This is pseudo code for how MVS handles dependencies

From: `mvs vendor <lang>` (when no lang, loops over discovered)

### 1. Read Root Module

1. Read MVS / mod files
    1. `.mvsconfig` if exists (this will configure paths and behaviors for this module)
    1. `<lang>.mod` file, EXIT if does not exist, warn to init
    1. `<lang>.sum` file, if exists
1. If compare(sum, mod) == Same / Valid
    1. Check content is ok
    1. If not, EXIT and WARN
1. For each entry in {mod - sum}
    1. Fetch: cache -> remote
    1. CALL: ReadDepsModule





### FN: ReadDepsModule


1. Read MVS / mod files
    1. `.mvsconfig` if exists (this will configure paths and behaviors for this module)
    1. `<lang>.mod` file, EXIT if does not exist, warn to init
    1. `<lang>.sum` file, if exists
1. 




