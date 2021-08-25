set main=%~dp0
"%main%nssm.exe" stop MAGateway
"%main%nssm.exe" stop MAAD
"%main%nssm.exe" stop MADB
"%main%nssm.exe" stop MAEXCEL
"%main%nssm.exe" stop MAPDF
"%main%nssm.exe" stop MAFAF
"%main%nssm.exe" stop MARWS
"%main%nssm.exe" stop MASDP
"%main%nssm.exe" stop MARA


"%main%nssm.exe" remove MAGateway confirm
"%main%nssm.exe" remove MAAD confirm
"%main%nssm.exe" remove MADB confirm
"%main%nssm.exe" remove MAEXCEL confirm
"%main%nssm.exe" remove MAPDF confirm
"%main%nssm.exe" remove MAFAF confirm
"%main%nssm.exe" remove MARWS confirm
"%main%nssm.exe" remove MASDP confirm
"%main%nssm.exe" remove MARA confirm