
# Introduction
This is a sub-module.
To create a release for this sub-module:
<pre>
git tag modules/remote/v1.9.9
git push origin modules/remote/v1.9.9
</pre>

> **Note 1**: This module depends on rsc.io/sampler v1.3.1
<pre>
module github.com/philongct/golessions/modules/remote

go 1.12

<b>require rsc.io/sampler v1.3.1</b>
</pre>

> **Note 2**: Sub-module is not recommended. You should follow single module per repository

