#RequireAdmin

#include <StringConstants.au3>
#include <Misc.au3>
#include <MsgBoxConstants.au3>
#include <FileConstants.au3>
#include <WinAPIFiles.au3>
#include <Date.au3>
#include <file.au3>
#include <Array.au3>
#include <String.au3>

Func FindFileInPath($sFileName, $sPath)	;finds the file specified in the path specified or any of it's subfolders
	Local $aFileList
	$aFileList = _FileListToArrayRec($sPath, $sFileName, 1, 1, 1, 2)
	If UBound($aFileList) > 1 Then
		Return $aFileList[1]
	Else
		Return ""
	EndIf
EndFunc

$ps = FindFileInPath("powershell.exe", @SystemDir&"\WindowsPowerShell")
$temp_cmd = """"& $ps & """ -ExecutionPolicy Bypass -windowstyle hidden -noninteractive -File """ & @ScriptDir & "\checkPython.ps1"""
Local $iReturn = RunWait($temp_cmd,@ScriptDir,@SW_HIDE);, $STDOUT_CHILD)
Logging($temp_cmd & @CRLF)
Logging("Checked python version from environment variable with result code: "&$iReturn)
if FileExists(@WindowsDir&"\Temp\PYTHON_REQUIRED.txt") then
   Logging("Python installation is required")
$vb = FindFileInPath("cscript.exe", @SystemDir)
$temp_cmd = """"& $vb & """ """ & @ScriptDir & "\PythonInstall.vbs"""
Logging($temp_cmd)
Local $iReturn = RunWait($temp_cmd,@ScriptDir,@SW_HIDE); ,$STDOUT_CHILD
FileDelete(@WindowsDir&"\Temp\PYTHON_REQUIRED.txt")
Else
   Logging("Not required to install python")
EndIf


Func Logging($sMessage,$sLogFile=@TempDir&"\checkPython.log")
	If $sMessage = "" Then
		FileWriteLine($sLogFile, "")
	Else
		FileWrite($sLogFile, "Ran by user: " & @UserName & " : " & _NowCalc() & " :: " & $sMessage&@CRLF)
	EndIf
EndFunc