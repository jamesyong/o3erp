/*******************************************************************************
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *******************************************************************************/
package o3erp.plugin;

import java.io.File;

import org.ofbiz.base.container.Container;
import org.ofbiz.base.container.ContainerConfig;
import org.ofbiz.base.container.ContainerException;
import org.ofbiz.base.util.Debug;

import ro.fortsoft.pf4j.DefaultPluginManager;
import ro.fortsoft.pf4j.PluginManager;

/**
 * Container Implementation for Plug-in
 * 
 * @author james.yong
 * 
 */
public class PluginContainer implements Container {
	
	public static final String module = PluginContainer.class.getName();
	
	protected String configFileLocation = null;
	protected String containerName;
	
	PluginManager pluginManager;

	@Override
    public void init(String[] args, String name, String configFile)
			throws ContainerException {
		
		this.containerName = name;
		this.configFileLocation = configFile;
	}

	@Override
    public boolean start() throws ContainerException {
		
        // get the container config
        ContainerConfig.Container cfg = ContainerConfig.getContainer(containerName, configFileLocation);
        String pluginPath = ContainerConfig.getPropertyValue(cfg, "plugin-path", System.getProperty("ofbiz.home")+"/framework/plugin/plugins");

        Debug.logInfo("Plugin Path: "+pluginPath, module);
        
        pluginManager = new DefaultPluginManager(new File(pluginPath));
        pluginManager.loadPlugins();
        pluginManager.startPlugins();
        return false;
	}

	@Override
	public void stop() throws ContainerException {
        pluginManager.stopPlugins();
	}

	@Override
	public String getName() {
        return containerName;
	}

}
