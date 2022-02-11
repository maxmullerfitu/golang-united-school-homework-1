package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	message := emoji.Sprintf("Hello :world_map:! !")
	//fmt.Println(message)
	return message
}
