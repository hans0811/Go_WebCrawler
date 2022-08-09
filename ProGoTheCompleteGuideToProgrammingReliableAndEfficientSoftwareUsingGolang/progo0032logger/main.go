package main

import (
	"PROGO0032LOGGER/config"
	log "PROGO0032LOGGER/loggin"
	"PROGO0032LOGGER/services"
)

func writeMessage(logger log.Logger, cfg config.Configuration) {

	section, ok := cfg.GetSection("main")
	
	if(ok) {
		message, ok := section.GetString("message")

		if(ok){
			logger.Info(message)
		}else{
			logger.Panic("Cannot find configuration setting")
		}
	}else{
		logger.Panic("Config section not found")
	}

}

func main(){
	//writeMessage()

	services.RegisterDefaultServices()

	services.Call(writeMessage)

	// var cfg config.Configuration
	// services.GetService(&cfg)

	// var logger log.Logger
	// services.GetService(&logger)

	//writeMessage(logger, cfg)


	// var err error

	// cfg, err = config.Load("config.json")
	// if(err != nil){
	// 	panic(err)
	// }

	// var logger log.Logger = log.NewDefaultLogger(cfg)



	//writeMessage(logger, cfg)
}