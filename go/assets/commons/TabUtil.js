define(["dojox/layout/ContentPane", "dojo/topic", "dijit/registry", "dijit/layout/TabContainer"], 
	function(ContentPane, topic, registry){
	    return {
			setSubTabContent: function(pTabContainerId, item){
				var pane = registry.byId(pTabContainerId+":"+item.id);
				if (pane==null){
					if (item.urlType=='iframe'){
						pane = new ContentPane({
					        id: pTabContainerId+":"+item.id,
					        title: item.id,
							content: '<iframe src="'+item.url+'" style="border: 0; width: 100%; height: 100%"></iframe>',
							closable:true,
							onClose: function(){
							  topic.publish("tab/close", item.id);
	  			              return true;
					        }
					    });
					} else {
						pane = new ContentPane({
							id: pTabContainerId+":"+item.id,
					        title: item.id,
							href: item.url,
							closable:true,
							onClose: function(){
							  topic.publish("tab/close", item.id);
	  			              return true;
					        }
					    });
					}
					registry.byId(pTabContainerId).addChild(pane);
				}
				registry.byId(pTabContainerId).selectChild(pane);
			},
			setTabContent: function (item, hasSubTabContent){
				
				if (typeof hasSubTabContent == "undefined"){
					var hasSubTabContent = true;
				}
				if (hasSubTabContent){
					var tabContainer = registry.byId(item.id);
					if (tabContainer==null){
						tabContainer = new dijit.layout.TabContainer({
							id: item.id,
							title: item.name,
							// doLayout: false,
							nested: true,
							closable:true,
							onClose: function(){
							  topic.publish("tab/close", item.id);
	  			              return true;
					        }
						});
						// add the new sub tab container to our contentTabs widget
						registry.byId("contentTabs").addChild(tabContainer);
						
						if (item.urlType=='iframe'){
							var pane = new ContentPane({
						        id: item.id+":list",
						        title: "List",
								content: '<iframe src="'+item.url+'" style="border: 0; width: 100%; height: 100%"></iframe>'
						    });
							tabContainer.addChild(pane);
						} else {
							var pane = new ContentPane({
								id: item.id+":list",
						        title: "List",
								href: item.url
						    });
							tabContainer.addChild(pane);
						}
					}
					registry.byId("contentTabs").selectChild(tabContainer);
				} else {
					var pane = registry.byId(item.id);
					if (pane==null){
						if (item.urlType=='iframe'){
							pane = new ContentPane({
						        id: item.id,
						        title: item.name,
								content: '<iframe src="'+item.url+'" style="border: 0; width: 100%; height: 100%"></iframe>',
								closable:true,
								onClose: function(){
								  topic.publish("tab/close", item.id);
		  			              return true;
						        }
						    });
						} else {
							pane = new ContentPane({
								id: item.id,
						        title: item.name,
								href: item.url,
								closable:true,
								onClose: function(){
								  topic.publish("tab/close", item.id);
		  			              return true;
						        }
						    });
						}
						registry.byId("contentTabs").addChild(pane);
						registry.byId("contentTabs").selectChild(pane);
					} 
				}
		    },
			closeCurrentTab: function (tabContainerId){
				var tc = dijit.byId(tabContainerId);
				if (tc!=null){
					var select = tc.selectedChildWidget;
					if (select!=null){
						var id = select.id;
						// onclose is not called when tab is closed programmatically, so we repeat onclose content here 
						topic.publish("tab/close", id);
		
						tc.removeChild(select);
						dijit.byId(id).destroy();
					}
				}
			}
		};
	}	
);

function setSubTabContent(pTabContainerId,item) {
    require(["commons/TabUtil"],function(TabUtil){
        TabUtil.setSubTabContent( pTabContainerId, item );
    });
}
