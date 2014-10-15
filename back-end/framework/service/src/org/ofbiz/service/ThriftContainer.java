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
package org.ofbiz.service;

import o3erp.thrift.BaseService;
import o3erp.thrift.impl.BaseServiceHandler;

import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TServer.Args;
import org.apache.thrift.server.TSimpleServer;
import org.apache.thrift.server.TThreadPoolServer;
import org.apache.thrift.transport.TSSLTransportFactory;
import org.apache.thrift.transport.TSSLTransportFactory.TSSLTransportParameters;
import org.apache.thrift.transport.TServerSocket;
import org.apache.thrift.transport.TServerTransport;
import org.ofbiz.base.container.Container;
import org.ofbiz.base.container.ContainerConfig;
import org.ofbiz.base.container.ContainerException;
import org.ofbiz.base.util.Debug;
import org.ofbiz.base.util.UtilValidate;
import org.ofbiz.entity.Delegator;
import org.ofbiz.entity.DelegatorFactory;

/**
 * Container Implementation for Thrift
 * 
 * @author jamesyong
 * 
 */
public class ThriftContainer implements Container {
	
	public static final String module = ThriftContainer.class.getName();

	protected String configFileLocation = null;
	protected String containerName;
	protected String name;
	protected int port = 9090;
	protected boolean useThreadPool = true;
	protected boolean isSecure = true;
	protected String pathToKeyStore;
	protected BaseService.Processor processor;
	
	protected Delegator delegator;
	
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
        this.delegator = DelegatorFactory.getDelegator(ContainerConfig.getPropertyValue(cfg, "delegator-name", "default"));
        
        // get the app-name
        ContainerConfig.Container.Property appName = cfg.getProperty("app-name");
        if (appName == null || UtilValidate.isEmpty(appName.value)) {
            throw new ContainerException("Invalid app-name defined in container configuration");
        } else {
            this.name = appName.value;
        }
        
        // get the port
        ContainerConfig.Container.Property pPort = cfg.getProperty("port");
        if (pPort == null || UtilValidate.isEmpty(pPort.value)) {
            throw new ContainerException("Invalid port defined in container configuration");
        } else {
            try {
                this.port = Integer.parseInt(pPort.value);
            } catch (Exception e) {
                throw new ContainerException("Invalid port defined in container configuration; not a valid int");
            }
        }
        
        // useThreadPool
        ContainerConfig.Container.Property pUseThreadPool = cfg.getProperty("use-thread-pool");
        if (pUseThreadPool == null || UtilValidate.isEmpty(pUseThreadPool.value)) {
            throw new ContainerException("Invalid use-thread-pool defined in container configuration");
        } else {
            try {
                this.useThreadPool = Boolean.valueOf(pUseThreadPool.value);
            } catch (Exception e) {
                throw new ContainerException("Invalid use-thread-pool defined in container configuration; not a valid boolean");
            }
        }
        
        // isSecure
        ContainerConfig.Container.Property pIsSecure = cfg.getProperty("is-secure");
        if (pIsSecure == null || UtilValidate.isEmpty(pIsSecure.value)) {
            throw new ContainerException("Invalid is-secure defined in container configuration");
        } else {
            try {
                this.isSecure = Boolean.valueOf(pIsSecure.value);
            } catch (Exception e) {
                throw new ContainerException("Invalid is-secure defined in container configuration; not a valid boolean");
            }
        }
        
        // pathToKeyStore
        ContainerConfig.Container.Property pPathToKeyStore = cfg.getProperty("path-to-key-store");
        if (pPathToKeyStore == null || UtilValidate.isEmpty(pPathToKeyStore.value)) {
            throw new ContainerException("Invalid path-to-key-store defined in container configuration");
        } else {
            try {
                this.pathToKeyStore = pPathToKeyStore.value;
            } catch (Exception e) {
                throw new ContainerException("Invalid path-to-key-store defined in container configuration; not a valid String");
            }
        }

        
        processor = new BaseService.Processor(new BaseServiceHandler(delegator));
		if (this.isSecure){
			Runnable secure = new Runnable() {
		        public void run() {
		        	secure(processor);
		        }
		    };
		    new Thread(secure).start();
		} else {
			Runnable simple = new Runnable() {
				public void run() {
					simple(processor);
				}
			};
			new Thread(simple).start();
		}
		return true;
	}

	@Override
	public void stop() throws ContainerException {

	}

	@Override
	public String getName() {
		return containerName;
	}

	public void simple(o3erp.thrift.BaseService.Processor processor) {
		try {
			TServerTransport serverTransport = new TServerSocket(port);
			TServer server = null;
			
			if (useThreadPool) /* multithreaded server */{
				server = new TThreadPoolServer(
						new TThreadPoolServer.Args(serverTransport).processor(processor));
			} else {
				server = new TSimpleServer(
						new Args(serverTransport).processor(processor));
			}
			
			Debug.logInfo("Starting the simple server...", module);

			server.serve();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	public void secure(BaseService.Processor processor) {
	    try {
	      /*
	       * Use TSSLTransportParameters to setup the required SSL parameters. In this example
	       * we are setting the keystore and the keystore password. Other things like algorithms,
	       * cipher suites, client auth etc can be set. 
	       */
	      TSSLTransportParameters params = new TSSLTransportParameters();
	      // The Keystore contains the private key
	      params.setKeyStore(pathToKeyStore, "thrift", null, null);

	      /*
	       * Use any of the TSSLTransportFactory to get a server transport with the appropriate
	       * SSL configuration. You can use the default settings if properties are set in the command line.
	       * Ex: -Djavax.net.ssl.keyStore=.keystore and -Djavax.net.ssl.keyStorePassword=thrift
	       * 
	       * Note: You need not explicitly call open(). The underlying server socket is bound on return
	       * from the factory class. 
	       */
	      TServerTransport serverTransport = TSSLTransportFactory.getServerSocket(port, 0, null, params);
	      TServer server = null;
			
		  if (useThreadPool) /* multithreaded server */{
			  server = new TThreadPoolServer(new TThreadPoolServer.Args(serverTransport).processor(processor));	
		  } else {
			  server = new TSimpleServer(new Args(serverTransport).processor(processor));
		  }
	      
	      Debug.logInfo("Starting the secure server...", module);
	      server.serve();
	    } catch (Exception e) {
	      e.printStackTrace();
	    }
	  }

}
