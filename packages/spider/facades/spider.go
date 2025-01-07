package facades

import (
	"log"

	"goravel/packages/spider"
	"goravel/packages/spider/contracts"
)

func Spider() contracts.Spider {
	instance, err := spider.App.Make(spider.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Spider)
}
