param (
    [string]$Password,
    [string]$Username,
    [string]$Url,
    [string]$Name,
    [string]$UserPrincipalName,
    [string]$NewPassword,
    [string]$Database,
    [string]$FirstName,
    [string]$LastName,
    [string]$DisplayName,
    [string]$Ou
)
Try
{
    [System.Threading.Thread]::CurrentThread.CurrentUICulture = "en-US"
    Set-ExecutionPolicy RemoteSigned
    $pass = ConvertTo-SecureString -AsPlainText $Password -Force
    $Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass
    $Session = New-PSSession -ConfigurationName Microsoft.Exchange -ConnectionUri $Url -Authentication Kerberos -Credential $Cred
    Import-PSSession $Session -DisableNameChecking 
    if($Ou -ne ""){
        if($Database -ne ""){
            New-Mailbox -Name $Name -UserPrincipalName $UserPrincipalName -Password (ConvertTo-SecureString -String $NewPassword -AsPlainText -Force) -Database $Database -FirstName $FirstName -LastName $LastName -DisplayName $DisplayName -OrganizationalUnit $Ou 
        }else{
            New-Mailbox -Name $Name -UserPrincipalName $UserPrincipalName -Password (ConvertTo-SecureString -String $NewPassword -AsPlainText -Force) -FirstName $FirstName -LastName $LastName -DisplayName $DisplayName -OrganizationalUnit $Ou 
        }        
    }else{
        if($Database -ne ""){
            New-Mailbox -Name $Name -UserPrincipalName $UserPrincipalName -Password (ConvertTo-SecureString -String $NewPassword -AsPlainText -Force) -Database $Database -FirstName $FirstName -LastName $LastName -DisplayName $DisplayName
        }else{
            New-Mailbox -Name $Name -UserPrincipalName $UserPrincipalName -Password (ConvertTo-SecureString -String $NewPassword -AsPlainText -Force) -FirstName $FirstName -LastName $LastName -DisplayName $DisplayName 
        }
    }
}
Catch
{
    Write-Output "Error Message:"
    Write-Output $Error[0].Exception
}