package model

type Profile struct {
	RemoteRepo string
	LocalRepo  string
	UserName   string
	Password   string
}

var CurrentProfile = Profile{}
