package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/jamesyong/o3erp/go/helper"
	"github.com/jamesyong/o3erp/go/sessions"
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
		
		<menu-item id="apAgreements" name="AccountingUiLabels#AccountingAgreements">
			<url type='iframe'>/ap/control/FindAgreement</url>
		</menu-item>
        <menu-item id="apInvoices" name="AccountingUiLabels#AccountingInvoicesMenu">
			<url type='iframe'>/ap/control/FindApInvoices</url>
		</menu-item>
        <menu-item id="apPayments" name="AccountingUiLabels#AccountingPaymentsMenu">
			<url type='iframe'>/ap/control/findPayments</url>
		</menu-item>
        <menu-item id="apPaymentGroups" name="AccountingUiLabels#AccountingApPaymentGroupMenu">
			<url type='iframe'>/ap/control/FindApPaymentGroups</url>
		</menu-item>
        <menu-item id="apFindVendors" name="AccountingUiLabels#AccountingApPageTitleFindVendors">
			<url type='iframe'>/ap/control/findVendors</url>
		</menu-item>
        <menu-item id="apReport" name="AccountingUiLabels#AccountingReports">
			<url type='iframe'>/ap/control/listReports</url>
		</menu-item>
		
    </menu>
	<menu id="ArAppBar" name="AR" directory="true" p="ACCOUNTING_VIEW">
		<url type='iframe'>/ar/control/main</url>
		<menu-item id="arAgreements" name="AccountingUiLabels#AccountingAgreements">
			<url type='iframe'>/ar/control/FindAgreement</url>
		</menu-item>
        <menu-item id="arInvoices" name="AccountingUiLabels#AccountingInvoicesMenu">
			<url type='iframe'>/ar/control/findInvoices</url>
		</menu-item>
        <menu-item id="arPayments" name="AccountingUiLabels#AccountingPaymentsMenu">
			<url type='iframe'>/ar/control/findPayments</url>
		</menu-item>
        <menu-item id="arPaymentGroups" name="AccountingUiLabels#AccountingArPaymentGroupMenu">
			<url type='iframe'>/ar/control/FindArPaymentGroups</url>
		</menu-item>
        <menu-item id="arReports" name="AccountingUiLabels#AccountingReports">
			<url type='iframe'>/ar/control/ListReports</url>
		</menu-item>
    </menu>
	<menu name="Accounting" directory="true" p="ACCOUNTING_VIEW">
		<url type='iframe'>/accounting/control/main</url>
		<menu-item name="CommonMain">
            <url type='iframe'>/accounting/control/main</url>
        </menu-item>
		<menu-item id="invoices" name="AccountingUiLabels#AccountingInvoicesMenu">
            <url type='iframe'>/accounting/control/findInvoices</url>
        </menu-item>
		<menu-item id="payments" name="AccountingUiLabels#AccountingPaymentsMenu">
            <url type='iframe'>/accounting/control/findPayments</url>
        </menu-item>
		<menu-item id="PaymentGroup" name="AccountingUiLabels#AccountingPaymentGroup">
            <url type='iframe'>/accounting/control/FindPaymentGroup</url>
        </menu-item>
		<menu-item id="transaction" p="MANUAL_PAYMENT | ACCOUNTING_CREATE" name="AccountingUiLabels#AccountingTransactions">
            <url type='iframe'>/accounting/control/FindGatewayResponses</url>
        </menu-item>
		<menu-item id="PaymentGatewayConfig" p="PAYPROC_ADMIN | ACCOUNTING_ADMIN" name="AccountingUiLabels#AccountingPaymentGatewayConfig">
            <url type='iframe'>/accounting/control/FindPaymentGatewayConfig</url>
        </menu-item>
        <menu-item id="billingaccount" name="AccountingUiLabels#AccountingBillingMenu">
            <url type='iframe'>/accounting/control/FindBillingAccount</url>
        </menu-item>
        <menu-item id="FindFinAccount" name="AccountingUiLabels#AccountingFinAccount">
            <url type='iframe'>/accounting/control/FinAccountMain</url>
        </menu-item>
        <menu-item id="TaxAuthorities" name="AccountingUiLabels#AccountingTaxAuthorities">
            <url type='iframe'>/accounting/control/FindTaxAuthority</url>
        </menu-item>
        <menu-item id="agreements" name="AccountingUiLabels#AccountingAgreements">
            <url type='iframe'>/accounting/control/FindAgreement</url>
        </menu-item>
        <menu-item id="ListFixedAssets" name="AccountingUiLabels#AccountingFixedAssets">
            <url type='iframe'>/accounting/control/ListFixedAssets</url>
        </menu-item>
        <menu-item id="GlobalGLSettings" name="AccountingUiLabels#AccountingGlobalGLSettings">
            <url type='iframe'>/accounting/control/globalGLSettings</url>
        </menu-item>
        <menu-item id="companies" name="AccountingUiLabels#AccountingOrgGlSettings">
            <url type='iframe'>/accounting/control/ListCompanies</url>
        </menu-item>		
    </menu>
	<menu id="CatalogAppBar" name="Catalog" directory="true" p="CATALOG_VIEW">
		<url type='iframe'>/catalog/control/main</url>
		
        <menu-item id="pCatalogs" name="ProductUiLabels#ProductCatalogs">
			<url type='iframe'>/catalog/control/FindCatalog</url>
		</menu-item>
        <menu-item id="pCategories" name="ProductUiLabels#ProductCategories">
			<url type='iframe'>/catalog/control/FindCategory</url>
		</menu-item>
        <menu-item id="pProducts" name="ProductUiLabels#ProductProducts">
			<url type='iframe'>/catalog/control/FindProduct</url>
		</menu-item>
        <menu-item id="pFeaturecats" name="ProductUiLabels#ProductFeatureCats">
			<url type='iframe'>/catalog/control/EditFeatureCategories</url>
		</menu-item>
        <menu-item id="pPromos" name="ProductUiLabels#ProductPromos">
			<url type='iframe'>/catalog/control/FindProductPromo</url>
		</menu-item>
        <menu-item id="pPricerules" name="ProductUiLabels#ProductPriceRules">
			<url type='iframe'>/catalog/control/FindProductPriceRules</url>
		</menu-item>
        <menu-item id="pStore" name="ProductUiLabels#ProductStores">
			<url type='iframe'>/catalog/control/FindProductStore</url>
		</menu-item>
        <menu-item id="pStoreGroup" name="ProductUiLabels#ProductProductStoreGroups">
			<url type='iframe'>/catalog/control/ListParentProductStoreGroup</url>
		</menu-item>
        <menu-item id="pThesaurus" name="ProductUiLabels#ProductThesaurus">
			<url type='iframe'>/catalog/control/editKeywordThesaurus</url>
		</menu-item>
        <menu-item id="pReviews" name="ProductUiLabels#ProductReviews">
			<url type='iframe'>/catalog/control/FindReviews</url>
		</menu-item>
        <menu-item id="pConfigs" name="ProductUiLabels#ProductConfigItems">
			<url type='iframe'>/catalog/control/FindProductConfigItems</url>
		</menu-item>
        <menu-item id="pSubscription" name="ProductUiLabels#ProductSubscriptions">
			<url type='iframe'>/catalog/control/FindSubscription</url>
		</menu-item>
        <menu-item id="pShipping" name="ProductUiLabels#ProductShipping">
			<url type='iframe'>/catalog/control/ListShipmentMethodTypes</url>
		</menu-item>
        <menu-item id="pImagemanagement" name="ProductUiLabels#ImageManagement">
			<url type='iframe'>/catalog/control/Imagemanagement</url>
		</menu-item>

    </menu>
	<menu name="Content" directory="true" p="CONTENT_VIEW">
		<url type='iframe'>/content/control/main</url>
    </menu>
	<menu id="FacilityAppBar" name="Facility" directory="true" p="FACILITY_VIEW">
		<url type='iframe'>/facility/control/main</url>
		<menu-item id="fFacility" name="ProductUiLabels#ProductFacilities">
			<url type='iframe'>/facility/control/FindFacility</url>
        </menu-item>			
        <menu-item id="fFacilityGroup" name="ProductUiLabels#ProductFacilityGroups">
			<url type='iframe'>/facility/control/FindFacilityGroup</url>
        </menu-item>			
        <menu-item id="fInventoryItemLabel" name="ProductUiLabels#ProductInventoryItemLabels">
			<url type='iframe'>/facility/control/FindInventoryItemLabels</url>
        </menu-item>			
        <menu-item id="fShipmentGatewayConfig" name="ProductUiLabels#FacilityShipmentGatewayConfig" p="PAYPROC_ADMIN">
            <url type='iframe'>/facility/control/FindShipmentGatewayConfig</url>
        </menu-item>
        <menu-item id="fShipment" name="ProductUiLabels#ProductShipments">
			<url type='iframe'>/facility/control/FindShipment</url>
		</menu-item>
        <menu-item id="fReports" name="CommonReports">
			<!--TODO required facilityId, so add dropdown to select facility -->
            <url type='iframe'>/facility/control/InventoryReports</url>
        </menu-item>
    </menu>
	<menu name="HR" directory="true" p="HUMANRES_VIEW">
		<url type='iframe'>/humanres/control/main</url>
		<menu-item name="CommonMain">
            <url type='iframe'>/humanres/control/main</url>
        </menu-item>
        	<menu-item id="Employees" name="HumanResUiLabels#HumanResEmployees">
            <url type='iframe'>/humanres/control/findEmployees</url>
        </menu-item>
		<menu-item id="Employment" name="HumanResUiLabels#HumanResEmployment">
            <url type='iframe'>/humanres/control/FindEmployments</url>
        </menu-item>
        <menu-item id="EmplPosition" name="HumanResUiLabels#HumanResEmployeePosition">
            <url type='iframe'>/humanres/control/FindEmplPositions</url>
        </menu-item>
        <menu-item id="PerfReview" name="HumanResUiLabels#HumanResPerfReview">
            <url type='iframe'>/humanres/control/FindPerfReviews</url>
        </menu-item>
        <menu-item id="EmplSkills" name="HumanResUiLabels#HumanResSkills">
            <url type='iframe'>/humanres/control/FindPartySkills</url>
        </menu-item>
        <menu-item id="PartyQual" name="HumanResUiLabels#HumanResPartyQualification">
            <url type='iframe'>/humanres/control/FindPartyQuals</url>
        </menu-item>
        <menu-item id="Recruitment" name="HumanResUiLabels#HumanResRecruitment">
            <url type='iframe'>/humanres/control/FindJobRequisitions</url>
        </menu-item>
        <menu-item id="Training" name="HumanResUiLabels#HumanResTraining">
            <url type='iframe'>/humanres/control/TrainingCalendar</url>
        </menu-item>
        <menu-item id="EmploymentApp" name="HumanResUiLabels#HumanResEmploymentApp">
            <url type='iframe'>/humanres/control/FindEmploymentApps</url>
        </menu-item>
        <menu-item id="PartyResume" name="HumanResUiLabels#HumanResPartyResume">
            <url type='iframe'>/humanres/control/FindPartyResumes</url>
        </menu-item>
        <menu-item id="Leave" name="HumanResUiLabels#HumanResEmplLeave">
            <url type='iframe'>/humanres/control/FindEmplLeaves</url>
        </menu-item>
        <menu-item id="GlobalHRSettings" name="HumanResUiLabels#HumanResGlobalHRSettings">
            <url type='iframe'>/humanres/control/globalHRSettings</url>
        </menu-item>
    </menu>
	<menu name="Manufacturing" directory="true" p="MANUFACTURING_VIEW">
		<url type='iframe'>/manufacturing/control/main</url>
		
		<menu-item id="mJobshop" name="ManufacturingUiLabels#ManufacturingJobShop" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/FindProductionRun</url>
        </menu-item>
        <menu-item id="mRouting" name="ManufacturingUiLabels#ManufacturingRouting" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/FindRouting</url>
        </menu-item>
        <menu-item id="mRoutingTask" name="ManufacturingUiLabels#ManufacturingRoutingTask" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/FindRoutingTask</url>
        </menu-item>
        <menu-item id="mCalendar" name="ManufacturingUiLabels#ManufacturingCalendar" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/FindCalendar</url>
        </menu-item>
        <menu-item id="mCosts" name="ManufacturingUiLabels#ManufacturingCostCalcs" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/EditCostCalcs</url>
        </menu-item>
        <menu-item id="mBom" name="ManufacturingUiLabels#ManufacturingBillOfMaterials" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/FindBom</url>
        </menu-item>
        <menu-item id="mMrp" name="ManufacturingUiLabels#ManufacturingMrp" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/FindInventoryEventPlan</url>
        </menu-item>
        <menu-item id="mShipmentPlans" name="ManufacturingUiLabels#ManufacturingShipmentPlans" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/WorkWithShipmentPlans</url>
        </menu-item>
        <menu-item id="mReports" name="ManufacturingUiLabels#ManufacturingReports" p="MANUFACTURING_CREATE">
            <url type='iframe'>/manufacturing/control/ManufacturingReports</url>
        </menu-item>
		
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

		labelMap, err := helper.RunThriftService(helper.GetMessageMapFunction(userLoginId.(string), labels))
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
