 param (
    [Parameter(Mandatory=$true)][string]$file
 )

Set-AuthenticodeSignature $file -Certificate (Get-ChildItem Cert:\CurrentUser\My -CodeSigningCert)
