$tmp = $($env:temp)+"\upgradelogic.log" #New-TemporaryFile
function WriteLog($sMsg)
{
	$sMsg | Out-File $tmp -Append
}

#Set-Property -Name PYTHON_REQUIRED -Value 1
$PYTHON_REQUIRED=$true
$pythonVersion = $(python --version)

if ($pythonVersion -eq $null)
{
  WriteLog "Python not installed"
  #$true | Out-File $($($env:windir)+"\temp\PYTHON_REQUIRED.txt")
  #Exit
}
elseif($pythonVersion.startswith('Python')){
  $pythonVersionNumber = $pythonVersion -replace "^\w*\s"
  if ($pythonVersionNumber -ne $null){
  if([System.Version]$pythonVersionNumber -ge [System.Version]"3.7.0"){
    #Set-Property -Name PYTHON_REQUIRED -Value 0
    $PYTHON_REQUIRED=$false
  }
  }
}

if($PYTHON_REQUIRED -eq $true)
{
    $true | Out-File $($($env:windir)+"\temp\PYTHON_REQUIRED.txt")
}
else{
    echo "Not required"
}




