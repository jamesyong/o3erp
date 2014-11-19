package o3erp.plugin.extension;

import java.util.Map;
import java.util.Set;

import ro.fortsoft.pf4j.ExtensionPoint;

public interface IExpression extends ExtensionPoint {
	/**
	 * get the variable names found in an expression
	 * @param expr
	 * @return
	 */
	public Set<String> getVariables(String sExpr);
	
	public String evalExpr(String sExpr, Map<String,Boolean> variables);
}



    

