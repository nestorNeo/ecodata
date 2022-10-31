### Instructions 
Users will use jwt to auth the rest api, 
so each user or device must have a token when security enable see 

```

 .\security.exe -email test@ecodata.leofigy.xyz -key C:\Users\leofi\repos\ecodata\security\authpriv
TOKEN ------------------------
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZWNvZGF0YS5sZW9maWd5Lnh5eiIsImJpcnRoZGF5Ijo4Mjc0NTI4MDAsImV4cCI6MTY3NTEzODI4MCwiaXNzIjoiRnJpZW5kcyBvZiBHbyJ9.v-b-qyG6K6PAguXEy86Kjn4zAjBHCYlmS42YzRVe3kbrjgwGDd9qQ8-2woh0ie_O308Fqt6gDfIoV0Vougb1P-oz-yNkwhiSLCz7tNc5QbfDilQ8-7MirdwkFY8ot5iX0IUDw4sB1Auso55U_B1lBvBhQ-Q2X3CA4f6fq81nRZ-3c1CRcGsOxD_wJ7DDULh92CeuMc0Cn6O_G29hVetlDV6Sb6i76BCN_qjI5MI1QWwL3puY_JAbTMkq28DIbGx6ni9Rq2Hg72jHQyrfakeS0-_uqf-WsXF7gELlsz4ewMDabXhw6BBywxaXkvvR2mivnh7ZYPMhs41F9vopI9eRhA
TOKEN ------------------------
map[birthday:8.274528e+08 email:test@ecodata.leofigy.xyz exp:1.67513828e+09 iss:Friends of Go]

```

sample usage
```
 curl localhost:443/api/v1/ -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZWNvZGF0YS5sZW9maWd5Lnh5eiIsImJpcnRoZGF5Ijo4Mjc0NTI4MDAsImV4cCI6MTY3NTEzODI4MCwiaXNzIjoiRnJpZW5kcyBvZiBHbyJ9.v-b-qyG6K6PAguXEy86Kjn4zAjBHCYlmS42YzRVe3kbrjgwGDd9qQ8-2woh0ie_O308Fqt6gDfIoV0Vougb1P-oz-yNkwhiSLCz7tNc5QbfDilQ8-7MirdwkFY8ot5iX0IUDw4sB1Auso55U_B1lBvBhQ-Q2X3CA4f6fq81nRZ-3c1CRcGsOxD_wJ7DDULh92CeuMc0Cn6O_G29hVetlDV6Sb6i76BCN_qjI5MI1QWwL3puY_JAbTMkq28DIbGx6ni9Rq2Hg72jHQyrfakeS0-_uqf-WsXF7gELlsz4ewMDabXhw6BBywxaXkvvR2mivnh7ZYPMhs41F9vopI9eRhA"

```