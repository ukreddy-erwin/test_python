set main=C:\Program Files\MateAgent\

"%main%nssm.exe" install MAGateway "%main%grpc.gateway\cmd\service\service.exe"
"%main%nssm.exe" start MAGateway

"%main%nssm.exe" install MAAD "%main%ad\cmd\cmd.exe"
"%main%nssm.exe" start MAAD

"%main%nssm.exe" install MADB "%main%database\cmd\cmd.exe"
"%main%nssm.exe" start MADB

"%main%nssm.exe" install MAEXCEL "%main%excel\cmd\cmd.exe"
"%main%nssm.exe" start MAEXCEL

"%main%nssm.exe" install MAPDF "%main%pdf\cmd\cmd.exe"
"%main%nssm.exe" start MAPDF

"%main%nssm.exe" install MAFAF "%main%faf\cmd\cmd.exe"
"%main%nssm.exe" start MAFAF

"%main%nssm.exe" install MARWS "%main%rws\cmd\cmd.exe"
"%main%nssm.exe" start MARWS

"%main%nssm.exe" install MASDP "%main%sdp\cmd\cmd.exe"
"%main%nssm.exe" start MASDP

"%main%nssm.exe" install MARA "%main%recordingagent\cmd\cmd.exe"
"%main%nssm.exe" start MARA

"%main%nssm.exe" install MAEWS "%main%ews\cmd\cmd.exe"
"%main%nssm.exe" start MAEWS

"%main%nssm.exe" install MAMD "%main%md\cmd\cmd.exe"
"%main%nssm.exe" start MAMD

"%main%nssm.exe" install MADA "%main%desktopagent\cmd\cmd.exe"
"%main%nssm.exe" start MADA