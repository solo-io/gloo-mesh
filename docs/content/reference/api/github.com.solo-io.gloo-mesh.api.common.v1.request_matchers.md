
---

title: "request_matchers.proto"

---

## Package : `common.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for request_matchers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## request_matchers.proto


## Table of Contents
  - [HeaderMatcher](#common.mesh.gloo.solo.io.HeaderMatcher)
  - [StatusCodeMatcher](#common.mesh.gloo.solo.io.StatusCodeMatcher)

  - [StatusCodeMatcher.Comparator](#common.mesh.gloo.solo.io.StatusCodeMatcher.Comparator)






<a name="common.mesh.gloo.solo.io.HeaderMatcher"></a>

### HeaderMatcher
Describes a matcher against HTTP request headers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | string |  | Specify the name of the header in the request. |
  | value | string |  | Specify the value of the header. If the value is absent a request that has the name header will match, regardless of the header’s value. |
  | regex | bool |  | Specify whether the header value should be treated as regex. |
  | invertMatch | bool |  | If set to true, the result of the match will be inverted. Defaults to false.<br>Examples:<br>- name=foo, invert_match=true: matches if no header named `foo` is present - name=foo, value=bar, invert_match=true: matches if no header named `foo` with value `bar` is present - name=foo, value=``\d{3}``, regex=true, invert_match=true: matches if no header named `foo` with a value consisting of three integers is present. |
  





<a name="common.mesh.gloo.solo.io.StatusCodeMatcher"></a>

### StatusCodeMatcher
Describes a matcher against HTTP response status codes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | uint32 |  | The status code value to match against. |
  | comparator | [common.mesh.gloo.solo.io.StatusCodeMatcher.Comparator]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.common.v1.request_matchers#common.mesh.gloo.solo.io.StatusCodeMatcher.Comparator" >}}) |  | The comparison type used for matching. |
  




 <!-- end messages -->


<a name="common.mesh.gloo.solo.io.StatusCodeMatcher.Comparator"></a>

### StatusCodeMatcher.Comparator


| Name | Number | Description |
| ---- | ------ | ----------- |
| EQ | 0 | Strict equality. |
| GE | 1 | Greater than or equal to. |
| LE | 2 | Less than or equal to. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

