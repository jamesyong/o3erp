package helper

import (
	"encoding/xml"
	"log"
)

type MenuGroup struct {
	XMLName xml.Name `xml:"menuGroup"`
	Id      string
	Menus   []Menu `xml:"menu" json:",omitempty"`
}

type UrlGroup struct {
	XMLName struct{} `xml:"url"`
	Type    string   `xml:"type,attr"`
	Href    string   `xml:",chardata"`
}

type Menu struct {
	XMLName   xml.Name   `xml:"menu" json:"-"`
	Id        string     `xml:"id,attr" json:"id"`
	Name      string     `xml:"name,attr" json:"name"`
	P         string     `xml:"p,attr" json:"p"`
	Url       UrlGroup   `xml:"url" json:"url"`
	Directory bool       `xml:"directory,attr" json:"directory"`
	MenuItems []MenuItem `xml:"menu-item" json:",omitempty"`
}

type MenuItem struct {
	XMLName   xml.Name   `xml:"menu-item" json:"-"`
	Id        string     `xml:"id,attr" json:"id"`
	Name      string     `xml:"name,attr" json:"name"`
	P         string     `xml:"p,attr" json:"p"`
	Url       UrlGroup   `xml:"url" json:"url"`
	Directory bool       `xml:"directory,attr" json:"directory"`
	MenuItems []MenuItem `xml:"menu-item" json:",omitempty"`
}

func GetMenuExtraInfo(menus []Menu, userLoginId string) (map[string]string, map[string]string, error) {

	permissionSet := make(map[string]struct{})
	labelSet := make(map[string]struct{})
	iterateMenuGetInfo(menus, &permissionSet, &labelSet)

	// convert permission from set to array
	permissions := []string{}
	for k := range permissionSet {
		permissions = append(permissions, k)
	}
	// convert permission from set to array
	labels := []string{}
	for k := range labelSet {
		labels = append(labels, k)
	}

	permissionMap, err := RunThriftService(GetHasPermissionFunction(userLoginId, permissions))
	if err != nil {
		log.Println("error: ", err)
	}

	labelMap, err := RunThriftService(GetMessageMapFunction(userLoginId, labels))
	if err != nil {
		log.Println("error: ", err)
	}

	return permissionMap, labelMap, err
}

func iterateMenuGetInfo(menus []Menu, permissionMap *map[string]struct{}, labelMap *map[string]struct{}) {
	for _, menu := range menus {
		if menu.P != "" {
			(*permissionMap)[menu.P] = struct{}{}
		}
		if menu.Name != "" {
			(*labelMap)[menu.Name] = struct{}{}
		}
		if menu.MenuItems != nil {
			iterateMenuItemGetInfo(menu.MenuItems, permissionMap, labelMap)
		}
	}
}

func iterateMenuItemGetInfo(menuItems []MenuItem, permissionMap *map[string]struct{}, labelMap *map[string]struct{}) {
	for _, menuItem := range menuItems {
		if menuItem.P != "" {
			(*permissionMap)[menuItem.P] = struct{}{}
		}
		if menuItem.Name != "" {
			(*labelMap)[menuItem.Name] = struct{}{}
		}
		if menuItem.MenuItems != nil {
			iterateMenuItemGetInfo(menuItem.MenuItems, permissionMap, labelMap)
		}
	}
}
