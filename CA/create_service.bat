set main=%~dp0

set service_name=run_ad_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_database_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_excel_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_faf_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_pdf_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_recordingagent_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_rws_service
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"

set service_name=run_gateway
sc create %service_name% displayname= %service_name% error= severe binpath= "%main%services\%service_name%.bat"
