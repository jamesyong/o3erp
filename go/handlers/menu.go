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

func MenuHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := `<menuGroup>
	<menu name="AP" directory="true" p="ACCOUNTING_VIEW">
		<menu-item id="apAgreements" name="AccountingUiLabels#AccountingAgreements">
			<url type='iframe'>/ap/control/FindAgreement</url>
		</menu-item>
		<menu-item id="apAgreements" name="Agreements (Work in progress)">
			<url>/acctg_agreement</url>
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
		<menu-item name="Overview">
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
		<menu-item name="Overview">
			<url>/catalog/control/main</url>
		</menu-item>
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
	<menu id="MarketingAppBar" name="Marketing" directory="true" p="MARKETING_VIEW">
		<menu-item id="mkgDataSource" name="DataSource">
			<url type='iframe'>/marketing/control/FindDataSource</url>
		</menu-item>
        <menu-item id="mkgCampaign" name="MarketingUiLabels#MarketingCampaign">
			<url type='iframe'>/marketing/control/FindMarketingCampaign</url>
		</menu-item>
        <menu-item id="mkgTracking" name="MarketingUiLabels#MarketingTracking">
			<url type='iframe'>/marketing/control/FindTrackingCode</url>
		</menu-item>
        <menu-item id="mkgSegment" name="MarketingUiLabels#MarketingSegment">
			<url type='iframe'>/marketing/control/FindSegmentGroup</url>
		</menu-item>
        <menu-item id="mkgContactList" name="MarketingUiLabels#MarketingContactList">
			<url type='iframe'>/marketing/control/FindContactLists</url>
		</menu-item>
        <menu-item id="mkgReports" name="MarketingUiLabels#MarketingReports">
			<url type='iframe'>/marketing/control/MarketingReport</url>
		</menu-item>
    </menu>
	<menu name="Order" directory="true" p="ORDERMGR_VIEW">
		<menu-item id="request" name="OrderUiLabels#OrderRequests" p="ORDERMGR_VIEW | ORDERMGR_PURCHASE_VIEW">
            <url type='iframe'>/ordermgr/control/FindRequest</url>
        </menu-item>
        <menu-item id="quote" name="OrderUiLabels#OrderOrderQuotes" p="ORDERMGR_VIEW | ORDERMGR_PURCHASE_VIEW">
            <url type='iframe'>/ordermgr/control/FindQuote</url>
        </menu-item>
        <menu-item id="orderlist" name="OrderUiLabels#OrderOrderList" p="ORDERMGR_VIEW">
            <url type='iframe'>/ordermgr/control/orderlist</url>
        </menu-item>
        <menu-item id="findorders" name="OrderUiLabels#OrderFindOrder" p="ORDERMGR_VIEW">
            <url type='iframe'>/ordermgr/control/findorders</url>
        </menu-item>
        <menu-item id="orderentry" name="OrderUiLabels#OrderOrderEntry" p="ORDERMGR_CREATE | ORDERMGR_PURCHASE_CREATE">
            <url type='iframe'>/ordermgr/control/orderentry</url>
        </menu-item>
        <menu-item id="return" name="OrderUiLabels#OrderOrderReturns" p="ORDERMGR_RETURN">
            <url type='iframe'>/ordermgr/control/findreturn</url>
        </menu-item>
        <menu-item id="requirement" name="OrderUiLabels#OrderRequirements" p="ORDERMGR_VIEW | ORDERMGR_ROLE_VIEW">
            <url type='iframe'>/ordermgr/control/FindRequirements</url>
        </menu-item>
        <menu-item id="orderReports" name="CommonReports">
            <url type='iframe'>/ordermgr/control/OrderPurchaseReportOptions</url>
        </menu-item>
        <menu-item id="orderStats" name="CommonStats">
            <url type='iframe'>/ordermgr/control/orderstats</url>
        </menu-item>
    </menu>
	<menu name="Party" directory="true" p="PARTYMGR_VIEW">
		<menu-item id="ptyFind" name="PartyUiLabels#PartyParties">
			<url type='iframe'>/partymgr/control/findparty</url>
		</menu-item>
        <menu-item id="ptyMycomm" name="PartyUiLabels#PartyMyCommunications">
			<url type='iframe'>/partymgr/control/MyCommunicationEvents</url>
		</menu-item>
        <menu-item id="ptyComm" name="PartyUiLabels#PartyCommunications">
			<url type='iframe'>/partymgr/control/FindCommunicationEvents</url>
		</menu-item>
        <menu-item id="ptyVisits" name="PartyUiLabels#PartyVisits">
			<url type='iframe'>/partymgr/control/findVisits</url>
		</menu-item>
        <menu-item id="ptyLoggedinusers" name="PartyUiLabels#PartyLoggedInUsers">
			<url type='iframe'>/partymgr/control/listLoggedInUsers</url>
		</menu-item>
        <menu-item id="ptyClassification" name="PartyUiLabels#PartyClassifications">
			<url type='iframe'>/partymgr/control/showclassgroups</url>
		</menu-item>
        <menu-item id="ptySecurity" name="CommonSecurity" p="PARTYMGR_VIEW">
            <url type='iframe'>/partymgr/control/FindSecurityGroup</url>
        </menu-item>
        <menu-item id="addrmap" name="PartyUiLabels#PageTitleAddressMatchMap">
			<url type='iframe'>/partymgr/control/addressMatchMap</url>
		</menu-item>
        <menu-item id="partyinv" name="PartyUiLabels#PartyInvitation">
			<url type='iframe'>/partymgr/control/partyInvitation</url>
		</menu-item>
		
    </menu>
	<menu name="SFA" directory="true" p="SFA_VIEW">
		<menu-item id="sfaAccounts" name="MarketingUiLabels#SfaAcccounts">
			<url type='iframe'>/sfa/control/FindAccounts</url>
		</menu-item>
        <menu-item id="sfaContacts" name="MarketingUiLabels#SfaContacts">
			<url type='iframe'>/sfa/control/FindContacts</url>
		</menu-item>
        <menu-item id="sfaLeads" name="MarketingUiLabels#SfaLeads">
			<url type='iframe'>/sfa/control/FindLeads</url>
		</menu-item>
        <menu-item id="sfaCompetitors" name="MarketingUiLabels#SfaCompetitors">
			<url type='iframe'>#</url>
		</menu-item>
        <menu-item id="sfaEvents" name="MarketingUiLabels#SfaEvents">
			<url type='iframe'>/sfa/control/Events</url>
		</menu-item>
        <menu-item id="sfaDocuments" name="MarketingUiLabels#SfaDocuments">
			<url type='iframe'>#</url>
		</menu-item>
        <menu-item id="sfaForecast" name="MarketingUiLabels#SfaForecasts">
			<url type='iframe'>/sfa/control/FindSalesForecast</url>
		</menu-item>
        <menu-item id="sfaOpportunities" name="MarketingUiLabels#SfaOpportunities">
			<url type='iframe'>/sfa/control/FindSalesOpportunity</url>
		</menu-item>
			
    </menu>
	<menu name="WorkEffort" directory="true" p="WORKEFFORTMGR_VIEW">
		<menu-item id="weTask" name="WorkEffortUiLabels#WorkEffortTaskList">
			<url type='iframe'>/workeffort/control/mytasks</url>
		</menu-item>
        <menu-item id="weCalendar" name="WorkEffortUiLabels#WorkEffortCalendar">
			<url type='iframe'>/workeffort/control/calendar</url>
		</menu-item>
        <menu-item id="weMytime" name="WorkEffortUiLabels#WorkEffortTimesheetMyTime">
			<url type='iframe'>/workeffort/control/MyTimesheets</url>
		</menu-item>
        <menu-item id="weRequest" name="WorkEffortUiLabels#WorkEffortRequestList">
			<url type='iframe'>/workeffort/control/requestlist</url>
		</menu-item>
        <menu-item id="weWorkeffort" name="WorkEffortUiLabels#WorkEffortWorkEffort">
			<url type='iframe'>/workeffort/control/FindWorkEffort</url>
		</menu-item>
        <menu-item id="weTimesheet" name="WorkEffortUiLabels#WorkEffortTimesheet">
			<url type='iframe'>/workeffort/control/FindTimesheet</url>
		</menu-item>
        <menu-item id="weUserJobs" name="WorkEffortUiLabels#WorkEffortJobList">
			<url type='iframe'>/workeffort/control/UserJobs</url>
		</menu-item>
        <menu-item id="WorkEffortICalendar" name="WorkEffortUiLabels#WorkEffortICalendar">
			<url type='iframe'>/workeffort/control/FindICalendars</url>
		</menu-item>
    </menu>
	<menu name="Business Intelligence" directory="true" p="BI_VIEW">
		<url type='iframe'>/bi/control/main</url>
    </menu>
	<menu name="WebTools" directory="true" p="WEBTOOLS_VIEW">
		<menu-item id="wtOverview" name="Overview">
			<url type='iframe'>/webtools/control/main</url>
		</menu-item>
		<menu-item id="wtLogging" name="WebtoolsUiLabels#WebtoolsLogging">
            <url type='iframe'>/webtools/control/LogView</url>
        </menu-item>
        <menu-item id="wtCache" name="WebtoolsUiLabels#WebtoolsCacheMaintenance">
            <url type='iframe'>/webtools/control/FindUtilCache</url>
        </menu-item>
        <menu-item id="wtArtifact" name="WebtoolsUiLabels#WebtoolsArtifactInfo">
            <url type='iframe'>/webtools/control/ArtifactInfo</url>
        </menu-item>
        <menu-item id="wtEntity" name="WebtoolsUiLabels#WebtoolsEntityEngine">
            <url type='iframe'>/webtools/control/entitymaint</url>
        </menu-item>
        <menu-item id="wtService" name="WebtoolsUiLabels#WebtoolsServiceEngineTools">
            <url type='iframe'>/webtools/control/ServiceList</url>
        </menu-item>
        <menu-item id="wtImportExport" name="WebtoolsUiLabels#WebtoolsImportExport">
            <url type='iframe'>/webtools/control/xmldsdump</url>
        </menu-item>
        <menu-item id="wtStats" name="WebtoolsUiLabels#WebtoolsStatistics">
            <url type='iframe'>/webtools/control/StatsSinceStart</url>
        </menu-item>
        <menu-item id="wtConfiguration" name="WebtoolsUiLabels#WebtoolsConfiguration">
            <url type='iframe'>/webtools/control/myCertificates</url>
        </menu-item>
        <menu-item id="wtGeoManagement" name="WebtoolsUiLabels#WebtoolsGeoManagement">
            <url type='iframe'>/webtools/control/FindGeo</url>
        </menu-item>        
        <menu-item id="wtPortalAdmin" name="WebtoolsUiLabels#WebtoolsAdminPortalPage" p="PORTALPAGE_ADMIN">
            <url type='iframe'>/webtools/control/FindPortalPage</url>
        </menu-item>
        <menu-item id="wtSecurity" name="CommonSecurity" p="SECURITY_VIEW">
            <url type='iframe'>/webtools/control/security</url>
        </menu-item>
        <menu-item id="wtLayoutDemo" name="WebtoolsUiLabels#WebtoolsLayoutDemo">
            <url type='iframe'>/webtools/control/WebtoolsLayoutDemo</url>
        </menu-item>
    </menu>
	</menuGroup>`

	menusGroup := []helper.MenuGroup{}
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

		permissionMap, labelMap, err := helper.GetMenuExtraInfo(menusGroup[0].Menus, userLoginId.(string))
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

func iterateMenu(menus []helper.Menu, buffer *bytes.Buffer, parent string, counter *int, permissionMap map[string]string, labelMap map[string]string) {
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

func iterateMenuItem(menuItems []helper.MenuItem, buffer *bytes.Buffer, parent string, counter *int, permissionMap map[string]string, labelMap map[string]string) {
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

func btos(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
