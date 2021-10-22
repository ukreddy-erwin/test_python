Session.Property("CA_InstallPython")=Session.Property("INSTALLDIR")
Session.Property("CA_pipmodules")=Session.Property("INSTALLDIR")
Session.Property("CA_DeleteServiceVBS")=Session.Property("INSTALLDIR")
Session.Property("CA_CreateServiceVBS")=Session.Property("INSTALLDIR")

Session.Property("CA_UpdateYaml")=Session.Property("INSTALLDIR") & "grpc.gateway\config\appConfigDev.yaml" + ";" & Session.Property("INSTALLDIR") & "grpc.gateway\config\appConfigProd.yaml" & ";" & Session.Property("AU_CODE") & ";"   & Session.Property("AU_URL") & ";" & Session.Property("AU_PORT")


