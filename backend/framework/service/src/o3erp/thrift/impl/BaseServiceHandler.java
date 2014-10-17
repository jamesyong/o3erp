package o3erp.thrift.impl;

import java.util.HashMap;
import java.util.Locale;
import java.util.Map;

import org.apache.thrift.TException;
import org.ofbiz.base.util.UtilHttp;
import org.ofbiz.base.util.UtilMisc;
import org.ofbiz.base.util.UtilProperties;
import org.ofbiz.base.util.UtilValidate;
import org.ofbiz.entity.Delegator;
import org.ofbiz.entity.GenericValue;
import org.ofbiz.service.GenericServiceException;
import org.ofbiz.service.LocalDispatcher;
import org.ofbiz.service.ModelService;
import org.ofbiz.service.ServiceContainer;
import org.ofbiz.service.ServiceUtil;

import o3erp.thrift.BaseService;

public class BaseServiceHandler implements BaseService.Iface {
	
	protected Delegator delegator;

	
	public BaseServiceHandler(Delegator delegator){
		this.delegator = delegator;
	}

	@Override
	public Map<String,String> userLogin(String loginName, String loginPwd) throws TException {
	
		Map<String,String> resultMap = new HashMap();
		if (!UtilValidate.isEmpty(loginName) && !UtilValidate.isEmpty(loginPwd)){
			try {
				LocalDispatcher dispatcher = ServiceContainer.getLocalDispatcher(this.delegator.getDelegatorName(), delegator);

				Map result = dispatcher.runSync("userLogin", UtilMisc.toMap("login.username", loginName, "login.password", loginPwd));
				if (ModelService.RESPOND_SUCCESS.equals(result.get(ModelService.RESPONSE_MESSAGE))) {

					resultMap.put(ModelService.RESPONSE_MESSAGE, ModelService.RESPOND_SUCCESS);
					String successMessage = (String)result.get(ModelService.SUCCESS_MESSAGE);
					if (UtilValidate.isEmpty(successMessage)){
						successMessage = "Login Successful";
					}
					resultMap.put(ModelService.SUCCESS_MESSAGE, successMessage);
					
					GenericValue userLogin = (GenericValue)result.get("userLogin");
					resultMap.put("partyId", userLogin.getString("partyId"));
					
				} else {
					System.out.println(456);
					resultMap.put(ModelService.RESPONSE_MESSAGE, ModelService.RESPOND_ERROR);
					resultMap.put(ModelService.ERROR_MESSAGE, ServiceUtil.getErrorMessage(result));
				}
			} catch (GenericServiceException e) {
				resultMap.put(ModelService.RESPONSE_MESSAGE, ModelService.RESPOND_ERROR);
				resultMap.put(ModelService.ERROR_MESSAGE, e.getMessage());
			}
		}
				
		return resultMap;
	}

}
