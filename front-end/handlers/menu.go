package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/julienschmidt/httprouter"
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
	Role      string     `xml:"role,attr" json:"role"`
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
	<menu name="AP" directory="true">
		<url type='iframe'>/ap/control/main</url>
    </menu>
	<menu name="AR" directory="true">
		<url type='iframe'>/ar/control/main</url>
    </menu>
	<menu name="Accounting" directory="true">
		<url type='iframe'>/accounting/control/main</url>
    </menu>
	<menu name="Catalog" directory="true">
		<url type='iframe'>/catalog/control/main</url>
    </menu>
	<menu name="Content" directory="true">
		<url type='iframe'>/content/control/main</url>
    </menu>
	<menu name="Facility" directory="true">
		<url type='iframe'>/facility/control/main</url>
    </menu>
	<menu name="HR" directory="true">
		<url type='iframe'>/humanres/control/main</url>
    </menu>
	<menu name="Manufacturing" directory="true">
		<url type='iframe'>/manufacturing/control/main</url>
    </menu>
	<menu name="Marketing" directory="true">
		<url type='iframe'>/marketing/control/main</url>
    </menu>
	<menu name="Order" directory="true">
		<url type='iframe'>/ordermgr/control/main</url>
    </menu>
	<menu name="Party" directory="true">
		<url type='iframe'>/partymgr/control/main</url>
    </menu>
	<menu name="SFA" directory="true">
		<url type='iframe'>/sfa/control/main</url>
    </menu>
	<menu name="Work Effort" directory="true">
		<url type='iframe'>/workeffort/control/main</url>
    </menu>
	<menu name="Business Intelligence" directory="true">
		<url type='iframe'>/bi/control/main</url>
    </menu>
	<menu name="Web Tools" directory="true">
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

	counter := 1
	var buffer bytes.Buffer
	buffer.WriteString("{ id:'root', name:'root'}")
	iterateMenu(menusGroup[0].Menus, &buffer, "root", &counter)
	w.Header().Set("Content-Type", "text/json")
	w.Write([]byte("[" + buffer.String() + "]"))
}

func iterateMenu(menus []Menu, buffer *bytes.Buffer, parent string, counter *int) {
	for _, menu := range menus {
		buffer.WriteString(",{")
		if menu.Id == "" {
			*counter = *counter + 1
			menu.Id = strconv.Itoa(*counter)
		}
		buffer.WriteString(" id:'" + menu.Id + "'")
		if menu.Name != "" {
			buffer.WriteString(", name:'" + menu.Name + "'")
		}
		if menu.Role != "" {
			buffer.WriteString(", role:'" + menu.Role + "'")
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
		if menu.MenuItems != nil {
			iterateMenuItem(menu.MenuItems, buffer, menu.Id, counter)
		}
	}
}

func iterateMenuItem(menuItems []MenuItem, buffer *bytes.Buffer, parent string, counter *int) {
	for _, menuItem := range menuItems {
		buffer.WriteString(",{")
		if menuItem.Id == "" {
			*counter = *counter + 1
			menuItem.Id = strconv.Itoa(*counter)
		}
		buffer.WriteString(" id:'" + menuItem.Id + "'")

		if menuItem.Name != "" {
			buffer.WriteString(", name:'" + menuItem.Name + "'")
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
		if menuItem.MenuItems != nil {
			iterateMenuItem(menuItem.MenuItems, buffer, menuItem.Id, counter)
		}
	}
}

func btos(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
