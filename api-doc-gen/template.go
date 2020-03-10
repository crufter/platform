package main

var serviceTemplate = `
---
weight: 11
title: {{ .serviceName }}
---

# {{ .serviceName }}

{{ range $rpc := .rpcs }}
## {{ $rpc.Name }}

` + "```" + `go
package main

import "github.com/micro/clients"

func main() {
	api := auth.Authorize("meowmeowmeow")

	_ = api.GetKittens()
}
` + "```" + `

` + "```" + `ruby
require 'kittn'

api = Kittn::APIClient.authorize!('meowmeowmeow')
api.kittens.get
` + "```" + `

` + "```" + `python
import kittn

api = kittn.authorize('meowmeowmeow')
api.kittens.get()
` + "```" + `

` + "```" + `shell
curl "http://example.com/api/kittens"
  -H "Authorization: meowmeowmeow"
  ` + "```" + `

> The above command returns JSON structured like this:

` + "```" + `json
{ 1111s }
` + "```" + `

  ` + "```" + `javascript
const kittn = require('kittn');

let api = kittn.authorize('meowmeowmeow');
let kittens = api.kittens.get();
` + "```" + `

> The above command returns JSON structured like this:

` + "```" + `json
{ sadasas }
` + "```" + `

{{ commentLines $rpc.Comment }}

### Request Parameters

Name |  Type | Description
--------- | --------- | ---------
{{ range $field := (getNormalFields $rpc.RequestType) }}{{ $field.Name }} | {{ $field.Type }} | {{ commentLines $field.Comment }}
{{ end }}

### Response Parameters

Name |  Type | Description
--------- | --------- | ---------
{{ range $field := (getNormalFields $rpc.ReturnsType) }}{{ $field.Name }} | {{ $field.Type }} | {{ commentLines $field.Comment }}
{{ end }}


<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>
{{ end }}
`
