config, err := loadConfig()
if err != nil {
	log.Fatalf("Could not load configuration: %v", err.Error())
}

err = loadGridsFS(config)
if err != nil {
	log.Fatalf("Could not load databases: %v", err.Error())
}

err = loadPaths(config)
if err != nil {
	log.Fatalf("Could not load paths: %v", err.Error())
}