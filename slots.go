package main

import (
	"fmt"
)

type slot struct {
	Id              string
	AssignedBadge   string
	AvailableBadges map[int]*microBadge
}

type microBadge struct {
	Id          string
	Name        string
	Description string
	ImgURL      string
	Category    string
}

func (mb *microBadge) String() string {
	return fmt.Sprintf("[Id: %s, Name: %s,Description: %s, ImgURL: %s, Category: %s]", mb.Id, mb.Name, mb.Description, mb.ImgURL, mb.Category)
}

func (mb *microBadge) SetImg(img string) {
	mb.ImgURL = img
}
