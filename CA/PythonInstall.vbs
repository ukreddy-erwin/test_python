Dim oFso,oShell,sSetup,sParams,iRet,sLine,sLogName,sLogDir,sMasterLog,oMasterLog,sSuccessfulCodes,folder,python_folder,temp_folder

Set oFso = CreateObject("Scripting.FileSystemObject")
Set oShell = CreateObject("WScript.Shell")

Const iTidyDivLen = 127

folder = "C:\Program Files\MateAgent\"
'folder = Session.Property("CustomActionData")
python_folder = folder&"Tools"
'temp_folder = folder&"temp"
temp_folder = oFso.GetParentFolderName(WScript.ScriptFullName)
'RunInstallation temp_folder&"\python-3.7.9-amd64.exe","/quiet InstallAllUsers=1 PrependPath=1 Include_test=0 TargetDir="""&python_folder&"""","0,3010",True



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