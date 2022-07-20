### Quick Installation

To install the Picsmize library, Here is how you can install using command line,

```bash
$ go get github.com/picsmize/picsmize-go
```

### Quick Example

This Picsmize PHP library allows all the operations available with the [Picsmize API](https://picsmize.com/document/introduction). The following example uses image `fetch`, `compress`, `resize` and `filter` with different mode and get the output file directly with `toJSON()` method:

```go
package main

import (
	"fmt"

	"github.com/picsmize/picsmize-go"
)

func main() {

	pics, err := picsmize.Init("your-api-key")
	if err != nil {
		panic(err)
	}

	/**
	* Use of Fetch() method
	*/

	res, err := pics.Fetch("https://www.example.com/image.jpg").

		/**
		* Use of Compress() method with low mode
		*/

		Compress(picsmize.Options{
			"level": "low",
		}).

		/**
		* Use of Resize() method with auto mode
		* and width set to 400
		*/

		Resize("auto", picsmize.Options{
			"width": 400,
		}).

		/**
		* Use of Filter() blur method with gaussian mode
		* and value set to 10
		*/

		Filter("auto", picsmize.Options{
			"mode": "gaussian",
			"value": 10
		}).

		/**
		* Call ToJSON() on the final step and return the JSON response
		*/

		ToJSON()

	if err != nil {
		panic(err)
	}

	/**
	* You'll find the full JSON metadata within the `res` variable.
	*/

	fmt.Println(res)
}
```