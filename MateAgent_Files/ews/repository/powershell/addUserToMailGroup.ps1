param (
    [string]$Password,
    [string]$Username,
    [string]$Url,
    [string]$Member,
    [string]$GroupName
)
Try
{
    [System.Threading.Thread]::CurrentThread.CurrentUICulture = "en-US"
    Set-ExecutionPolicy RemoteSigned
    $pass = ConvertTo-SecureString -AsPlainText $Password -Force
    $Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass
    $Session = New-PSSession -ConfigurationName Microsoft.Exchange -ConnectionUri $Url -Authentication Kerberos -Credential $Cred
    Import-PSSession $Session -DisableNameChecking 
    Add-DistributionGroupMember -Identity $GroupName -Member $Member
}
Catch
{
    Write-Output "Error Message:"
    Write-Output $Error[0].Exception
}