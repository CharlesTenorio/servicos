import (
	"github.com/CharlesTenorio/servicos/models"
	"net/http"
	"github.com/labstack/echo"
)

func GetTimes(c echo.Context) error {
	db := db.DbManager()
	time := []models.Time{}
	db.Find(&time)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, time)
}