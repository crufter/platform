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

  ` + "```" + `javascript
const kittn = require('kittn');

let api = kittn.authorize('meowmeowmeow');
let kittens = api.kittens.get();
` + "```" + `

> The above command returns JSON structured like this:

` + "```" + `json
{ sadasas }
` + "```" + `

This endpoint retrieves all kittens.

### HTTP Request

` + "`" + `GET http://example.com/api/kittens` + "`" + `

### Query Parameters

Parameter | Default | Description
--------- | ------- | -----------
include_cats | false | If set to true, the result will also include cats.
available | true | If set to false, the result will include kittens that have already been adopted.

<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>
{{ end }}
`
