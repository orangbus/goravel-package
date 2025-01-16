package facades

import (
	"github.com/orangbus/axios"
	"github.com/orangbus/axios/contracts"
	"log"
)

func Axios() contracts.Axios {
	instance, err := axios.App.Make(axios.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Axios)
}
