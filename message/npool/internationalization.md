# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/internationalization.proto](#npool/internationalization.proto)
    - [AddLangRequest](#internationalization.v1.AddLangRequest)
    - [AddLangResponse](#internationalization.v1.AddLangResponse)
    - [CreateMessageRequest](#internationalization.v1.CreateMessageRequest)
    - [CreateMessageResponse](#internationalization.v1.CreateMessageResponse)
    - [CreateMessagesRequest](#internationalization.v1.CreateMessagesRequest)
    - [CreateMessagesResponse](#internationalization.v1.CreateMessagesResponse)
    - [GetLangsRequest](#internationalization.v1.GetLangsRequest)
    - [GetLangsResponse](#internationalization.v1.GetLangsResponse)
    - [GetMessageByMessageIDRequest](#internationalization.v1.GetMessageByMessageIDRequest)
    - [GetMessageByMessageIDResponse](#internationalization.v1.GetMessageByMessageIDResponse)
    - [GetMessagesRequest](#internationalization.v1.GetMessagesRequest)
    - [GetMessagesResponse](#internationalization.v1.GetMessagesResponse)
    - [Lang](#internationalization.v1.Lang)
    - [Message](#internationalization.v1.Message)
    - [UpdateMessageRequest](#internationalization.v1.UpdateMessageRequest)
    - [UpdateMessageResponse](#internationalization.v1.UpdateMessageResponse)
    - [UpdateMessagesRequest](#internationalization.v1.UpdateMessagesRequest)
    - [UpdateMessagesResponse](#internationalization.v1.UpdateMessagesResponse)
    - [VersionResponse](#internationalization.v1.VersionResponse)
  
    - [Internationalization](#internationalization.v1.Internationalization)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/internationalization.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/internationalization.proto



<a name="internationalization.v1.AddLangRequest"></a>

### AddLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Lang | [string](#string) |  |  |
| Logo | [string](#string) |  |  |






<a name="internationalization.v1.AddLangResponse"></a>

### AddLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization.v1.Lang) |  |  |






<a name="internationalization.v1.CreateMessageRequest"></a>

### CreateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization.v1.Message) |  |  |






<a name="internationalization.v1.CreateMessageResponse"></a>

### CreateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization.v1.Message) |  |  |






<a name="internationalization.v1.CreateMessagesRequest"></a>

### CreateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization.v1.Message) | repeated |  |






<a name="internationalization.v1.CreateMessagesResponse"></a>

### CreateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization.v1.Message) | repeated |  |






<a name="internationalization.v1.GetLangsRequest"></a>

### GetLangsRequest







<a name="internationalization.v1.GetLangsResponse"></a>

### GetLangsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Lang](#internationalization.v1.Lang) | repeated |  |






<a name="internationalization.v1.GetMessageByMessageIDRequest"></a>

### GetMessageByMessageIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| MessageID | [string](#string) |  |  |






<a name="internationalization.v1.GetMessageByMessageIDResponse"></a>

### GetMessageByMessageIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization.v1.Message) |  |  |






<a name="internationalization.v1.GetMessagesRequest"></a>

### GetMessagesRequest







<a name="internationalization.v1.GetMessagesResponse"></a>

### GetMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization.v1.Message) | repeated |  |






<a name="internationalization.v1.Lang"></a>

### Lang



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Lang | [string](#string) |  |  |
| Logo | [string](#string) |  |  |






<a name="internationalization.v1.Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| MessageID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="internationalization.v1.UpdateMessageRequest"></a>

### UpdateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization.v1.Message) |  |  |






<a name="internationalization.v1.UpdateMessageResponse"></a>

### UpdateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization.v1.Message) |  |  |






<a name="internationalization.v1.UpdateMessagesRequest"></a>

### UpdateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization.v1.Message) | repeated |  |






<a name="internationalization.v1.UpdateMessagesResponse"></a>

### UpdateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization.v1.Message) | repeated |  |






<a name="internationalization.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="internationalization.v1.Internationalization"></a>

### Internationalization
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#internationalization.v1.VersionResponse) | Method Version |
| AddLang | [AddLangRequest](#internationalization.v1.AddLangRequest) | [AddLangResponse](#internationalization.v1.AddLangResponse) |  |
| GetLangs | [GetLangsRequest](#internationalization.v1.GetLangsRequest) | [GetLangsResponse](#internationalization.v1.GetLangsResponse) |  |
| CreateMessage | [CreateMessageRequest](#internationalization.v1.CreateMessageRequest) | [CreateMessageResponse](#internationalization.v1.CreateMessageResponse) |  |
| CreateMessages | [CreateMessagesRequest](#internationalization.v1.CreateMessagesRequest) | [CreateMessagesResponse](#internationalization.v1.CreateMessagesResponse) |  |
| UpdateMessage | [UpdateMessageRequest](#internationalization.v1.UpdateMessageRequest) | [UpdateMessageResponse](#internationalization.v1.UpdateMessageResponse) |  |
| UpdateMessages | [UpdateMessagesRequest](#internationalization.v1.UpdateMessagesRequest) | [UpdateMessagesResponse](#internationalization.v1.UpdateMessagesResponse) |  |
| GetMessages | [GetMessagesRequest](#internationalization.v1.GetMessagesRequest) | [GetMessagesResponse](#internationalization.v1.GetMessagesResponse) |  |
| GetMessageByMessageID | [GetMessageByMessageIDRequest](#internationalization.v1.GetMessageByMessageIDRequest) | [GetMessageByMessageIDResponse](#internationalization.v1.GetMessageByMessageIDResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

