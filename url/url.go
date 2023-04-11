package url

import (
	"github.com/NaufalDekha002/dekha-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                        //ujicoba panggil package musik
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/peneliti", controller.GetPeneliti)
	page.Get("/hasil teliti", controller.GetHasilTeliti)
	page.Get("/all", controller.GetAll)
}
