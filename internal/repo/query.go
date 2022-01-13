package repo

const (
	insertCar                = `INSERT INTO "cars" ("car_type","name","color","speed_range","create_time","last_updated","id") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`
	insertFeature            = `INSERT INTO "features" ("name","car_id","id") VALUES ($1,$2,$3) ON CONFLICT ("id") DO UPDATE SET "car_id"="excluded"."car_id" RETURNING "id"`
	selectFromCarsWithID     = `SELECT * FROM "cars" WHERE id = $1`
	selectFromFeaturesWithID = `SELECT * FROM "features" WHERE "features"."car_id" = $1`
	selectFromCars           = `SELECT * FROM "cars" WHERE "cars"."car_type" = $1 AND "cars"."name" = $2 AND "cars"."color" = $3 AND "cars"."speed_range" = $4`
)
