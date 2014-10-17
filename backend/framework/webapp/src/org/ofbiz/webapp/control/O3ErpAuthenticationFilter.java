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
package org.ofbiz.webapp.control;

import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;

public class O3ErpAuthenticationFilter implements Filter {
	 
    public void destroy() {
                 
    }
 
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse,
            FilterChain filterChain) throws IOException, ServletException {

        if(servletRequest instanceof HttpServletRequest) {

            HttpServletRequest httpServletRequest = (HttpServletRequest)servletRequest;

            O3ErpRequestWrapper request = new O3ErpRequestWrapper(httpServletRequest, httpServletRequest.getHeader("REMOTE_USER"));

            filterChain.doFilter(request, servletResponse);
        } else {
            //otherwise, continue on in the chain with the ServletRequest and ServletResponse objects
            filterChain.doFilter(servletRequest, servletResponse);
        }      
         
        return;
    }
 
    public void init(FilterConfig filterConfig) throws ServletException {
         
    }
 
}
