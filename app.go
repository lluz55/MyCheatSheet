package main

// App ...
type App struct {
	Version string
	Time    int64
	Cheats  []Cheats
}

type Cheats struct {
	Name   string
	Notes  []Cheat
	Active bool
}

type Cheat struct {
	Name        string
	Description string
	Tags        []string
}
