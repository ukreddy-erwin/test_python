param (
    [string]$Password,
    [string]$Username,
	[string]$Url,
	[string]$Script
)
Try
{
   Set-ExecutionPolicy RemoteSigned
    $pass = ConvertTo-SecureString -AsPlainText $Password -Force
    $Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass
    $Session = New-PSSession -ConfigurationName Microsoft.Exchange -ConnectionUri http://Questsrv.itherotest.com/PowerShell/ -Authentication Kerberos -Credential $Cred
    Import-PSSession $Session -DisableNameChecking
    $Script
}
Catch
{
    Write-Output $Error[0].Exception
}

