
SELECT t.* FROM tenants t INNER JOIN apttenants at ON ( t.tenantID == at.tenantID )

