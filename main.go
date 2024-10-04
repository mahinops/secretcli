package main

func main() {
	secrets := Secrets{}
	cmdFlags := NewCommandFlags()
	cmdFlags.Execute(&secrets)
}
