## Notes for "Lets Go Further"

### 3 - Sending JSON responses
- ```func Marshal(v any) ([]byte, error)``` -> pass native Go object and returns JSON representation in []byte
- ```json.Encoder``` -> JSON is created and written to the http.ResponseWriter in a single step, which means there’s no opportunity to set HTTP response headers conditionally
- All struct attributes in Movie are exported (pubic), so that they are visible to encoding/json
- Struct tag directives
  - ```omitempty``` directive hides a field in the JSON output if and only if the struct field value is empty
  - ```-``` never show particular field in JSON
  - ```string``` converts to JSON string (only works for int*, uint*, float* or bool)
- ```json.MarshalIndent(data, "", "\t")``` -> Indents for raw response (but larger in bytes + time)
- Enveloping responses
- When Go encodes JSON, it looks to see if the type has a ```MarshalJSON() ([]byte, error)``` method implemented on it
- ```func (r Runtime) MarshalJSON() ([]byte, error)``` -> value reciver r instead of pointer. The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.
  - could also create custom MarshalJSON on Movie struct
