package handlers

import (
	"encoding/xml"
	"fmt"
	"github.com/jamesyong/o3erp/go/helper"
	"github.com/jamesyong/o3erp/go/sessions"
	"github.com/jamesyong/o3erp/go/templating"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func DashboardViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	labels := []string{}
	urls := []string{}

	data := `<menuGroup>	<menu>
			<menu-item name="AP" p="ACCOUNTING_VIEW">
				<url>/ap/control/main</url>
			</menu-item>
			<menu-item name="AR" p="ACCOUNTING_VIEW">
				<url>/ar/control/main</url>
			</menu-item>
			<menu-item name="HR" p="HUMANRES_VIEW">
				<url>/humanres/control/main</url>
			</menu-item>
			<menu-item name="Marketing" p="MARKETING_VIEW">
				<url>/marketing/control/main</url>
			</menu-item>
			<menu-item name="Order" p="ORDERMGR_VIEW">
				<url>/ordermgr/control/main</url>
			</menu-item>
			<menu-item name="SFA" p="SFA_VIEW">
				<url>/sfa/control/main</url>
			</menu-item>
			<menu-item name="WorkEffort" p="WORKEFFORTMGR_VIEW">
				<url>/workeffort/control/main</url>
			</menu-item>
		</menu></menuGroup>`

	menusGroup := []helper.MenuGroup{}
	err := xml.Unmarshal([]byte(data), &menusGroup)
	if err != nil {
		fmt.Printf("error during unmarshal: %menus", err)
		return
	}

	session, _ := sessions.SessionStore.Get(r, "session-name")
	userLoginId := session.Values[sessions.USER_LOGIN_ID]
	if userLoginId != nil {

		permissionMap, labelMap, err := helper.GetMenuExtraInfo(menusGroup[0].Menus, userLoginId.(string))
		if err != nil {
			log.Println("error: ", err)
		}

		iterateList(menusGroup[0].Menus, permissionMap, labelMap, &urls, &labels)
	}

	mapDetail := make(map[string]interface{})
	mapDetail["url"] = urls[0]

	mapDetail["Labels"] = labels
	mapDetail["Links"] = urls
	templating.Render.HTML(w, http.StatusOK, "dashboard", mapDetail)
}

func iterateList(menus []helper.Menu, permissionMap map[string]string, labelMap map[string]string, urls *[]string, labels *[]string) {
	for _, menu := range menus {
		if menu.MenuItems != nil {
			iterateListItem(menu.MenuItems, permissionMap, labelMap, urls, labels)
		}
	}
}

func iterateListItem(menuItems []helper.MenuItem, permissionMap map[string]string, labelMap map[string]string, urls *[]string, labels *[]string) {
	for _, menuItem := range menuItems {
		if menuItem.P == "" || permissionMap[menuItem.P] == "true" {
			*labels = append(*labels, labelMap[menuItem.Name])
			*urls = append(*urls, menuItem.Url.Href)
		}
	}
}
