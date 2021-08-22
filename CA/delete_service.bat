set service_name=run_ad_service
net stop %service_name%
sc delete %service_name%

set service_name=run_database_service
net stop %service_name%
sc delete %service_name%

set service_name=run_excel_service
net stop %service_name%
sc delete %service_name%

set service_name=run_faf_service
net stop %service_name%
sc delete %service_name%

set service_name=run_pdf_service
net stop %service_name%
sc delete %service_name%

set service_name=run_recordingagent_service
net stop %service_name%
sc delete %service_name%

set service_name=run_rws_service
net stop %service_name%
sc delete %service_name%

set service_name=run_gateway
net stop %service_name%
sc delete %service_name%
