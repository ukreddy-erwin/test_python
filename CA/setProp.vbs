Session.Property("CA_UpdateYaml")=Session.Property("INSTALLDIR") & "grpc.gateway\config\appConfigDev.yaml" + ";" & Session.Property("INSTALLDIR") & "grpc.gateway\config\appConfigProd.yaml" & ";" & Session.Property("AU_CODE") & ";"   & Session.Property("AU_URL") & ";" & Session.Property("AU_PORT")

MsgBox(Session.Property("CA_UpdateYaml"))
MsgBox(Session.Property("AU_CODE"))
MsgBox(Session.Property("AU_URL"))
MsgBox(Session.Property("AU_PORT"))