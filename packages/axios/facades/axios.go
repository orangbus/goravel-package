package facades

import (
	"log"

	"goravel/packages/axios"
	"goravel/packages/axios/contracts"
)

func Axios() contracts.Axios {
	instance, err := axios.App.Make(axios.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Axios)
}
