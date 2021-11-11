param (
    [string]$Password,
    [string]$Username,
    [string]$Url,
    [string]$User,
    [string]$IsItPermanent
)
Try
{
    [System.Threading.Thread]::CurrentThread.CurrentUICulture = "en-US"
    $Permanent = [System.Convert]::ToBoolean($IsItPermanent)
    Set-ExecutionPolicy RemoteSigned
    $pass = ConvertTo-SecureString -AsPlainText $Password -Force
    $Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass
    $Session = New-PSSession -ConfigurationName Microsoft.Exchange -ConnectionUri $Url -Authentication Kerberos -Credential $Cred
    Import-PSSession $Session -DisableNameChecking 
    Remove-Mailbox -Identity $User -Permanent $Permanent -Confirm:$false
}
Catch
{
    Write-Output "Error Message:"
    Write-Output $Error[0].Exception
}