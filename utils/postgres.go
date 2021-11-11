package utils

import(
	"fmt"
)

var PGConnection = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d",
	"localhost", "", "", "", 5432,
)