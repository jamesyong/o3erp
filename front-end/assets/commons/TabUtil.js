define(["dojox/layout/ContentPane", "dojo/topic"], 
	function(ContentPane, topic){
	    return {
			setTabContent: function (registry, tabContainerId, item){
		        var pane = dijit.byId(item.id)
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
				    // add the new pane to our contentTabs widget
				    registry.byId("contentTabs").addChild(pane);
				}
				registry.byId("contentTabs").selectChild(pane);
		    },
			closeCurrentTab: function (tabContainerId){
				var tc = dijit.byId(tabContainerId);
				var id = tc.selectedChildWidget.id;
				// onclose is not called when tab is closed programmatically, so we repeat onclose content here 
				topic.publish("tab/close", id);

				tc.removeChild(tc.selectedChildWidget);
				dijit.byId(id).destroy();
			}
		};
	}	
);