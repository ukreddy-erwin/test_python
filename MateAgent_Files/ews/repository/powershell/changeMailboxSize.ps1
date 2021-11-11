param (
    [string]$Password,
    [string]$Username,
    [string]$Url,
    [string]$User,
    [string]$WarningQuota,
    [string]$SendQuota,
    [string]$ReceiveQuota
)
Try
{
    [System.Threading.Thread]::CurrentThread.CurrentUICulture = "en-US"
    $WQuata,$SQuota,$RQuota

    if($WarningQuota.Contains("mb")){
	$WQuata = [int64]$WarningQuota.Replace('mb','') * 1MB
    } elseif($WarningQuota.Contains("kb")){
        $WQuata = [int64]$WarningQuota.Replace('kb','') * 1KB
    } elseif($WarningQuota.Contains("gb")){
        $WQuata = [int64]$WarningQuota.Replace('gb','') * 1GB
    }
    if($SendQuota.Contains("mb")){
	$SQuota = [int64]$SendQuota.Replace('mb','') * 1MB
    } elseif($SendQuota.Contains("kb")){
        $SQuota = [int364]$SendQuota.Replace('kb','') * 1KB
    } elseif($SendQuota.Contains("gb")){
        $SQuota = [int64]$SendQuota.Replace('gb','') * 1GB
    }
    if($ReceiveQuota.Contains("mb")){
	$RQuota = [int64]$ReceiveQuota.Replace('mb','') * 1MB
    } elseif($ReceiveQuota.Contains("kb")){
        $RQuota = [int64]$ReceiveQuota.Replace('kb','') * 1KB
    } elseif($ReceiveQuota.Contains("gb")){
        $RQuota = [int64]$ReceiveQuota.Replace('gb','') * 1GB
    }
    Set-ExecutionPolicy RemoteSigned
    $pass = ConvertTo-SecureString -AsPlainText $Password -Force
    $Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass
    $Session = New-PSSession -ConfigurationName Microsoft.Exchange -ConnectionUri $Url -Authentication Kerberos -Credential $Cred
    Import-PSSession $Session -DisableNameChecking 
    Set-Mailbox -Identity $User -IssueWarningQuota $WQuata -ProhibitSendQuota $SQuota -ProhibitSendReceiveQuota $RQuota -UseDatabaseQuotaDefaults $false
}
Catch
{
    Write-Output "Error Message:"
    Write-Output $Error[0].Exception
}