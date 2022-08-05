select cm.CUST_IDNO, cm.CASE_NO, cm.STATUS,
		replace(convert(varchar, cm.CREATED_TIME, 112),'-','') 
        replace(convert(varchar, cm.CREATED_TIME ,108),':','') as CREATED_TIME 
from CMN_CASE_MASTER cm
join pln_case_info ci on ci.CASE_NO = cm.CASE_NO 
where cm.case_type = 'A' and ci.NTB_FLAG = 'Y'
and convert(date, cm.CLOSE_TIME) = convert(date, dateadd(day,-1, ?)), DateUtils.getCurrentDate());

select cm.CUST_IDNO, cm.CASE_NO,  cm.STATUS,  
replace(convert(varchar, cm.CREATED_TIME, 112),'-','') +  replace(convert(varchar, cm.CREATED_TIME ,108),':','') as CREATED_TIME   
from CMN_CASE_MASTER cm  
join pln_case_info ci on ci.CASE_NO = cm.CASE_NO   where cm.case_type = 'A'  and ci.NTB_FLAG = 'Y' 
and convert(date, cm.CLOSE_TIME) = convert(date, dateadd(day,-1, 'Wed Jul 27 14:22:13 CST 2022'))


