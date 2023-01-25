package main

import "fmt"

type BaseFile interface {
	GetName() string
	PrintList()
}

type File struct {
	name string
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) PrintList() {
	fmt.Println(f.name)
}

type Directory struct {
	includedFiles []BaseFile
	name          string
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) Add(entryFile BaseFile) {
	d.includedFiles = append(d.includedFiles, entryFile)
}

func (d *Directory) PrintList() {
	fmt.Println()
	for _, file := range d.includedFiles {
		fmt.Print(d.GetName() + "/")
		file.PrintList()
	}
}

func main() {
	music := &Directory{name: "MUSIC"}
	maherZain := &Directory{name: "Maher Zain"}

	track1 := &File{name: "Rahmatun lil alameen.mp3"}
	track2 := &File{name: "InsyaAllah.mp3"}
	track3 := &File{name: "Ramadhan.mp3"}

	music.Add(maherZain)

	maherZain.Add(track1)
	maherZain.Add(track2)
	maherZain.Add(track3)

	music.PrintList()
}