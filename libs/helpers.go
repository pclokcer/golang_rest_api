package libs

import "github.com/gin-gonic/gin"

func Find(array []string, val string) (int, bool) {
	for i, item := range array {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func GetLocalizer(c *gin.Context) interface{} {
	localizer, ok := c.Get("localizer")

	if ok {
		return localizer
	}

	return localizer
}
