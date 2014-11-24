package org.ofbiz.base.util;

import java.util.List;

import o3erp.plugin.PluginContainer;
import ro.fortsoft.pf4j.ExtensionPoint;
import ro.fortsoft.pf4j.PluginManager;

public class UtilPlugin {
	
	public static <T extends ExtensionPoint> T get(Class<T> extensionClass){//
		
		PluginManager pluginManager = PluginContainer.getPluginManager();
		/**
		 * check if we have plugin to evaluate logical expression
		 */
		List<T> exprList = pluginManager.getExtensions(extensionClass);
		if (UtilValidate.isNotEmpty(exprList)){
			return exprList.get(0);
		}
		return null;
	}

}
