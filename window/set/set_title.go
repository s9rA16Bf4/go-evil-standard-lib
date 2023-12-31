package set

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Sets the title of the window that appears
func Title(new_title string, data_object *json.Json_t) {
	new_title = gotools.EraseDelimiter(new_title, []string{"\""}, -1)

	data_object.Set_title(new_title)

}
