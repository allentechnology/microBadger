package main

import (
	"fmt"
)

type slot struct {
	Id              string
	AssignedBadge   string
	AvailableBadges map[string]*microBadge
}

type microBadge struct {
	Id          string
	Name        string
	Description string
	ImgURL      string
	Category    string
	Selected    bool
}

func (mb *microBadge) UpdateMB(newMB *microBadge) {
	mb.Id = newMB.Id
	mb.Name = newMB.Name
	mb.Description = newMB.Description
	mb.ImgURL = newMB.ImgURL
	mb.Category = newMB.Category
}

func (mb *microBadge) String() string {
	return fmt.Sprintf("[Id: %s, Name: %s,Description: %s, ImgURL: %s, Category: %s]", mb.Id, mb.Name, mb.Description, mb.ImgURL, mb.Category)
}

func (mb *microBadge) SetImg(img string) {
	mb.ImgURL = img
}
