package main

import (
	"fmt"
	"sort"
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
}

func (mb *microBadge) String() string {
	return fmt.Sprintf("[Id: %s, Name: %s,Description: %s, ImgURL: %s, Category: %s]", mb.Id, mb.Name, mb.Description, mb.ImgURL, mb.Category)
}

func (mb *microBadge) SetImg(img string) {
	mb.ImgURL = img
}

type ByDescription []*microBadge

func (mb ByDescription) Len() int {
	return len(mb)
}

func (mb ByDescription) Swap(i, j int) {
	mb[i], mb[j] = mb[j], mb[i]
}

func (mb ByDescription) Less(i, j int) bool {
	return mb[i].Description < mb[j].Description
}

func mbSort(mb []*microBadge) {
	sort.Sort(ByDescription(mb))
}
