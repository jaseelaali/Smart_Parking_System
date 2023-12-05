package dep

import (
	"github.com/google/wire"
	"github.com/jaseelaali/Smart_Parking_System/config"
	"github.com/jaseelaali/Smart_Parking_System/database"
)

func InitializeApi(cfg config.Config) {
	wire.Build(
		database.ConnectDatabase(config),
	)
}
