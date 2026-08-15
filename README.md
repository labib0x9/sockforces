# sockforces
thinking....


# GitHub Setup
### Creating Organization
```
https://github.com/settings/organizations -> New Organization
https://github.com/<org-name> to check if it is created
```
### Creating App
```
https://github.com/organizations/<org-name>/settings/apps -> New GitHub App

goto smee.io, create a channel
paste it on Webhook URL

open terminal
openssl rand -hex 20
paste it on secret

save both webhook link and secret

Repository Permission -> Administration -> Read and Write
Repository Permission -> Contents -> Read and Write
Organization Permissions -> Administration -> Read and Write
Subscribe to events -> mark Push
Where can this GitHub App be installed -> Only on this account
```

### Configuring App
```
https://github.com/organizations/<org-name>/settings/apps/<app-name> -> Get App ID
https://github.com/organizations/<org-name>/settings/apps/<app-name> -> Generate a Private Key -> download the .pem file

https://github.com/organizations/<org-name>/settings/apps/<app-name>/installations -> you will see org name and install -> click install -> you will be redireted to 
https://github.com/organizations/<org-name>/settings/installations/<installation id> -> save installation id
```

### Setup Lab's Repository
```
repo -> settings -> mark Template repository
```

### Webhook Setup
```
if lost webhook secret : goto https://github.com/organizations/<org-name>/settings/apps/<app-name> & set a new one

forward using smee client:
npm install -g smee-client
smee -u https://smee.io/xxxxxxxxx -t http://localhost:8080/submissions/hook
```