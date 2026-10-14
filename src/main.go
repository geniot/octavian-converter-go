package main

func main() {
	tk := NewToolkitWithoutResourcePath()
	println(tk.Version())
}
