# Protocol Documentation
<a name="top"/>

## Table of Contents
* [proto-kvs.proto](#proto-kvs.proto)
 * [DeleteRequest](#proto.DeleteRequest)
 * [DeleteResponse](#proto.DeleteResponse)
 * [Entry](#proto.Entry)
 * [GetRequest](#proto.GetRequest)
 * [GetResponse](#proto.GetResponse)
 * [ListRequest](#proto.ListRequest)
 * [ListResponse](#proto.ListResponse)
 * [ListResponse.StoreEntry](#proto.ListResponse.StoreEntry)
 * [SetRequest](#proto.SetRequest)
 * [SetResponse](#proto.SetResponse)
 * [WatchRequest](#proto.WatchRequest)
 * [Kvs](#proto.Kvs)
* [Scalar Value Types](#scalar-value-types)

<a name="proto-kvs.proto"/>
<p align="right"><a href="#top">Top</a></p>

## proto-kvs.proto



<a name="proto.DeleteRequest"/>
### DeleteRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) | optional |  |


<a name="proto.DeleteResponse"/>
### DeleteResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="proto.Entry"/>
### Entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) | optional |  |
| value | [string](#string) | optional |  |


<a name="proto.GetRequest"/>
### GetRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) | optional |  |


<a name="proto.GetResponse"/>
### GetResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) | optional |  |


<a name="proto.ListRequest"/>
### ListRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="proto.ListResponse"/>
### ListResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| store | [ListResponse.StoreEntry](#proto.ListResponse.StoreEntry) | repeated |  |


<a name="proto.ListResponse.StoreEntry"/>
### ListResponse.StoreEntry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) | optional |  |
| value | [string](#string) | optional |  |


<a name="proto.SetRequest"/>
### SetRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) | optional |  |
| value | [string](#string) | optional |  |


<a name="proto.SetResponse"/>
### SetResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="proto.WatchRequest"/>
### WatchRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prefix | [string](#string) | optional |  |





<a name="proto.Kvs"/>
### Kvs


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#proto.GetRequest) | [GetResponse](#proto.GetResponse) |  |
| Set | [SetRequest](#proto.SetRequest) | [SetResponse](#proto.SetResponse) |  |
| Delete | [DeleteRequest](#proto.DeleteRequest) | [DeleteResponse](#proto.DeleteResponse) |  |
| List | [ListRequest](#proto.ListRequest) | [ListResponse](#proto.ListResponse) |  |
| Watch | [WatchRequest](#proto.WatchRequest) | [Entry](#proto.Entry) |  |



<a name="scalar-value-types"/>
## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double"/> double |  | double | double | float |
| <a name="float"/> float |  | float | float | float |
| <a name="int32"/> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64"/> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32"/> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64"/> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32"/> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64"/> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32"/> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64"/> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32"/> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64"/> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool"/> bool |  | bool | boolean | boolean |
| <a name="string"/> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes"/> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
