package mapstruct

import "testing"
import "assert"

type Blog struct {
  Author string
  Title string
  Body string
  Rating int
}

func TestStructToMap(t *testing.T) {
  b := Blog{
    Author: "steven",
    Title: "stuff is cool",
    Body: "yep! truly indeed.",
    Rating: 5, // out of 10, tho, so...
  }

  m := ConvertToMap(b)

  assert.DeepEquals(t, m, map[string]interface{} {
    "Author": "steven",
    "Title": "stuff is cool",
    "Body": "yep! truly indeed.",
    "Rating": 5,
  })
}

func TestMapToStruct(t *testing.T) {
  m := map[string]interface{} {
    "Author": "steven",
    "Title": "stuff is cool",
    "Body": "yep! truly indeed.",
    "Rating": 5,
  }

  var b Blog
  ConvertToStruct(m, &b)

  assert.DeepEquals(t, b, Blog{
    Author: "steven",
    Title: "stuff is cool",
    Body: "yep! truly indeed.",
    Rating: 5, // out of 10, tho, so...
  })
}
