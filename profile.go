package whatsapp

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gabstv/go-whatsapp/binary"
)

// Pictures must be JPG 640x640 and 96x96, respectively
func (wac *Conn) UploadProfilePic(image, preview []byte) (<-chan string, error) {
	tag := fmt.Sprintf("%d.--%d", time.Now().Unix(), wac.msgCount*19)
	n := binary.Node{
		Description: "action",
		Attributes: map[string]string{
			"type":  "set",
			"epoch": strconv.Itoa(wac.msgCount),
		},
		Content: []interface{}{
			binary.Node{
				Description: "picture",
				Attributes: map[string]string{
					"id":   tag,
					"jid":  wac.info.Wid,
					"type": "set",
				},
				Content: []binary.Node{
					{
						Description: "image",
						Attributes:  nil,
						Content:     image,
					},
					{
						Description: "preview",
						Attributes:  nil,
						Content:     preview,
					},
				},
			},
		},
	}
	return wac.writeBinary(n, profile, 136, tag)
}
