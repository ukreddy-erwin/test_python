Function RunInstallation()
{
	Log ""
	$sSetup = $sSetup.Trim()
	$sParams = $sParams.Trim()
	If ($sSetup -eq "")
    {
		Log "Setup not specified. Aborting..."
		Return 1
	}

	If (($sSetup -eq "msiexec.exe") -or ($sSetup -eq "cmd.exe"))
    {
		$sSetup = "$System32\$sSetup"
    }
	ElseIf (($sSetup.SubString(1, 2) -ne ":\") -and ($sSetup.SubString(0, 2) -ne "\\"))
    {
		$sSetup = $ScriptDir + "\" + $sSetup
	}
    
	If (!(Test-Path "$sSetup"))
    {
		Log $sSetup + " not found. Aborting..."
		Return 1
	}

	$sErrCodesFile = $ScriptDir + "\" + "ErrorCodes.ini"
	If (!(Test-Path $sErrCodesFile))
    {
		Log "Error Codes file, '$sErrCodesFile', not found. Aborting..."
		Return 1
	}

	If ($sSetup.Contains(" ") -eq $True) {$sSetup = '"' + $sSetup + '"'}

	If (($sParams -eq $Null) -or ($sParams -eq ""))
    {
		Log "Running Command: '$sSetup'"
		$iRet = Start-Process -FilePath $sSetup -NoNewWindow -PassThru -Wait -WindowStyle "Hidden"
    }
	Else
    {
		Log "Running Command: $sSetup $sParams"
		$iRet = Start-Process "$sSetup" -ArgumentList "$sParams" -NoNewWindow -PassThru -Wait
	}

	Log "Return Code from '$sSetup': $($iRet.ExitCode)"

	$sDesc = "" + $(ReadIniValue "$sErrCodesFile" "$sErrSection" "$($iRet.ExitCode)")

	If ($sDesc -eq "")
    {
		$sDesc = "No Description found for the Return Code"
	}
    Else
    {
	   Log "Return Code Description: $sDesc"
    }
    
	If (Test-Path $sLog)
    {
        #set-variable -name aLogs -scope 1 -value ($aLogs + $sLog)
        $global:aLogs = ($aLogs + $sLog)
	}

    $sSuccessfulCodes = ",$sSuccessfulCodes,"
	If ($sSuccessfulCodes.Contains(",$($iRet.ExitCode),") -eq $True)
    {
        Log "Return code category: Successful"
	}
    Else
    {
		Log "Return code category: Failure."
        Log $sTidyDiv
        Return $($iRet.ExitCode)
	}
}

Function Log($sMsg)
{
    $ErrorActionPreference = "SilentlyContinue"
    $sMasterLog = "" + $sMasterLog
    If ($sMasterLog -eq "")
    {
        $sLogDir = "" + $sLogDir
        $sLogName = "" + $sLogName
        If ($sLogDir -eq "") {$sLogDir = $ScriptDir + "\"}
        If (!(Test-Path $sLogDir))
        {
            New-Item -Path $sLogDir -ItemType directory
            If ($? -eq $False)
            {
                $sLogDir = $env:Temp + "\"
            }
        }
        If ($sLogName -eq "") {$sLogName = $ScriptName.Substring(0, $ScriptName.Length - 4) + "_Master_" + $(CurrentDateTime) + ".log"}
        $sMasterLog = $sLogDir + $sLogName
    }
    Add-Content $sMasterLog "$(Get-Date) : $sMsg"
}


#$folder = "C:\Program Files\MateAgent"
#$python_folder = $folder+"\Tools"
#$temp_folder = $folder+"\temp"

$folder = Get-Property -Name CustomActionData
$python_folder= $folder+"\Tools"
$temp_folder = $folder+"temp"

RunInstallation $temp_folder+"\python-3.7.9-amd64.exe" "/quiet InstallAllUsers=1 PrependPath=1 Include_test=0 TargetDir="""+$python_folder+"""" "0,3010" True


