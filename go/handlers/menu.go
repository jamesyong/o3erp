package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/jamesyong/o3erp/frontend/helper"
	"github.com/jamesyong/o3erp/frontend/sessions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
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

func MenuHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := `<menuGroup>
	<menu name="AP" directory="true" p="ACCOUNTING_VIEW">
		<url type='iframe'>/ap/control/main</url>
    </menu>
	<menu name="AR" directory="true" p="ACCOUNTING_VIEW">
		<url type='iframe'>/ar/control/main</url>
    </menu>
	<menu name="Accounting" directory="true" p="ACCOUNTING_VIEW">
		<url type='iframe'>/accounting/control/main</url>
    </menu>
	<menu name="Catalog" directory="true" p="CATALOG_VIEW">
		<url type='iframe'>/catalog/control/main</url>
    </menu>
	<menu name="Content" directory="true" p="CONTENT_VIEW">
		<url type='iframe'>/content/control/main</url>
    </menu>
	<menu name="Facility" directory="true" p="FACILITY_VIEW">
		<url type='iframe'>/facility/control/main</url>
    </menu>
	<menu name="HR" directory="true" p="HUMANRES_VIEW">
		<url type='iframe'>/humanres/control/main</url>
    </menu>
	<menu name="Manufacturing" directory="true" p="MANUFACTURING_VIEW">
		<url type='iframe'>/manufacturing/control/main</url>
    </menu>
	<menu name="Marketing" directory="true" p="MARKETING_VIEW">
		<url type='iframe'>/marketing/control/main</url>
    </menu>
	<menu name="Order" directory="true" p="ORDERMGR_VIEW">
		<url type='iframe'>/ordermgr/control/main</url>
    </menu>
	<menu name="Party" directory="true" p="PARTYMGR_VIEW">
		<url type='iframe'>/partymgr/control/main</url>
    </menu>
	<menu name="SFA" directory="true" p="SFA_VIEW">
		<url type='iframe'>/sfa/control/main</url>
    </menu>
	<menu name="WorkEffort" directory="true" p="WORKEFFORTMGR_VIEW">
		<url type='iframe'>/workeffort/control/main</url>
    </menu>
	<menu name="Business Intelligence" directory="true" p="BI_VIEW">
		<url type='iframe'>/bi/control/main</url>
    </menu>
	<menu name="WebTools" directory="true" p="WEBTOOLS_VIEW">
		<url type='iframe'>/webtools/control/main</url>
    </menu>
	</menuGroup>`

	menusGroup := []MenuGroup{}
	err := xml.Unmarshal([]byte(data), &menusGroup)
	if err != nil {
		fmt.Printf("error during unmarshal: %menus", err)
		return
	}

	// print in xml
	/*
		output, err := xml.MarshalIndent(menus, "  ", "    ")
		if err != nil {
			fmt.Printf("error during marshal: %v\n", err)
		}
		os.Stdout.Write(output) */

	session, _ := sessions.SessionStore.Get(r, "session-name")
	userLoginId := session.Values[sessions.USER_LOGIN_ID]
	if userLoginId != nil {

		permissionSet := make(map[string]struct{})
		labelSet := make(map[string]struct{})
		iterateMenuGetInfo(menusGroup[0].Menus, &permissionSet, &labelSet)

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

		permissionMap, err := helper.RunThriftService(helper.GetHasPermissionFunction(userLoginId.(string), permissions))
		if err != nil {
			log.Println("error: ", err)
		}
		labelMap, err := helper.RunThriftService(helper.GetMessageMapFunction(userLoginId.(string), "CommonUiLabels", labels))
		if err != nil {
			log.Println("error: ", err)
		}

		counter := 1
		var buffer bytes.Buffer
		buffer.WriteString("{ id:'root', name:'root'}")
		iterateMenu(menusGroup[0].Menus, &buffer, "root", &counter, permissionMap, labelMap)
		w.Header().Set("Content-Type", "text/json")
		w.Write([]byte("[" + buffer.String() + "]"))

	}

}

func iterateMenu(menus []Menu, buffer *bytes.Buffer, parent string, counter *int, permissionMap map[string]string, labelMap map[string]string) {
	for _, menu := range menus {
		if menu.P == "" || permissionMap[menu.P] == "true" {
			buffer.WriteString(",{")
			if menu.Id == "" {
				*counter = *counter + 1
				menu.Id = strconv.Itoa(*counter)
			}
			buffer.WriteString(" id:'" + menu.Id + "'")
			if menu.Name != "" {
				buffer.WriteString(", name:'" + labelMap[menu.Name] + "'")
			}
			if menu.P != "" {
				buffer.WriteString(", p:'" + menu.P + "'")
			}
			if &menu.Url != nil {
				buffer.WriteString(", url:'" + menu.Url.Href + "'")
				if menu.Url.Type != "" {
					buffer.WriteString(", urlType:'" + menu.Url.Type + "'")
				}
			}
			if menu.Directory != false {
				buffer.WriteString(", directory:'" + btos(menu.Directory) + "'")
			}
			buffer.WriteString(", parent:'" + parent + "'")
			buffer.WriteString("}")
		}
		if menu.MenuItems != nil {
			iterateMenuItem(menu.MenuItems, buffer, menu.Id, counter, permissionMap, labelMap)
		}
	}
}

func iterateMenuItem(menuItems []MenuItem, buffer *bytes.Buffer, parent string, counter *int, permissionMap map[string]string, labelMap map[string]string) {
	for _, menuItem := range menuItems {
		if menuItem.P == "" || permissionMap[menuItem.P] == "true" {
			buffer.WriteString(",{")
			if menuItem.Id == "" {
				*counter = *counter + 1
				menuItem.Id = strconv.Itoa(*counter)
			}
			buffer.WriteString(" id:'" + menuItem.Id + "'")

			if menuItem.Name != "" {
				buffer.WriteString(", name:'" + labelMap[menuItem.Name] + "'")
			}
			if menuItem.P != "" {
				buffer.WriteString(", p:'" + menuItem.P + "'")
			}
			if &menuItem.Url != nil {
				buffer.WriteString(", url:'" + menuItem.Url.Href + "'")
				if menuItem.Url.Type != "" {
					buffer.WriteString(", urlType:'" + menuItem.Url.Type + "'")
				}
			}
			if menuItem.Directory != false {
				buffer.WriteString(", directory:'" + btos(menuItem.Directory) + "'")
			}
			buffer.WriteString(", parent:'" + parent + "'")
			buffer.WriteString("}")
		}
		if menuItem.MenuItems != nil {
			iterateMenuItem(menuItem.MenuItems, buffer, menuItem.Id, counter, permissionMap, labelMap)
		}
	}
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

func btos(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
