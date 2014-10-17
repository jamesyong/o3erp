define(["dojo/topic", "dijit/registry", "dojo/ready"], 
	function(topic, registry, ready){
	    return {
			
			newPage: function(tabId, tabMenuList){
				var hands=[];
				
				this.runLocal = function (func){
					return function(){
						if (arguments[0]==tabId){
							console.log("arguments[0]:"+tabId);
							func();
						}
					}
				}
				this.connect = function(key, func, isLocal){
					if (isLocal === undefined || isLocal){
						hands[hands.length] = topic.subscribe(key, this.runLocal(func));
					} else {
						hands[hands.length] = topic.subscribe(key, func);
					}
				}
				
				this.init = function(){
					
					ready(function(){
						
						registry.byId("btnNew").set("disabled", true);
						registry.byId("btnSave").set("disabled", true);
						registry.byId("btnCancel").set("disabled", true);
						
						tabMenuList.map(function(tabMenu){
							switch(tabMenu) {
							    case "new":{
									registry.byId("btnNew").set("disabled", false);
							    } break;
							    case "save":{
									registry.byId("btnSave").set("disabled", false);	        
								} break;
								case "cancel":{
									registry.byId("btnCancel").set("disabled", false);	        
								} break;
								case "all":{
							        registry.byId("btnNew").set("disabled", false);
									registry.byId("btnSave").set("disabled", false);
									registry.byId("btnCancel").set("disabled", false);
								} break;
							}
							
						});
						
					});
					this.connect("tab/menu", function(){
						registry.byId("btnNew").set("disabled", true);
						registry.byId("btnSave").set("disabled", true);
						registry.byId("btnCancel").set("disabled", true);
						
						tabMenuList.map(function(tabMenu){
							switch(tabMenu) {
							    case "new":{
									registry.byId("btnNew").set("disabled", false);
							    } break;
							    case "save":{
									registry.byId("btnSave").set("disabled", false);	        
								} break;
								case "cancel":{
									registry.byId("btnCancel").set("disabled", false);	        
								} break;
								case "all":{
							        registry.byId("btnNew").set("disabled", false);
									registry.byId("btnSave").set("disabled", false);
									registry.byId("btnCancel").set("disabled", false);
								} break;
							}
							
						});
					});
					
					
					
			        this.connect("tab/close", function(){
						console.log("tab/close:"+tabId);
						//unsubscribe when tab is closed
						hands.map(function(handler){handler.remove();});
			    	});
			    }
			    this.init();
								
				return this;
			}

		};
	}	
);

