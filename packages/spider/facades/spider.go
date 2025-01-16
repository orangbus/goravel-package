package facades

import (
	"log"

	"github.com/orangbus/goravel-spider"
	"github.com/orangbus/goravel-spider/contracts"
)

func Spider() contracts.Spider {
	instance, err := spider.App.Make(spider.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Spider)
}
