Dim oFso,oShell,sSetup,sParams,iRet,sLine,sLogName,sLogDir,sMasterLog,oMasterLog,sSuccessfulCodes,id,main,service_name,service_path,rev,rev1

Set oFso = CreateObject("Scripting.FileSystemObject")
Set oShell = CreateObject("WScript.Shell")

Const iTidyDivLen = 127

main = Session.Property("CustomActionData")
'main = oFso.GetParentFolderName(WScript.ScriptFullName) & "\"
'service_name = "MAGateway"
'service_path = main&"gateway\cmd\service\service.exe"

InstallService "MAGateway",main&"grpc.gateway\cmd\service\service.exe"
InstallService "MAAD",main&"ad\cmd\cmd.exe"
InstallService "MADB",main&"database\cmd\cmd.exe"
InstallService "MAEXCEL",main&"excel\cmd\cmd.exe"
InstallService "MAPDF",main&"pdf\cmd\cmd.exe"
InstallService "MAFAF",main&"faf\cmd\cmd.exe"
InstallService "MARWS",main&"rws\cmd\cmd.exe"
InstallService "MASDP",main&"sdp\cmd\cmd.exe"
InstallService "MARA",main&"recordingagent\cmd\cmd.exe"
'New services
InstallService "MAEWS",main&"ews\cmd\cmd.exe"
InstallService "MAMD",main&"md\cmd\cmd.exe"
InstallService "MADA",main&"desktopagent\cmd\cmd.exe"

Sub InstallService(service_name,service_path)
 rev = RunInstallation(main&"nssm.exe","install "&service_name&" """&service_path&"""","0,3010",False)
 Log "Service install status "&rev
 rev1 = RunInstallation(main&"nssm.exe","start """&service_name&"""","0,3010",False)
 Log "service start status "&rev1
End Sub

Function RunInstallation(sSetup, sParams, sSuccessfulCodes, bExitOnFailure)
	Log ""
	Log String(iTidyDivLen, "-")
	sSetup = Trim(sSetup)
	sParams = Trim(sParams)
	If sSetup = "" Then
		Log "Setup not specified. Aborting..."
		WScript.Quit(1)
	End If

	If (LCase(sSetup) = "msiexec.exe") Or (LCase(sSetup) = "cmd.exe")  Or (LCase(sSetup) = "wscript.exe")  Or (LCase(sSetup) = "cscript.exe") Then
		sSetup = oShell.ExpandEnvironmentStrings("%WinDir%") & "\System32\" & sSetup
	ElseIf (Mid(sSetup, 2, 2) <> ":\") And (Left(sSetup, 2) <> "\\") Then
		sSetup = Left(WScript.ScriptFullName, InStrRev(WScript.ScriptFullName, "\")) & sSetup
	End If

	If Not oFso.FileExists(sSetup) Then
		Log sSetup & " not found. Aborting..."
		WScript.Quit(1)
	End If

	If InStr(sSetup, " ") <> 0 Then sSetup = """" & sSetup & """"

	If IsEmpty(sParams) Then
		Log "Running Command: " & sSetup
		iRet = oShell.Run(sSetup, 0, True)
	Else
		Log "Running Command: " & sSetup & " " & sParams
		iRet = oShell.Run(sSetup & " " & sParams, 0, True)
	End If

	Log String(iTidyDivLen, "-")
	Log "Return Code from " & sSetup & ": " & iRet
	Log String(iTidyDivLen, "-")

	If InStr("," & sSuccessfulCodes & ",", "," & iRet & ",") = 0 Then
		Log "Return code category: Failure."
		If bExitOnFailure Then
			WScript.Quit(iRet)
		End If
	Else
		Log "Return code category: Successful"
	End If

	Log String(iTidyDivLen, "-")
	Log ""

	RunInstallation = iRet
End Function

Sub Log(sLine)
	On Error Resume Next
	If Session.Property("ProductCode") = "" Then
		sLogDir = "" & sLogDir
		sLogName = "" & sLogName
		If sLogDir = "" Then sLogDir = oFso.GetParentFolderName(WScript.ScriptFullName)
		If Not oFso.FolderExists(sLogDir) Then
			oShell.Run "cmd.exe /c MD """ & sLogDir & """", 0, True
			If Not oFso.FolderExists(sLogDir) Then
				sLogDir = oShell.ExpandEnvironmentStrings("%TEMP%")
			End If
		End If
		If sLogName = "" Then sLogName = Left(WScript.ScriptName, Len(WScript.ScriptName) - 4) & "_Master_" & Replace(Replace(Replace(Now, " ", "_"), ":", "."), "/", ".") & ".log"
		sMasterLog = sLogDir & "\" & sLogName
		Set oMasterLog = oFso.OpenTextFile(sMasterLog, 8, True)
		Err.Clear
		oMasterLog.WriteLine Now & " : " & sLine
		oMasterLog.Close
		If Err.Number <> 0 Then WScript.Quit
		oMasterLog.Close
		Set oMasterLog = Nothing
	Else
		Dim oRec
		Set oRec = Session.Installer.CreateRecord(1)
		oRec.StringData(1) = Now & " : " & sLine
		Session.Message &H04000000, oRec
	End If
End Sub