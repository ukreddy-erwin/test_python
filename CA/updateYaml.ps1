#$MyDir = [System.IO.Path]::GetDirectoryName($myInvocation.MyCommand.Definition)
#$text = Get-Content -Path @($($MyDir+"\\input.yml"))

$tmp = $($env:temp)+"\CustomDataMateAgent.log" #New-TemporaryFile
function WriteLog($sMsg)
{
	$sMsg | Out-File $tmp -Append
}

$propertyset = Get-Property -Name CustomActionData
WriteLog "Data received to update: $propertyset"
$propertyset = $propertyset.Split(";")

$text1 = Get-Content -Path @($propertyset[0])
$text2 = Get-Content -Path @($propertyset[1])

$au_code = $propertyset[2]#"JohnSOn and Johnson"
$au_url = $propertyset[3]#"http://google.com"
$au_port = $propertyset[4]#"5678"


#echo "---Updating content---"
#$text.automCenterConfig.code = "value1"
#$text.automCenterConfig.path = "value2"
#$text.automCenterConfig.ext = "value3"

#$txt = $text | Select-String -Pattern 'gatewayGrpcConfig: &gatewayGrpcConfig'

function updateFile($yaml_file,$text)
{
 
WriteLog "Updating $yaml_file" 
 
 $match = $false
$count = 0

foreach($line in $text) {
    
    if($line.trim().StartsWith('automCenterConfig:')){
        $match = $true
        break
    }
    $count++
    }
    
    if($match){
echo $count
echo $text[$count]
$text[$count+1]="  code: $au_code"
$text[$count+2]="  url: $au_url"
$text[$count+3]="  port: $au_port"
}
else
{
  echo "----------No matches------"
  #$text = $text + "automCenterConfig: &automCenterConfig"+"  code: $au_code"+"  url: $au_url"+"  port: $au_port"
  $text =  "automCenterConfig: &automCenterConfig`r"+"  code: $au_code`r"+"  url: $au_url`r"+"  port: $au_port`r" + ($text -Join "`n")
}

$text | Out-File "$yaml_file"

}

updateFile $($propertyset[0]) $text1
updateFile $($propertyset[1]) $text2


