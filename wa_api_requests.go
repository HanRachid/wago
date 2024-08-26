package main
import "fmt"
var wa_id = "hello"
var url string = fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages",wa_id)



func SendTextMessage(){
	fmt.Println(url)
}
