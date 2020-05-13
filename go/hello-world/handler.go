package function

import (
	"fmt"
)

func Handle(req []byte) string {
	return fmt.Sprintf("Hello Rex: %s", string(req))
}
