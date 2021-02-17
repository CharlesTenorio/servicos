import (
	"github.com/CharlesTenorio/servicos/models"
	"net/http"
	"github.com/labstack/echo"
)

func GetLigas(c echo.Context) error {
	db := db.DbManager()
	liga := []models.Liga{}
	db.Find(&liga)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, liga)
}