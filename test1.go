/*testing
*/

package controllers

import (
	"fmt"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "HIIII"})

}
func PerformPostJsonRequest() {
	const myurl = "https://localhost:8000/post"
	//fake json payload
	requestBody := strings.NewReader(`
	
        {
			"coursename":"let's go with golang",
			"price":0,
			"platform":"github.com/sindhusweety"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
