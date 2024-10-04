package main

func main() {
	secrets := Secrets{}
	storage := NewStorage[Secrets]("./secrets.json")
	storage.Load(&secrets)
	cmdFlags := NewCommandFlags()
	cmdFlags.Execute(&secrets)
	storage.Save(secrets)
}
